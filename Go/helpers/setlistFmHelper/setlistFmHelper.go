package setlistFmHelper

import (
	"net/http"
	"os"
)

func GetFormattedUrlParamString(path, params string) string {
	toReturn := os.Getenv("SETLISTFM_URL") + path + "?" + params
	return toReturn
}

func GetHeaders() http.Header {
	headers := http.Header{}
	headers.Add("accept", "application/json")
	headers.Add("x-api-key", os.Getenv("SETLISTFM_KEY"))

	return headers
}
