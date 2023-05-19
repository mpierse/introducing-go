package main

import (
	"fmt"
	"math"
)

func main() {

	circle := Circle{0, 0, 5}
	rectangle := Rectangle{0, 0, 10, 10}
	fmt.Println("The perimeter and area of circle: ")
	fmt.Println(rectangle.perimeter())
	fmt.Println(rectangle.area())
	fmt.Println("The perimeter and area of rectangle: ")
	fmt.Println(circle.perimeter())
	fmt.Println(circle.area())
	fmt.Println("The total perimeter and area: ")
	fmt.Println(totalPerimeter(&circle, &rectangle))
	fmt.Println(totalArea(&circle, &rectangle))
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) area() float64 {
	return (r.x2 - r.x1) * (r.y2 - r.y1)
}

func (r *Rectangle) perimeter() float64 {
	return 2 * ((r.x2 - r.x1) + (r.y2 - r.y1))
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

/* Exercise Answers
1. A method is a function with a receiver, which means you can call the method with the .
operator on the struct type. A function is a standalone function that can be called from
anywhere.
2. You might use an embedded anonymous field if the relationship is a "is a" instead of 'has a'. You
may also want to be able to use the embedded field's methods on the struct.
3. See code above
*/
