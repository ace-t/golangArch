package main

import "fmt"

type Animal struct {
	name string
	size int
}

/*  이 부분이 제일 중요하다고 생각한다.
   (i Animal) 은 receiver라고 불리우며 메소드와 Animal Struct와 연결 시켜주는 부분이다.
   그래서 main부분에서 intro.introduce() 형태로 호출이 가능하다. 대박!
   receiver에는 2가지가 있다.
   1) Value receiver : struct의 data를 복사하여 전달.
      그래서 struct의 필드값이 변경 되더라도 호출자의 데이터는 변경되지 않는다.
   2) Point receiver : 포인터를 전달.

*/
func (i Animal)introduceValue() string {
	i.size++
	return fmt.Sprintf("이 동물의 이름은 %s 이며 크기는 %d이다.", i.name, i.size)
}

func (i *Animal)introducePoint() string {
	i.size++
	return fmt.Sprintf("이 동물의 이름은 %s 이며 크기는 %d이다.", i.name, i.size)
}

/*
	Step02에서는 Method에 대한 코드를 작성하여 본다.
	Go는 struct가 필드만을 가지며, 메소드는 별도로 분리되어 정의된다.
 */
func main() {

	// Value receiver
	intro := Animal{"만순이", 5}
	animalIntro := intro.introduceValue()
	fmt.Println(animalIntro, intro.size) // Value receiver라서 값이 변경되지 않는다.
	// 이 동물의 이름은 만순이 이며 크기는 6이다. 5

	// Point receiver
	intro3 := Animal{"토끼", 15}
	animalIntro = intro3.introducePoint()
	fmt.Println(animalIntro, intro3.size)  // Point receiver라서 값이 변경 된다.
	// 이 동물의 이름은 토끼 이며 크기는 16이다. 16
}
