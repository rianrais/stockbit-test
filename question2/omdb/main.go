package omdb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"question2/utilities"

	"github.com/julienschmidt/httprouter"
)

func Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get query params "searchword" and "pagination"
	q := r.URL.Query()
	title := q.Get("searchword")
	page := q.Get("pagination")

	if title == "" || page == "" {
		utilities.ErrRes(w, http.StatusText(400), http.StatusBadRequest, "Required parameters is missing!")
	}

	curPage, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Error when determining page: ", err.Error())
		utilities.ErrRes(w, "Page must be a number / Error when determining page", http.StatusBadRequest, err.Error())
	}

	url := SetOmdbURL(title, page)
	omdbData, err := GetOmdbData(url)
	if err != nil {
		utilities.ErrRes(w, "Error when getting data from OMDB API", http.StatusInternalServerError, err.Error())
	}

	totalPage := GetTotalPage(omdbData.TotalResult)

	/*
		Logging each call (I'm not sure what to do here based on the question asked).
		Usually, we can store in bulk / per data to DB or other processes - based on business/product's goal.
		In the following code I will simply log the data to the terminal.

		And since you asked for use of concurrency with Golang, I will use
		this chance to demonstrate the use go func() and sync.WaitGroup.
		I also print out the index to show that it will not log based on index (ascending) like traditional for loop.
	*/
	var wg sync.WaitGroup
	wg.Add(len(omdbData.Search))
	for idx, movie := range omdbData.Search {
		go func(movie SearchResult, idx int) {
			defer wg.Done()
			fmt.Printf(
				"Index in array: %d\nTitle: %s\nYear: %s\nIMDB ID: %s\nType: %s\nPoster: %s\n\n",
				idx,
				movie.Title,
				movie.Year,
				movie.ImdbID,
				movie.Type,
				movie.Poster,
			)
		}(movie, idx)
	}
	wg.Wait()

	resp := GetResponse{
		CurrentPage: curPage,
		TotalPage:   totalPage,
		Data:        omdbData.Search,
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
