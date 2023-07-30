package data

// MetaData represents meta-data object
type MetaData struct {
	Site        string   `json:"site"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
