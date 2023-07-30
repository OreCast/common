package data

// Site represents Site object returned from discovery service
type Site struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
