package utils

import (
	"fmt"
	"net/http"
	"regexp"
)

var baseAppHostRegex = regexp.MustCompile(`^(api\.|auth\.):(.+)$`)

func GetBaseUrl(r *http.Request) string {
	return fmt.Sprintf("%s://%s", r.URL.Scheme, r.URL.Host)
}

func GetAppUrl(r *http.Request) string {
	baseHost := baseAppHostRegex.ReplaceAll([]byte(r.URL.Host), []byte("$2"))
	return fmt.Sprintf("%s://%s", r.URL.Scheme, string(baseHost))
}

func GetApiUrl(r *http.Request) string {
	baseHost := baseAppHostRegex.ReplaceAll([]byte(r.URL.Host), []byte("api.$2"))
	return fmt.Sprintf("%s://%s", r.URL.Scheme, string(baseHost))
}

func GetAuthUrl(r *http.Request) string {
	baseHost := baseAppHostRegex.ReplaceAll([]byte(r.URL.Host), []byte("auth.$2"))
	return fmt.Sprintf("%s://%s", r.URL.Scheme, string(baseHost))
}
