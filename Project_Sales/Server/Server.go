package main

import (
	model "Project_Sales/Model"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type Product struct {
	idproduct   int
	nameproduct string
	price       int
}

var ProductList []Product

var OrderList []*model.AfterOrder

type SalesService struct{}

//============================FUNC MEMBUAT ORDER BARU==============================
func (SalesService) NewOrder(ctx context.Context, order *model.OrderProduct) (*model.AfterOrder, error) {
	idproduct, _ := strconv.Atoi(order.GetIdproduct())
	newIdOrder := strconv.Itoa(rand.Intn(10000))
	selectedproduct := ProductList[idproduct-1]
	bill := strconv.Itoa(selectedproduct.price)
	fmt.Println(selectedproduct)
	neworder := model.AfterOrder{
		Idorder:        newIdOrder,
		Idproduct:      order.GetIdproduct(),
		Nameproduct:    selectedproduct.nameproduct,
		Bill:           bill,
		Status_Payment: "Waiting to Payment",
	}

	OrderList = append(OrderList, &neworder)
	fmt.Println(OrderList)
	return &neworder, nil
}

//============================MEMANGGIL SEMUA TRANSAKSI==============================
func (SalesService) ListOrder(ctx context.Context, empty *empty.Empty) (*model.OrderList, error) {
	NewOrderList := model.OrderList{
		OrderList: OrderList,
	}
	return &NewOrderList, nil
}

//============================PAYMENT==============================
func (SalesService) Payment(ctx context.Context, payment *model.InputPayment) (*model.OutputPayment, error) {
	// idtrans := payment.GetIdorder()
	// idnotf := payment.GetNoTransfer()
	Contoh := OrderList[1]
	msg := "Yey"
	Newreturn := model.OutputPayment{
		Transaction: Contoh,
		Message:     msg,
	}
	return &Newreturn, nil
}

func main() {
	IsiProduk()
	log.Println("Server Running --> PORT", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failedd to listen: %v", err)
	}

	s := grpc.NewServer()
	var salesService SalesService
	model.RegisterSalesServiceServer(s, salesService)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Print(s.Serve(lis))
}

func IsiProduk() {
	a := Product{1, "Sumplemen Peninggi Badan", 700000}
	b := Product{2, "Skincare Glowing", 500000}
	c := Product{3, "Sumplemen Penggemuk Berat Badan", 400000}
	ProductList = append(ProductList, a)
	ProductList = append(ProductList, b)
	ProductList = append(ProductList, c)
}
