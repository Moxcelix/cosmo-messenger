package projections

import "time"

type MessageDemo struct {
	ID        string    `json:"id" bson:"_id"`
	ChatID    string    `json:"chat_id" bson:"chat_id"`
	Content   string    `json:"content" bson:"content"`
	IsReply   bool      `json:"is_reply" bson:"is_reply"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	SenderId  string    `json:"sender_id" bson:"sender_id"`
}
