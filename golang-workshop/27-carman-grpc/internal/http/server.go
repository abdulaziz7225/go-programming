package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	carman "golang.source-fellows.com/training/carman/v6/internal"
	"golang.source-fellows.com/training/carman/v6/internal/model"
)

//go:generate go tool mockgen -source=../repository.go -destination=../mocks/repository_mock.go -package=mocks

func StartServer(repository carman.CarRepository) error {
	r := gin.Default()

	r.GET("/ping", handlePing)

	authorized := r.Group("/api", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
	}))

	authorized.GET("/cars", handleGetCars(repository))
	authorized.POST("/audi", handleAddNewCar(repository))
	return r.Run()
}

func handlePing(context *gin.Context) {
	context.Data(http.StatusOK, "text/plain", []byte("pong"))
}

func handleGetCars(repository carman.CarRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		xtraceId := c.GetHeader("x-trace-id")
		ctx := context.WithValue(c.Request.Context(), "x-trace-id", xtraceId)

		// ctx := c.Request.Context()
		cars, err := repository.GetAllCars(ctx)
		
		if err != nil {
			// Replace log.Println(err) with logrus:
			logrus.WithContext(ctx).WithError(err).Error("Failed to fetch cars from repository")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Negotiate(http.StatusOK, gin.Negotiate{
			Offered: []string{gin.MIMEJSON, gin.MIMEXML},
			Data:    cars,
		})
	}
}

func handleAddNewCar(repository carman.CarRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		xtraceId := c.GetHeader("x-trace-id")
		ctx := context.WithValue(c.Request.Context(), "x-trace-id", xtraceId)

		var newAudi model.Audi

		if err := c.ShouldBindJSON(&newAudi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repository.AddCar(ctx, &newAudi); err != nil {
			logrus.WithContext(ctx).WithError(err).Error("Failed to save new car")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save car"})
			return
		}

		c.JSON(http.StatusCreated, newAudi)
	}
}
