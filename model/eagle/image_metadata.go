package eagle

// ImageMetadata, representing the structure of the JSON data
type ImageMetadata struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
	// Image Created Time
	Btime int64 `json:"btime"`
	Mtime int64 `json:"mtime"`
	// Image Modified Time
	Ext              string   `json:"ext"`
	Tags             []string `json:"tags"`
	Folders          []string `json:"folders"`
	IsDeleted        bool     `json:"isDeleted"`
	URL              string   `json:"url"`
	Annotation       string   `json:"annotation"`
	ModificationTime int64    `json:"modificationTime"`
	Height           int      `json:"height"`
	Width            int      `json:"width"`
	LastModified     int64    `json:"lastModified"`
	Star             int      `json:"star"`
}
