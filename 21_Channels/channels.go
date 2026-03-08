package main

import (
	"fmt"
	"time"
)
// func sum(res chan int , num1 int , num2 int){
// 	numR := num1+num2
// 	res<- numR



// }
// // // func pricessNum(numChan chan string ){
// // 	fmt.Println("priccssing" , <- numChan)

// // }
// func task(done chan bool){
// 	defer func(){done <- true}()
// 	fmt.Print("processing,.........")
// 	// done <- true
// }


func emilSender(email <-chan string , done chan <-bool){
	defer func() { done <- true}()
	for e:= range email{
		fmt.Println("senfing emil" , e)
		time.Sleep(time.Second)
	}
}
func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func(){
		chan1<-10
		
	}()
	go func(){
		chan2<-"pong"
		
	}()
	for i:= 0 ; i<2 ; i++{
		select{
		case chanval1:= <- chan1:
			fmt.Println("received data",chanval1)
		case chanVal2:= <- chan2:
			fmt.Println("received data",chanVal2)
		}
	}
















// 	emailChan := make(chan  string , 100)
// 	done := make(chan bool)
//    go emilSender(emailChan , done )

//    for i := 1; i <= 5 ; i++ {
// 		email := fmt.Sprintf("%d@gmail.com", i)
// 		emailChan <- email
// 	}
// 	close(emailChan)
//         <-done // blocking

	// done := make(chan bool)
	// go task(done)

	// <-done//block


	// res:= make (chan int)

	// go sum(res , 4 ,5)

	// re:=<- res
	// fmt.Println(re)


// meesgageChain := make(chan string)
// 	go pricessNum(meesgageChain)
// 	meesgageChain<-"ping"
// 	time.Sleep(time.Second*2)
	
}