package chat_infrastructure

import (
	"context"
	"errors"
	"time"

	chat_domain "main/internal/domain/chat"
	"main/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChatRepository struct {
	collection *mongo.Collection
	logger     pkg.Logger
}

func NewChatRepository(db pkg.MongoDatabase, logger pkg.Logger) chat_domain.ChatRepository {
	return &ChatRepository{
		collection: db.Collection("chats"),
		logger:     logger,
	}
}

func (r *ChatRepository) Create(chat *chat_domain.Chat) error {
	if chat.ID == "" {
		chat.ID = primitive.NewObjectID().Hex()
	}
	now := time.Now()
	chat.CreatedAt = now
	chat.UpdatedAt = now

	_, err := r.collection.InsertOne(context.Background(), chat)
	return err
}

func (r *ChatRepository) GetByID(id string) (*chat_domain.Chat, error) {
	var chat chat_domain.Chat
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&chat)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &chat, nil
}

func (r *ChatRepository) GetByMember(userID string, offset, limit int) ([]*chat_domain.Chat, error) {
	filter := bson.M{"members": userID}

	opts := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(limit)).
		SetSort(bson.D{
			{Key: "updated_at", Value: -1},
		})

	cursor, err := r.collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var chats []*chat_domain.Chat
	if err := cursor.All(context.Background(), &chats); err != nil {
		return nil, err
	}

	return chats, nil
}

func (r *ChatRepository) Update(chat *chat_domain.Chat) error {
	chat.UpdatedAt = time.Now()
	update := bson.M{
		"$set": bson.M{
			"name":       chat.Name,
			"members":    chat.Members,
			"updated_at": chat.UpdatedAt,
		},
	}
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": chat.ID}, update)
	return err
}

func (r *ChatRepository) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
