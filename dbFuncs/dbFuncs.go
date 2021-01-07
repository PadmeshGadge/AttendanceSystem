package dbFuncs
import(
	"database/sql"
	_"mysql"
	// "fmt"
)
var Db *sql.DB

func Connect() {
	var err error
	Db,err = sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{
		panic("Database error")
	}
}