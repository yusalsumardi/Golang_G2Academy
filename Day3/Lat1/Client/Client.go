package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Parkir struct {
	id   int64  `json:"id"`
	tipe string `json:"tipe"`
}

type Error struct {
	Status  int64
	Message string
}

func main() {
	http.HandleFunc("/Masuk", ParkIn)   //MAPPING
	http.HandleFunc("/Keluar", ParkOut) //MAPPING
	fmt.Println("Client Running at Port 8082...")
	http.ListenAndServe(":8083", nil)
}

//=========================REQUEST MASUK============================
func ParkIn(w http.ResponseWriter, r *http.Request) {
	var client = &http.Client{}
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		w.WriteHeader(200)
		request, _ := http.NewRequest("POST", "http://localhost:8082/ParkIn", nil) //SEND REQUEST TO SERVER
		request.Header.Set("Content-Type", "application/json")
		resp, _ := client.Do(request)
		defer resp.Body.Close()
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(string(responseData)))
	} else {
		error := Error{400, "Method Not Allowed"}
		js, err := json.Marshal(error)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(400)
		w.Write(js)
	}
}

//=========================REQUEST KELUAR============================
func ParkOut(w http.ResponseWriter, r *http.Request) {
	var client = &http.Client{}
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		w.WriteHeader(200)
		resp2, _ := ioutil.ReadAll(r.Body)
		request, _ := http.NewRequest("POST", "http://localhost:8082/ParkOut", bytes.NewBuffer(resp2)) //SEND REQUEST AND BODY TO SERVER
		request.Header.Set("Content-Type", "application/json")
		resp, _ := client.Do(request)
		defer resp.Body.Close()
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(string(responseData)))
	} else {
		error := Error{400, "Method Not Allowed"}
		js, err := json.Marshal(error)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(400)
		w.Write(js)
	}
}
