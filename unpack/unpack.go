package unpack

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	testStrings := [...]string{"a4bc2d5e", "abcd", "45", "7q", "q3", "a11"}

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

func Unpack(str string) (string, error) {
	result := strings.Builder{}

	charsSlice := regexp.MustCompile(`(?m)\D+`).FindAllString(str, -1)
	numsSlice := regexp.MustCompile(`(?m)\d+`).FindAllString(str, -1)

	if is, _ := regexp.MatchString(`^\d`, str); is == true {
		return "", fmt.Errorf("invalid string: %s", str)
	}

	if len(charsSlice) == 0 {
		return "", fmt.Errorf("invalid string: %s", str)
	}

	for i := 0; i < len(charsSlice); i++ {
		result.WriteString(charsSlice[i])

		if i < len(numsSlice) {
			num, err := strconv.Atoi(numsSlice[i])

			if err == nil {
				result.WriteString(strings.Repeat(
					charsSlice[i][len(charsSlice[i])-1:],
					num-1))
			}
		}
	}

	return result.String(), nil
}
