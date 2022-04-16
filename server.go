package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Event struct {
	Type string                 `json:"type"`
	Args map[string]interface{} `json:"args"`
}

func readLoop(c *websocket.Conn) {
	log.Println("start readLoop")
	for {
		var event Event
		err := c.ReadJSON(&event)

		if err != nil {
			log.Println("connection is closed")
			log.Println(err)
			c.Close()
			break

		}

		log.Println(event)
		log.Println(event.Type)
		log.Println(event.Args)
		log.Println(event.Args["turn"])

		c.WriteMessage(0, []byte("hello world"))

		log.Println("end forLoop")
	}

	log.Println("end readLoop")
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Websocket connected successfully")

	go readLoop(conn)
}
