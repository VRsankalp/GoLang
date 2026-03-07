package main

import "fmt"

func add(a int, b int) int {
	return a + b

}
// multiple return vales

func lang()( string , int , bool	){
	return "hello" , 34 , true
}
func proce(fn func(a int )int){
	fn(1)
}
// main function
func main() {
	res := add(3, 5)
	fmt.Println(res)
	 l , _ , m :=lang()
	 fmt.Println(l , m   )
	 
	 
	 fn := func(a int )int{
		return 2
	 }

	 proce(fn)

}