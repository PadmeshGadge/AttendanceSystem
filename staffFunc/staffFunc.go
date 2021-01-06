package staffFunc
import(
	"fmt"
	"time"
	"attendSys/lecture"
	"attendSys/genKey"
)
func Open(subject,uname,key string){
	fmt.Println("Login successful!")
	time.Sleep(1*time.Second)
	n,i := 0,0
	for i<1{
		fmt.Println("\n1.Start Lecture\t2.Check attendance\t\t4.Logout\n")
		fmt.Scanln(&n)
		switch (n) {
			case 1:lecture.Check_lec(subject,uname,key)
			case 2:lecture.Check_attn(subject,uname,key)
			default :{
				genKey.Removekey(uname)
				i++
			}
		}
	}
	time.Sleep(1*time.Second)
}