package message_infrastructure

import (
	"context"
	"errors"
	"time"

	message_domain "main/internal/domain/message"
	"main/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageRepository struct {
	collection *mongo.Collection
	logger     pkg.Logger
}

func NewMessageRepository(db pkg.MongoDatabase, logger pkg.Logger) message_domain.MessageRepository {
	return &MessageRepository{
		collection: db.Collection("messages"),
		logger:     logger,
	}
}

func (r *MessageRepository) CreateMessage(message *message_domain.Message) error {
	if message.ID == "" {
		message.ID = primitive.NewObjectID().Hex()
	}
	now := time.Now()
	message.CreatedAt = now
	message.UpdatedAt = now
	_, err := r.collection.InsertOne(context.Background(), message)
	return err
}
func (r *MessageRepository) GetMessageById(id string) (*message_domain.Message, error) {
	var message message_domain.Message
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&message)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &message, nil
}

func (r *MessageRepository) UpdateMessage(message *message_domain.Message) error {
	message.UpdatedAt = time.Now()
	update := bson.M{
		"$set": bson.M{
			"reply_to":   message.ReplyTo,
			"content":    message.Content,
			"updated_at": message.UpdatedAt,
		},
	}
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": message.ID}, update)
	return err
}

func (r *MessageRepository) DeleteMessage(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *MessageRepository) GetMessagesByChatId(
	chatId string, offset, limit int) (*message_domain.MessageList, error) {

	total, err := r.collection.CountDocuments(context.Background(), bson.M{
		"chat_id": chatId,
	})
	if err != nil {
		return nil, err
	}

	opts := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(limit)).
		SetSort(bson.D{
			{Key: "created_at", Value: -1},
		})

	cursor, err := r.collection.Find(context.Background(), bson.M{
		"chat_id": chatId,
	}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var messages []*message_domain.Message
	if err := cursor.All(context.Background(), &messages); err != nil {
		return nil, err
	}

	return &message_domain.MessageList{
		Messages: messages,
		Total:    int(total),
		Offset:   offset,
		Limit:    limit,
	}, nil
}
