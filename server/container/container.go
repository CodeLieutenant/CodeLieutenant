package container

import (
	"context"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"

	"github.com/malusev998/malusev998/server/config"
	"github.com/malusev998/malusev998/server/repositories"
	"github.com/malusev998/malusev998/server/services"
	"github.com/malusev998/malusev998/server/utils"
)

type Container struct {
	Ctx    context.Context
	Logger zerolog.Logger
	DB     *pgxpool.Pool
	Config *config.Config

	subscriptionRepository repositories.Subscribe
	postRepository         repositories.Post
	contactRepository      repositories.Contact

	contactService      services.ContactService
	postService         services.PostService
	subscriptionService services.SubscribeService

	validator  *validator.Validate
	translator ut.Translator
	session    *session.Store
	urlSigner  utils.URLSigner
}


func (c *Container) Close() error {
	c.DB.Close()

	return nil
}
