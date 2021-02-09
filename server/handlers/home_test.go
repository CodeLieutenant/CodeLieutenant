package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/malusev998/template/jet"
	"github.com/stretchr/testify/require"

	"github.com/malusev998/malusev998/server/handlers"
)

func TestHome(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	abs, _ := filepath.Abs("../views")

	engine := jet.NewFileSystem(http.Dir(abs), ".jet")
	engine.Debug(true)
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

	bytes, _ := ioutil.ReadAll(res.Body)
	str := string(bytes)
	assert.Contains(str, "Dusan Malusev - Home")
	year := strconv.Itoa(time.Now().Year())
	assert.Contains(str, year)
	assert.Equal(http.StatusOK, res.StatusCode)
}
