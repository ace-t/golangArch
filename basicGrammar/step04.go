package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"bytes"
	"net/url"
	"encoding/json"
)

type Person struct {
	Name string
	Age int
}

func main(){
	//httpGet1()
	//httpGet2()
	//httpPost1()
	//httpPostForm()
	//httpPostJson()
	httpPostRequestJson()

}


// HTTP GET - request시에 header를 변경한다던지 등의 세밀한 컨트롤을 할 수 없다는 단점이 있다.
func httpGet1(){
	// GET
	resp, err := http.Get("http://www.daum.net")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// result
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err.Error()
	}
	fmt.Printf("%s\n", string(data))
}

func httpGet2(){
	req, err := http.NewRequest("GET", "http://www.naver.com", nil)
	if err != nil {
		panic(err)
	}

	// 필요시 헤더 추가
	req.Header.Add("User-Agent","Crawler")

	// client 객체에서 Request를 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes) // 바이트를 문자열로
	fmt.Println(str)

}

func httpPost1(){
	reqBody := bytes.NewBufferString("Post Test!")
	resp, err := http.Post("http://www.daum.net", "text/plain", reqBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Response 체크
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

func httpPostForm(){
	resp, err := http.PostForm("http://www.daum.net", url.Values{"Name":{"ParkTaeha"}, "Age":{"10"}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func httpPostJson(){
	person := Person{"Taeha", 20}
	pBytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(pBytes)
	resp, err := http.Post("http://www.daum.net", "application/json", buff)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func httpPostRequestJson(){
	person := Person{"ParkTaeha", 10}
	pBytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(pBytes)

	// Requset 객체 생성
	req, err := http.NewRequest("POST", "http://www.daum.net", buff)
	if err != nil {
		panic(err)
	}

	// Content-Type header 추가
	req.Header.Add("Content-Type", "application/json")

	// Client객체에서 Request 실행
	//client := &http.Client{}
	client := http.Client{}  // addressable하면 자동으로 &를 붙여준다. 
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Response 체크
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println("hahaha")
		println(str)
	}
}

