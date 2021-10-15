package main

import (
	"fmt"
	"strings"
)

// findFirstStringInBracket() will find words inside first full parentheses
func findFirstStringInBracket(str string) string { 
	if (len(str) == 0) {
		return ""
	}

	firstBracketIndex := strings.Index(str,"(")
	// No first bracket found
	if firstBracketIndex == -1 {
		return ""
	}

	runes := []rune(str) 
	wordsAfterFirstBracket := string(runes[firstBracketIndex : len(str)])

	closingBracketIndex := strings.Index(wordsAfterFirstBracket,")") 
	// No closing bracket found
	if closingBracketIndex == -1 {
		return ""
	}

	runes = []rune(wordsAfterFirstBracket)
	return string(runes[1 : closingBracketIndex])
}

// alternativeFindFirstStringInBracket() will find words inside first parentheses if we know the string will have no Unicode text (without using rune)
func alternativeFindFirstStringInBracket(str string) string {
	// Find opening and closing parenthesis index
	opIndex := strings.Index(s, "(")
	closeIndex := strings.Index(s, ")")

	// If either index not found, return empty string
	if opIndex == -1 || closeIndex == -1 {
		return ""
	}

	return s[opIndex+1 : closeIndex]
}

func main() {
	firstString := "Ini array (tapi bohong, ini string)"
	secondString := "String ini tidak memiliki (bracket penutup"
	thirdString := "Error: token has expired. (code: 403)"

	fmt.Println(findFirstStringInBracket(firstString)) // will return: tapi bohong, ini string
	fmt.Println(findFirstStringInBracket(secondString)) // will return nothing
	fmt.Println(findFirstStringInBracket(thirdString)) // will return: "code: 403"

	// Alternatif function without using rune:
	fmt.Println(alternativeFindFirstStringInBracket(firstString)) // will return: tapi bohong, ini string
	fmt.Println(alternativeFindFirstStringInBracket(secondString)) // will return nothing
	fmt.Println(alternativeFindFirstStringInBracket(thirdString)) // will return: "code: 403"
}