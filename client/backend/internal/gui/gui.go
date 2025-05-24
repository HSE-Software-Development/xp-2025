package gui

import (
	"log"
	"net/http"

	"github.com/HSE-Software-Development/xp-2025/client/backend/internal/server"
	"github.com/HSE-Software-Development/xp-2025/client/backend/internal/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // В продакшене нужно ограничить домены!
	},
}

type Client struct {
	conn *websocket.Conn
	room string
	name string
}

type Request struct {
	Type   string `json:"type"` // join, message, leave
	Room   string `json:"room"`
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

var clients = make(map[*Client]bool)

func RunServer() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	log.Println("Сервер запущен на :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка сервера: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	client := &Client{conn: ws}
	clients[client] = true

	for {
		var req Request
		err := ws.ReadJSON(&req)
		if err != nil {
			log.Printf("Ошибка: %v", err)
			break
		}

		switch req.Type {
		case "join":
			client.room = req.Room
			client.name = req.Sender
			msg := utils.Message{
				Room:   req.Room,
				Sender: "Система",
				Text:   req.Sender + " присоединился к чату",
			}
			server.Join(req.Room)
			server.SendMessage(msg)
		case "message":
			msg := utils.Message{
				Room:   req.Room,
				Sender: req.Sender,
				Text:   req.Text,
			}
			server.Join(req.Room)
			server.SendMessage(msg)
		}
	}
}

func handleMessages() {
	for {
		msg := <-server.ReceivedMessages

		for client := range clients {
			if client.room == msg.Room {
				err := client.conn.WriteJSON(msg)
				if err != nil {
					log.Printf("Ошибка: %v", err)
					client.conn.Close()
					delete(clients, client)
				}
			}
		}
	}
}
