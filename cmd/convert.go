package cmd

import (
	"debris/core"
	"debris/model/billfish"
	"debris/model/eagle"
	"debris/model/pixcall"
	"debris/pkg"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/mozillazg/go-pinyin"
	"github.com/spf13/cobra"
	"golang.org/x/exp/rand"
	// 显式导入 gorm 包
)

var (
	chineseRegex = regexp.MustCompile("[\u4e00-\u9fa5]") // 提前编译正则表达式，提高效率, 仅编译一次
	defaultArgs  = pinyin.NewArgs()                      // 默认参数，只需要创建一次
)

var (
	fromLibraryType, toLibraryType, inputFolder, outputFolder string
)

var rootCmd = &cobra.Command{
	Use:   "debris",
	Short: "Converts data from one format to another",
	Long:  `Converts data from one format (e.g., billfish) to another (e.g., pixcall).`, //init: Initialize a new lib-bridge project
	// 示例用法 需要与 Flags 对应
	Example: `
  debris --from billfish --to pixcall --input ./oldFolder.billfish --output ./newFolder
  debris -f billfish -t pixcall -i ./oldFolder.billfish -o ./newFolder
`,
	Run: func(cmd *cobra.Command, args []string) {
		Convert(cmd, args) // 调用 covert 函数
	},
}

func Execute() {
	fmt.Println("Welcome to lib-bridge")

	rootCmd.Flags().StringVarP(&fromLibraryType, "from", "f", "", "Source format (e.g., billfish)")
	rootCmd.Flags().StringVarP(&toLibraryType, "to", "t", "", "Destination format (e.g., pixcall)")
	rootCmd.Flags().StringVarP(&inputFolder, "input", "i", "", "Input file path")
	rootCmd.Flags().StringVarP(&outputFolder, "output", "o", "", "Output file path")
	//Mark 属性设置成必须
	rootCmd.MarkFlagRequired("from")
	rootCmd.MarkFlagRequired("to")
	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("output")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Convert 函数，处理数据库转换逻辑
func Convert(cmd *cobra.Command, args []string) {
	// 检查输入参数是否有效
	if fromLibraryType == "" || toLibraryType == "" || inputFolder == "" || outputFolder == "" {
		log.Println("Error: Missing required arguments.")
		cmd.Help() // 显示帮助信息
		return
	}

	// 打印配置信息，方便调试
	fmt.Printf("Converting from: %s\n", fromLibraryType)
	fmt.Printf("Converting to: %s\n", toLibraryType)
	fmt.Printf("Input folder: %s\n", inputFolder)
	fmt.Printf("Output folder: %s\n", outputFolder)

	// 执行转换逻辑
	switch fromLibraryType {
	case "billfish":
		switch toLibraryType {
		case "pixcall":
			err := billfishToPixcall(inputFolder, outputFolder)
			if err != nil {
				log.Fatalf("Error converting from Billfish to Pixcall: %v\n", err)
			}
			fmt.Println("Successfully converted from Billfish to Pixcall!")
		case "eagle":
			err := billfishToEagle(inputFolder, outputFolder)
			if err != nil {
				log.Fatalf("Error converting from Billfish to Eagle: %v\n", err)
			}
			fmt.Println("Successfully converted from Billfish to Eagle!")
		default:
			log.Fatalf("Error: Unsupported conversion from %s to %s\n", fromLibraryType, toLibraryType)
		}
	case "pixcall":
		switch toLibraryType {
		case "billfish":
			// 实现 pixcallToBillfish 函数并调用
			// err := pixcallToBillfish(inputFile, outputFile)
			// if err != nil {
			// 	log.Fatalf("Error converting from Pixcall to Billfish: %v\n", err)
			// }
			fmt.Println("Successfully converted from Pixcall to Billfish!")
		default:
			log.Fatalf("Error: Unsupported conversion from %s to %s\n", fromLibraryType, toLibraryType)
		}
	default:
		log.Fatalf("Error: Unsupported source format: %s\n", fromLibraryType)
	}
}

/*
Billfish to Pixcall
@params

- billfishPath: existing billfish library folder path
- pixcallPath: new pixcall library folder path
*/
func billfishToPixcall(billfishPath string, pixcallPath string) (err error) {
	billfishDB := core.GetDB(filepath.Join(billfishPath, ".bf", "billfish.db")) // 使用 filepath.Join

	err = os.MkdirAll(filepath.Join(pixcallPath, ".pixcall", "database"), 0700) // 使用 filepath.Join
	if err != nil {
		return err
	}

	pixcallDB := core.GetDB(filepath.Join(pixcallPath, ".pixcall", "database", "main.db")) // 使用 filepath.Join

	// Migrate the schema
	err = pixcallDB.AutoMigrate(
		&pixcall.BoardEntry{},
		&pixcall.Board{},
		&pixcall.Entry{},
		&pixcall.Exif{},
		&pixcall.Folders{},
		&pixcall.Kvs{},
		&pixcall.Tag{},
		&pixcall.Media{},
		&pixcall.RemoteEvents{},
		&pixcall.TagGroups{},
	)
	if err != nil {
		return err // 检查 AutoMigrate 的错误
	}

	// Copy all Tags
	var billfishTags []billfish.Tag
	result := billfishDB.Model(&billfish.Tag{}).Find(&billfishTags) // 直接使用 result
	if result.Error != nil {
		return fmt.Errorf("querying billfish tags: %w", result.Error)
	}

	var pixcallTags []pixcall.Tag

	for _, tag := range billfishTags {
		var tagNamePinyin string

		if chineseRegex.MatchString(tag.Name) { // 使用正则进行匹配
			pinyinSlice := pinyin.Pinyin(tag.Name, defaultArgs) // 使用复用的参数
			for _, pinyinItem := range pinyinSlice {
				tagNamePinyin += strings.Join(pinyinItem, "")
			}
		}

		firstLetter := determineFirstLetter(tag.Name, tagNamePinyin)

		pixcallTags = append(pixcallTags, pixcall.Tag{ // 构建切片
			ID:         randomNumber(),
			Name:       tag.Name,
			Pinyin:     tagNamePinyin,
			Category:   firstLetter,
			ExternalID: (string)(tag.ID),
		})
	}

	// 批量插入 pixcall tags
	result = pixcallDB.Create(&pixcallTags)
	if result.Error != nil {
		fmt.Println("Error creating pixcall tags:", result.Error)
		return result.Error
	}

	// 新建 Pixcall 根文件夹
	err = pixcallDB.Create(&pixcall.Folders{
		EntryID:     1,
		FileCount:   0,
		FolderCount: 0,
		FileSize:    0,
		Ranking:     0,
	}).Error
	if err != nil {
		return fmt.Errorf("creating pixcall root folder: %w", err)
	}

	// 新建 Pixcall 根 Entry
	err = pixcallDB.Create(&pixcall.Entry{
		Name:        "Pixcall",
		ID:          1,
		ParentID:    -1,
		Kind:        0,
		Size:        0,
		ContentType: "unknown",
		IsHidden:    0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Ranking:     9000000000000000000,
		Status:      64,
		IsDeleted:   0,
	}).Error
	if err != nil {
		return fmt.Errorf("creating pixcall root entry: %w", err)
	}

	// 处理 files
	var pixcallEntries []pixcall.Entry
	var billfishFiles []billfish.File
	result = billfishDB.Preload("MaterialUserData").Preload("Tags").Model(&billfish.File{}).Find(&billfishFiles)
	if result.Error != nil {
		return fmt.Errorf("querying billfish files for eagle conversion func: %w", result.Error)
	}

	for _, file := range billfishFiles {
		var fileNamePinyin string

		if chineseRegex.MatchString(file.Name) {
			pinyinSlice := pinyin.Pinyin(file.Name, defaultArgs)
			for _, pinyinItem := range pinyinSlice {
				fileNamePinyin += strings.Join(pinyinItem, "")
			}
		}

		var billfishTagIDs []string
		for _, tag := range file.Tags {
			billfishTagIDs = append(billfishTagIDs, (string)(tag.ID))
		}
		var tags []pixcall.Tag
		result = pixcallDB.Model(&pixcall.Tag{}).Where("external_id IN (?)", billfishTagIDs).Find(&tags)
		if result.Error != nil {
			return fmt.Errorf("querying pixcall tags for external IDs: %w", result.Error)
		}

		var pixcallIDs []string
		for _, tag := range tags {
			pixcallIDs = append(pixcallIDs, (string)(tag.ID))
		}

		pixcallEntries = append(pixcallEntries, pixcall.Entry{
			ID:          randomNumber(),
			Name:        file.Name,
			ParentID:    1,
			Kind:        1,
			Size:        0,
			ContentType: "unknown",
			IsHidden:    0,
			NamePinyin:  fileNamePinyin,
			CreatedAt:   time.UnixMicro(file.Born),
			UpdatedAt:   time.UnixMicro(file.Mtime),
			Tags:        strings.Join(pixcallIDs, "|"),
			Ranking:     1,
			Status:      2069,
			IsDeleted:   0,
		})
	}

	result = pixcallDB.Create(&pixcallEntries)
	if result.Error != nil {
		return fmt.Errorf("creating pixcall files: %w", result.Error)
	}

	// // Copy all Folders (注释掉的代码如果不需要就删除，保留在这里只是为了展示删除线)
	// billfishFolders := billfishDB.Model(&billfish.Folder{}).Find(&billfish.Folder{})
	// for _, folder := range billfishFolders {
	// 	pixcallDB.Create(&pixcall.Folders{Name: folder.Name, Color: folder.Color})
	// }

	// // Copy all Entries (注释掉的代码如果不需要就删除，保留在这里只是为了展示删除线)
	// billfishEntries := billfishDB.Model(&billfish.Entry{}).Find(&billfish.Entry{})
	// for _, entry := range billfishEntries {
	// 	pixcallDB.Create(&pixcall.Entries{Name: entry.Name, Color: entry.Color})
	// 	pixcallDB.Create(&pixcall.BoardEntry{BoardID: entry.BoardID, EntryID: entry.ID, EntryKind: 1})
	// }

	return nil
}

// Copy files and sub-folders in one folder to another folder
func CopyFoldersAndFiles(originalPath string, newPath string) error {
	return filepath.Walk(originalPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(originalPath, path)
		if err != nil {
			return err
		}

		destinationPath := filepath.Join(newPath, relativePath)

		if info.IsDir() {
			return os.MkdirAll(destinationPath, info.Mode())
		}

		// Skip special files like sockets, pipes, etc.  Only copy regular files and symlinks.
		if !info.Mode().IsRegular() && info.Mode()&os.ModeSymlink == 0 {
			return nil
		}

		// Symlink handling
		if info.Mode()&os.ModeSymlink != 0 {
			// Read the target of the symbolic link.
			target, err := os.Readlink(path)
			if err != nil {
				return err
			}
			// Create the symbolic link in the destination directory.
			return os.Symlink(target, destinationPath)
		}

		return copyFile(path, destinationPath)
	})
}

/*
Billfish to Pixcall
@params

- billfishPath: existing billfish library folder path
- pixcallPath: new pixcall library folder path
*/
func billfishToEagle(billfishPath string, eagleLibraryPath string) (err error) {
	billfishDB := core.GetDB(filepath.Join(billfishPath, ".bf", "billfish.db")) // 使用 filepath.Join

	err = os.MkdirAll(filepath.Join(eagleLibraryPath, "images"), 0700) // 使用 filepath.Join
	if err != nil {
		return err
	}

	// Query all Billfish Tags
	var billfishTags []billfish.Tag
	result := billfishDB.Model(&billfish.Tag{}).Find(&billfishTags) // 直接使用 result
	if result.Error != nil {
		return fmt.Errorf("querying billfish tags: %w", result.Error)
	}

	// Map: tag id->name
	var tagIdMap = make(map[string]string, 0)

	for _, tag := range billfishTags {
		tagIdMap[(string)(tag.ID)] = strings.Trim(tag.Name, " ")
	}

	// Get all available billfish Folders
	// No need to get recycled folders
	var billFishFolders []billfish.Folder
	result = billfishDB.Model(&billfish.Folder{}).Where("is_recycle = ?", 0).Find(&billFishFolders)
	if result.Error != nil {
		return fmt.Errorf("querying billfish folders (is_recycle = false) : %w", result.Error)
	}

	// Convert billfish Folders to Eagle Folders
	// here we use a map to store the converted folders
	// because we need to reference the converted folder uids later
	var eagleFoldersMap map[string]*eagle.Folder = make(map[string]*eagle.Folder)
	// billfish folderID->folder Map
	var billfishFoldersMap map[string]*billfish.Folder = make(map[string]*billfish.Folder)
	for _, folder := range billFishFolders {
		var folderID string = pkg.GenerateEagleRandomID("M")
		eagleFoldersMap[folderID] = &eagle.Folder{
			ID:   folderID,
			Name: folder.Name,
		}
		billfishFoldersMap[(string)(folder.ID)] = &folder
	}

	// 获取 Billfish files
	var billfishFiles []billfish.File
	result = billfishDB.Preload("MaterialUserData").Preload("Tags").Model(&billfish.File{}).Find(&billfishFiles)
	if result.Error != nil {
		return fmt.Errorf("querying billfish files: %w", result.Error)
	}

	for _, file := range billfishFiles {
		// If file is deleted, the parent folder id is -1
		if file.Pid == -1 {
			continue
		}

		// file's tags
		var tagNames []string = make([]string, 0)
		for _, tag := range file.Tags {
			tagName := tagIdMap[(string)(tag.ID)]
			tagNames = append(tagNames, tagName)
		}

		var imageID string = pkg.GenerateEagleRandomID("M")
		var fileName, ext = pkg.SplitFilename(file.Name)

		var imageFolderPath = filepath.Join(eagleLibraryPath, "images", imageID+".info")
		// generate image info folder
		err = os.MkdirAll(imageFolderPath, 0700) // 使用 filepath.Join
		if err != nil {
			return err
		}

		// Get Full Path of file
		folderNames, err := GetFullPath(file.Pid, billfishFoldersMap)
		// 构建完整路径，使用 "/" 作为分隔符
		fullPath := strings.Join(folderNames, "/")
		// copy image file to this folder
		err = copyFile(filepath.Join(billfishPath, fullPath, file.Name), filepath.Join(imageFolderPath, fileName+"."+ext))
		if err != nil {
			return err
		}

		eagleImage := &eagle.ImageMetadata{
			ID:               imageID,
			Name:             fileName,
			Size:             file.FileSize,
			Btime:            file.Born * 1000,
			Mtime:            file.Mtime * 1000,
			Ext:              ext,
			Tags:             tagNames,
			Folders:          []string{},
			IsDeleted:        false,
			URL:              file.MaterialUserData.Origin,
			Annotation:       file.MaterialUserData.Note,
			ModificationTime: file.Mtime * 1000,
			LastModified:     file.Mtime * 1000,
			Star:             file.MaterialUserData.Score,
		}

		// save metadata json
		jsonData, err := json.MarshalIndent(eagleImage, "", "    ")

		if err != nil {
			fmt.Println("编码 JSON 错误:", err)
			return err
		}

		// 3. 写入文件
		filePath := filepath.Join(imageFolderPath, "metadata.json") // 文件名
		file, err := os.Create(filePath)                            // 创建文件
		if err != nil {
			fmt.Println("创建文件错误:", err)
			return err
		}
		defer file.Close() // 确保文件在使用完毕后关闭

		_, err = file.Write(jsonData) // 写入数据
		if err != nil {
			fmt.Println("写入文件错误:", err)
			return err
		}
	}

	eagleLibraryMetadata := &eagle.LibraryMetadata{
		SmartFolders:       make([]any, 0),
		Folders:            make([]eagle.Folder, 0),
		QuickAccess:        make([]any, 0),
		TagsGroups:         make([]any, 0),
		ModificationTime:   1,
		ApplicationVersion: "4.0.0",
	}

	// 2. 将结构体编码为 JSON 数据
	jsonData, err := json.MarshalIndent(eagleLibraryMetadata, "", "    ")

	if err != nil {
		fmt.Println("编码 JSON 错误:", err)
		return
	}

	// 3. 写入文件
	filePath := eagleLibraryPath + "/" + "metadata.json" // 文件名
	file, err := os.Create(filePath)                     // 创建文件
	if err != nil {
		fmt.Println("创建文件错误:", err)
		return
	}
	defer file.Close() // 确保文件在使用完毕后关闭

	_, err = file.Write(jsonData) // 写入数据
	if err != nil {
		fmt.Println("写入文件错误:", err)
		return
	}

	return nil
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}

// determineFirstLetter 提取首字母
func determineFirstLetter(tagName, tagNamePinyin string) string {
	if len(tagNamePinyin) > 0 {
		return string(tagNamePinyin[0])
	}

	if len(tagName) > 0 {
		firstByte := tagName[0] // 获取 tag.Name 的第一个字节

		// 检查首字母是大写字母
		if 'A' <= firstByte && firstByte <= 'Z' {
			return string(firstByte)
		}
	}

	return "A" // 默认值
}

// 随机数
func randomNumber() int64 {
	// Seed the random number generator
	rand.Seed((uint64)(time.Now().UnixNano()))

	// Generate a random 18-digit number
	// The lower bound is 10^17 and the upper bound is 10^18
	lower := int64(1e17)
	upper := int64(1e18)

	randomNumber := rand.Int63n(upper-lower) + lower
	return randomNumber
}

// GetFullPath 获取文件的完整路径
func GetFullPath(filePid int64, billfishFoldersMap map[string]*billfish.Folder) ([]string, error) {
	var pathComponents []string
	innerFilePid := filePid

	// 防止无限循环，设置最大迭代次数 (例如，防止数据异常导致的环路)
	maxIterations := 100
	iterations := 0

	for innerFilePid != 0 && iterations < maxIterations { // 假设根目录的 Pid 是 "0" (字符串类型)
		iterations++
		folder, ok := billfishFoldersMap[(string)(innerFilePid)] // Pid 假设已经是 string 类型
		if !ok {
			return nil, fmt.Errorf("父文件夹 PID '%d' 未在 billfishFoldersMap 中找到", innerFilePid)
		}
		pathComponents = append(pathComponents, folder.Name)
		innerFilePid = folder.Pid
	}

	if iterations >= maxIterations {
		return nil, fmt.Errorf("达到最大迭代次数，可能存在循环引用或路径过深")
	}

	// 反转路径组件，得到正确的顺序 (从根到文件所在目录)
	reversedComponents := make([]string, len(pathComponents))
	for i := range pathComponents {
		reversedComponents[i] = pathComponents[len(pathComponents)-1-i]
	}

	return reversedComponents, nil
}
