package main

import "fmt"

func chnN(num *int)  {
	*num = 1
	fmt.Println(num)

}
func main() {
num:=12
chnN(&num)
fmt.Println(num)
}