package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"umkmgubernur/internal/dto"
	"umkmgubernur/internal/handler"
	"umkmgubernur/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockAuthService struct {
	mock.Mock
}

func (m *mockAuthService) RegisterUser(req dto.RegisterRequestDTO) (*models.User, error) {
	args := m.Called(req)
	if user, ok := args.Get(0).(*models.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestRegister_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(mockAuthService)
	h := handler.NewAuthHandler(mockService)
	r.POST("/auth/register", h.Register)

	reqBody := `{"fullName":"Test User","email":"test@example.com","password":"password123"}`
	mockService.On("RegisterUser", mock.Anything).Return(map[string]string{"id": "1"}, nil)

	req, _ := http.NewRequest("POST", "/auth/register", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

// Add more tests for validation, conflict, and error cases
