package data

// S3 represent S3 storage record
type S3 struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
	UseSSL       bool
}
