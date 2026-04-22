package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CatFactResponse struct{
	Fact string `json:"fact"`
	Length int `json:"length"`
}

func writeJSON(w http.ResponseWriter, status int, data any){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func fetchCatFact()(CatFactResponse,error){
	url:= "http://catfact.ninja/fact"

	resp,err := http.Get(url)
	if err != nil{
		return CatFactResponse{},err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return CatFactResponse{},fmt.Errorf("external api failed: %s",resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil{
		return CatFactResponse{},err
	}

	var data CatFactResponse

	if err :=json.Unmarshal(bodyBytes,&data); err != nil{
		return CatFactResponse{},err
	}

	return data,nil
}

func externalHandler(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodGet{
			writeJSON(w , http.StatusMethodNotAllowed,map[string]any{
				"ok": false,
				"error":"only get method will be allowed",
			})
			return
		}

		data,err := fetchCatFact()
		if err != nil{
			writeJSON(w,http.StatusBadGateway,map[string]any{
				"ok":false,
				"error":"failed to fetch data",
			})
			return
		}
		writeJSON(w,http.StatusOK,map[string]any{
			"ok":true,
			"timeStamp":time.Now().UTC(),
			"external":map[string]any{
				"source": "Catfact.ninja",
				"fact": data.Fact,
				"length": data.Length,
			},
		})
}



func main() {

	http.HandleFunc("/external",externalHandler)
	err := http.ListenAndServe(":8080",nil)
	fmt.Println(err)
}