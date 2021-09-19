package utils

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

var ErrInvalidUrl = errors.New("signature is invalid")

type SignerMock struct {
	mock.Mock
}

func (signer *SignerMock) Sign(url string, args ...interface{}) (string, error) {
	ar := signer.Called(url, args)

	return ar.String(0), ar.Error(1)
}

func (signer *SignerMock) Verify(url string) error {
	args := signer.Called(url)

	return args.Error(0)
}
