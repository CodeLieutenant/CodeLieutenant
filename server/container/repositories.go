package container

import (
	"github.com/malusev998/malusev998/server/repositories"
	"github.com/malusev998/malusev998/server/repositories/contact"
	"github.com/malusev998/malusev998/server/repositories/posts"
	"github.com/malusev998/malusev998/server/repositories/subscribe"
)

func (c *Container) GetSubscriptionRepository() repositories.Subscribe {
	if c.subscriptionRepository == nil {
		c.subscriptionRepository = subscribe.New(c.GetDatabasePool())
	}

	return c.subscriptionRepository
}

func (c *Container) GetPostRepository() repositories.Post {
	if c.postRepository == nil {
		c.postRepository = posts.New(c.GetDatabasePool())
	}

	return c.postRepository
}

func (c *Container) GetContactRepository() repositories.Contact {
	if c.contactRepository == nil {
		c.contactRepository = contact.New(c.GetDatabasePool())
	}

	return c.contactRepository
}
