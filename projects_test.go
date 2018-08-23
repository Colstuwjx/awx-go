package awx

import (
	"reflect"
	"testing"
	"time"
)

func TestListProjects(t *testing.T) {
	var (
		expectListProjectsResponse = []*Project{
			{
				ID:   4,
				Type: "project",
				URL:  "/api/v2/projects/4/",
				Related: &Related{
					CreatedBy:    "/api/v2/users/1/",
					ModifiedBy:   "/api/v2/users/1/",
					ObjectRoles:  "/api/v2/projects/4/object_roles/",
					Copy:         "/api/v2/projects/4/copy/",
					AccessList:   "/api/v2/projects/4/access_list/",
					Schedules:    "/api/v2/projects/4/schedules/",
					Organization: "/api/v2/organizations/1/",
				},
				SummaryFields: &Summary{
					Organization: &OrgnizationSummary{
						ID:          1,
						Name:        "Default",
						Description: "",
					},
					CreatedBy: &ByUserSummary{
						ID:        1,
						Username:  "admin",
						FirstName: "admin",
						LastName:  "admin",
					},
					ModifiedBy: &ByUserSummary{
						ID:        1,
						Username:  "admin",
						FirstName: "admin",
						LastName:  "admin",
					},
					ObjectRoles: &ObjectRoles{
						AdminRole: &ApplyRole{
							ID:          14,
							Description: "Can manage all aspects of the project",
							Name:        "Admin",
						},

						UseRole: &ApplyRole{
							ID:          16,
							Description: "Can manage all aspects of the project",
							Name:        "Use",
						},

						UpdateRole: &ApplyRole{
							ID:          17,
							Description: "May update project or inventory or group using the configured source update system",
							Name:        "Update",
						},

						ReadRole: &ApplyRole{
							ID:          15,
							Description: "May view settings for the project",
							Name:        "Read",
						},
					},

					UserCapabilities: &UserCapabilities{
						Edit:     true,
						Start:    true,
						Copy:     true,
						Schedule: true,
						Delete:   true,
					},
				},
				Created: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-06-28T16:31:15.809617Z")
					return t
				}(),

				Modified: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-06-28T16:31:15.923732Z")
					return t
				}(),
				Name:                  "Demo Project",
				Description:           "",
				LocalPath:             "",
				ScmType:               "git",
				ScmURL:                "https://github.com/ansible/ansible-tower-samples",
				ScmClean:              false,
				ScmDeleteOnUpdate:     false,
				Status:                "never updated",
				ScmDeleteOnNextUpdate: false,
				ScmUpdateOnLaunch:     true,
				ScmUpdateCacheTimeout: 0,
				LastUpdateFailed:      false,
			},
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, _, err := awx.ProjectService.ListProjects(map[string]string{
		"name": "Demo Project",
	})

	if err != nil {
		t.Fatalf("ListProjects err: %s", err)
	} else if !reflect.DeepEqual(result[0], expectListProjectsResponse[0]) {
		t.Logf("expected: %v", expectListProjectsResponse[0])
		t.Logf("result: %v", result[0])
		t.Fatal("ListProjects response is not as expected")
	} else {
		t.Log("ListProjects passed!")
	}
}

func TestCreateProject(t *testing.T) {
	var (
		expectCreateProjectResponse = &Project{
			ID:   4,
			Type: "project",
			URL:  "/api/v2/projects/4/",
			Related: &Related{
				CreatedBy:    "/api/v2/users/1/",
				ModifiedBy:   "/api/v2/users/1/",
				ObjectRoles:  "/api/v2/projects/4/object_roles/",
				Copy:         "/api/v2/projects/4/copy/",
				AccessList:   "/api/v2/projects/4/access_list/",
				Schedules:    "/api/v2/projects/4/schedules/",
				Organization: "/api/v2/organizations/1/",
			},
			SummaryFields: &Summary{
				Organization: &OrgnizationSummary{
					ID:          1,
					Name:        "Default",
					Description: "",
				},
				CreatedBy: &ByUserSummary{
					ID:        1,
					Username:  "admin",
					FirstName: "admin",
					LastName:  "admin",
				},
				ModifiedBy: &ByUserSummary{
					ID:        1,
					Username:  "admin",
					FirstName: "admin",
					LastName:  "admin",
				},
				ObjectRoles: &ObjectRoles{
					AdminRole: &ApplyRole{
						ID:          14,
						Description: "Can manage all aspects of the project",
						Name:        "Admin",
					},

					UseRole: &ApplyRole{
						ID:          16,
						Description: "Can manage all aspects of the project",
						Name:        "Use",
					},

					UpdateRole: &ApplyRole{
						ID:          17,
						Description: "May update project or inventory or group using the configured source update system",
						Name:        "Update",
					},

					ReadRole: &ApplyRole{
						ID:          15,
						Description: "May view settings for the project",
						Name:        "Read",
					},
				},

				UserCapabilities: &UserCapabilities{
					Edit:     true,
					Start:    true,
					Copy:     true,
					Schedule: true,
					Delete:   true,
				},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-06-28T16:31:15.809617Z")
				return t
			}(),

			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-06-28T16:31:15.923732Z")
				return t
			}(),
			Name:                  "Demo Project",
			Description:           "",
			LocalPath:             "",
			ScmType:               "git",
			ScmURL:                "https://github.com/ansible/ansible-tower-samples",
			ScmClean:              false,
			ScmDeleteOnUpdate:     false,
			Status:                "never updated",
			ScmDeleteOnNextUpdate: false,
			ScmUpdateOnLaunch:     true,
			ScmUpdateCacheTimeout: 0,
			LastUpdateFailed:      false,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.ProjectService.CreateProject(map[string]interface{}{
		"name":        "TestProject",
		"description": "Test project",
	}, map[string]string{})

	if err != nil {
		t.Fatalf("CreateInventory err: %s", err)
	} else if !reflect.DeepEqual(result, expectCreateProjectResponse) {
		t.Logf("expected: %v", expectCreateProjectResponse)
		t.Logf("result: %v", result)
		t.Fatal("CreateProject resp not as expected")
	} else {
		t.Log("CreateProject passed!")
	}
}
