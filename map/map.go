package main

import (
	"fmt"
	"strings"
)

func main() {
	wc := NewWordCount()
	wc = wc.AddWord("I am am am")
	wc = wc.AddWord("I am am am")

	wc.PrintAll()
}

type WordsCount map[string]WordCount

type WordCount struct {
	Count  int
	Length int
}

func NewWordCount() WordsCount {
	return make(WordsCount)
}

func (wc WordsCount) AddWord(s string) WordsCount {
	for _, str := range strings.Fields(s) {
		wc[str] = WordCount{
			Count:  wc[str].Count + 1,
			Length: len(str),
		}
	}
	return wc
}

func (wc WordsCount) MaxWordCount() (string, int) {
	return "", 0
}

func (wc WordsCount) MinWordCount() (string, int) {
	return "", 0
}

func (wc WordsCount) LongestWord() (string, int) {
	return "", 0
}

func (wc WordsCount) ShortestWord() (string, int) {
	return "", 0
}

func (wc WordsCount) PrintAll() {
	for key, value := range wc {
		fmt.Println("Key:", key, "Value:", value)
	}
}
