package queries

type JobQueries interface {
	GetJobsByFilter(filter map[string]interface{}) interface{}
}
