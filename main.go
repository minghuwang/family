package main

import (
	"fmt"
)

//func TestFib(t *testing.T){
//	Fib(10)
//}

// Define a new data type "Triangle"
type Triangle struct {
	base, height float32
}
type TriangleStub struct {
	base, height float32
}

func (t TriangleStub) Area() float32 {
	return 0.5555
}

// Define a new data type "Square"
type Square struct {
	length float32
}

// Define a new data type "Rectangle"
type Rectangle struct {
	length, width float32
}

// Define a new data type "Circle"
type Circle struct {
	radius float32
}

// A method for type "Triangle"
func (t Triangle) Area() float32 {
	return 0.5 * t.base * t.height
}

// A method for type "Square"
func (l Square) Area() float32 {
	return l.length * l.length
}

// A method for type "Rectangle"
func (r Rectangle) Area() float32 {
	return r.length * r.width
}

// A method for type "Circle"
func (c Circle) Area() float32 {
	return 3.14 * (c.radius * c.radius)
}

// Define an interface as achieve abstraction
type Area interface {
	Area() float32
}

func test(a Area) {
	fmt.Println(a.Area())

}

func main() {
	fmt.Println("Hello Azure!")
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
