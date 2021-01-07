package genKey
import(
	"math/rand"
	"time"
	"sync"
)
const l = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type keyStore struct {
	m 	sync.RWMutex
	keymap 	map[string]string
}

var Store *keyStore;

func init () {
	Store = &keyStore{keymap:map[string]string{}}
}
func (s *keyStore)Setkey(user string){
	s.m.Lock()
	defer s.m.Unlock()
	rand.Seed(time.Now().UnixNano())
	b:=make([]byte,10)
	for i:=range b{
		b[i]=l[rand.Intn(len(l))]
	}
	s.keymap[user] = string(b)
}

func (s *keyStore) Getkey(user string)string{
	s.m.RLock()
	defer s.m.RUnlock()
	return s.keymap[user]
}

func (s *keyStore)Removekey(user string){
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.keymap,user)
}