package main

import "fmt"
// multiple parameter in functions 
func sum(nums ...int) int {
	to := 0
	for _, num := range nums {
		to += num
	}
	return to
}
func main() {
	num:=[]int{21,3,1,1,1,23,}
	res := sum(num...)
	fmt.Println(res)

}
