package spotifyHelper

import (
	"math/rand"
	"net/http"
	"os"
)

func GetFormattedUrlParamString(path, params string) string {
	toReturn := os.Getenv("SPOTIFY_URL") + path + "?" + params
	return toReturn
}

func GetHeaders() http.Header {
	headers := http.Header{}
	headers.Add("Content-Type", "application/x-www-form-urlencoded")

	return headers
}

func GenerateRandomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}
