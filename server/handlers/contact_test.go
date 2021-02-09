package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/malusev998/template/jet"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/handlers"
	"github.com/malusev998/malusev998/server/models"
)

type contactServiceMock struct {
	mock.Mock
}

func (c *contactServiceMock) AddMessage(ctx context.Context, contactDto dto.Contact) (models.Contact, error) {
	args := c.Called(ctx, contactDto)

	return args.Get(0).(models.Contact), args.Error(1)
}

func TestContactPage(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	abs, _ := filepath.Abs("../views")

	engine := jet.NewFileSystem(http.Dir(abs), ".jet")
	engine.Debug(true)
	engine.AddFunc("now", time.Now)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	home := handlers.Contact{}
	app.Get("/contact", home.Index)

	req := httptest.NewRequest(http.MethodGet, "/contact", nil)

	res, err := app.Test(req)

	assert.NoError(err)
	assert.NotNil(res)

	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	str := string(bytes)
	assert.Contains(str, "Dusan Malusev - Contact")
	year := strconv.Itoa(time.Now().Year())
	assert.Contains(str, year)
	assert.Equal(http.StatusOK, res.StatusCode)
}

func TestMessageSuccess(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app := fiber.New()

	service := &contactServiceMock{}

	contact := handlers.Contact{
		Service: service,
	}

	contactDto := dto.Contact{
		Name:    "Test",
		Email:   "test@test.com",
		Subject: "Test Subject",
		Message: "Test Message",
	}

	ct := models.Contact{
		Model: models.Model{
			ID: 1,
		},
		Name:    "Test",
		Email:   "test@test.com",
		Subject: "Test Subject",
		Message: "Test Message",
	}

	service.On("AddMessage", mock.Anything, contactDto).Return(ct, nil).Times(2)

	app.Post("/contact", contact.Message)

	body, _ := json.Marshal(ct)

	t.Run("ReturnJSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/contact", bytes.NewReader(body))
		req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		req.Header.Add(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
		req.Header.Add(fiber.HeaderXRequestedWith, "XMLHttpRequest")
		res, err := app.Test(req)
		assert.NoError(err)
		assert.NotNil(res)

		defer res.Body.Close()
		body, _ = ioutil.ReadAll(res.Body)

		assert.Equal(http.StatusCreated, res.StatusCode)

		var contactFromResponse models.Contact
		assert.NoError(json.Unmarshal(body, &contactFromResponse))
		assert.EqualValues(ct, contactFromResponse)
	})

	t.Run("Redirect", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/contact", bytes.NewReader(body))
		req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		req.Header.Add(fiber.HeaderAccept, fiber.MIMETextHTML)
		res, err := app.Test(req)
		assert.NoError(err)
		assert.NotNil(res)

		defer res.Body.Close()
		body, _ = ioutil.ReadAll(res.Body)
		assert.Equal(http.StatusFound, res.StatusCode)
	})
	service.AssertExpectations(t)
}

func TestMessageInternalError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app := fiber.New()

	service := &contactServiceMock{}

	contact := handlers.Contact{
		Service: service,
	}

	contactDto := dto.Contact{
		Name:    "Test",
		Email:   "test@test.com",
		Subject: "Test Subject",
		Message: "Test Message",
	}

	service.On("AddMessage", mock.Anything, contactDto).Return(models.Contact{}, errors.New("Server error")).Once()
	app.Post("/contact", contact.Message)
	body, _ := json.Marshal(contactDto)
	req := httptest.NewRequest(http.MethodPost, "/contact", bytes.NewReader(body))
	req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	req.Header.Add(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	req.Header.Add(fiber.HeaderXRequestedWith, "XMLHttpRequest")
	res, err := app.Test(req)

	assert.NoError(err)
	assert.NotNil(res)

	assert.Equal(http.StatusInternalServerError, res.StatusCode)
}

func TestMessageInvalidPayload(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app := fiber.New()

	service := &contactServiceMock{}

	contact := handlers.Contact{
		Service: service,
	}

	app.Post("/contact", contact.Message)

	req := httptest.NewRequest(http.MethodPost, "/contact", bytes.NewReader([]byte{1}))
	req.Header.Add(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	res, err := app.Test(req)

	assert.NoError(err)
	assert.NotNil(res)

	assert.Equal(http.StatusUnprocessableEntity, res.StatusCode)
}
