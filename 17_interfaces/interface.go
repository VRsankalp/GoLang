package main

import "fmt"

type paymenter interface{
	pay(amount float32)

}

type payment struct{
	gateWay paymenter
}

func (p payment) makePayment(amount float32) {
	// razorPaymentGw := razorpay{}
	// stripePaymemtGw := stripe{}
	// razorPaymentGw.pay(amount)
	p.gateWay.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment " , amount)
}

//striver method//////////////////////////
type stripe struct{}
func(s stripe)pay(amount float32){
	fmt.Println("making payment with stripve" , amount)
}

func main() {
	newPayment :=  payment{
		gateWay: stripe{},
	}
	newPayment.makePayment(100)	 

}