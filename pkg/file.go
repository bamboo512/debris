package pkg

import "path/filepath"

// 获取文件名（不含扩展名）和扩展名
func SplitFilename(filename string) (name, ext string) {
	ext = filepath.Ext(filename)[1:]
	name = filename[:len(filename)-len(ext)-1]
	return name, ext
}
