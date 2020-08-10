package main

import (
	"context"
	"fmt"
	"log"

	model "Lat1/ProtoMod"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	URL = "localhost:8081"
)

func main() {
	//Set Connection
	con, err := grpc.Dial(URL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer con.Close()
	newserv := model.NewParkingServiceClient(con)
	var input int

	for input != 99 {
		fmt.Println("====Menu====") //Pilihan Menu
		fmt.Println("1. Masuk Parkir")
		fmt.Println("2. Keluar Parkir")
		fmt.Println("99. EXIT")
		fmt.Println("Pilih Menu:")
		fmt.Scanln(&input) //Scan menu
		switch input {
		case 1:
			newId, newTime := GetId(newserv)
			fmt.Println("Id :", newId)
			fmt.Println("Time In :", newTime)
			input = 0
		case 2:
			var (
				id, platno, tipe string
			)
			fmt.Println("Input your ID Parking")
			fmt.Scan(&id)
			fmt.Println("Input Vehicle type [motor/mobil]")
			fmt.Scan(&tipe)
			fmt.Println("Input Vehicle Number")
			fmt.Scan(&platno)
			datain := &model.InputData{
				Id:     id,
				Tipe:   tipe,
				Platno: platno,
			}
			result := CheckOut(newserv, datain)
			fmt.Println(result)
			input = 0
		}
	}
}

func GetId(mod model.ParkingServiceClient) (string, string) {
	resp, err := mod.ParkIn(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatalf(err.Error())
	}
	return resp.Id, resp.Time
}

func CheckOut(c model.ParkingServiceClient, input *model.InputData) string {
	resp, err := c.ParkOut(context.Background(), input)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return "Id :" + resp.Id + "\n" +
		"Platno :" + resp.Platno + "\n" +
		"Duration :" + resp.Duration + "\n" +
		"Message :" + resp.Message
}
