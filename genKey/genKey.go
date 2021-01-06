package genKey
import(
	"math/rand"
	"time"
)
const l = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var keymap = make(map[string]string)

func Setkey(user string){
	rand.Seed(time.Now().UnixNano())
	b:=make([]byte,10)
	for i:=range b{
		b[i]=l[rand.Intn(len(l))]
	}
	keymap[user] = string(b)
}

func Getkey(user string)string{
	return keymap[user]
}

func Removekey(user string){
	delete(keymap,user)
}