package main

import "fmt"

func main() {
	// initalize
	var nums []int
	fmt.Println(nums)


	//initnalize using make

	var num = make([]int , 2 , 5)
	fmt.Println(len(num))
	fmt.Println(cap(num))
	num = append(num, 1)
	num = append(num, 3)
	num = append(num, 4)
	num = append(num, 1)
	num = append(num, 22)
	fmt.Println(num)
	fmt.Println(cap(num))

	// cpoy function 
	var nums2 = make([]int  , len(num))
	copy(nums2 , num)
	fmt.Println(nums2)


}