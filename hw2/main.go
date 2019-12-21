package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testStrings := [...]string{"a4bc2d5e", "abcd", "45", "7q", "q3"}

	for _, testString := range testStrings {
		fmt.Println("input:  " + testString)
		resultString, err := Unpack(testString)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("result: " + resultString)
		}
	}
}

func Unpack(str string) (result string, err error) {
	var (
		lastChar  string
		numString string
	)

	slice := strings.Split(str, "")

	for _, char := range slice {
		if _, err := strconv.ParseInt(char, 10, 64); err == nil {
			numString += char

			if lastChar == "" {
				return "", errors.New("invalid string")
			}
		} else {
			lastChar = char
			result += char
			continue
		}

		if numString != "" {
			count, _ := strconv.ParseInt(numString, 10, 64)
			result += strings.Repeat(lastChar, int(count) - 1)
			numString = ""
		}
	}

	return
}
