package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ClientSet map[*websocket.Conn]bool

func main() {
	var clientSet = make(ClientSet)

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWS(clientSet, w, r)
	})

	log.Println("Websocket server started...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("server starting error")
	}
}
