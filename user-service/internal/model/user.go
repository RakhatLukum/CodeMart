package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User — модель для MongoDB
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}
