package main
import (
	"log"
	"encoding/gob"
	"time"
	"os"
	"fmt"
)
type user struct{
	Name string
	age	int
	CreatedAt	time.Time}
func main(){
        users:=[]user{
        user{Name:"ayoub",age:21,CreatedAt:time.Now()},
        user{Name:"hamza",age:19,CreatedAt:time.Now()},
        }
        f,err:=os.Create("user.gob")
        if err!=nil{
                log.Fatal(err)
        }
        defer f.Close()
        enc:=gob.NewEncoder(f)
        fmt.Println(enc)
	fmt.Println(f)
	fmt.Println(users)
        if err:=enc.Encode(users);err!=nil{ 
        log.Fatal(err) 
        }
	fmt.Println(err)
}