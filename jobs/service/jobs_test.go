package service

import (
	"context"
	"errors"
	"grabjobs/domain"
	"grabjobs/domain/mocks/repository"
	"grabjobs/jobs/repository/queries"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByFilter(t *testing.T) {
	as := assert.New(t)
	jobRepo := &repository.JobRepositoryMock{}
	t.Run("happy path: Succesfully get jobs", func(t *testing.T) {
		filter := make(map[string]interface{})
		filter["radius"] = int64(1000)
		filter["title"] = "Accountant"
		filter["lng"] = float64(103.851959)
		filter["lat"] = float64(1.290270)
		expectedResult := []domain.Jobs{
			{
				Title: "ACCOUNTANT ASSISTANT",
			},
			{
				Title: "Financial Accountant Iron Ore with a global MNC (7 to 11 yrs required)",
			},
		}
		jobRepo.On("GetByFilter", mock.Anything, mock.Anything).Return(expectedResult, nil).Once()
		jobService := NewJobService(jobRepo, queries.MongoQuery{})
		jobs, err := jobService.GetByFilter(context.Background(), filter)
		as.NoError(err)
		as.Equal(len(jobs), 2)
		jobRepo.AssertExpectations(t)
	})

	t.Run("error: record not found ", func(t *testing.T) {
		filter := make(map[string]interface{})
		filter["radius"] = int64(1000)
		filter["title"] = "Accountantssss"
		filter["lng"] = float64(103.851959)
		filter["lat"] = float64(1.290270)
		jobRepo.On("GetByFilter", mock.Anything, mock.Anything).Return([]domain.Jobs{}, nil).Once()
		jobService := NewJobService(jobRepo, queries.MongoQuery{})
		jobs, err := jobService.GetByFilter(context.Background(), filter)
		as.Error(err)
		as.Nil(jobs)
		jobRepo.AssertExpectations(t)
	})

	t.Run("system error: Database failed ", func(t *testing.T) {
		filter := make(map[string]interface{})
		filter["radius"] = int64(1000)
		filter["title"] = "Accountant"
		filter["lng"] = float64(103.851959)
		filter["lat"] = float64(1.290270)
		jobRepo.On("GetByFilter", mock.Anything, mock.Anything).Return([]domain.Jobs{}, errors.New("something failed")).Once()
		jobService := NewJobService(jobRepo, queries.MongoQuery{})
		jobs, err := jobService.GetByFilter(context.Background(), filter)
		as.Error(err)
		as.Nil(jobs)
		jobRepo.AssertExpectations(t)
	})
}
