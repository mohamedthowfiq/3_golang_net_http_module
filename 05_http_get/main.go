package main

import (
	"fmt"
	"net/http"
)

func main(){
	url := "https://jsonplaceholder.typicode.com/todos"

	resp, err := http.Get(url)
	if err !=nil{
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status code",resp.StatusCode)
	fmt.Println("Status",resp.Status)

 }