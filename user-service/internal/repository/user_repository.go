package repository

import (
	"context"
	"errors"
	"time"

	"user-service/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrEmailTaken   = errors.New("email already taken")
	ErrInvalidCreds = errors.New("invalid email or/or password")
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(client *mongo.Client, dbName, collName string) *UserRepository {
	coll := client.Database(dbName).Collection(collName)
	// уникальный индекс по email
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return &UserRepository{coll: coll}
}

func (r *UserRepository) Create(ctx context.Context, u *model.User) (string, error) {
	res, err := r.coll.InsertOne(ctx, u)
	if err != nil {
		// проверка на дублирование
		if we, ok := err.(mongo.WriteException); ok {
			for _, writeErr := range we.WriteErrors {
				if writeErr.Code == 11000 {
					return "", ErrEmailTaken
				}
			}
		}
		return "", err
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func (r *UserRepository) FindByEmailAndPassword(ctx context.Context, email, password string) (*model.User, error) {
	var u model.User
	err := r.coll.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&u)
	if err == mongo.ErrNoDocuments {
		return nil, ErrInvalidCreds
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindByID(ctx context.Context, hexID string) (*model.User, error) {
	oid, err := primitive.ObjectIDFromHex(hexID)
	if err != nil {
		return nil, ErrUserNotFound
	}
	var u model.User
	err = r.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&u)
	if err == mongo.ErrNoDocuments {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}
