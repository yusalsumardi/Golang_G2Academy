package main

import (
	"context"
	"fmt"
	"log"

	model "Project_Sales/Model"

	"github.com/golang/protobuf/ptypes/empty"
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
			fmt.Println("4. Perata Gigi")
			fmt.Println("Select a Product:")
			var inputproduct string
			fmt.Scanln(&inputproduct)
			neworder := model.OrderProduct{
				Idproduct: inputproduct,
			}
			result, msg := NewOrder(newserv, &neworder)
			fmt.Println("\n=================================================")
			fmt.Println("Id Order\t: " + result.Idorder)
			fmt.Println("Id Product\t: " + result.Idproduct)
			fmt.Println("Product Name\t: " + result.Nameproduct)
			fmt.Println("Price\t\t: Rp." + result.Bill)
			fmt.Println("Payment Status\t: " + result.Status_Payment)
			fmt.Println("No_Transfer\t: " + result.No_Tf)
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
			fmt.Println(msg)
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++\n")
		case 2:
			orderlist := GetOrderList(newserv)
			fmt.Println(orderlist)
		case 3:
			var inputOrderID, inputNoTransfer string
			fmt.Println("====Payment====") //Pilihan Menu
			fmt.Println("Input Order ID :")
			fmt.Scanln(&inputOrderID)
			fmt.Println("Input No Transfer :")
			fmt.Scanln(&inputNoTransfer)
			newpayment := model.InputPayment{
				Idorder:    inputOrderID,
				NoTransfer: inputNoTransfer,
			}
			transaction, msg := SetPayment(newserv, &newpayment)
			fmt.Println("\n=================================================")
			fmt.Println("Id Order\t: " + transaction.Idorder)
			fmt.Println("Id Product\t: " + transaction.Idproduct)
			fmt.Println("Product Name\t: " + transaction.Nameproduct)
			fmt.Println("Price\t\t: Rp." + transaction.Bill)
			fmt.Println("Payment Status\t: " + transaction.Status_Payment)
			fmt.Println("No_Transfer\t: " + transaction.No_Tf)
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
			fmt.Println(msg)
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++\n")
		}
	}

}

func NewOrder(mod model.SalesServiceClient, order *model.OrderProduct) (*model.Transaction, string) {
	resp, err := mod.NewOrder(context.Background(), order)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return resp.Transaction, resp.Message
}

func GetOrderList(mod model.SalesServiceClient) string {
	resp, err := mod.ListOrder(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatalf(err.Error())
	}
	OrderList := resp.GetOrderList()
	for _, value := range OrderList {
		fmt.Printf("------------------------------------------\n")
		fmt.Printf("Id Order : %s\n"+
			"Id Product : %s\n"+
			"Product Name : %s\n"+
			"Price : Rp.%s\n"+
			"Payment Status : %s\n"+
			"No_Tf : %s\n", value.Idorder, value.Idproduct, value.Nameproduct, value.Bill, value.Status_Payment, value.No_Tf)
		fmt.Printf("------------------------------------------\n")

	}
	return "DONE\n"
}

func SetPayment(mod model.SalesServiceClient, payment *model.InputPayment) (*model.Transaction, string) {
	resp, err := mod.Payment(context.Background(), payment)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return resp.Transaction, resp.Message
}
