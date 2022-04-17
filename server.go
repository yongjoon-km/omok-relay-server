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

func readLoop(c *websocket.Conn, clientSet ClientSet) {
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

		for client := range clientSet {
			client.WriteJSON(event)
		}
	}
	log.Println("readLoop is closed")
}

func handleWS(client ClientSet, w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client[conn] = true

	log.Println("Websocket connected successfully connection is")

	go readLoop(conn, client)
}
