package omdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"question2/model"
	"strconv"
)

// This is used to simulate .env file
var (
	BaseURL string = "http://www.omdbapi.com/"
	OmdbKey string = "faf7e5bb"
)

// will return model.Response if successfull to be processed as log.
func GetOmdbData(url string) (model.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error when hitting ", url, ": ", err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error when reading response.Body: ", err)
	}

	var res model.Response
	json.Unmarshal(responseData, &res)
	for i := 0; i < len(res.Search); i++ {
		fmt.Println(res.Search[i].Title)
	}

	return res, err
}

// will return string that will be used to determine the URL to hit OMDB.
func SetOmdbURL(currentPage string, searchParam string) string {
	searchParam = url.QueryEscape(searchParam)
	urlString := fmt.Sprintf("%s/?apikey=%s&page=%s&s=%s", BaseURL, OmdbKey, currentPage, searchParam)

	return urlString
}

// will return total page to determine how many pages.
func GetTotalPage(totalRes string) int {
	totalResult, _ := strconv.ParseFloat(totalRes, 64)
	totalPage := int(math.Ceil(totalResult / 10))

	return totalPage
}
