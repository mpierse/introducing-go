package main

import "fmt"

func main() {
	multiplesOfThree()
	fizzBuzz()
}

// multiplesOfThree prints all multiples of 3 between 1 and 100
func multiplesOfThree() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

/* Exercise Answers
1. The sample program will print "Small" because i = 10, which is not greater than 10
2. See func multiplesOfThree() above
3. See func fizzBuzz() above
*/
