package main

import (
	"fmt"
	"time"
)

func goroutineFunc(name string){
	for i :=0; i < 3; i++ {
		fmt.Println(name, "****", i)
	}
}

func goroutineRanageFor(){
	// 함수 밖 변수 선언
	names := []string{"에이스티","애니자바","굿택","틴톨"}
	for idx, name := range names {
		name := name  // 이 소스에서 가장 핵심적인 부분 이다. range안의 범위에서는 namme이 모두 같은 주소로 나오게 된다.
		println(&name, name)
		go func(){  // goroutine 익명함수
			println(name,idx)
			goroutineFunc(name)
		}()
	}

	//fmt.Scanln()
	time.Sleep(time.Second*3)

	/*
	result : 같은 주소에 담겨진다.
	0xc420043f00   에이스티
    0xc420043f00   애니자바
    0xc420043f00   굿택
    0xc420043f00   틴톨

	name := name 선언 후에는 주소값이 다르다.
	0xc42000e1e0 에이스티
	0xc42000e1f0 애니자바
	0xc42000e200 굿택
	0xc42000e210 틴톨

	 */
}


func goroutineGenFor(){
	// 함수 밖 변수 선언
	names := []string{"에이스티","애니자바","굿택","틴톨"}
	for i:=0; i < len(names); i++{
		i := i  // ranage나 일반적인 for나 동일하다.
		go func(){  // goroutine 익명함수
			println(i,names[0])
			println(i,names[1])
		}()
	}
	time.Sleep(time.Second*3)

	/*
	result
		1 에이스티
		1 애니자바
		2 에이스티
		2 애니자바
		3 에이스티
		3 애니자바
		0 에이스티
		0 애니자바
	 */

}

// go 루틴에 대해서! 알아보자~
// go는 디폴트로 하나의 CPU를 사용. 여러개의 CPU 사용한다면 다중 CPU에서 병렬처리를할 수 있도록 아래와 같은 함수를 호출하여야 한다.
// runtime.GOMAXPROCS(cpu수)함수를 호출.
func main() {
	// 함수를 동기적으로 수행
	//goroutineFunc("태하-Sync")

	// 함수를 비동기적으로 수행.
	//go goroutineFunc("미애-Async")
	//go goroutineFunc("한전-Async")
	//go goroutineFunc("영금-Async")

	//goroutineRanageFor()
	goroutineGenFor()



}


