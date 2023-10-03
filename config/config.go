package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// OAuthRecord defines OAuth provider's credentials
type OAuthRecord struct {
	Provider     string `mapstructure:"provider"`      // name of the provider
	ClientID     string `mapstructure:"client_id"`     // client id
	ClientSecret string `mapstructure:"client_secret"` // client secret
}

// WebServer represents common web server configuration
type WebServer struct {
	Base      string `mapstructure:"base"`       // base URL
	LogFile   string `mapstructure:"log_file"`   // server log file
	Port      int    `mapstructure:"port"`       // server port number
	Verbose   int    `mapstructure:"verbose"`    // verbose output
	StaticDir string `mapstructure:"static_dir"` // speficy static dir location

	// middleware server parts
	LimiterPeriod string `mapstructure:"rate"` // limiter rate value

	// proxy parts
	XForwardedHost      string `mapstructure:"X-Forwarded-Host"`       // X-Forwarded-Host field of HTTP request
	XContentTypeOptions string `mapstructure:"X-Content-Type-Options"` // X-Content-Type-Options option

	// TLS server parts
	RootCAs     string   `mapstructure:"rootCAs"`      // server Root CAs path
	ServerCrt   string   `mapstructure:"server_cert"`  // server certificate
	ServerKey   string   `mapstructure:"server_key"`   // server certificate
	DomainNames []string `mapstructure:"domain_names"` // LetsEncrypt domain names
}

// Frontend stores frontend configuration parameters
type Frontend struct {
	WebServer

	// OAuth parts
	OAuth []OAuthRecord `mapstructure:"oauth"` // oauth configurations

	// captcha parts
	CaptchaSecretKey string `mapstructure:"captchaSecretKey"` // re-captcha secret key
	CaptchaPublicKey string `mapstructure:"captchaPublicKey"` // re-captcha public key
	CaptchaVerifyUrl string `mapstructure:"captchaVerifyUrl"` // re-captcha verify url
}

// Encryption represents encryption configuration parameters
type Encryption struct {
	Secret string `mapstructure:"secret"`
	Cipher string `mapstructure:"cipher"`
}

// Discovery represents discovery service configuration
type Discovery struct {
	WebServer
	Encryption
}

// MetaData represents metadata service configuration
type MetaData struct {
	WebServer
}

// DataManagement represents data-management service configuration
type DataManagement struct {
	WebServer
}

// DataBookkeeping represents data-bookkeeping service configuration
type DataBookkeeping struct {
	WebServer

	DBFile             string `mapstructure:"dbfile"`               // dbs db file with secrets
	MaxDBConnections   int    `mapstructure:"max_db_connections"`   // maximum number of DB connections
	MaxIdleConnections int    `mapstructure:"max_idle_connections"` // maximum number of idle connections
}

// Authz represents authz service configuration
type Authz struct {
	WebServer
	Encryption

	DbUri        string `mapstructure:"dburi"` // database URI
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	Domain       string `mapstructure:"domain"`
}

// OreCastConfig represents orecast configuration
type OreCastConfig struct {
	Frontend
	Discovery
	MetaData
	DataManagement
	DataBookkeeping
	Authz
}

func ParseConfig(cfile string) (OreCastConfig, error) {
	var config OreCastConfig
	arr := strings.Split(cfile, ".")
	cname := arr[0]
	cext := arr[1]
	viper.SetConfigName(cname)
	viper.SetConfigType(cext)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		msg := err.Error()
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			msg = fmt.Sprintf("%s file not found", cfile)
		} else {
			// Config file was found but another error was produced
			msg = fmt.Sprintf("unable to parse %s, error", cfile, err)
		}
		return config, errors.New(msg)
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}

/*
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cfile := "orecast.json"
	config, err := ParseConfig(cfile)
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
	fmt.Printf("Frontend %+v\n", config.Frontend)
	fmt.Printf("Authz %+v\n", config.Authz)
}
*/
