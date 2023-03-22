package main

import "fmt"

//this is a comment

/* This is also a comment
it can go across multiple lines
*/

func main() {
	fmt.Println("Hello, World")
	fmt.Println("Hello, my name is Tiffany")
}

/* Exercise Answers
1. White space is new lines, spaces, and tabs. It is for readability and mostly ignored by Go
2. A comment is ignored by the Go compiler and is for the benifit of humans, see examples above
3. Files in the "fmt" package would begin with "package fmt"
4. To use the exit function from the os package you would need to import "os" and access the Exit function by writing "os.Exit(0)" where the argument is an int exit code
5. The code is modified to answer question 5
*/
