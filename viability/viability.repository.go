package viability

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Database   string
	Connection *mongo.Client
}

type Point struct {
	Long float64
	Lat  float64
}

const COLLECTION = "splitters"

func (m *Repository) Find(searchDistance float64, searchPoint Point) ([]SplitterSchema, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()

	cur, err := m.Connection.Database(m.Database).Collection(COLLECTION).Find(ctx, bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": bson.A{searchPoint.Long, searchPoint.Lat},
				},
				"$maxDistance": searchDistance,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	var splitters []SplitterSchema
	for cur.Next(ctx) {
		var splitter SplitterSchema
		err := cur.Decode(&splitter)

		if err != nil {
			return nil, err
		}

		splitters = append(splitters, splitter)
	}

	return splitters, nil
}
