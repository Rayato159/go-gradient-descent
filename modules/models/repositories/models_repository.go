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

func (mr *modelsRep) GetData(ctx context.Context, getType string, ratio float64) ([]entities.Data, error) {
	ctx = context.WithValue(ctx, entities.ModelsRep, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsRep).(int64)))

	count, err := mr.Db.Collection(getType).CountDocuments(ctx, bson.D{})
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("error, can't count data collection with an error: %v", err.Error())
	}

	var groupStage bson.D
	switch getType {
	case "train_data":
		groupStage = bson.D{{
			"$sample", bson.D{{
				"size", int64(float64(count) * float64(ratio)),
			}},
		}}
	case "test_data":
		groupStage = bson.D{{
			"$sample", bson.D{{
				"size", int64(float64(count) * float64(1-ratio)),
			}},
		}}
	default:
		groupStage = bson.D{{
			"$sample", bson.D{{
				"size", count,
			}},
		}}
	}

	cursor, err := mr.Db.Collection(getType).Aggregate(ctx, mongo.Pipeline{groupStage})
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("error, can't aggregate data collection with an error: %v", err.Error())
	}

	data := make([]entities.Data, 0)
	if err = cursor.All(ctx, &data); err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("error, can't cursor a data with an error: %v", err.Error())
	}
	return data, nil
}
