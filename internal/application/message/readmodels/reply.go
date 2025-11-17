package readmodels

import user_readmodels "main/internal/application/user/readmodels"

type Reply struct {
	ID      string
	Content string
	Sender  *user_readmodels.Sender
}
