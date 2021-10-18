package omdb

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type TestURL struct {
	Input []string
	Want  string
}

func TestSetOmdbUrl(t *testing.T) {
	tests := []TestURL{
		{
			Input: []string{"robin", "2"},
			Want:  "http://www.omdbapi.com/?apikey=faf7e5bb&page=2&s=robin",
		}, {
			Input: []string{"la la land", "1"},
			Want:  "http://www.omdbapi.com/?apikey=faf7e5bb&page=1&s=la+la+land",
		}, {
			Input: []string{"goblinwati", "100"},
			Want:  "http://www.omdbapi.com/?apikey=faf7e5bb&page=100&s=goblinwati",
		},
	}

	t.Log("Given the need to return string when creating OMDB URL to hit, the URL string should return string with query escape.\n")
	for _, val := range tests {
		eval := SetOmdbURL(val.Input[0], val.Input[1])
		if !reflect.DeepEqual(val.Want, eval) {
			t.Errorf(
				"Testing SetOmdbUrl() failed. Expect: %s - Got: %s\n\n",
				val.Want,
				eval,
			)
		}
	}
}

type TestTotalPage struct {
	Input string
	Want  int
}

func TestGetTotalPage(t *testing.T) {
	tests := []TestTotalPage{
		{
			Input: "101",
			Want:  11,
		}, {
			Input: "100",
			Want:  10,
		}, {
			Input: "666",
			Want:  67,
		},
	}

	t.Log("Given the need to get total page. Return total pages based on how many datas there are (assuming there is only 10 data per page).")

	for _, val := range tests {
		eval := GetTotalPage(val.Input)
		if !reflect.DeepEqual(val.Want, eval) {
			t.Errorf(
				"Testing GetTotalPage() failed. Expect: %d - Got: %d\n\n",
				val.Want,
				eval,
			)
		}
	}
}

type TestOmdbData struct {
	Input string
	Want  bool
}

func TestGetOmdbData(t *testing.T) {
	tests := []TestOmdbData{
		{
			Input: "Avenger",
			Want:  true,
		}, {
			Input: "Batman",
			Want:  true,
		},
	}

	t.Log("Given definite existing movie. This test should at least return one movie contained with name inputted.")

	for _, val := range tests {
		urlString := fmt.Sprintf("http://www.omdbapi.com/?apikey=faf7e5bb&page=1&s=%s", val.Input)
		eval, _ := GetOmdbData(urlString)

		isExist := strings.Contains(eval.Search[0].Title, val.Input)
		if !isExist {
			t.Errorf(
				"Testing GetOmdbData() failed. Expect \"%s\" to contain \"%s\".",
				eval.Search[0].Title,
				val.Input,
			)
		}
	}
}
