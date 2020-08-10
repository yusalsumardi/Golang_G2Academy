package main

import (
	model "Project_Sales/Model"
	"context"
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

var OrderList []*model.Transaction

type SalesService struct{}

//============================FUNC MEMBUAT ORDER BARU==============================
func (SalesService) NewOrder(ctx context.Context, order *model.OrderProduct) (*model.OutTransaction, error) {
	idproduct, _ := strconv.Atoi(order.GetIdproduct())
	newIdOrder := strconv.Itoa(rand.Intn(10000))
	sumproduct := len(ProductList)
	var msg string
	var neworder model.Transaction
	if idproduct <= sumproduct {
		selectedproduct := ProductList[idproduct-1] //Set Product yang dipilih
		bill := strconv.Itoa(selectedproduct.price)
		// fmt.Println(selectedproduct)
		tempneworder := model.Transaction{
			Idorder:        newIdOrder,
			Idproduct:      order.GetIdproduct(),
			Nameproduct:    selectedproduct.nameproduct,
			Bill:           bill,
			Status_Payment: "Waiting to Payment",
			No_Tf:          "",
		}
		neworder = tempneworder
		OrderList = append(OrderList, &neworder)
		msg = "Order Susscess!"
	} else {
		msg = "Order Failed: Product Number does not exist"
	}

	newoutput := model.OutTransaction{
		Transaction: &neworder,
		Message:     msg,
	}
	// fmt.Println(OrderList)
	return &newoutput, nil
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
	idtrans := payment.GetIdorder()
	idnotf := payment.GetNoTransfer()
	OrderListSelected := OrderList
	var msg string
	var newOut model.Transaction
	// OrderSeleced := *model.Transaction
	for _, value := range OrderListSelected {
		if idtrans == value.Idorder {
			if value.Status_Payment == "Waiting to Payment" {
				value.Status_Payment = "Paid"
				value.No_Tf = idnotf
				msg = "Transaction Success"
			} else {
				msg = "Transaction Failed: Transaction Has Been Paid"
			}
			TempOut := &model.Transaction{
				Idorder:        value.Idorder,
				Idproduct:      value.Idproduct,
				Nameproduct:    value.Nameproduct,
				Bill:           value.Bill,
				Status_Payment: value.Status_Payment,
				No_Tf:          value.No_Tf,
			}
			newOut = *TempOut
		}
	}
	newoutput := model.OutputPayment{
		Transaction: &newOut,
		Message:     msg,
	}
	return &newoutput, nil
}

func main() { //SET SERVER BARU
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
	d := Product{4, "Perata Gigi", 10000}
	ProductList = append(ProductList, a)
	ProductList = append(ProductList, b)
	ProductList = append(ProductList, c)
	ProductList = append(ProductList, d)
}
