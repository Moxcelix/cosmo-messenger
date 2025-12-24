package models

import "time"

type Attachment struct {
	ID        string    `json:"id" bson:"_id"`
	Type      string    `json:"type" bson:"type"`
	URL       string    `json:"url" bson:"url"`
	Filename  string    `json:"filename,omitempty" bson:"filename,omitempty"`
	Size      int64     `json:"size,omitempty" bson:"size,omitempty"`
	MimeType  string    `json:"mime_type,omitempty" bson:"mime_type,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}
