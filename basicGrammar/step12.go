package main

import (
	"os"
	"fmt"
)

/*

	defer는 defer를 호출하는 함수가 리턴하기 직전에 실행하게 된다.

 */
func deferFunc(){
	f, err := os.Open("config/test.txt")
	if err != nil{
		panic(err)
	}

	// deferFunc() 마지막에 파일 close() 실행
	defer f.Close()

	// file 읽기
	bytes := make([]byte, 1024)
	len, err := f.Read(bytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(len) // hi everyone.. : 13
}

/*
	panic()는 현재 함수를 즉시 멈추고 현재 함수에 defer함수들을 모두 실행한 후 즉시 리턴한다.
 	이런 실행방식은 상위 함수에도 동일하게 적용되고, 계속 콜스택을 타고 올라가며 적용된다.
	그리고 마지막에는 프로그램이 에러를 내고 종료하게 된다.
 */
func panicFunc(){
	openFile("noText.txt")
}

func openFile(txtName string){

	// defere 함수. panic 호출시 실행됨
	defer func() {
		// Go 내장함수인 recover()함수는 panic 함수에 의한 패닉상태를 다시 정상상태로 되돌리는 함수
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(txtName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

/*
	defer, panic, recover 함수에 대해서 알아보자.

 */
func main() {
	deferFunc()
	panicFunc()
}
