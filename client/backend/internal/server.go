package server

import (
	"github.com/HSE-Software-Development/xp-2025/internal/manager"
	"github.com/HSE-Software-Development/xp-2025/internal/utils"
)



var ReceivedMessages = make(chan utils.Message) // Буфер для сообщений

var Manager, err = manager.New([]string{"127.0.0.1:9092"})

func SendMessage(message utils.Message) error{
	ReceivedMessages <- message // Отправляем сообщение в канал ????? проверить надо ли
	return Manager.Send(message)
}
func JoinTo(room string) error {
	return Manager.Subscribe(room, ReceivedMessages)
}

//только создает, не подключается
func Create(room string) error {
	return Manager.CreateTopic(room)
}

func CreateAndJoin(room string) error {
	err := Manager.CreateTopic(room)
	if err != nil {
		return nil
	}
	return Manager.Subscribe(room, ReceivedMessages)
}
