package mappers

import (
	"main/internal/application/chat/dto"
	"main/internal/application/chat/readmodels"
	user_dto "main/internal/application/user/dto"
)

type ChatCollectionMapper struct{}

func NewChatCollectionMapper() *ChatCollectionMapper {
	return &ChatCollectionMapper{}
}

func (m *ChatCollectionMapper) MapToDTO(
	readModelList *readmodels.ChatList,
	chatNames map[string]string,
) *dto.ChatCollection {

	chatItems := make([]*dto.ChatItem, len(readModelList.Chats))

	for i, chatReadModel := range readModelList.Chats {
		chatItems[i] = m.mapChatToDTO(chatReadModel, chatNames[chatReadModel.ID])
	}

	return &dto.ChatCollection{
		Chats: chatItems,
		Meta: &dto.ScrollingMeta{
			HasPrev: readModelList.HasPrev,
			HasNext: readModelList.HasNext,
		},
	}
}

func (m *ChatCollectionMapper) mapChatToDTO(
	chatReadModel *readmodels.ChatWithLastMessage,
	chatName string,
) *dto.ChatItem {
	return &dto.ChatItem{
		ID:          chatReadModel.ID,
		Name:        chatName,
		Type:        chatReadModel.Type,
		LastMessage: m.mapLastMessageToDTO(chatReadModel.LastMessage),
	}
}

func (m *ChatCollectionMapper) mapLastMessageToDTO(lastMessage *readmodels.LastMessage) *dto.LastMessage {
	if lastMessage == nil {
		return nil
	}

	return &dto.LastMessage{
		ID:        lastMessage.ID,
		Content:   lastMessage.Content,
		Timestamp: lastMessage.SentAt,
		Sender:    m.mapSenderToDTO(lastMessage.Sender),
	}
}

func (m *ChatCollectionMapper) mapSenderToDTO(sender *readmodels.LastMessageSender) *user_dto.Sender {
	if sender == nil {
		return nil
	}

	return &user_dto.Sender{
		ID:   sender.ID,
		Name: sender.Name,
	}
}
