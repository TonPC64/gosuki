package fizzbuzz_test

import (
	"testing"

	"github.com/TonPC64/gosuki/fizzbuzz"
)

func TestFizzBuzzCase1(t *testing.T) {
	if fizzbuzz.FizzBuzz(1) != 1 {
		t.Errorf("FizzBuzz(1) should be equal 1")
	}
}

func TestFizzBuzzCase2(t *testing.T) {
	if fizzbuzz.FizzBuzz(2) != 2 {
		t.Errorf("FizzBuzz(2) should be equal 2")
	}
}

func TestFizzBuzzCase3(t *testing.T) {
	if fizzbuzz.FizzBuzz(3) != "fizz" {
		t.Errorf("FizzBuzz(3) should be equal fizz")
	}
}

func TestFizzBuzzCase5(t *testing.T) {
	if fizzbuzz.FizzBuzz(5) != "buzz" {
		t.Errorf("FizzBuzz(5) should be equal buzz")
	}
}

func TestFizzBuzzCase15(t *testing.T) {
	if fizzbuzz.FizzBuzz(15) != "fizzbuzz" {
		t.Errorf("FizzBuzz(15) should be equal fizzbuzz")
	}
}
