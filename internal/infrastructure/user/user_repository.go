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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
	logger     pkg.Logger
}

func NewUserRepository(db pkg.MongoDatabase, logger pkg.Logger) userservice.UserRepository {
	_, err := db.Collection("users").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		logger.Error("failed to create unique index on username", err)
	}

	return &UserRepository{
		collection: db.Collection("users"),
		logger:     logger,
	}
}

func (r *UserRepository) GetUserById(userId string) (*userservice.User, error) {
	var schema User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": userId}).Decode(&schema)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	user := mapSchemaToDomain(schema)
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*userservice.User, error) {
	var schema User
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&schema)
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

func (r *UserRepository) DeleteUserByUsername(username string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"username": username})
	return err
}

func (r *UserRepository) DeleteUserById(userId string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": userId})
	return err
}

func (r *UserRepository) UpdateUser(user *userservice.User) error {
	user.UpdatedAt = time.Now()
	update := bson.M{
		"$set": bson.M{
			"name":          user.Name,
			"username":      user.Username,
			"password_hash": user.PasswordHash,
			"bio":           user.Bio,
			"updated_at":    user.UpdatedAt,
		},
	}
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, update)
	return err
}

func (r *UserRepository) GetUsersByRange(offset, limit int) (*userservice.UsersList, error) {
	total, err := r.collection.CountDocuments(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	opts := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{"created_at", -1}})

	cursor, err := r.collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var schemas []User
	if err := cursor.All(context.Background(), &schemas); err != nil {
		return nil, err
	}

	var users []*userservice.User
	for _, schema := range schemas {
		user := mapSchemaToDomain(schema)
		users = append(users, &user)
	}

	return &userservice.UsersList{
		Users:  users,
		Total:  int(total),
		Offset: offset,
		Limit:  limit,
	}, nil
}

func mapSchemaToDomain(schema User) userservice.User {
	return userservice.User{
		ID:           schema.ID,
		Name:         schema.Name,
		Username:     schema.Username,
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
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		Bio:          user.Bio,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
