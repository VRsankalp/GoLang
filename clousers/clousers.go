package main

import "fmt"

func count() func() int {
	cot := 0
	return func() int {
		cot += 1
		return cot
	}
}
func main() {
	inc := count()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}