package repository_test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type urlRepositorySuite struct {
	suite.Suite
	ctx context.Context
}

func TestUrlRepository(t *testing.T) {
	suite.Run(t, new(urlRepositorySuite))
}
