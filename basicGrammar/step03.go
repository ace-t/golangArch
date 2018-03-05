package main

import (
	"math"
	"fmt"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64{return r.height*r.width}
func (r Rect) perimeter() float64{return 2*(r.width+r.height)}

// Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64{return math.Pi * c.radius * c.radius}
func (c Circle) perimeter() float64{return 2*math.Pi*c.radius}

/*
	Step03에서는 Interface에 대한 코드를 작성하여 본다.
	Go는 struct가 필드만을 가지며, 메소드는 별도로 분리되어 정의된다.
    Interface는 메소드들의 집합체이다.
 */
func main() {
	r := Rect{10,20}
	c :=Circle{10}
	showArea(r,c)

	var x interface{}
	x = 1
	x = "Terry"
	printIt(x)

	var t interface{} = 1000
	h := t
	fmt.Println(h)
	a := t.(int)
	fmt.Println(a)


}

func showArea(shapes ...Shape){
	for _, s := range shapes {
		a := s.area()
		println(a)
	}
}

func printIt(v interface{}){
	fmt.Println(v)
}

// interface에 있는 메소드들은 모구 구현해줘야 함.