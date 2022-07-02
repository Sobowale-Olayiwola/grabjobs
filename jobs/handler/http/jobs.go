package http

import (
	"context"
	"errors"
	"grabjobs/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	JobAndLocationService domain.JobService
}

func NewJobHandler(router *gin.Engine, j domain.JobService) {
	handler := &JobHandler{
		JobAndLocationService: j,
	}
	api := router.Group("/api/v1")
	api.GET("/jobs-locations/near-by", handler.GetJobByFilter)
}

func (j *JobHandler) GetJobByFilter(c *gin.Context) {
	lat, err := strconv.ParseFloat(c.Query("lat"), 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	lng, err := strconv.ParseFloat(c.Query("lng"), 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	radius, err := strconv.ParseInt(c.Query("radius"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	// Converting Kilometres to metres
	radius = radius * 1000
	title := c.Query("title")
	filter := make(map[string]interface{})
	filter["radius"] = radius
	filter["title"] = title
	filter["lng"] = lng
	filter["lat"] = lat
	ctx := context.TODO()
	jobs, err := j.JobAndLocationService.GetByFilter(ctx, filter)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrRecordNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusFound, gin.H{"payload": jobs})
}
