package main

import (
	"fmt"
	"strings"
)

func GetTopNWords(words map[string]int, n int) map[string]int {
	result := make(map[string]int)
	currentLength := 0

	for word, cnt := range words {
		if currentLength < n {
			result[word] = cnt
			currentLength += 1
		} else {
			for rWord, rCnt := range result {
				if rCnt < cnt {
					delete(result, rWord)
					result[word] = cnt
					break
				}
			}
		}
	}

	return result
}

func GetNumberOfWords(text string) map[string]int {
	words := make(map[string]int)

	splitText := strings.Split(text, " ")

	for _, word := range splitText {
		word = strings.ToLower(word)
		word = strings.ReplaceAll(word, ".", "")
		word = strings.ReplaceAll(word, ",", "")
		if _, ok := words[word]; ok {
			words[word] += 1
		} else {
			words[word] = 0
		}
	}

	return GetTopNWords(words, 10)
	//return words
}

func main() {
	text := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."

	result := GetNumberOfWords(text)
	for word, cnt := range result {
		fmt.Printf("%s | %d\n", word, cnt)
	}
}
