package user

import (
	"errors"
	"fmt"
	"go-tutorial/db"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetHandler(t *testing.T) {
	database, err := db.Connect()
	assert.Nil(t, err)
	repo := NewRepository(database)
	service := NewService(repo)
	handler := NewHandler(service)

	app := fiber.New()
	app.Get("/users/:id", handler.Get)

	id, err := repo.Create(Model{Name: "test", Email: "test@mail.com"})
	assert.Nil(t, err)
	req := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", id), nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

func TestGetHandler_ReturnsErrorFromService(t *testing.T) {
	database, err := db.Connect()
	assert.Nil(t, err)
	repo := NewRepository(database)
	service := &MockService{}
	service.On("Get", mock.AnythingOfType("uint")).Return(nil, errors.New("Bir hata oluştu"))
	handler := NewHandler(service)

	app := fiber.New()
	app.Get("/users/:id", handler.Get)

	id, err := repo.Create(Model{Name: "test", Email: "test@mail.com"})
	assert.Nil(t, err)
	req := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", id), nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	body, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	fmt.Println(string(body))
	assert.Equal(t, 400, res.StatusCode)
}
