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
