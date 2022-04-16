package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	fmt.Println("hello world")

	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/ws", handleWS)
	http.ListenAndServe(":8080", nil)

}
