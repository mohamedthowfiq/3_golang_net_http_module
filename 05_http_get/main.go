package main

import (
	"fmt"
	"net/http"
)

func main(){
	url := "https://jsonplaceholder.typicode.com/todos"

	resp, error := http.Get(url)
	if error !=nil{
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("States code",req.status)

 }