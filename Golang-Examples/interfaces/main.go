package main

import "fmt"

// Shape is an interface that defines the area method
type Shape interface {
    area() float64
}

// Rectangle is a struct that represents a rectangle
type Rectangle struct {
    width  float64
    height float64
}

// Circle is a struct that represents a circle
type Circle struct {
    radius float64
}

// area is a method that calculates the area of a rectangle
func (r Rectangle) area() float64 {
    return r.width * r.height
}

// area is a method that calculates the area of a circle
func (c Circle) area() float64 {
    return 3.14 * c.radius * c.radius
}

// getArea is a function that takes a Shape interface as a parameter
func getArea(s Shape) float64 {
    return s.area()
}

func main() {
    r := Rectangle{width: 5, height: 7}
    c := Circle{radius: 3}

    fmt.Println("Area of rectangle:", getArea(r))
    fmt.Println("Area of circle:", getArea(c))
}