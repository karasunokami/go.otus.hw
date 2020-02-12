package twf

import (
	"errors"
	"sort"
	"strings"
)

func Find(text string, size int) (result []string, err error) {
	if text == "" {
		return nil, errors.New("text parameter cannot be empty")
	}

	text = strings.Replace(text, ",", "", -1)
	text = strings.Replace(text, ".", "", -1)

	wordsCounts := make(map[string]int)

	for _, word := range strings.Split(text, " ") {
		if word == "-" {
			continue
		}

		word = strings.ToLower(word)
		if wordsCounts[word] == 0 {
			wordsCounts[word] = 1
			continue
		}

		wordsCounts[word] += 1
	}

	type keyValue struct {
		Key   string
		Value int
	}

	var keyValueSlice []keyValue

	for k, v := range wordsCounts {
		keyValueSlice = append(keyValueSlice, keyValue{k, v})
	}

	sort.Slice(keyValueSlice, func(i, j int) bool {
		return keyValueSlice[i].Value > keyValueSlice[j].Value
	})

	for i := 0; i < size; i++ {
		if i < len(keyValueSlice) {
			result = append(result, keyValueSlice[i].Key)
		} else {
			break
		}
	}

	return result, nil
}
