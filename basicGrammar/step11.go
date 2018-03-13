package main

import (
	"fmt"
	"time"
)


func chBasic(){
	// 정수형 채널을 생성.
	ch := make(chan int)

	go func(){
		ch <- 123 // 채널에 123을 보낸다.
	}()

	var i int
	i = <- ch // 채널로 부터 123을 받는다.

	fmt.Println(i)
}

func chFeature(){
	done := make(chan bool)

	go func(){
		for i:=0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	var pb bool
	// go routine이 끝나야 처리 됨.
	pb = <- done
	fmt.Println(pb)
}

/*
  channel에는 2가지 채널이 있다.
  1) Unbuffered Channel
    - chBasic(), chFeatuer() 모두 Unbuffered Channel이다.
  2) buffered Channel : 수신자가 받을 준비가 되어있지 않을지라도 지정된 버퍼만큼 데이터를 보내고 계속 다른 일을할 수 있다.
    - make(chan type, N)의 형태이며 두번째 파라미터 N에 사용할 버퍼 캣수를 넣는다.


 */
func chKinds(){
	//ch := make(chan int)
	//ch <- 1   // 수신 루틴이 없기 때문에 데드락이 걸린다. fatal error: all goroutines are asleep - deadlock!
	//fmt.Println(<-ch)

	chBuffered := make(chan int,3)
	chBuffered <- 1
	chBuffered <- 2
	chBuffered <- 3

	close(chBuffered) // close를 해줘야 3번째 채널까지임을 for에서 알수가 있다.
	                  // 해주지 않으면 fatal error: all goroutines are asleep - deadlock! 가 난다.
	                  // close()시 송신은 불가 하지만 수신은 가능.

	for chBuffer := range chBuffered {
		fmt.Println(chBuffer)
	}

	if _, success := <- chBuffered; !success {  // 첫번째는 채널 메시지, 두번째는 수신이 제대로 되었는가? 데이터가 없으면 false
		fmt.Println("더이상 데이터 없음.", success)
	}

}

func chParam(){
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)
}

func sendChan(ch chan<- string){   // chan<-
	ch <- "ACE-T"
	// t := <-ch // send only라서 오류 남.
}

func receiveChan(ch <-chan string){ // <-chan
	data := <- ch
	fmt.Println(data)
}

/*
	select문은 복수 채널들을 기다리면서 준비된(데이터를 보내온) 채널을 실행하는 기능.
    즉, 여러개의 case문에서 각각 다른 채널을 기다리다가 준비가 된 채널의 case를 실행한다.
    준비가 되지 않았으면 계속 대기하게 되지만 default문이 있으면 case문 채널이 준비되지 않더라도 계속 대기하지 않고 바로 default문을 실행한다.
 */
func chSelect(){
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

	EXIT:
		for {
			select{
			case <-done1:
				fmt.Println("run1 완료!")

			case <-done2:
				fmt.Println("run2 완료!")
				break EXIT
			}
		}
}

func run1(done chan bool){
	time.Sleep(1 * time.Second)
	done <- true

}

func run2(done chan bool){
	time.Sleep(2 * time.Second)
	done <- true
}

/*

  - go 채널
    * 채널을 통하여 데이터를 주고 받는 통로라고 볼수 있음.
    * make() 함수를 통해 미리 생성되어야 한다.
    * 채널 연산자(<-)를 통해 데이터를 보내고 받는다.
    * 채널은 흔히 go루틴들 사이 데이터를 주고 받는데 사용되는데 상대편이 준비될 때까지 채널에서 대기함으로써
      별도의 lock을 걸지 않고 데이터를 동기화 한다.
    *

 */
func main() {
	// chBasic()   // 데이터를 연산자를 통해 주고 받는것을 알수 있음.
	// chFeature() // go채널은 수신자와 송신자가 서로를 기다리는 속성이 있음.
	// chKinds()
	// chParam()
	chSelect()


}
