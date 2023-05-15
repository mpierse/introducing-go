package main

import "fmt"

func main() {
	var y = []float64{98, 93, 77, 82, 83}
	fmt.Println(average(y))

	// Closure example - a function that references non-local variables
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(half(1))
	fmt.Println(greatest(1, 2, 3, 4, 5))

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	fmt.Println(fib(10))

	pointer := 1.5
	square(&pointer)
	fmt.Println("this is squared: ")
	fmt.Println(pointer)

	val1 := 1
	val2 := 2
	swap(&val1, &val2)
	fmt.Println(val1, val2)
	// Examples of defer, panic and recover from panic
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")

}

func average(xs []float64) float64 {
	var total float64 = 0
	for _, val := range xs {
		total += val
	}
	return total / float64(len(xs))
}

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func greatest(args ...int) int {
	greatest := 0
	for _, val := range args {
		if val > greatest {
			greatest = val
		}
	}
	return greatest
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fib(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

/* Exercise Answers
1. The function signature of sum would look like: func sum(xs []int) int
2. See func half() above
3. See func greatest() above
4. See func makeOddGenerator() above
5. See func fib() above
6. defer is used to delay the execution of a statement until the surrounding function returns.
   panic is used to cause a run time error intentionally.
   recover is used to regain control of a panicking goroutine.
7. You can get the memory address of a variable by using the & operator, e.g. &x
8. You assign a value to a pointer by using the * operator, e.g. *x = 5
9. You create a new pointer variable by using the * operator or new keyword, e.g. x := new(int)
10. After running the program, the value of x is 2.25
11. See func swap() above
*/
