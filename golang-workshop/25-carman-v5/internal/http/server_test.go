package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"

	"golang.source-fellows.com/training/carman/v5/internal/mocks" 
	"golang.source-fellows.com/training/carman/v5/internal/model"
)

func TestHandleAllCar_success(t *testing.T) {
	// 1. Initialize the gomock controller
	ctrl := gomock.NewController(t)

	// 2. Create the mock repository
	mockRepo := mocks.NewMockCarRepository(ctrl)

	// 3. Configure the mock's expected behavior
	expectedCars := []model.Car{}

	// mockRepo.EXPECT().GetAllCars().Return(expectedCars, nil).Times(1)
	mockRepo.EXPECT().GetAllCars(gomock.Any()).Return(expectedCars, nil).Times(1)

	// 4. Set up the Gin router in Test Mode
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/api/cars", handleGetCars(mockRepo))

	// 5. Create a simulated HTTP request
	req, err := http.NewRequest(http.MethodGet, "/api/cars", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// 6. Create a ResponseRecorder to capture what Gin sends back
	w := httptest.NewRecorder()

	// 7. Execute the request through the Gin engine
	r.ServeHTTP(w, req)

	// 8. Assert the results
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}