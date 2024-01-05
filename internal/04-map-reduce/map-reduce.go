package main

import (
	"fmt"
	"strings"
)

// https://go.dev/play/p/CU8lt4mIflo

var lines = []string{
	"Lorem Ipsum is simply dummy text of the printing and typesetting industry",
	"Lorem Ipsum has been the industry's standard dummy text ever since the",
	"when an unknown printer took a galley of type and scrambled it to make a type specimen book",
	"It has survived not only five centuries but also the leap into electronic typesetting remaining essentially unchanged. It was popularised in the with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
}

func main() {
	linesChan := make(chan string)
	wordsChan := make([]chan string, 26)
	wordCountChan := make(chan map[string]int)
	go func() {
		for _, line := range lines {
			linesChan <- strings.ToLower(line)
		}
	}()

	// TODO: (mappers) start a worker pool to read from linesChan, split the string
	//       and send to the correct wordsChan

	//  take the first letter of each word
	//  idx := int('a' - firstLetter)  hash function for the words

	// TODO: (reducers) have a worker pool to read from wordsChan form the wordCount
	//       in a map[string]int and once all words have been read send the map to
	//       the consumer thread

	go func() {
		for counts := range wordCountChan {
			fmt.Printf("counts received:\n%v", counts)
		}
	}()
}
