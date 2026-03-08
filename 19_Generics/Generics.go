package main

import "fmt"

func pS[T int | string | bool](item []T) {
	for _, num := range item {
		fmt.Println(num)
	}
}
func main() {
	val := []int{2,4,21,4}
	sval := []string{"hi" , "hello"}
	bval := []bool{true , false}
	pS(val)
	pS(sval)
	pS(bval )

}