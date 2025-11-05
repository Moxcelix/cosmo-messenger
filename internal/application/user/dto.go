package user_application

var (
	DeletedSender = &Sender{
		Name: "DELETED",
	}
)

type Sender struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserReview struct {
	ID           string  `json:"id"`
	Username     string  `json:"username"`
	Name         string  `json:"name"`
	Bio          string  `json:"bio"`
	DirectChatId *string `json:"direct_chat_id,omitempty"`
}
