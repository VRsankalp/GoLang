package main

import (
	"fmt"
	"time"
)
//structure embeedinh
type customer struct{
	cid string
	cName string
}
type order struct {
	id        string
	amount    float32
	status    string
	creadteAt time.Time
	customer
}
// constructor  in go no constr
func newOrder(id string, amount float32, status string)*order{
	myOrder:=order{
		id:id,
		amount: amount,
		status: status,
		// creadteAt :time.Now(),
	}
	return &myOrder
}

// reciver type
func (o *order) chnageStstaus(status string){
	o.status=status
}


func main(){
	// myOrder:=order{
	// 	id:"12",
	// 	amount: 59,
	// 	status: "recives",
	// 	creadteAt :time.Now(),
	// }
myOrder:=newOrder("1",32323,"paseed")
fmt.Println(myOrder)
// myOrder.chnageStstaus("56")
fmt.Println(myOrder.amount)

}