package staffFunc
import(
	"fmt"
	"time"
	"attendSys/lecture"
	"attendSys/student"
)
func Open(subject string){
	fmt.Println("Login successful!")
	time.Sleep(1*time.Second)
	n,i := 0,0
	for i<1{
		fmt.Println("\n1.Start Lecture\t2.Check attendance\t3.Add Student\n4.Remove Student\t5.Edit Student\t6.Logout\n")
		fmt.Scanln(&n)
		switch (n) {
			case 1:lecture.Check_lec(subject)
			case 2:lecture.Check_attn(subject)
			case 3:student.Addstudent()
			case 4:student.Removestudent()
			case 5:student.EditStudent()
			default :i++
		}
	}
	time.Sleep(1*time.Second)
}