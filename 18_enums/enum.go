package  main
//enum tyoes 
import"fmt"
type orderStatus int

const(
	recives  orderStatus =iota
	confirmed  
	prepared
	deliversd
)
func ChnageOder(status  orderStatus){
fmt.Println(status)
}
func main(){
	ChnageOder(confirmed)

}
