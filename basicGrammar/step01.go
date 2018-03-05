package main

import "fmt"


// Struct는 Go에서 Custom Data Type를 표현하는데 사용 된다.
// 선언은 type문을 사용한다. 뒤에 학습할 interface도 동일하다.
type person struct {
	name string
	age int
}

func newPerson() *person {
	p := person{}
	p.name = "초기화"
	p.age = 0
	return &p
}

/*
	Step01에서는 Struct에 대한 코드를 작성하여 본다.
 */
func main() {
	// person 객체 생성
	p := person{}
	// filed 값 설정
	p.name = "박태하"
	p.age = 100
	// 출력
	fmt.Println(p)

	// 객체 생성에는 여러가지고 있는데 또 다른 방법은 new를 사용하는 것이다.
	p2  := new(person)  // pointer를 전달.
	p2.name = "태하팍"
	p2.age = 20
	fmt.Println(*p2)

	// 또는 바로 데이터를 초기화 하여 생성하는 방법이 있다.
	p3 := person{"태하하", 30}
	fmt.Println(p3)
	p4 := person{name:"박태하핫", age:10}
	fmt.Println(p4)

	// 생성자 함수 호출 : 필드가 사용하기 전에 초기화가 되어야하는 경우에 사용되어진다.
	p5 := newPerson()
	p5.age = 1
	fmt.Println(*p5)
}




