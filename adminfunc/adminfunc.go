package adminfunc
import(
	"fmt"
	"time"
	"database/sql"
	_"mysql"
	"os"
)
type staff struct{
	Name string
	Subject string
	Total_lec int
}
func Open(){
	fmt.Println("Login successful!")
	time.Sleep(1*time.Second)
	n,i := 0,0
	for i<1{
		fmt.Println("\n1.Staff details\t2.Add Staff\n3.Remove Staff\t4.Logout\n")
		fmt.Scanln(&n)
		switch (n) {
			case 1:Getstaff()
			case 2:Addstaff()
			case 3:Removestaff()
			case 4:i++
		}
	}
	time.Sleep(1*time.Second)
}
func Getstaff(){
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	results,err := db.Query("SELECT Name,Subject,Total_lec FROM staff_tbl")
	if err!=nil{throwError(err)}
	defer results.Close()
	var arrstaff []staff
	for i:=0;results.Next();i++{
		var s staff
		results.Scan(&s.Name,&s.Subject,&s.Total_lec)
		arrstaff = append(arrstaff,s)
	}
	fmt.Println("Name\t\tSubject\t\tLectures")
	fmt.Println("-------------------------------------------------")
	for i:=range arrstaff{
		fmt.Printf("\n%v\t%v\t\t%v\n",arrstaff[i].Name,arrstaff[i].Subject,arrstaff[i].Total_lec)
	}
}
func Addstaff(){
	var uname,pass,name,subj string
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	fmt.Print("Enter username: ")
	fmt.Scanln(&uname)
	fmt.Print("Enter password: ")
	fmt.Scanln(&pass)
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter subject: ")
	fmt.Scanln(&subj)
	_,err = db.Query("INSERT INTO staff_tbl (username,password,Name,subject,Total_lec) VALUES (?,?,?,?,0)",uname,pass,name,subj)
	if err!=nil{throwError(err)}
	fmt.Println("STAFF ADDED\n")
	Getstaff()
}
func Removestaff(){
	var uname string
	fmt.Println("Enter username to be removed")
	fmt.Scanln(&uname)
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	_,err = db.Query("DELETE FROM staff_tbl WHERE username=?",uname)
	if err!=nil{throwError(err)}
	fmt.Println("Staff user",uname,"deleted")
}
func throwError(err error){
	fmt.Println("Database error")
	os.Exit(1)
}