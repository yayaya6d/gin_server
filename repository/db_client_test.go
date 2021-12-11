package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbClientTestSuite struct {
	suite.Suite
}

func TestDBClientSuite(t *testing.T) {
	suite.Run(t, new(dbClientTestSuite))
}

func (s *dbClientTestSuite) TestDBCLient_InitDBClient_ReturnNoError() {
	// arrange
	ctx := context.Background()
	opt := options.Client().ApplyURI("mongodb://localhost:27017")

	// act
	err := InitDBClient(ctx, *opt)
	defer DisconnectDBClient()

	// assert
	s.NoError(err)
}

func (s *dbClientTestSuite) TestDBCLient_GetDBClientAfterInit_ReturnNoError() {
	// arrange
	_ = InitDBClient(context.Background(), *options.Client().ApplyURI("mongodb://localhost:27017"))
	defer DisconnectDBClient()

	// act
	dbClient, err := GetDBClient()

	// assert
	s.NotNil(dbClient)
	s.NoError(err)
}

func (s *dbClientTestSuite) TestDBCLient_DisconnectDBClientAfterInit_ReturnNoError() {
	// arrange
	_ = InitDBClient(context.Background(), *options.Client().ApplyURI("mongodb://localhost:27017"))
	defer DisconnectDBClient()

	// act
	err := DisconnectDBClient()

	// assert
	s.NoError(err)
}

func (s *dbClientTestSuite) TestDBCLient_DisconnectDBClientBeforeInit_ReturnNoError() {
	// act
	err := DisconnectDBClient()

	// assert
	s.NoError(err)
}
