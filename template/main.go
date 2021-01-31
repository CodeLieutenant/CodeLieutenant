package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/malusev998/template/jet"
)

const (
	Ext = "html"

	Index   = "index"
	Contact = "contact"
	About   = "about"
)

var (
	render    bool
	templates = []string{Index, Contact, About}
)

func renderTemplates(engine *jet.Engine) {
	var wg sync.WaitGroup
	errCh := make(chan error, len(templates))
	data := fiber.Map{
		"Title": "Dusan Malusev - About me",
	}

	wg.Add(len(templates))
	for _, t := range templates {
		go func(wg *sync.WaitGroup, t string) {
			defer wg.Done()

			path := filepath.Join(".", fmt.Sprintf("%s.%s", t, Ext))

			f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
			if err != nil {
				errCh <- err
				return
			}

			defer f.Close()

			if err := engine.Render(f, t, data); err != nil {
				errCh <- err
			}
		}(&wg, t)
	}

	wg.Wait()

	close(errCh)

	for err := range errCh {
		log.Printf("Error while rendering templates to html: %v", err)
	}
}

func main() {
	flag.BoolVar(&render, "prerender", false, "Prerender Templates to HTML")
	flag.Parse()
	engine := jet.NewFileSystem(http.Dir("./views"), ".jet")
	engine.AddFunc("now", time.Now)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	if render {
		renderTemplates(engine)
		return
	}

	app.Static("/public/img", "./img/")
	app.Static("/public", "./dist/")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render(Index, fiber.Map{
			"Title": "Dusan Malusev - Index Page",
		})
	})

	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.Render(Contact, fiber.Map{
			"Title": "Dusan Malusev - Contact Page",
		})
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render(About, fiber.Map{
			"Title": "Dusan Malusev - About me",
		})
	})

	log.Fatalf("Error while starting the fiber application: %v", app.Listen(":3000"))
}
