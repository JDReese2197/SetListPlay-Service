package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	helper "jdreese/setlistplay/helpers/spotifyHelper"
	"log"
	"net/http"
	"net/url"
	"os"
)

func HandleSpotifyLogin(c echo.Context) error {
	params := url.Values{}
	params.Add("client_id", os.Getenv("SPOTIFY_CLIENTID"))
	params.Add("response_type", "code")
	params.Add("redirect_uri", os.Getenv("SPOTIFY_REDIRECT"))
	params.Add("state", helper.GenerateRandomString(16))
	params.Add("scope", "playlist-modify-public playlist-modify-private")
	params.Add("show_dialog", "true")
	paramString := params.Encode()

	url := helper.GetFormattedUrlParamString("/authorize", paramString)

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("Error in request: %+v\n", err)
	}
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error in response: %+v\n", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error closing response.Body ready:\n %s", err)
		}
	}(response.Body)

	var body []byte
	if response.StatusCode == http.StatusOK {
		bodyBites, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		body = bodyBites
	}

	return c.HTMLBlob(http.StatusOK, body)
}
