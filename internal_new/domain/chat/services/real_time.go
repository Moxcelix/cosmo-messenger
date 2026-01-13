package services

import (
	"main/internal_new/domain/chat/factories"
	"main/internal_new/domain/chat/models"
)

type RealTimeService struct {
	publisher Publisher
	factory   *factories.CollectionProjectionFactory
}

func NewRealTimeService(
	publisher Publisher,
	factory *factories.CollectionProjectionFactory,
) *RealTimeService {
	return &RealTimeService{
		publisher: publisher,
		factory:   factory,
	}
}

func (s *RealTimeService) NewMessage(recipientsId []string, collection *models.Collection) error {
	for _, recipientId := range recipientsId {
		collection, err := s.factory.ProjectCollection(recipientId, collection)
		if err != nil {
			return err
		}

		if err = s.publisher.NewMessage(recipientId, collection); err != nil {
			return err
		}
	}

	return nil
}
