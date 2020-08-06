package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var timePark = make([]int64, 0)
var park = make([]string, 0)

type Parking struct {
	Id   int64  `json:"id"`
	Tipe string `json:"tipe"`
	Plat string `json:"plat"`
}

type Send struct {
	Status  int64
	Id      string
	Message string
}

func ParkIn(w http.ResponseWriter, r *http.Request) {
	send_id := Send{200, ParkInServ(), "Masukan ID parkir untuk keluar"}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	js, _ := json.Marshal(send_id)
	w.Write(js)
}

func ParkInServ() string {
	lenBef := len(park)
	timePark = append(timePark, time.Now().Unix())
	park = append(park, strconv.Itoa(rand.Intn(1000000)))
	return park[lenBef]
}

func ParkOut(w http.ResponseWriter, r *http.Request) {
	resp, _ := ioutil.ReadAll(r.Body)
	var userParking Parking
	err := json.Unmarshal(resp, &userParking)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := servParkOut(userParking.Tipe, userParking.Plat, strconv.Itoa(int(userParking.Id)))
	var result = Send{200, strconv.Itoa(int(userParking.Id)), res}
	js, _ := json.Marshal(result)
	w.Write(js)
}

func servParkOut(tipe, plat, id_parkir string) string {
	parkir := false
	var timeParkir int64
	for key, val := range park {
		if id_parkir == val {
			parkir = true
			timeParkir = timePark[key]
			RemoveIndexInt64(&timePark, key)
			RemoveIndexStr(&park, key)
			break
		}
	}
	if !parkir {
		return "ID Parkir tidak ditemukan"
	}
	tarif := 0
	tempcounttime := time.Now().Unix() - timeParkir
	counttime := int(tempcounttime)
	switch tipe {
	case "motor":
		if counttime >= 1 {
			tarif = 3000
		}
		tarif = tarif + ((counttime - 1) * 2000)
	case "mobil":
		if counttime >= 1 {
			tarif = 5000
		}
		tarif = tarif + ((counttime - 1) * 3000)
	default:
		return "Tipe yang anda masukan tidak tersedia"
	}
	return "Waktu parkir anda: " + strconv.Itoa(counttime) + " detik, maka tarif parkir anda: Rp." + strconv.Itoa(tarif)
}

func RemoveIndexInt64(arr *[]int64, index int) {
	newArr := *arr
	*arr = append(newArr[:index], newArr[index+1:]...)
}
func RemoveIndexStr(arr *[]string, index int) {
	newArr := *arr
	*arr = append(newArr[:index], newArr[index+1:]...)
}

func main() {

	http.HandleFunc("/ParkIn", ParkIn)
	http.HandleFunc("/ParkOut", ParkOut)

	fmt.Println("running server...")
	http.ListenAndServe(":8082", nil)
}
