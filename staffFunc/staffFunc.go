package staffFunc
import(
	"fmt"
	"attendSys/lecture"
	"attendSys/genKey"
	"attendSys/student"
)
func Open(subject,uname,key string){
	fmt.Println("Login successful!")
	n,i := 0,0
	for i<1{
		fmt.Println("\n1.Start Lecture\t2.Check attendance\t3.Edit student details\t4.Logout\n")
		fmt.Scanln(&n)
		switch (n) {
			case 1:lecture.Check_lec(subject,uname,key)
			case 2:lecture.Check_attn(subject,uname,key)
			case 3:student.EditStudent(uname,key)
			default :{
				g := genKey.Store
				g.Removekey(uname)
				i++
			}
		}
	}
}