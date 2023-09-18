package data

// Site represents Site object returned from discovery service
type Site struct {
	Name         string `json:"name" binding:"required"`
	URL          string `json:"url" binding:"required"`
	Endpoint     string `json:"endpoint" binding:"required"`
	AccessKey    string `json:"access_key" binding:"required"`
	AccessSecret string `json:"access_secret" binding:"required"`
	UseSSL       bool   `json:"use_ssl"`
	Description  string `json:"description"`
}
