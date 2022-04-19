package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Room map[string]ClientSet

type ClientSet map[*websocket.Conn]bool

func main() {
	var room = make(Room)

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		paths := strings.Split(r.URL.Path, "/")
		var roomHash = paths[len(paths)-1]
		client, ok := room[roomHash]
		if !ok {
			log.Println("new client")
			client = make(ClientSet)
			room[roomHash] = client
		}

		if len(room[roomHash]) >= 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		handleWS(client, w, r)
	})

	log.Println("Websocket server started...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("server starting error")
	}
}
