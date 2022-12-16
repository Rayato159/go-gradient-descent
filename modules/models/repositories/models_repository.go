package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"www.github.com/Rayato159/go-gradient-descent/modules/entities"
	"www.github.com/Rayato159/go-gradient-descent/pkg/utils"
)

type modelsRep struct {
	Db *mongo.Database
}

func NewModelsRepository(db *mongo.Database) *modelsRep {
	return &modelsRep{
		Db: db,
	}
}

func (mr *modelsRep) GetData(ctx context.Context, getTypeQuery string) ([]entities.Data, error) {
	ctx = context.WithValue(ctx, entities.ModelsRep, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsRep).(int64)))

	cursor, err := mr.Db.Collection("train_data").Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("error, can't aggregate data collection with an error: %v", err.Error())
	}
	data := make([]entities.Data, 0)
	if err = cursor.All(ctx, &data); err != nil {
		panic(err)
	}
	return data, nil
}
