package omdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

// This is used to simulate .env file
var (
	BaseURL string = "http://www.omdbapi.com"
	OmdbKey string = "faf7e5bb"
)

// will return Response if successfull to be processed as log.
func GetOmdbData(url string) (OmdbResponse, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error when hitting ", url, ": ", err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error when reading response.Body: ", err)
	}

	var res OmdbResponse
	json.Unmarshal(responseData, &res)

	return res, err
}

// will return string that will be used to determine the URL to hit OMDB based on parameters from the API.
func SetOmdbURL(searchParam string, page string) string {
	searchParam = url.QueryEscape(searchParam)
	urlString := fmt.Sprintf("%s/?apikey=%s&page=%s&s=%s", BaseURL, OmdbKey, page, searchParam)

	return urlString
}

// will return total page to determine how many pages.
func GetTotalPage(totalRes string) int {
	totalResult, _ := strconv.ParseFloat(totalRes, 64)
	totalPage := int(math.Ceil(totalResult / 10))

	return totalPage
}
