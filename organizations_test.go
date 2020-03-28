package awx

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type OrganizationsTestSuite struct {
	suite.Suite
	testOrg *Organization
	client  *AWX
}

func TestProviders(t *testing.T) {
	suite.Run(t, &OrganizationsTestSuite{})
}

func (o *OrganizationsTestSuite) SetupTest() {
	awxUser := "admin"
	awxPassword := "password"
	c := &http.Client{}
	o.client = NewAWX("http://localhost", awxUser, awxPassword, c)
	o.T().Logf("AWX: %+v", o.client)
}

func (o *OrganizationsTestSuite) TestCreateOrganization() {
	oc := o.client.OrganizationService

	data := map[string]interface{}{"name": "Testing", "description": "A Test organization"}

	org, err := oc.CreateOrganization(data, map[string]string{})
	require.Nil(o.T(), err)
	// Make sure that the org was created by making sure an ID was set
	require.NotNil(o.T(), org.ID)
	o.testOrg = org
}

func (o *OrganizationsTestSuite) TestListOrganizations() {
	oc := o.client.OrganizationService

	org, _, err := oc.ListOrganizations(map[string]string{})
	require.Nil(o.T(), err)
	// Make sure we get a non nil list back
	require.NotEmpty(o.T(), org)
	o.T().Logf("%+v", org)
}
