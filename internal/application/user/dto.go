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
