package container

import (
	"github.com/malusev998/malusev998/server/services"
	"github.com/malusev998/malusev998/server/services/contact"
	"github.com/malusev998/malusev998/server/services/subscribe"
)

//func (c *Container) GetPostService() services.PostService {
//	if c.postService == nil {
//		c.postService = posts.New(c.GetDatabasePool(), c.GetValidator())
//	}
//
//	return c.postService
//}

func (c *Container) GetContactService() services.ContactService {
	if c.contactService == nil {
		c.contactService = contact.New(c.GetContactRepository(), c.GetValidator())
	}

	return c.contactService
}

func (c *Container) GetSubscriptionService() services.SubscribeService {
	if c.contactService == nil {
		c.subscriptionService = subscribe.New(c.GetSubscriptionRepository(), c.GetValidator())
	}

	return c.subscriptionService
}


// func (c *Container) GetEmailService() email.Interface {
// 	service, err := email.NewEmailService(email.Config{
// 		Addr:     fmt.Sprintf("%s:%d", c.Config.SMTP.Host, c.Config.SMTP.Port),
// 		From:     fmt.Sprintf("%s <%s>", c.Config.SMTP.From.Name, c.Config.SMTP.From.Email),
// 		Auth:     smtp.PlainAuth("", c.Config.SMTP.Username, c.Config.SMTP.Password, c.Config.SMTP.Host),
// 		TLS:      &tls.Config{},
// 		Logger:   c.Logger,
// 		PoolSize: 0,
// 		Senders:  0,
// 	})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	return service
// }