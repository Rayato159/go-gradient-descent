package repositories

import "go.mongodb.org/mongo-driver/mongo"

type modelsRep struct {
	Db *mongo.Database
}

func NewModelsRepsository(db *mongo.Database) *modelsRep {
	return &modelsRep{
		Db: db,
	}
}
