package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y = [5]float64{98, 93, 77, 82, 83}
	var total float64 = 0
	for _, val := range y {
		total += val
	}
	fmt.Println(total / float64(len(y)))

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice1 = []int{1, 2, 3}
	slice2 = make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	z := make(map[string]int)
	z["key"] = 10
	fmt.Println(z)

	smallestNumInList()
}

func smallestNumInList() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	smallestNum := x[0]
	for _, val := range x {
		if smallestNum > val {
			smallestNum = val
		}
	}
	fmt.Println(smallestNum)
}

/* Exercise Answers
1. In an array or slice named x you could access the 4th element by invoking x[3]
2. The length of a slice created using make([]int, 3, 9) would be 3
3. See func smallestNumInList() for how to find the smallest number in this list
*/
