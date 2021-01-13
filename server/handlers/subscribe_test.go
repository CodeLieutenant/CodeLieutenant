package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/malusev998/dusanmalusev/dto"
	"github.com/malusev998/dusanmalusev/handlers"
	"github.com/malusev998/dusanmalusev/models"
)

type subscribeServiceMock struct {
	mock.Mock
}

func (c *subscribeServiceMock) Subscribe(ctx context.Context, subscribeDto dto.Subscription) (models.Subscription, error) {
	args := c.Called(ctx, subscribeDto)

	return args.Get(0).(models.Subscription), args.Error(1)
}

func (c *subscribeServiceMock) Unsubscribe(ctx context.Context, id uint64) error {
	args := c.Called(ctx, id)
	return args.Error(0)
}

func TestSubscribeSuccess(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app := fiber.New()

	service := &subscribeServiceMock{}

	subscribe := handlers.Subscribe{
		Service: service,
	}

	subscribeDto := dto.Subscription{
		Name:    "Test",
		Email:   "test@test.com",
	}

	subscription := models.Subscription{
		Model: models.Model{
			ID: 1,
		},
		Name:    "Test",
		Email:   "test@test.com",
	}

	service.On("Subscribe", mock.Anything, subscribeDto).Return(subscription, nil).Times(2)

	app.Post("/subscribe", subscribe.Subscribe)

	body, _ := json.Marshal(subscription)

	t.Run("ReturnJSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/subscribe", bytes.NewReader(body))
		req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		req.Header.Add(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
		req.Header.Add(fiber.HeaderXRequestedWith, "XMLHttpRequest")
		res, err := app.Test(req)
		assert.NoError(err)
		assert.NotNil(res)

		defer res.Body.Close()
		body, _ = ioutil.ReadAll(res.Body)

		assert.Equal(http.StatusCreated, res.StatusCode)

		var contactFromResponse models.Subscription
		assert.NoError(json.Unmarshal(body, &contactFromResponse))
		assert.EqualValues(subscription, contactFromResponse)
	})

	t.Run("Redirect", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/subscribe", bytes.NewReader(body))
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

func TestSubscribeInternalError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app := fiber.New()

	service := &subscribeServiceMock{}

	subscribe := handlers.Subscribe{
		Service: service,
	}

	subscribeDto := dto.Subscription{
		Name:    "Test",
		Email:   "test@test.com",
	}

	service.On("Subscribe", mock.Anything, subscribeDto).Return(models.Subscription{}, errors.New("Server error")).Once()
	app.Post("/subscribe", subscribe.Subscribe)
	body, _ := json.Marshal(subscribeDto)
	req := httptest.NewRequest(http.MethodPost, "/subscribe", bytes.NewReader(body))
	req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	req.Header.Add(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	req.Header.Add(fiber.HeaderXRequestedWith, "XMLHttpRequest")
	res, err := app.Test(req)

	assert.NoError(err)
	assert.NotNil(res)

	assert.Equal(http.StatusInternalServerError, res.StatusCode)
}

func TestSubscribeInvalidPayload(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app := fiber.New()

	service := &subscribeServiceMock{}

	subscribe := handlers.Subscribe{
		Service: service,
	}

	app.Post("/subscribe", subscribe.Subscribe)

	req := httptest.NewRequest(http.MethodPost, "/subscribe", bytes.NewReader([]byte{1}))
	req.Header.Add(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	res, err := app.Test(req)

	assert.NoError(err)
	assert.NotNil(res)

	assert.Equal(http.StatusUnprocessableEntity, res.StatusCode)
}
