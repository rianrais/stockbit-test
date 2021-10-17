package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"

	"question2/model"
)

func getOmbdApi(url string) (model.Response, error) {
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

func main() {
	baseURL := "http://www.omdbapi.com/"
	omdbKey := "faf7e5bb"
	currentPage := "1"
	searchParam := "la land"
	urlString := fmt.Sprintf("%s/?apikey=%s&page=%s&s=%s", baseURL, omdbKey, currentPage, url.QueryEscape(searchParam))

	res, err := getOmbdApi(urlString)
	if err != nil {
		log.Println(err)
	}

	totalResult, _ := strconv.ParseFloat(res.TotalResult, 64)
	totalPage := int(math.Ceil(totalResult / 10))
	// fmt.Printf("Total result: %s\nTotal page: %d\n", res.TotalResult, totalPage)

	// Use sync.WaitGroup and adding totalPage into it
	var wg sync.WaitGroup
	wg.Add(totalPage - 1)

	// Use go func() to hit OMDB API concurrently
	if totalPage > 1 {
		go func() {
			for i := 2; i <= totalPage; i++ {
				defer wg.Done()
				currentPage = strconv.Itoa(i)
				urlString = fmt.Sprintf("%s/?apikey=%s&page=%s&s=%s", baseURL, omdbKey, currentPage, searchParam)
				getOmbdApi(urlString)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Done!")
}
