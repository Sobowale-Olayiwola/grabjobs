package domain

import (
	"context"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Jobs struct {
	Title    string `json:"title" bson:"title"`
	Location GeoJSON
}

// GeoJSON is an embedded struct to be used by Jobs struct for location based
// on MongoDB documentation to be able to keep track of the coordinates and
// easily determine nearby distances around the coordinates
// check https://www.mongodb.com/docs/manual/geospatial-queries/#std-label-geospatial-geojson
// for more reference.
type GeoJSON struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type JobService interface {
	GetByFilter(ctx context.Context, filter map[string]interface{}) ([]Jobs, error)
}

type JobRepository interface {
	GetByFilter(ctx context.Context, query interface{}) ([]Jobs, error)
}
