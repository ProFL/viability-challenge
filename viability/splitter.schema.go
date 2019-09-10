package viability

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PointSchema struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Type        string             `json:"type" bson:"type"`
	Coordinates []float64          `json:"coordinates" bson:"coordinates"`
}

type SplitterSchema struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Location PointSchema        `json:"location" bson:"location"`
}
