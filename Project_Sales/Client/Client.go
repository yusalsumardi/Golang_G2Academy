package main

import (
	"context"
	"fmt"
	"log"

	model "Project_Sales/Model"

	"google.golang.org/grpc"
)

const (
	URL = "localhost:8080"
)

func main() {
	//===============Set Connection=================
	con, err := grpc.Dial(URL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer con.Close()
	newserv := model.NewSalesServiceClient(con)

	//===============Menu Client======================
	var input int
	for input != 99 {
		fmt.Println("====Menu====") //Pilihan Menu
		fmt.Println("1. Buy a Product")
		fmt.Println("2. Check Transaction")
		fmt.Println("3. Payment")
		fmt.Println("99. EXIT")
		fmt.Println("Pilih Menu:")
		fmt.Scanln(&input) //Scan menu
		switch input {
		case 1:
			fmt.Println("====Our Product====") //Pilihan Menu
			fmt.Println("1. Sumplemen Peninggi Badan")
			fmt.Println("2. Skincare Glowing")
			fmt.Println("3. Sumplemen Penggemuk Berat Badan")
			fmt.Println("Select a Product:")
			var inputproduct string
			fmt.Scanln(&inputproduct)
			neworder := model.OrderProduct{
				Idproduct: inputproduct,
			}
			result := NewOrder(newserv, &neworder)
			fmt.Println("============================")
			fmt.Println(result)
			fmt.Println("============================")
		case 2:

		case 3:
		}
	}

}

func NewOrder(mod model.SalesServiceClient, order *model.OrderProduct) string {
	resp, err := mod.NewOrder(context.Background(), order)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return "Id Order\t: " + resp.Idorder + "\n" +
		"Id Product\t: " + resp.Idproduct + "\n" +
		"Product Name\t: " + resp.Nameproduct + "\n" +
		"Price\t\t: " + resp.Bill + "\n" +
		"Payment Status\t: " + resp.Status_Payment
}
