package main
import(
	"fmt"
	"attendSys/staffFunc"
	"attendSys/genKey"
	"attendSys/dbFuncs"
)
type User struct{
	uname string
	pass string
	subject string
}

func main(){
	var user User
	var uname,pass string
	fmt.Println("Enter Username")
	fmt.Scanln(&uname)
	fmt.Println("Enter Password")
	fmt.Scanln(&pass)

	dbFuncs.Connect()
	defer dbFuncs.Db.Close()
	results,err := dbFuncs.Db.Query("SELECT `username`,`subject` FROM staff_tbl WHERE username=? AND password=?",uname,pass)
	if err!=nil{
		throwError(err)
	}else{
		results.Next()
		results.Scan(&user.uname,&user.subject)
		results.Close()
	}

	if user.uname != ""{
		g := genKey.Store
		g.Setkey(user.uname)
		key := g.Getkey(user.uname)
		staffFunc.Open(user.subject,user.uname,key)
	}else{
		fmt.Println("\nIncorrect credentials!")
		return
	}
}
func throwError(err error){
	defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
	}()
	panic(err)
}