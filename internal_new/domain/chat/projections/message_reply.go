package projections

import "time"

type MessageReply struct {
	ID        string    `json:"id" bson:"_id"`
	ChatID    string    `json:"chat_id" bson:"chat_id"`
	SenderID  string    `json:"sender_id" bson:"sender_id"`
	Content   string    `json:"content" bson:"content"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
}
