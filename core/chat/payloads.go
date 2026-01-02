package chat

import "time"

type WebSocketMessage struct {
	Type string `json:"Type"`
	Data interface{} `json:"Data"`
}

type ChatMessageData struct {
	FullName string `json:"FullName"`
	Message string `json:"Message"`
	SentAt time.Time `json:"SentAt"`
}
