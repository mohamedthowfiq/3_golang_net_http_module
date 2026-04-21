package main

import(
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter,r *http.Request){

	_,_ = w.Write([]byte("welcome try to /hello?name=thowfiq"))

}

func helloHandler(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")
	if name ==""{
		name = "Guest"
	}
	_,_ = w.Write([]byte(name))
}



func main(){

	http.HandleFunc("/",rootHandler)
	http.HandleFunc("/hello",helloHandler)
	err:=http.ListenAndServe(":8080",nil)
	fmt.Print(err)
}