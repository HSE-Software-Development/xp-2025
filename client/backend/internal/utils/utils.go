package utils

type Message struct {
	Room   string `json:"room"`
	Sender string `json:"sender"`
	Text   string `json:"text"`
}
