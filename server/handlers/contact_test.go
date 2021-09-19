package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/malusev998/malusev998/server/__mocks__/services/contact"
	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/handlers"
	"github.com/malusev998/malusev998/server/models"
)

func TestMessage(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app := fiber.New()

	service := new(contact.ServiceMock)

	contactHandler := handlers.Contact{
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

	app.Post("/contactHandler", contactHandler.Message)

	body, _ := json.Marshal(ct)

	t.Run("ReturnJSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/contactHandler", bytes.NewReader(body))
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
		req := httptest.NewRequest(http.MethodPost, "/contactHandler?redirect=/contactHandler", bytes.NewReader(body))
		req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		req.Header.Add(fiber.HeaderAccept, fiber.MIMETextHTML)
		res, err := app.Test(req)
		assert.NoError(err)
		assert.NotNil(res)

		defer res.Body.Close()
		body, _ = ioutil.ReadAll(res.Body)
		assert.Equal(fiber.StatusTemporaryRedirect, res.StatusCode)
		url, err := res.Location()

		assert.NoError(err)
		assert.Equal(url.String(), "/contactHandler?message=Message+sent&status=success")
	})

	service.AssertExpectations(t)
}

func TestMessageInternalError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app := fiber.New()

	service := new(contact.ServiceMock)

	contactHandler := handlers.Contact{
		Service: service,
	}

	contactDto := dto.Contact{
		Name:    "Test",
		Email:   "test@test.com",
		Subject: "Test Subject",
		Message: "Test Message",
	}

	service.On("AddMessage", mock.Anything, contactDto).
		Return(models.Contact{}, errors.New("server error")).
		Once()

	app.Post("/contactHandler", contactHandler.Message)
	body, _ := json.Marshal(contactDto)
	req := httptest.NewRequest(http.MethodPost, "/contactHandler", bytes.NewReader(body))
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

	service := new(contact.ServiceMock)

	contactHandler := handlers.Contact{
		Service: service,
	}

	app.Post("/contactHandler", contactHandler.Message)

	req := httptest.NewRequest(http.MethodPost, "/contactHandler", bytes.NewReader([]byte{1}))
	req.Header.Add(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	res, err := app.Test(req)

	assert.NoError(err)
	assert.NotNil(res)

	assert.Equal(http.StatusUnprocessableEntity, res.StatusCode)
}
