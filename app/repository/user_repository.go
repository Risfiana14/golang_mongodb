package repository

import (
	"context"
	"tugas8/app/model"
	"tugas8/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TIDAK ADA GLOBAL VARIABLE LAGI! ← INI YANG PENTING!

var userCollection *mongo.Collection // kita buat nanti setelah koneksi

// Inisialisasi collection setelah koneksi berhasil
func InitCollections() {
	userCollection = database.MongoDB.Collection("users")
}

// FindUserByUsername
func FindUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers (opsional)
func GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}