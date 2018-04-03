package main

import "fmt"

func main() {

	// 변수 선언
	// var 변수명 자료형
	var acet string
	acet = "박태하"

	// := 를 사용하여 짦은 선언을 할수 있다.
	th := "taeha"

	fmt.Println(acet)
	fmt.Println(th)

	// 변수 여러개 선언 하기
	var a,b,c,d,e,f,g string
	a = "a"
	b = "b"
	c = "c"
	d = "d"
	e = "e"
	f = "f"
	g = "g"

	fmt.Println(a + b + c + d + e + f + g)

	// var와 ()를 사용하면 변수 여러개를 선언하고 초기화할 수 있음.
	var (
		x, y int =10, 20
	)

	fmt.Println(x, y)

}
