package services

import "main/internal_new/domain/chat/projections"

type Publisher interface {
	NewMessage(recipientId string, collection *projections.CollectionProjection) error
}
