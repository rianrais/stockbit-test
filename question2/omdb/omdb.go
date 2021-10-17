package omdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"question2/model"
)

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

func SetOmdbURL(currentPage string) string {
	baseURL := "http://www.omdbapi.com/"
	omdbKey := "faf7e5bb"
	searchParam := "candyman"
	url := fmt.Sprintf("%s/?apikey=%s&page=%s&s=%s", baseURL, omdbKey, currentPage, searchParam)

	return url
}
