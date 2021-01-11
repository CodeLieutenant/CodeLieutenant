package services

import (
	"context"
	"testing"

	"github.com/malusev998/dusanmalusev/dto"
	"github.com/malusev998/dusanmalusev/tests"
	"github.com/stretchr/testify/suite"
)

type ContactServiceTest struct {
	tests.DBTestCase
}

func (c *ContactServiceTest) TestAddMessage_ValidationError() {
	contactDto := dto.Contact{
		Name:    "Test",
		Email:   "test@",
		Subject: "Test Subject",
		Message: "Test Message",
	}
	service := contactService{
		db:       c.DB,
		validate: tests.GetValidator(),
	}

	con, err := service.AddMessage(c.Ctx, contactDto)

	c.Error(err)
	c.Empty(con)

	row := c.DB.QueryRow(c.Ctx, "SELECT COUNT(id) FROM contacts;")
	var count uint64
	c.NoError(row.Scan(&count))
	c.Equal(uint64(0), count)
}

func (c *ContactServiceTest) TestAddMessage_Success() {
	contactDto := dto.Contact{
		Name:    "Test",
		Email:   "test@test.com",
		Subject: "Test Subject",
		Message: "Test Message",
	}
	service := contactService{
		db:       c.DB,
		validate: tests.GetValidator(),
	}

	con, err := service.AddMessage(c.Ctx, contactDto)

	c.NoError(err)

	c.NotEmpty(con.ID)
	c.Equal("Test", con.Name)
	c.Equal("test@test.com", con.Email)
	c.Equal("Test Subject", con.Subject)
	c.Equal("Test Message", con.Message)

	row := c.DB.QueryRow(c.Ctx, "SELECT COUNT(id) FROM contacts;")
	var count uint64
	c.NoError(row.Scan(&count))
	c.Equal(uint64(1), count)
}

func (s *ContactServiceTest) TearDownTest() {
	_, err := s.DB.Exec(s.Ctx, "DELETE FROM contacts;")
	if err != nil {
		s.FailNow("Failed cleaning up the table", err.Error())
	}
}

func TestContactService(t *testing.T) {
	t.Parallel()
	suite.Run(t, &ContactServiceTest{
		DBTestCase: tests.DBTestCase{
			Ctx:              context.Background(),
			ConnectionString: tests.GetConnectionString(),
		},
	})
}
