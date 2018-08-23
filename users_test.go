package awx

import (
	"reflect"
	"testing"
	"time"
)

func TestListUsers(t *testing.T) {
	var (
		expectListUsersResponse = []*User{
			{
				ID:   1,
				Type: 0,
				URL:  "/api/v2/users/1/",
				Related: &Related{
					AdminOfOrganizations: "/api/v2/users/1/admin_of_organizations/",
					Organizations:        "/api/v2/users/1/organizations/",
					Roles:                "/api/v2/users/1/roles/",
					AccessList:           "/api/v2/users/1/access_list/",
					Teams:                "/api/v2/users/1/teams/",
					Credentials:          "/api/v2/users/1/credentials/",
					ActivityStream:       "/api/v2/users/1/activity_stream/",
					Projects:             "/api/v2/users/1/projects/",
				},
				SummaryFields: &Summary{
					UserCapabilities: &UserCapabilities{
						Edit:   true,
						Delete: true,
					},
				},
				Created: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2017-09-06T02:55:30.492Z")
					return t
				}(),
				Username:        "admin",
				FirstName:       "",
				LastName:        "",
				Email:           "admin@example.com",
				IsSuperUser:     true,
				IsSystemAuditor: false,
				LdapDn:          "",
				ExternalAccount: nil,
			},
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, _, err := awx.UserService.ListUsers(map[string]string{
		"name": "Demo User",
	})

	if err != nil {
		t.Fatalf("ListUsers err: %s", err)
	} else if !reflect.DeepEqual(result, expectListUsersResponse) {
		t.Logf("expected: %v", expectListUsersResponse[0])
		t.Logf("result: %v", result[0])
		t.Fatal("ListUsers resp not as expected")
	} else {
		t.Log("ListUsers passed!")
	}
}

func TestCreateUser(t *testing.T) {
	var (
		expectCreateUserResponse = &User{
			ID:   1,
			Type: 0,
			URL:  "/api/v2/users/1/",
			Related: &Related{
				AdminOfOrganizations: "/api/v2/users/1/admin_of_organizations/",
				Organizations:        "/api/v2/users/1/organizations/",
				Roles:                "/api/v2/users/1/roles/",
				AccessList:           "/api/v2/users/1/access_list/",
				Teams:                "/api/v2/users/1/teams/",
				Credentials:          "/api/v2/users/1/credentials/",
				ActivityStream:       "/api/v2/users/1/activity_stream/",
				Projects:             "/api/v2/users/1/projects/",
			},
			SummaryFields: &Summary{
				UserCapabilities: &UserCapabilities{
					Edit:   true,
					Delete: true,
				},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2017-09-06T02:55:30.492Z")
				return t
			}(),
			Username:        "admin",
			FirstName:       "",
			LastName:        "",
			Email:           "admin@example.com",
			IsSuperUser:     true,
			IsSystemAuditor: false,
			LdapDn:          "",
			ExternalAccount: nil,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.UserService.CreateUser(map[string]interface{}{
		"name":         "TestUser",
		"description":  "for testing CreateUser api",
		"organization": 1,
		"kind":         "",
		"host_filter":  "",
		"variables":    "",
	}, map[string]string{})

	if err != nil {
		t.Fatalf("CreateUser err: %s", err)
	} else if !reflect.DeepEqual(result, expectCreateUserResponse) {
		t.Logf("expected: %v", expectCreateUserResponse)
		t.Logf("result: %v", result)
		t.Fatal("CreateUser response not as expected")
	} else {
		t.Log("CreateUser passed!")
	}
}
