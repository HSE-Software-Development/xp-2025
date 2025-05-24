package server

import (
	"github.com/HSE-Software-Development/xp-2025/client/backend/internal/manager"
	"github.com/HSE-Software-Development/xp-2025/client/backend/internal/utils"
)

var ReceivedMessages = make(chan utils.Message) // Буфер для сообщений

var Manager, err = manager.New([]string{"127.0.0.1:9092"})

func SendMessage(message utils.Message) error {
	ReceivedMessages <- message // Отправляем сообщение в канал ????? проверить надо ли
	return Manager.Send(message)
}

func Join(room string) error {
	Manager.CreateTopic(room)
	return Manager.Subscribe(room, ReceivedMessages)
}
