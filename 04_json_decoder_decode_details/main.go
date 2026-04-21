package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func writeJSON(w http.ResponseWriter, status int,data any){
	w.Header().Set("Content-type","Application/json")
	w.WriteHeader(status)
	_=json.NewEncoder(w).Encode(data)
}

type TestRequest struct{
	Name string`json:"name"`


}

func testHandler(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost{
		writeJSON(w, http.StatusMethodNotAllowed,map[string]any{
			"ok":false,
			"error":"Only POST is allowed",
		})
		return
	}

	defer r.Body.Close()
	var req TestRequest
	decoder :=json.NewDecoder(r.Body)
	err := decoder.Decode(&req);
	if  err != nil {
		writeJSON(w,http.StatusBadRequest, map[string]any{
			"ok": false,
			"error":"Invalid json format",
		})
		return
	}
	req.Name = strings.TrimSpace(req.Name)

	if req.Name ==""{
		writeJSON(w,http.StatusBadRequest,map[string]any{
			"ok": false,
			"error":"name must not empty!",
		})
		return
	}

	writeJSON(w,http.StatusOK,map[string]any{
		"ok":true,
		"data":req,
		"timestamp": time.Now().UTC(),
	})
}	

func main(){

	http.HandleFunc("/test",testHandler)

	err := http.ListenAndServe(":8080",nil)
	fmt.Println(err)
}