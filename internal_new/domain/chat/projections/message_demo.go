package projections

import "time"

type MessageDemo struct {
	ID       string    `json:"id" bson:"_id"`
	Content  string    `json:"content" bson:"content"`
	IsReply  bool      `json:"is_reply" bson:"is_reply"`
	SentAt   time.Time `json:"sent_at" bson:"sent_at"`
	SenderId string    `json:"sender_id" bson:"sender_id"`
}
