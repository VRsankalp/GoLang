package main

import (
	"fmt"
	// "time"
)
func sum(res chan int , num1 int , num2 int){
	numR := num1+num2
	res<- numR



}
// // func pricessNum(numChan chan string ){
// 	fmt.Println("priccssing" , <- numChan)

// }
func task(done chan bool){
	defer func(){done <- true}()
	fmt.Print("processing,.........")
	// done <- true
}
func main() {

	done := make(chan bool)
	go task(done)

	<-done//block


	res:= make (chan int)

	go sum(res , 4 ,5)

	re:=<- res
	fmt.Println(re)


// meesgageChain := make(chan string)
// 	go pricessNum(meesgageChain)
// 	meesgageChain<-"ping"
// 	time.Sleep(time.Second*2)
	
}