package main

import "fmt"

func RepeatLetter(r string, cnt int) string {
	var result string
	for i := 0; i < cnt; i++ {
		result += r
	}

	return result
}

func ParseString(text string) string {
	var currentLetter string
	var currentLetterCount int
	var escapeLetterFound bool
	var parseResult string

	for _, letterByte := range text {
		if escapeLetterFound {
			if currentLetter != "" {
				parseResult += currentLetter
			}
			currentLetter = string(letterByte)
			escapeLetterFound = false
			continue
		}

		if letterByte > '0'+0 && letterByte < '0'+9 {
			currentLetterCount = int(letterByte - '0')
			if currentLetter != "" {
				parseResult += RepeatLetter(currentLetter, currentLetterCount)
				currentLetter = ""
				currentLetterCount = 0
			} else {
				return "\"\" (некоректная строка)"
			}
		} else if string(letterByte) == "\\" {
			escapeLetterFound = true
			continue
		} else {
			if currentLetter != "" {
				parseResult += currentLetter
			}
			currentLetter = string(letterByte)
		}
	}

	if currentLetter != "" {
		parseResult += currentLetter
	}

	return parseResult
}

func main() {
	fmt.Printf("%s | %s\n", ParseString("a4bc2d5e"), "aaaabccddddde")
	fmt.Printf("%s | %s\n", ParseString("abcd"), "abcd")
	fmt.Printf("%s | %s\n", ParseString("45"), "\"\" (некорректная строка)")
	fmt.Printf("%s | %s\n", ParseString("qwe\\4\\5"), "qwe45")
	fmt.Printf("%s | %s\n", ParseString("qwe\\45"), "qwe44444")
	fmt.Printf("%s | %s\n", ParseString("qwe\\\\5"), "qwe\\\\\\\\\\")
}
