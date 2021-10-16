package main

import (
	"fmt"
	"sort"
	"strings"
)

// sortAnagram() will sort array of string based on anagram and print it
func sortAnagram(anagArr []string) {
	list := make(map[string][]string)

	for _, word := range anagArr {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}

	for _, words := range list {
		for idx, w := range words {
			fmt.Printf("'%s'", w)
			if idx+1 != len(words) {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
}

// sortStr() will sort any string in ascending matter
func sortStr(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

func main() {
	listOfStr := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	Anagram(listOfStr)
}
