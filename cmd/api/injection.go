package main

import (
	"grabjobs/internal/constants"
	"grabjobs/internal/middleware"
	_mongoJobsRepo "grabjobs/jobs/repository/mongo"
	_mongoJobsQueries "grabjobs/jobs/repository/queries"
	"net/http"

	_jobsService "grabjobs/jobs/service"

	_jobsHandler "grabjobs/jobs/handler/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func inject(d *DataSources) *gin.Engine {
	/*
	 * repository layer
	 */
	mongoJobRepo := _mongoJobsRepo.NewMongoJobRepository(d.DB, constants.JobsCollection)

	/*
	 * service layer
	 */
	jobService := _jobsService.NewJobService(mongoJobRepo, _mongoJobsQueries.MongoQuery{})

	router := gin.Default()
	router.Use(middleware.LoggerToFile())
	router.Use(cors.Default())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to grabjobs"})
	})
	/*
	 * handler layer
	 */
	_jobsHandler.NewJobHandler(router, jobService)
	return router
}
