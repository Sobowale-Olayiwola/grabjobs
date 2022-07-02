package repository

import (
	"context"
	"grabjobs/domain"

	"github.com/stretchr/testify/mock"
)

type JobRepositoryMock struct {
	mock.Mock
}

func (j *JobRepositoryMock) GetByFilter(ctx context.Context, query interface{}) ([]domain.Jobs, error) {
	output := j.Mock.Called(ctx, query)
	jobs := output.Get(0)
	err := output.Error(1)
	return jobs.([]domain.Jobs), err
}
