package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"jdreese/setlistplay/helpers/setlistFmHelper"
	sfmModels "jdreese/setlistplay/models/setlistFmModels"
	"log"
	"net/http"
)

// SetlistFmTest TODO - Delete method
func SetlistFmTest(c echo.Context) error {
	params := c.QueryParams()
	fmt.Printf("\nQueryParams: %s\n", params)

	paramString := params.Encode()
	url := setlistFmHelper.GetFormattedUrlParamString("/total/test/url", paramString)
	fmt.Printf("\nURL: %s\n", url)

	toReturn := "Done :)"
	return c.JSON(http.StatusOK, toReturn)
}

func HandleSetlistSearch(c echo.Context) error {
	params := c.QueryParams()
	paramString := params.Encode()
	url := setlistFmHelper.GetFormattedUrlParamString("/1.0/search/setlists", paramString)
	headers := setlistFmHelper.GetHeaders()

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error in request: %+v\n", err)
	}
	request.Header = headers

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

	var result sfmModels.SetlistResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, result)
}
