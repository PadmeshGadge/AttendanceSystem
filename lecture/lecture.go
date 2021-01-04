package lecture
import(
	"fmt"
	"database/sql"
	_"mysql"
	"os"
	"time"
	"strconv"
	"io"
)
type StudentAttn struct{
	Name string
	Roll_no int
	Subject string
	Attended int
	Total_lec int
	Percentage float64
}
func Check_lec(subject string){
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	results,err := db.Query("SELECT Subject FROM lecture_tbl WHERE Subject=?",subject)
	if err!=nil{throwError(err)}
	results.Next()
	var subj string
	results.Scan(&subj)
	if subj!=""{
		Newlec(subject)
	}else{
		Firstnewlec(subject)
	}
}
func Firstnewlec(subject string){
	var Present = make(map[int]string)
	fmt.Println("FIRST NEW LECTURE")
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	_,err = db.Query("UPDATE staff_tbl SET Total_lec=Total_lec+1 WHERE Subject=?",subject)
	results,err := db.Query("SELECT Total_lec FROM staff_tbl WHERE Subject=?",subject)
	if err!=nil{throwError(err)}
	results.Next()
	var total,roll_no int
	results.Scan(&total)
	results.Close()
	//Marking Attendance
	fmt.Println("Enter student's roll_no to mark attendance")
	i:=0
	for i<1{
		_,err = fmt.Scanln(&roll_no)
		if err!=nil{
			var discard string
			fmt.Scanln(&discard)
			i++
		}else{
			Present[roll_no] = "P"
			db.Query("INSERT INTO lecture_tbl(Subject, Total_lec, Roll_no, Attended) VALUES (?,?,?,1)",subject,total,roll_no)
		}
	}
	results,err = db.Query("SELECT Roll_no FROM student_tbl WHERE Roll_no NOT IN (SELECT Roll_no FROM lecture_tbl WHERE Subject=?)",subject)
	if err!=nil{throwError(err)}
	for results.Next() {
		results.Scan(&roll_no)
		db.Query("INSERT INTO lecture_tbl(Subject, Total_lec, Roll_no, Attended) VALUES (?,?,?,0)",subject,total,roll_no)
	}
	results.Close()
	Export(Present,subject)
}
func Newlec(subject string){
	var Present = make(map[int]string)
	fmt.Println("NEW LECTURE")
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	_,err = db.Query("UPDATE staff_tbl SET Total_lec=Total_lec+1 WHERE Subject=?",subject)
	results,err := db.Query("SELECT Total_lec FROM staff_tbl WHERE Subject=?",subject)
	if err!=nil{throwError(err)}
	results.Next()
	var total,roll_no int
	results.Scan(&total)
	results.Close()
	//Marking Attendance
	fmt.Println("Enter student's roll_no to mark attendance")
	i:=0
	for i<1{
		_,err = fmt.Scanln(&roll_no)
		if err!=nil{
			var discard string
			fmt.Scanln(&discard)
			i++
		}else{
			Present[roll_no] = "P"
			db.Query("UPDATE lecture_tbl SET Total_lec=?,Attended=Attended+1 WHERE Subject=? AND Roll_no=?",total,subject,roll_no)
		}
	}
	results,err = db.Query("SELECT roll_no FROM lecture_tbl WHERE Total_lec=? AND Subject=?",total-1,subject)
	if err!=nil{throwError(err)}
	for results.Next() {
		results.Scan(&roll_no)
		db.Query("UPDATE lecture_tbl SET Total_lec=? WHERE Subject=? AND Roll_no=?",total,subject,roll_no)
	}
	results.Close()
	Export(Present,subject)
}
func Check_attn(subject string){
	s1 := []StudentAttn{}
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	results,err := db.Query("SELECT s.Roll_no,s.Name,l.Subject,Attended,Total_lec,(Attended/Total_lec)*100 FROM lecture_tbl l, student_tbl s WHERE s.Roll_no = l.Roll_no AND Subject=?",subject)
	if err!=nil{throwError(err)}
	i:=0
	for results.Next(){
		var r StudentAttn
		results.Scan(&r.Roll_no,&r.Name,&r.Subject,&r.Attended,&r.Total_lec,&r.Percentage)
		i++
		s1 = append(s1,r)
	}
	fmt.Println("RollNo\tName\tSubject\t\tAttended  Total lectures\tPercentage")
	fmt.Println("------------------------------------------------------------------------------")
	for i:=range s1{
		fmt.Printf("\n%v\t%v\t%v\t\t  %v\t\t%v\t\t%.2f\n",s1[i].Roll_no,s1[i].Name,s1[i].Subject,s1[i].Attended,s1[i].Total_lec,s1[i].Percentage)
	}
}

func Export(Present map[int]string, subject string){
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	results,err := db.Query("SELECT Name,Roll_no FROM student_tbl")
	if err!=nil{throwError(err)}
	defer results.Close()
	t := time.Now()
	str := t.Format("2006-01-02-15H04M")
	str1 := "./"+str+"-attendance.txt"
	str += "\n"+subject+"\nName\tRollNo\tAttendance\n"
	file,err := os.Create(str1)
	for results.Next(){
		var roll int
		var name string
		results.Scan(&name,&roll)
		if _,ok := Present[roll];ok{
			str += name+"\t"+strconv.Itoa(roll)+"\tP\n"
		}else{
			str += name+"\t"+strconv.Itoa(roll)+"\tA\n"
		}
	}
	io.WriteString(file,str)
	fmt.Println("Exported to file")
	file.Close()
}
func throwError(err error){
	fmt.Println("Database error")
	os.Exit(1)
}