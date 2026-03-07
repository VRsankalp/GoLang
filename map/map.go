package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)
	// setting an elemenet

	m["name"] = "golang"
	m["hello"] = "babay"
	fmt.Println(m["name"])
	fmt.Println(m["hello "])

	ma:=make(map[string]int)
	ma["age"]=30
	fmt.Println(ma["phonr"])

/// length of  a map
fmt.Println(len(m))
//delete a function
delete(m ,"name")
fmt.Println(len(m))

// clear the map
clear(m)

// without make mp function

a:=map[string]int{"price": 30 , "phone": 3}
fmt.Println(a)

/// check emlemnt preseent
 
val,ok :=a["price"]
fmt.Println(val)
if(ok){
	fmt.Println("allOk")

} else{
	fmt.Print("not pk")
}
// comaprision of maps
fmt.Println(maps.Equal(a , a))


}