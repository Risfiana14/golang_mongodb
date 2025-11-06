package repository

import (
	"context"
	"time"
	"tugas8/app/model"
	"tugas8/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllPekerjaan ambil semua pekerjaan
func GetAllPekerjaan(ctx context.Context) ([]model.Pekerjaan, error) {
	var results []model.Pekerjaan
	collection := database.MongoDB.Collection("pekerjaan")
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

func GetPekerjaanByID(ctx context.Context, id string) (*model.Pekerjaan, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var p model.Pekerjaan
	collection := database.MongoDB.Collection("pekerjaan")
	if err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func GetPekerjaanByAlumniID(ctx context.Context, alumniID string) ([]model.Pekerjaan, error) {
	var results []model.Pekerjaan
	collection := database.MongoDB.Collection("pekerjaan")

	// Coba konversi alumniID ke ObjectID
	objID, err := primitive.ObjectIDFromHex(alumniID)
	filter := bson.M{"$or": []bson.M{
		{"alumni_id": objID},
		{"alumni_id": alumniID},
	}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}



func CreatePekerjaan(ctx context.Context, p model.Pekerjaan) (*model.Pekerjaan, error) {
	p.ID = primitive.NewObjectID()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	_, err := database.MongoDB.Collection("pekerjaan").InsertOne(ctx, p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func UpdatePekerjaan(ctx context.Context, id string, p model.Pekerjaan) (*model.Pekerjaan, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	p.UpdatedAt = time.Now()
	_, err = database.MongoDB.Collection("pekerjaan").UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": p})
	if err != nil {
		return nil, err
	}
	return GetPekerjaanByID(ctx, id)
}

func DeletePekerjaan(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = database.MongoDB.Collection("pekerjaan").DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
