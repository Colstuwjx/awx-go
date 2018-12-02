package awx

import (
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
	} else {
		checkAPICallResult(t, expectListUsersResponse, result)
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
		"name":       "TestUser",
		"username":   "test",
		"password":   "test",
		"first_name": "test",
		"last_name":  "test",
		"email":      "admin@example.com",
	}, map[string]string{})

	if err != nil {
		t.Fatalf("CreateUser err: %s", err)
	} else {
		checkAPICallResult(t, expectCreateUserResponse, result)
		t.Log("CreateUser passed!")
	}
}

func TestUpdateUser(t *testing.T) {
	var (
		expectUpdateUserResponse = &User{
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
			Email:           "admin@example.net",
			IsSuperUser:     true,
			IsSystemAuditor: false,
			LdapDn:          "",
			ExternalAccount: nil,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.UserService.UpdateUser(1, map[string]interface{}{
		"email": "admin@example.net",
	}, map[string]string{})

	if err != nil {
		t.Fatalf("UpdateUser err: %s", err)
	} else {
		checkAPICallResult(t, expectUpdateUserResponse, result)
		t.Log("UpdateUser passed!")
	}
}

func TestDeleteUser(t *testing.T) {
	var (
		expectDeleteUserResponse = &User{}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.UserService.DeleteUser(1)

	if err != nil {
		t.Fatalf("DeleteUser err: %s", err)
	} else {
		checkAPICallResult(t, expectDeleteUserResponse, result)
		t.Log("DeleteUser passed!")
	}
}

func TestUserGrantRole(t *testing.T) {

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	err := awx.UserService.GrantRole(1, 24)

	if err != nil {
		t.Fatalf("TestUserGrantRole err: %s", err)
	} else {
		checkAPICallResult(t, nil, nil)
		t.Log("TestUserGrantRole passed!")
	}
}

func TestUserRevokeRole(t *testing.T) {

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	err := awx.UserService.RevokeRole(1, 24)

	if err != nil {
		t.Fatalf("TestUserRevokeRole err: %s", err)
	} else {
		checkAPICallResult(t, nil, nil)
		t.Log("TestUserRevokeRole passed!")
	}
}
