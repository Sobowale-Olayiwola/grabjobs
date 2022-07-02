package service

import (
	"context"
	"grabjobs/domain"
	"grabjobs/jobs/repository/queries"
)

type jobService struct {
	jobRepository domain.JobRepository
	dbQueries     queries.JobQueries
}

func NewJobService(j domain.JobRepository, q queries.JobQueries) domain.JobService {
	return &jobService{jobRepository: j, dbQueries: q}
}

func (j *jobService) GetByFilter(ctx context.Context, filter map[string]interface{}) ([]domain.Jobs, error) {
	query := j.dbQueries.GetJobsByFilter(filter)
	jobs, err := j.jobRepository.GetByFilter(ctx, query)
	if err != nil {
		return nil, err
	}
	if len(jobs) == 0 {
		return nil, domain.ErrRecordNotFound
	}
	return jobs, nil
}
