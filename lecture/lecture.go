package lecture
import(
	"fmt"
	"database/sql"
	_"mysql"
	"os"
	"attendSys/genKey"
)
type StudentAttn struct{
	Name string
	Roll_no int
	Subject string
	Attended int
	Total_lec int
	Percentage float64
}
func Check_lec(subject,uname,key string){
	fmt.Println(key)
	if key == genKey.Getkey(uname){
		db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
		if err!=nil{throwError(err)}
		defer db.Close()
		results,err := db.Query("SELECT Subject FROM lecture_tbl WHERE Subject=?",subject)
		if err!=nil{throwError(err)}
		results.Next()
		var subj string
		results.Scan(&subj)
		if subj==""{
			Firstnewlec(subject,uname,key)
		}else{
			Newlec(subject,uname,key)
		}
	}else{fmt.Println("Wrong Key")}
}
func Firstnewlec(subject,uname,key string){
	if key == genKey.Getkey(uname){
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
				db.Query("INSERT INTO lecture_tbl(Subject, Total_lec, Roll_no, Attended) VALUES (?,?,?,1)",subject,total,roll_no)
			}
		}
		//If students were not present for first lecture
		results,err = db.Query("SELECT Roll_no FROM student_tbl WHERE Roll_no NOT IN (SELECT Roll_no FROM lecture_tbl WHERE Subject=?)",subject)
		if err!=nil{throwError(err)}
		for results.Next() {
			results.Scan(&roll_no)
			db.Query("INSERT INTO lecture_tbl(Subject, Total_lec, Roll_no, Attended) VALUES (?,?,?,0)",subject,total,roll_no)
		}
		results.Close()
	}else{fmt.Println("Wrong Key")}
}


func Newlec(subject,uname,key string){
	if key == genKey.Getkey(uname){
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
				db.Query("UPDATE lecture_tbl SET Total_lec=?,Attended=Attended+1 WHERE Subject=? AND Roll_no=?",total,subject,roll_no)
			}
		}
		//If students were not present for lecture
		results,err = db.Query("SELECT roll_no FROM lecture_tbl WHERE Total_lec=? AND Subject=?",total-1,subject)
		if err!=nil{throwError(err)}
		for results.Next() {
			results.Scan(&roll_no)
			db.Query("UPDATE lecture_tbl SET Total_lec=? WHERE Subject=? AND Roll_no=?",total,subject,roll_no)
		}
		results.Close()
	}else{fmt.Println("Wrong Key")}
}

func Check_attn(subject,uname,key string){
	fmt.Println(key)
	if key == genKey.Getkey(uname){
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
	}else{fmt.Println("Wrong Key")}
}

func throwError(err error){
	fmt.Println("Database error")
	os.Exit(1)
}