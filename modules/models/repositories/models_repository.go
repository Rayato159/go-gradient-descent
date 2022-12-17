package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (mr *modelsRep) InsertTrainResult(ctx context.Context, req *entities.TrainRes) error {
	ctx = context.WithValue(ctx, entities.ModelsRep, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsRep).(int64)))

	coll := mr.Db.Collection("records")
	req.Timestamp = time.Now()
	if _, err := coll.InsertOne(ctx, req); err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("error, can't insert a record")
	}
	return nil
}

func (mr *modelsRep) ClearData(ctx context.Context) error {
	ctx = context.WithValue(ctx, entities.ModelsRep, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsRep).(int64)))

	collTrain := mr.Db.Collection("train_data")
	collTest := mr.Db.Collection("test_data")

	if _, err := collTrain.DeleteMany(ctx, bson.D{}); err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("error, can't delete a train_data")
	}
	if _, err := collTest.DeleteMany(ctx, bson.D{}); err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("error, can't delete a test_data")
	}
	return nil
}

func (mr *modelsRep) ClearRecord(ctx context.Context) error {
	ctx = context.WithValue(ctx, entities.ModelsRep, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsRep).(int64)))

	coll := mr.Db.Collection("records")
	if _, err := coll.DeleteMany(ctx, bson.D{}); err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("error, can't delete a record")
	}
	return nil
}

func (mr *modelsRep) GetWeights(ctx context.Context) ([]float64, error) {
	ctx = context.WithValue(ctx, entities.ModelsRep, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsRep).(int64)))

	weights := make([]float64, 0)
	coll := mr.Db.Collection("records")

	docs := &entities.TrainRes{
		Weights: make([]float64, 0),
	}
	if err := coll.FindOne(
		ctx,
		bson.D{},
		options.FindOne().SetSort(
			bson.D{{"timestamp", -1}},
		),
	).Decode(&docs); err != nil {
		return nil, err
	}

	weights = append(weights, docs.Weights...)
	return weights, nil
}
