package eagle

/* Pack, only used in the exported .eaglepack file  */
type Pack struct {
	Images []PackImage `json:"images"`
}

/* Pack Image, only used in the exported .eaglepack file  */
type PackImage struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Btime int64  `json:"btime"`
	Mtime int64  `json:"mtime"`
	/* File Extension Name */
	Ext       string   `json:"ext"`
	Tags      []string `json:"tags"`
	Folders   []string `json:"folders"`
	IsDeleted bool     `json:"isDeleted"`
	/* URL */
	URL string `json:"url"`
	/* Remark */
	Annotation       string `json:"annotation"`
	ModificationTime int64  `json:"modificationTime"`
	/* Rating, must be between 0 and 5 and is integer */
	Star             int `json:"star"`
	Width            int `json:"width"`
	Height           int `json:"height"`
	ResolutionWidth  int `json:"resolutionWidth"`
	ResolutionHeight int `json:"resolutionHeight"`
	/* Video Duration, unit: second */
	Duration float64 `json:"duration"`
}
