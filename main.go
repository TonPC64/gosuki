package main

import (
	"errors"
	"fmt"
)

type stack []int

func main() {

	var value stack
	fmt.Println(value)

	value = value.push(1)
	fmt.Println(value)

	value, popval, err := value.pop()
	fmt.Println(value, popval, err)

	value, popval, err = value.pop()
	fmt.Println(value, popval, err)
}

func (s stack) push(a int) stack {
	return append(s, a)
}

func (s stack) pop() (stack, int, error) {
	if len(s) == 0 {
		return nil, -999, errors.New("can't pop stack")
	}
	return s[:len(s)-1], s[len(s)-1:][0], nil
}
