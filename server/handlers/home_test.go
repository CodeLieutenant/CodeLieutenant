package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/stretchr/testify/require"

	"github.com/malusev998/malusev998/handlers"
)

func TestHome(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	engine := html.New("../views", ".html")
	engine.AddFunc("now", time.Now)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	home := handlers.Home{}
	app.Get("/", home.Home)

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	res, err := app.Test(req)

	assert.NoError(err)
	assert.NotNil(res)

	defer res.Body.Close()

	assert.Equal(http.StatusOK, res.StatusCode)
	bytes, _ := ioutil.ReadAll(res.Body)
	str := string(bytes)
	assert.Contains(str, "Dusan Malusev - Home")
	year := strconv.Itoa(time.Now().Year())
	assert.Contains(str, year)
}
