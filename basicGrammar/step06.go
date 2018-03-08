package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func hello(){
	fmt.Println("TAEHA")
}

func useChiLib(){
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("root."))
	})

	http.ListenAndServe(":3333",r)

}

/*
 	go루틴을 사용해보자.
 */
func main() {
	go hello()  // go루틴은 메인함수와 동시에 실행 됨.
	            // go 키워드는 함수 여러개를 동시에 실행할 수 있음. go routine끼리 통신 가능.
	//fmt.Scanln()
	useChiLib()
}




