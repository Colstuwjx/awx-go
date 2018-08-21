package awx

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func TestListUsers(t *testing.T) {
	var (
		expectListUserResponse = []*User{
			{
				ID:   1,
				Type: "inventory",
				URL:  "/api/v2/users/1/",
				Related: &Related{
                    Organization: "/api/v1/users/1/organizations/",
                    Credentials: "/api/v1/users/1/credentials/",
                    Permissions: "/api/v1/users/1/permissions/",
                    ActivityStream: "/api/v1/users/1/activity_stream/",
                    Projects: "/api/v1/users/1/projects/",
                    Teams: "/api/v1/users/1/teams/"
				},
				
				Created: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-05-21T01:34:35.657185Z")
					return t
				}(),

				Modified: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-05-30T09:42:22.412749Z")
					return t
				}(),

			},
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, _, err := awx.UserService.ListUsers(map[string]string{
		"name": "Demo User",
	})
	if err != nil {
		log.Fatalf("ListInventories err: %s", err)
	}

	if !reflect.DeepEqual(result, expectListUsersResponse) {
		log.Fatalf("ListUsers resp not as expected, expected: %v, resp result: %v", expectListInventoriesResponse[0], result[0])
	}

	log.Println("ListUsers passed!")
}

func TestCreateUser(t *testing.T) {
	var (
		expectCreateInventoryResponse = &User{
			ID:   1,
				Type: "inventory",
				URL:  "/api/v2/users/1/",
				Related: &Related{
                    Organization: "/api/v2/users/1/organizations/",
                    Credentials: "/api/v2/users/1/credentials/",
                    Permissions: "/api/v2/users/1/permissions/",
                    ActivityStream: "/api/v2/users/1/activity_stream/",
                    Projects: "/api/v2/users/1/projects/",
                    Teams: "/api/v2/users/1/teams/"
				},
				
				Created: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-05-21T01:34:35.657185Z")
					return t
				}(),

				Modified: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-05-30T09:42:22.412749Z")
					return t
				}(),
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.UserService.CreateUser(map[string]interface{}{
		"name":         "TestInventory",
		"description":  "for testing CreateInventory api",
		"organization": 1,
		"kind":         "",
		"host_filter":  "",
		"variables":    "",
	}, map[string]string{})
	if err != nil {
		log.Fatalf("CreateInventory err: %s", err)
	}

	if !reflect.DeepEqual(result, expectCreateUserResponse) {
		log.Fatalf("CreateUser resp not as expected, expected: %v, resp result: %v", expectCreateInventoryResponse, result)
	}

	log.Println("CreateUser passed!")
}
