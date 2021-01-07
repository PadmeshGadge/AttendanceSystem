package student
import(
	"fmt"
	_"mysql"
	"os"
	"attendSys/genKey"
	"attendSys/dbFuncs"
)
var g = genKey.Store
func EditStudent(uname,key string){
	if key == g.Getkey(uname){
		var roll,n int
		fmt.Print("\nEnter roll no to update: ")
		fmt.Scanln(&roll)
		fmt.Println("\n1.Update Name\t2.Update Email")
		fmt.Scanln(&n)
		switch(n){
			case 1:{
				var str string
				fmt.Print("\nEnter new name: ")
				fmt.Scanln(&str)
				_,err := dbFuncs.Db.Query("UPDATE student_tbl SET Name=? WHERE roll_no=?",str,roll)
				if err!=nil{throwError(err)}
			}
			case 2:{
				var str string
				fmt.Print("\nEnter new email: ")
				fmt.Scanln(&str)
				_,err := dbFuncs.Db.Query("UPDATE student_tbl SET Email=? WHERE roll_no=?",str,roll)
				if err!=nil{throwError(err)}
			}
		}
		fmt.Println("Record updated successfully")
	}else{fmt.Println("Wrong Key")}
}
func throwError(err error){
	fmt.Println("Database error")
	os.Exit(1)
}