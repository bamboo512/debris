package eagle

type LibraryMetadata struct {
	Folders            []Folder      `json:"folders"`
	SmartFolders       []interface{} `json:"smartFolders"` // Assuming unspecified type
	QuickAccess        []interface{} `json:"quickAccess"`  // Assuming unspecified type
	TagsGroups         []interface{} `json:"tagsGroups"`   // Assuming unspecified type
	ModificationTime   int64         `json:"modificationTime"`
	ApplicationVersion string        `json:"applicationVersion"`
}

type Folder struct {
	ID               string        `json:"id"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	Children         []interface{} `json:"children"` // Assuming children are other Folders or similar
	ModificationTime int64         `json:"modificationTime"`
	Tags             []string      `json:"tags"`
	Password         string        `json:"password"`
	PasswordTips     string        `json:"passwordTips"`
}
