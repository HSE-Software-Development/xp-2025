package cli

import (
	"fmt"
	"strings"

	"github.com/HSE-Software-Development/xp-2025/client/backend/internal/server"
	"github.com/HSE-Software-Development/xp-2025/client/backend/internal/utils"
)

type Client struct {
	name string
	room string
}

var client *Client

func RunCLI() {
	var name, room string
	fmt.Print("Введите ваше имя: ")
	fmt.Scanln(&name)
	fmt.Print("Введите название комнаты: ")
	fmt.Scanln(&room)

	client = &Client{
		name: name,
		room: room,
	}

	go handleMessages()

	server.Join(room)
	sendMessage(room, "Система", fmt.Sprintf("%s присоединился чату", name))

	var input string
	for {
		fmt.Scanln(&input)
		if input[0] == '!' {
			words := strings.Fields(input)
			if len(words) == 2 && words[0] == "!switch" {
				sendMessage(client.room, "Система", fmt.Sprintf("%s покинул чат", name))
				client.room = words[1]
				server.Join(words[1])
				sendMessage(client.room, "Система", fmt.Sprintf("%s присоединился чату", name))
			} else if len(words) == 1 && words[0] == "!exit" {
				return
			}
		} else {
			sendMessage(client.room, client.name, input)
		}

	}
}

func sendMessage(room, sender, text string) {
	msg := utils.Message{
		Room:   room,
		Sender: sender,
		Text:   text,
	}
	server.SendMessage(msg)
}

func handleMessages() {
	for {
		msg := <-server.ReceivedMessages

		if msg.Room == client.room {
			fmt.Printf("<%s>: %s\n", msg.Sender, msg.Text)
		}
	}
}
