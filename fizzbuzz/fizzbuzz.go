package fizzbuzz

func FizzBuzz(i int) interface{} {
	if i%15 == 0 {
		return "fizzbuzz"
	} else if i%3 == 0 {
		return "fizz"
	} else if i%5 == 0 {
		return "buzz"
	}
	return i
}
