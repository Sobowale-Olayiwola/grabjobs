package repository

import (
	"context"
	"grabjobs/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoJob struct {
	db         *mongo.Database
	collection string
}

func NewMongoJobRepository(db *mongo.Database, collection string) domain.JobRepository {
	return &mongoJob{db: db, collection: collection}
}

func (j *mongoJob) GetByFilter(ctx context.Context, query interface{}) ([]domain.Jobs, error) {
	jobs := []domain.Jobs{}
	cursor, err := j.db.Collection(j.collection).Find(ctx, query)
	if err != nil {
		defer cursor.Close(ctx)
		return nil, err
	}
	if err := cursor.All(ctx, &jobs); err != nil {
		return nil, err
	}
	return jobs, nil
}
