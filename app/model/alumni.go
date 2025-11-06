package model

import "time"

type Alumni struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Nama      string    `bson:"nama" json:"nama"`
	Email     string    `bson:"email" json:"email"`
	NoHP      string    `bson:"no_hp" json:"no_hp"`
	Alamat    string    `bson:"alamat" json:"alamat"`
	Jurusan   string    `bson:"jurusan" json:"jurusan"`
	Angkatan  int       `bson:"angkatan" json:"angkatan"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
