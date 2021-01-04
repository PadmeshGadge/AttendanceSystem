package main
import(
	"fmt"
	"attendSys/admin"
	"attendSys/staff"
)
func main(){
	var n int
	i:=0
	for i<1{
		fmt.Println("\nSelect Login Type\n1.Admin\n2.Staff\n3.Exit\n")
		fmt.Scanln(&n)
		switch(n){
			case 1:admin.Login()
			case 2:staff.Login()
			case 3:i++
			default:i++
		}
	}
}