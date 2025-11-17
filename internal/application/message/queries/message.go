package queries

import "main/internal/application/message/readmodels"

type MessageQuery interface {
	Query(msgId string) (*readmodels.Message, error)
}
