package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	text := "textОлень - северное животное. В летнее время оленям в тайге жарко, а в горах даже в июле холодно. Олень как бы создан для северных просторов жёсткого ветра длинных морозных ночей. Олень легко бежит вперёд по тайге подминает под себя кусты переплывает быстрые реки. Олень не тонет потому, что каждая его шерстинка это длинная трубочка которую внутри наполняет воздух. Нос у оленя покрыт серебристой шёрсткой. Если бы шерсти на носу не было, олень бы его отморозил."

	result := Top10(text)

	fmt.Println(result)
}

func Top10(text string) (result []string) {
	text = strings.Replace(text, ",", "", -1)
	text = strings.Replace(text, ".", "", -1)

	wordsCounts := make(map[string] int)

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
		Key string
		Value int
	}

	var keyValueSlice []keyValue

	for k, v := range wordsCounts {
		keyValueSlice = append(keyValueSlice,  keyValue{k, v})
	}

	sort.Slice(keyValueSlice, func(i, j int) bool {
		return keyValueSlice[i].Value > keyValueSlice[j].Value
	})

	for i := 0; i < 10; i++ {
		if i < len(keyValueSlice) {
			result = append(result, keyValueSlice[i].Key)
		} else {
			break
		}
	}

	return result
}
