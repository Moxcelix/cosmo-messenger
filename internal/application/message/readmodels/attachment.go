package readmodels

import "time"

type Attachment struct {
	ID        string
	Type      string
	URL       string
	Filename  string
	Size      int
	MimeType  string
	CreatedAt time.Time
}
