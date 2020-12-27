package services

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ContactServiceTest struct {
	suite.Suite
}

func TestContactService(t *testing.T) {
	t.Parallel()
	suite.Run(t, &ContactServiceTest{})
}
