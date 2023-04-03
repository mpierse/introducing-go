package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1.0+1.0)

	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello, " + "World")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Printf("The answer is %d", computeExercise())
}

func computeExercise() int {
	return 32132 * 42452
}

/* Exercise Answers
1. Integers are stored as binary in a computer
2. The largest 8 digit binary number is 1111 1111; or 2^8 - 1; or 255
3. See func computeExercise() above
4. A string is a sequence of characters with a definite length used to represent text. The length of a string is found with the func len(s string)
5. The expression `(true && false) || (false && true) || !(false && false)` always evaluates to true
*/
