package contact

import (
	"context"
	"errors"
	"github.com/malusev998/malusev998/server/__mocks__/repositories/contact"
	"github.com/malusev998/malusev998/server/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/tests"
)

func TestAddMessage_ValidationError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	contactDto := dto.Contact{
		Name:    "Test",
		Email:   "test@",
		Subject: "Test Subject",
		Message: "Test Message",
	}
	repo := new(contact.RepositoryMock)

	service := service{
		repo:     repo,
		validate: tests.GetValidator(),
	}

	con, err := service.AddMessage(context.Background(), contactDto)

	assert.Error(err)
	assert.Empty(con)

	//row := c.DB.QueryRow(c.Ctx, "SELECT COUNT(id) FROM contacts;")
	//var count uint64
	//c.NoError(row.Scan(&count))
	//c.Equal(uint64(0), count)
}

func TestAddMessage_InsertFails(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	contactDto := dto.Contact{
		Name:    "Test",
		Email:   "test@test.com",
		Subject: "Test Subject",
		Message: "Test Message",
	}

	repo := new(contact.RepositoryMock)

	service := service{
		repo:     repo,
		validate: tests.GetValidator(),
	}

	repo.On("Insert", mock.Anything, contactDto).
		Once().
		Return(models.Contact{}, errors.New("insert failed"))

	con, err := service.AddMessage(context.Background(), contactDto)

	assert.Error(err)
	assert.Empty(con)
}

func TestAddMessage_Success(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	contactDto := dto.Contact{
		Name:    "Test",
		Email:   "test@test.com",
		Subject: "Test Subject",
		Message: "Test Message",
	}

	repo := new(contact.RepositoryMock)

	service := service{
		repo:     repo,
		validate: tests.GetValidator(),
	}

	contactModel := models.Contact{
		Model:   models.Model{ID: 1},
		Name:    "Test",
		Email:   "test@test.com",
		Subject: "Test Subject",
		Message: "Test Message",
	}

	repo.On("Insert", mock.Anything, contactDto).
		Once().
		Return(contactModel, nil)

	con, err := service.AddMessage(context.Background(), contactDto)

	assert.NoError(err)

	assert.NotEmpty(con.ID)
	assert.Equal(contactModel.Name, con.Name)
	assert.Equal(contactModel.Email, con.Email)
	assert.Equal(contactModel.Subject, con.Subject)
	assert.Equal(contactModel.Message, con.Message)

	//row := c.DB.QueryRow(c.Ctx, "SELECT COUNT(id) FROM contacts;")
	//var count uint64
	//c.NoError(row.Scan(&count))
	//c.Equal(uint64(1), count)
	repo.AssertExpectations(t)
}

//func (s *ContactServiceTest) TearDownTest() {
//	_, err := s.DB.Exec(s.Ctx, "DELETE FROM contacts;")
//	if err != nil {
//		s.FailNow("Failed cleaning up the table", err.Error())
//	}
//}

//func TestContactService(t *testing.T) {
//	t.Parallel()
//	suite.Run(t, &ContactServiceTest{
//		DBTestCase: tests.DBTestCase{
//			Ctx:              context.Background(),
//			ConnectionString: tests.GetConnectionString(),
//		},
//	})
//}
