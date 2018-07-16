package main

import (
	"fmt"
	"strings"
)

func main() {
	wc := NewWordCount()
	wc = wc.AddWord("I am")
	wc = wc.AddWord("I am")
	wc = wc.AddWord("am")
	fmt.Println(wc.MaxWordCount())
	fmt.Println(wc.MinWordCount())
	fmt.Println(wc.LongestWord())
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

func (wc WordsCount) MaxWordCount() ([]string, int) {
	counts := []int{}
	for _, v := range wc {
		counts = append(counts, v.Count)
	}
	max := maxInts(counts...)
	s := []string{}
	for key, v := range wc {
		if max == v.Count {
			s = append(s, key)
		}
	}
	return s, max
}

func (wc WordsCount) MinWordCount() ([]string, int) {
	counts := []int{}
	for _, v := range wc {
		counts = append(counts, v.Count)
	}
	min := minInts(counts...)
	s := []string{}
	for key, v := range wc {
		if min == v.Count {
			s = append(s, key)
		}
	}
	return s, min
}

func (wc WordsCount) LongestWord() ([]string, int) {
	counts := []int{}
	for _, v := range wc {
		counts = append(counts, v.Length)
	}
	max := maxInts(counts...)
	s := []string{}
	for key, v := range wc {
		if max == v.Length {
			s = append(s, key)
		}
	}
	return s, max
}

func (wc WordsCount) ShortestWord() ([]string, int) {
	counts := []int{}
	for _, v := range wc {
		counts = append(counts, v.Length)
	}
	min := minInts(counts...)
	s := []string{}
	for key, v := range wc {
		if min == v.Length {
			s = append(s, key)
		}
	}
	return s, min
}

func (wc WordsCount) PrintAll() {
	for key, value := range wc {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func maxInts(nums ...int) int {
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func minInts(nums ...int) int {
	minNum := nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}
