package utils

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// ReadToken function to either read file content or return given string
func ReadToken(r string) string {
	if _, err := os.Stat(r); err == nil {
		b, e := os.ReadFile(r)
		if e != nil {
			log.Fatalf("Unable to read data from file: %s, error: %s", r, e)
		}
		return strings.Replace(string(b), "\n", "", -1)
	}
	return r
}

// HttpClient provides cert/token aware HTTP client
func HttpClient() *http.Client {
	timeout := time.Duration(TIMEOUT) * time.Second
	if TIMEOUT > 0 {
		return &http.Client{Timeout: time.Duration(timeout)}
	}
	return &http.Client{}
}
