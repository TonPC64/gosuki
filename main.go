package main

import (
	"fmt"
)

func init() {
	println("init func()")
}

func main() {
	number := []int{}
	for i := 1; i <= 100; i++ {
		number = append(number, i)
	}
	sum, avg := calcOnlyPrime(number...)
	fmt.Printf("sum = %d, avg = %f\n", sum, avg)
}

func calcOnlyPrime(i ...int) (int, float32) {
	sum := 0
	count := 0
	for _, num := range i {
		if checkPrime(num) {
			sum += num
			count++
		}
	}
	avg := float32(sum) / float32(count)
	return sum, avg
}

func checkPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
