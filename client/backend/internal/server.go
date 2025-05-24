package server

type Message struct {
	Room   string
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

var ReceivedMessages = make(chan Message) // Буфер для сообщений

func SendMessage(message Message) {
	ReceivedMessages <- message // Отправляем сообщение в канал
}
