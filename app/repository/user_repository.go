package repository

import (
	"context"
	"tugas8/app/model"
	"tugas8/database"

	"go.mongodb.org/mongo-driver/bson"
)

// FindUserByUsername mencari user berdasarkan username di koleksi "users"
func FindUserByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User
	collection := database.MongoDB.Collection("users")
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return user, err
}

// GetAllUsers -> contoh fungsi tambahan untuk listing (opsional)
func GetAllUsers(ctx context.Context) ([]model.User, error) {
	var results []model.User
	collection := database.MongoDB.Collection("users")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}
