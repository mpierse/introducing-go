package main

import "fmt"

func main() {
	convertToCelsius()
	convertToMeters()
}

func convertToCelsius() {
	fmt.Println("Enter a temperature in Fahrenheit: ")
	var f float64
	fmt.Scan(&f)
	fmt.Println((f - 32) * 5 / 9)
	return
}

func convertToMeters() {
	fmt.Println("Enter a length in feet: ")
	var f float64
	fmt.Scan(&f)
	fmt.Println(f * 0.3048)
	return
}

/* Exercise Answers
1. You can create a variable with the var or const keywords or with the `:=` operator
2. After running `x := 5` and `x += 1`, the value of x is 6
3. Scope is the area of code where a variable is accessible, in Go scope is determined by the `{}` brackets
4. var defines a variable that can be reassinged, while const is a variable that cannot be changed
5. See the convertToCelsius function above
6. See the convertToMeters function above
*/
