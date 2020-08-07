package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	model "Lat1/ProtoMod"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	port = "8081"
)

type ParkingService struct{}

var Parking = make(map[string]time.Time)

func (ParkingService) ParkIn(ctx context.Context, empty *empty.Empty) (*model.UserIn, error) {
	newId := strconv.Itoa(rand.Intn(10000))
	newTime := time.Now()
	TimeString := newTime.String()
	Parking[newId] = newTime
	log.Println("Id :", newId, "\t", "Time :", newTime)
	message := &model.UserIn{
		Id:   newId,
		Time: TimeString,
	}
	return message, nil

}

func (ParkingService) ParkOut(ctx context.Context, UserInput *model.InputData) (*model.UserOut, error) {
	id := UserInput.Id
	setNewPark := Parking[id]
	tempcounttime := time.Now().Unix() - setNewPark.Unix()
	counttime := int(tempcounttime)
	strcounttime := strconv.Itoa(counttime)
	var tarif int
	tipe := UserInput.Tipe
	platno := UserInput.Platno
	msg := ""

	if _, found := Parking[id]; found {
		switch tipe {
		case "motor":
			if counttime >= 1 {
				tarif = 3000
			}
			tarif = tarif + ((counttime - 1) * 2000)
			tariffix := strconv.Itoa(tarif)
			msg = "Biaya Parkir anda: Rp." + tariffix
			delete(Parking, id)

		case "mobil":
			if counttime >= 1 {
				tarif = 5000
			}
			tarif = tarif + ((counttime - 1) * 3000)
			tariffix := strconv.Itoa(tarif)
			msg = "Biaya Parkir anda: Rp." + tariffix
			delete(Parking, id)
		}
	} else {
		msg = "Parking Id not found please input another Id"
	}
	UserOut := &model.UserOut{
		Id:       id,
		Platno:   platno,
		Duration: strcounttime,
		Message:  msg,
	}
	return UserOut, nil
}

func main() {
	log.Println("Server Running --> PORT", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failedd to listen: %v", err)
	}

	s := grpc.NewServer()
	var parkingService ParkingService
	model.RegisterParkingServiceServer(s, parkingService)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Print(s.Serve(lis))

}
