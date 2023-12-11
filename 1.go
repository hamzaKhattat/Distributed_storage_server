package main

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"net/http"
)

func main() {
	// var err error
	server, err := socketio.NewServer(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	server.On("connection", func(s socketio.Socket) {
		fmt.Println("Welcome in New Connection")
	})

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.Handle("/socket.io/", server)

	fmt.Println(http.ListenAndServe(":5000", nil))
}

/*package main
import (
	"fmt"
	socketio"github.com/googollee/go-socket.io"
	"net/http"
)
func main(){
//	var err error
	server ,err:=socketio.NewServer(nil)
	if err!=nil{
		fmt.Println(err)
	}
	server.On("connection",func(s socketio.Socket){
		fmt.Println("Welcome in New Connection")
	})
	fs:=http.FileServer(http.Dir("."))
	http.Handle("/",fs)
	http.Handle("/socket.io/",server)
	fmt.Println(http.ListenAndServe(":5000",nil))
}*/
