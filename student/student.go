package student
import(
	"fmt"
	"database/sql"
	_"mysql"
	"os"
	"time"
)
type students struct{
	Roll_no int
	Name string
	Gender string
	Email string
}
func Addstudent(){
	var roll_no int
	var name,gender,email string
	var subjects []string
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	results,err := db.Query("SELECT subject from staff_tbl")
	if err!=nil{throwError(err)}
	for i:=0;results.Next();i++{
		var str string
		results.Scan(&str)
		subjects = append(subjects,str)
	}
	results.Close()
	fmt.Print("Enter details of student\nRollno: ")
	fmt.Scanln(&roll_no)
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	fmt.Print("Gender: ")
	fmt.Scanln(&gender)
	fmt.Print("Email Id: ")
	fmt.Scanln(&email)
	_,err = db.Query("INSERT INTO student_tbl (roll_no,name,gender,email) VALUES (?,?,?,?)",roll_no,name,gender,email)
	for _,subject:=range subjects{
		results,err := db.Query("SELECT Subject,Total_lec FROM lecture_tbl WHERE Subject=?",subject)
		if err!=nil{throwError(err)}
		results.Next()
		var subj string
		var total int
		results.Scan(&subj,&total)
		if subj!=""{
			_,err = db.Query("INSERT INTO lecture_tbl (subject,total_lec,roll_no,attended) VALUES (?,?,?,0)",subject,total,roll_no)	
		}
	}
	results.Close()
	if err!=nil{throwError(err)}
	fmt.Println("STUDENT ADDED\n")
}
func ViewStudent(){
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	results,err := db.Query("SELECT Roll_no,Name,Gender,Email FROM student_tbl")
	if err!=nil{throwError(err)}
	var arr []students
	for i:=0;results.Next();i++{
		var s students
		results.Scan(&s.Roll_no,&s.Name,&s.Gender,&s.Email)
		arr = append(arr,s)
	}
	fmt.Println("Roll_no\tName\tGender\tEmailId")
	fmt.Println("-------------------------------------------------")
	for i:=range arr{
		fmt.Printf("\n%v\t%v\t%v\t\t%v\n",arr[i].Roll_no,arr[i].Name,arr[i].Gender,arr[i].Email)
	}
	results.Close()
}
func throwError(err error){
	fmt.Println("Database error")
	os.Exit(1)
}