package userservice_infrastructure

import (
	"context"
	"errors"
	"time"

	"main/internal/domain/userservice"
	"main/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
	logger     pkg.Logger
}

func NewUserRepository(db pkg.MongoDatabase, logger pkg.Logger) userservice.UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
		logger:     logger,
	}
}

func (r *UserRepository) GetUser(userId string) (*userservice.User, error) {
	objID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	var schema User
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&schema)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	user := mapSchemaToDomain(schema)
	return &user, nil
}

func (r *UserRepository) CreateUser(user *userservice.User) error {
	if user.ID == "" {
		user.ID = primitive.NewObjectID().Hex()
	}
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	doc := mapDomainToSchema(*user)
	_, err := r.collection.InsertOne(context.Background(), doc)
	return err
}

func (r *UserRepository) DeleteUser(userId string) error {
	objID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	return err
}

func (r *UserRepository) UpdateUser(user *userservice.User) error {
	objID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	user.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"name":          user.Name,
			"password_hash": user.PasswordHash,
			"bio":           user.Bio,
			"updated_at":    user.UpdatedAt,
		},
	}

	_, err = r.collection.UpdateByID(context.Background(), objID, update)
	return err
}

func mapSchemaToDomain(schema User) userservice.User {
	return userservice.User{
		ID:           schema.ID,
		Name:         schema.Name,
		PasswordHash: schema.PasswordHash,
		Bio:          schema.Bio,
		CreatedAt:    schema.CreatedAt,
		UpdatedAt:    schema.UpdatedAt,
	}
}

func mapDomainToSchema(user userservice.User) User {
	return User{
		ID:           user.ID,
		Name:         user.Name,
		PasswordHash: user.PasswordHash,
		Bio:          user.Bio,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
