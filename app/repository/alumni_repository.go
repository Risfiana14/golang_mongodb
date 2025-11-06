package repository

import (
	"context"
	"time"
	"tugas8/app/model"
	"tugas8/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllAlumni ambil semua alumni
func GetAllAlumni(ctx context.Context) ([]model.Alumni, error) {
	var alumni []model.Alumni
	collection := database.MongoDB.Collection("alumni")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &alumni); err != nil {
		return nil, err
	}
	return alumni, nil
}

func GetAlumniByID(ctx context.Context, id string) (*model.Alumni, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var a model.Alumni
	collection := database.MongoDB.Collection("alumni")
	if err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&a); err != nil {
		return nil, err
	}
	return &a, nil
}

func CreateAlumni(ctx context.Context, a model.Alumni) error {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	_, err := database.MongoDB.Collection("alumni").InsertOne(ctx, a)
	return err
}

func UpdateAlumni(ctx context.Context, id string, a model.Alumni) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	a.UpdatedAt = time.Now()
	_, err = database.MongoDB.Collection("alumni").UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": a})
	return err
}

func DeleteAlumni(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = database.MongoDB.Collection("alumni").DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
