package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Only GET is allowed",http.StatusMethodNotAllowed)
		return
	}

	_,_ = w.Write([]byte("Hello from GO net/http server"))
}

func main(){

	// to register route
	http.HandleFunc("/hello",helloHandler)
	fmt.Println("try goint to  8080 port ")

	// to listen particular port
	err:=http.ListenAndServe(":8080",nil)
	fmt.Println(err)
}