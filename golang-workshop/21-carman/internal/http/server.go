package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	carman "golang.source-fellows.com/training/carman/internal"
	"golang.source-fellows.com/training/carman/internal/model"
)

func StartServer(repository carman.CarRepository) error {
	r := gin.Default()
	r.GET("/ping", handlePing)
	r.GET("/api/cars", handleGetCars(repository))
	r.POST("/api/audi", handleAddNewCar(repository))
	return r.Run()
}

func handlePing(context *gin.Context) {
	context.Data(http.StatusOK, "text/plain", []byte("pong"))
}

func handleGetCars(repository carman.CarRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		cars, err := repository.GetAllCars()
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(200, cars)
	}
}

func handleAddNewCar(repository carman.CarRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newAudi model.Audi

		if err := c.ShouldBindJSON(&newAudi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repository.AddCar(&newAudi); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save car"})
			return
		}

		c.JSON(http.StatusCreated, newAudi)
	}
}
