package projections

import "time"

type MessageProjction struct {
	ID        string    `json:"id" bson:"_id"`
	ChatID    string    `json:"chat_id" bson:"chat_id"`
	SenderID  string    `json:"sender_id" bson:"sender_id"`
	Content   string    `json:"content" bson:"content"`
	ReplyToId string    `json:"reply_to_id,omitempty" bson:"reply_to_id,omitempty"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	IsEdited  bool      `bson:"is_edited" json:"is_edited"`
	ReplyOnly bool      `bson:"reply_only" json:"reply_only"`
}
