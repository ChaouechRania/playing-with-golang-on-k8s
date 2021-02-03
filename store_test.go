// +build integration

package store_test

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/suite"

	"github.com/figarocms/fcms-cv-completion-api/app/resumesHistory"
	hstore "github.com/figarocms/fcms-cv-completion-api/app/resumesHistory"
	"github.com/figarocms/fcms-cv-completion-api/pkg/store"
)

//go:generate mockery -name=Store -output ../internal/mocks -case=underscore

type storeSuite struct {
	suite.Suite
	store.Client
}

func (s *storeSuite) SetupSuite() {
	var err error
	s.Client, err = store.NewConfig().NewClient()
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *storeSuite) SetupTest() {
	s.Require().NoError(
		s.Client.DB().AutoMigrate(resumesHistory.Models()...).Error,
	)
}

func (s *storeSuite) TearDownTest() {
	s.Require().NoError(
		s.Client.DB().DropTableIfExists(resumesHistory.Models()...).Error,
	)
}

func TestStoreSuite(t *testing.T) {
	suite.Run(t, new(storeSuite))
}

func (s *storeSuite) TestStore_ListResumes() {
	st := hstore.NewStore(s.Client)
	list, err := st.ListResumes()
	s.NoError(err)
	s.Empty(list)
}
