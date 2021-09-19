package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/malusev998/malusev998/server/__mocks__/services/subscribe"
	"github.com/malusev998/malusev998/server/__mocks__/utils"
	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/handlers"
	"github.com/malusev998/malusev998/server/models"
)

const (
	BaseUrl             = "http://localhost:4000"
	FrontendRedirectUrl = "http://localhost:4000/redirect"
)

func TestUnsubscribe_InternalServerError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app := fiber.New()

	subscriptionService := new(subscribe.ServiceMock)
	signer := new(utils.SignerMock)

	subscribeModel := handlers.Subscribe{
		FrontendRedirectURL: FrontendRedirectUrl,
		BaseURL:             BaseUrl,
		SubscriptionService: subscriptionService,
		Signer:              signer,
	}

	app.Get("/unsubscribe", subscribeModel.Unsubscribe)

	url := "/unsubscribe?id=1"
	signer.On("Verify", fmt.Sprintf("%s%s", BaseUrl, url)).
		Once().
		Return(nil)

	subscriptionService.On("Unsubscribe",
		mock.MatchedBy(func(ctx context.Context) bool { return true }),
		uint64(1),
	).Once().
		Return(errors.New("invalid query"))

	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Add(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	req.Header.Add(fiber.HeaderXRequestedWith, "XMLHttpRequest")
	res, err := app.Test(req)
	assert.NoError(err)
	assert.NotNil(res)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	bodyStr := string(body)

	assert.Equal(fiber.StatusInternalServerError, res.StatusCode)
	assert.Equal("An error has occurred", bodyStr)
	subscriptionService.AssertExpectations(t)
	signer.AssertExpectations(t)

}

func TestUnsubscribe_UnsubscribeSuccess(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app := fiber.New()

	subscriptionService := new(subscribe.ServiceMock)
	signer := new(utils.SignerMock)

	subscribeModel := handlers.Subscribe{
		FrontendRedirectURL: FrontendRedirectUrl,
		BaseURL:             BaseUrl,
		SubscriptionService: subscriptionService,
		Signer:              signer,
	}

	app.Get("/unsubscribe", subscribeModel.Unsubscribe)

	url := "/unsubscribe?id=1"
	signer.On("Verify", fmt.Sprintf("%s%s", BaseUrl, url)).
		Once().
		Return(nil)

	subscriptionService.On("Unsubscribe",
		mock.MatchedBy(func(ctx context.Context) bool { return true }),
		uint64(1)).
		Once().
		Return(nil)

	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Add(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	req.Header.Add(fiber.HeaderXRequestedWith, "XMLHttpRequest")
	res, err := app.Test(req)
	assert.NoError(err)
	assert.NotNil(res)

	defer res.Body.Close()

	redirect, err := res.Location()
	assert.NoError(err)

	assert.Equal(fiber.StatusTemporaryRedirect, res.StatusCode)
	assert.Equal(FrontendRedirectUrl+"?message=Successfully+unsubscribed&status=success", redirect.String())
	subscriptionService.AssertExpectations(t)
	signer.AssertExpectations(t)
}

func TestUnsubscribe_InvalidSignature(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app := fiber.New()

	subscriptionService := new(subscribe.ServiceMock)
	signer := new(utils.SignerMock)

	subscribeModel := handlers.Subscribe{
		FrontendRedirectURL: FrontendRedirectUrl,
		BaseURL:             BaseUrl,
		SubscriptionService: subscriptionService,
		Signer:              signer,
	}

	app.Get("/unsubscribe", subscribeModel.Unsubscribe)

	signer.On("Verify", fmt.Sprintf("%s/unsubscribe", BaseUrl)).
		Once().
		Return(utils.ErrInvalidUrl)

	req := httptest.NewRequest(http.MethodGet, "/unsubscribe", nil)
	req.Header.Add(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	req.Header.Add(fiber.HeaderXRequestedWith, "XMLHttpRequest")
	res, err := app.Test(req)
	assert.NoError(err)
	assert.NotNil(res)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	bodyStr := string(body)

	assert.Equal(fiber.StatusBadRequest, res.StatusCode)
	assert.Equal("Invalid URL", bodyStr)

	subscriptionService.AssertExpectations(t)
	signer.AssertExpectations(t)
}

func TestSubscribeSuccess(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app := fiber.New()

	service := new(subscribe.ServiceMock)

	subscribeModel := handlers.Subscribe{
		SubscriptionService: service,
	}

	subscribeDto := dto.Subscription{
		Name:  "Test",
		Email: "test@test.com",
	}

	subscription := models.Subscription{
		Model: models.Model{
			ID: 1,
		},
		Name:  "Test",
		Email: "test@test.com",
	}

	service.On("Subscribe", mock.Anything, subscribeDto).Return(subscription, nil).Times(2)

	app.Post("/subscribeModel", subscribeModel.Subscribe)

	body, _ := json.Marshal(subscription)

	t.Run("ReturnJSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/subscribeModel", bytes.NewReader(body))
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
		req := httptest.NewRequest(http.MethodPost, "/subscribeModel?redirect=/subscribeModel", bytes.NewReader(body))
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
		assert.Equal(url.String(), "/subscribeModel?message=Message+sent&status=success")
	})
	service.AssertExpectations(t)
}

func TestSubscribeInternalError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app := fiber.New()

	service := new(subscribe.ServiceMock)

	subscribeModel := handlers.Subscribe{
		SubscriptionService: service,
	}

	subscribeDto := dto.Subscription{
		Name:  "Test",
		Email: "test@test.com",
	}

	service.On("Subscribe", mock.Anything, subscribeDto).Return(models.Subscription{}, errors.New("Server error")).Once()
	app.Post("/subscribeModel", subscribeModel.Subscribe)
	body, _ := json.Marshal(subscribeDto)
	req := httptest.NewRequest(http.MethodPost, "/subscribeModel", bytes.NewReader(body))
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

	service := new(subscribe.ServiceMock)

	subscribeModel := handlers.Subscribe{
		SubscriptionService: service,
	}

	app.Post("/subscribeModel", subscribeModel.Subscribe)

	req := httptest.NewRequest(http.MethodPost, "/subscribeModel", bytes.NewReader([]byte{1}))
	req.Header.Add(fiber.HeaderAccept, fiber.MIMEApplicationJSON)
	res, err := app.Test(req)

	assert.NoError(err)
	assert.NotNil(res)

	assert.Equal(http.StatusUnprocessableEntity, res.StatusCode)
}
