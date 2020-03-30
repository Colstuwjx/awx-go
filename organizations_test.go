package awx

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type OrganizationsTestSuite struct {
	suite.Suite
	client  *AWX
	service *OrganizationService
}

func TestProviders(t *testing.T) {
	suite.Run(t, &OrganizationsTestSuite{})
}

func (suite *OrganizationsTestSuite) SetupAllSuite() {
	rand.Seed(time.Now().UnixNano())
}

func (suite *OrganizationsTestSuite) SetupTest() {
	suite.client = NewAWX("http://localhost", "admin", "password", &http.Client{})
	suite.service = suite.client.OrganizationService
	suite.T().Logf("AWX: %+v", suite.client)
}

func (suite *OrganizationsTestSuite) TestCreateDeleteOrganization() {
	name := fmt.Sprintf("Testing-%v", rand.Int())
	org, err := suite.service.CreateOrganization(map[string]interface{}{
		"name": name,
		"description": "A Test organization",
	}, map[string]string{})
	suite.Nil(err)
	suite.NotNil(org.ID) // Make sure that the org was created by making sure an ID was set
	suite.Equal(org.Name, name)
	org, err = suite.client.OrganizationService.DeleteOrganization(org.ID)
	suite.Nil(err)
}

func (suite *OrganizationsTestSuite) TestUpdateOrganization() {
	description := fmt.Sprintf("This is a test %v", time.Now().String())
	org, err := suite.service.UpdateOrganization(1, map[string]interface{}{
		"description": description,
	}, map[string]string{})
	suite.Nil(err)
	suite.Equal(org.Description, description)
}

func (suite *OrganizationsTestSuite) TestListOrganizations() {
	org, _, err := suite.service.ListOrganizations(map[string]string{})
	suite.Nil(err)
	suite.NotEmpty(org)
	suite.Equal(org[0].Type, "organization")
	suite.Equal(org[0].Name, "Default")
	suite.T().Logf("%+v", org)
}
