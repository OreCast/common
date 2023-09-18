package data

// MetaData represents meta-data object
type MetaData struct {
	ID          string   `json:"id"`
	Site        string   `json:"site"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Bucket      string   `json:"bucket"`
}
