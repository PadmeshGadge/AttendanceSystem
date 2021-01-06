package main
import(
	"fmt"
	"database/sql"
	_ "mysql"
	"attendSys/staffFunc"
	"attendSys/genKey"
	"os"
)
type User struct{
	uname string
	pass string
	subject string
}
func main(){
	var uname,pass string
	fmt.Println("Enter Username")
	fmt.Scanln(&uname)
	fmt.Println("Enter Password")
	fmt.Scanln(&pass)
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	results,err := db.Query("SELECT `username`,`subject` FROM staff_tbl WHERE username=? AND password=?",uname,pass)
	if err!=nil{throwError(err)}
	defer results.Close()
	var user User
	results.Next()
	results.Scan(&user.uname,&user.subject)
	if user.uname != ""{
		genKey.Setkey(user.uname)
		key := genKey.Getkey(user.uname)
		staffFunc.Open(user.subject,user.uname,key)
	}else{
		fmt.Println("\nIncorrect credentials!")
		return
	}
}
func throwError(err error){
	fmt.Println("Database error")
	os.Exit(1)
}
