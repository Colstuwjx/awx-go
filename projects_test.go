package awx

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func TestListProjects(t *testing.T) {
	var (
		expectListProjectsResponse = []*Project{
			{
				ID: 4,
				Type: "project",
				URL: "/api/v2/projects/4/",
				Related: &Related{
					CreatedBy: "/api/v2/users/1/",
					ModifiedBy: "/api/v2/users/1/",
					ObjectRoles: "/api/v2/projects/4/object_roles/",
					Copy: "/api/v2/projects/4/copy/",
					ProjectUpdates: "/api/v2/projects/4/project_updates/",
					Update: "/api/v2/projects/4/update/",
					AccessList: "/api/v2/projects/4/access_list/",
					Teams: "/api/v2/projects/4/teams/",
					SCMInventorySources: "/api/v2/projects/4/scm_inventory_sources/",
					InventoryFiles: "/api/v2/projects/4/inventories/",
					Schedules: "/api/v2/projects/4/schedules/",
					Playbooks: "/api/v2/projects/4/playbooks/",
					ActivityStream: "/api/v2/projects/4/activity_stream/",
					Organization: "/api/v2/organizations/1/"
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
							FirstName: "",
							LastName:  "",
						},
						ModifiedBy: &ByUserSummary{
							ID:        1,
							Username:  "admin",
							FirstName: "",
							LastName:  "",
						},
						ObjectRoles: &ObjectRoles{
							UseRole: &ApplyRole{
								ID:          23,
								Description: "Can manage all aspects of the project",
								Name:        "Use",
							},
	
							AdminRole: &ApplyRole{
								ID:          21,
								Description: "Can use the project in a job template",
								Name:        "Admin",
							},
	
							UpdateRole: &ApplyRole{
								ID:          24,
								Description: "May update project or inventory or group using the configured source update system",
								Name:        "Update",
							},
	
							ReadRole: &ApplyRole{
								ID:          22,
								Description: "May view settings for the project",
								Name:        "Read",
							},
						},
						
						UserCapabilities: &UserCapabilities{
							Edit:   true,
							Copy:   true,
							Start:  true,
							Schedule:  true,
							Delete: true,
						},
					},
					Created: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2018-05-21T01:34:35.657185Z")
						return t
					}(),
	
					Modified: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2018-05-30T09:42:22.412749Z")
						return t
					}(),
					Name: "Demo Project",
					Description: "",
				}]
			},
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, _, err := awx.ProjectService.ListProjects(map[string]string{
		"name": "Demo Project",
	})
	if err != nil {
		log.Fatalf("ListProjects err: %s", err)
	}

	if !reflect.DeepEqual(result, expectListProjectsResponse) {
		log.Fatalf("ListProjects resp not as expected, expected: %v, resp result: %v", expectListProjectsResponse[0], result[0])
	}

	log.Println("ListProjects passed!")
}

func TestCreateProject(t *testing.T) {
	var (
		expectCreateProjectResponse = &CreateProject{
			ID: 4,
			Type: "project",
			URL: "/api/v2/projects/4/",
			Related: &Related{
				CreatedBy: "/api/v2/users/1/",
				ModifiedBy: "/api/v2/users/1/",
				ObjectRoles: "/api/v2/projects/4/object_roles/",
				Copy: "/api/v2/projects/4/copy/",
				ProjectUpdates: "/api/v2/projects/4/project_updates/",
				Update: "/api/v2/projects/4/update/",
				AccessList: "/api/v2/projects/4/access_list/",
				Teams: "/api/v2/projects/4/teams/",
				SCMInventorySources: "/api/v2/projects/4/scm_inventory_sources/",
				InventoryFiles: "/api/v2/projects/4/inventories/",
				Schedules: "/api/v2/projects/4/schedules/",
				Playbooks: "/api/v2/projects/4/playbooks/",
				ActivityStream: "/api/v2/projects/4/activity_stream/",
				Organization: "/api/v2/organizations/1/"
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
						FirstName: "",
						LastName:  "",
					},
					ModifiedBy: &ByUserSummary{
						ID:        1,
						Username:  "admin",
						FirstName: "",
						LastName:  "",
					},
					ObjectRoles: &ObjectRoles{
						UseRole: &ApplyRole{
							ID:          23,
							Description: "Can manage all aspects of the project",
							Name:        "Use",
						},

						AdminRole: &ApplyRole{
							ID:          21,
							Description: "Can use the project in a job template",
							Name:        "Admin",
						},

						UpdateRole: &ApplyRole{
							ID:          24,
							Description: "May update project or inventory or group using the configured source update system",
							Name:        "Update",
						},

						ReadRole: &ApplyRole{
							ID:          22,
							Description: "May view settings for the project",
							Name:        "Read",
						},
					},
					
					UserCapabilities: &UserCapabilities{
						Edit:   true,
						Copy:   true,
						Start:  true,
						Schedule:  true,
						Delete: true,
					},
				},
				Created: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-05-21T01:34:35.657185Z")
					return t
				}(),

				Modified: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-05-30T09:42:22.412749Z")
					return t
				}(),
				Name: "Demo Project",
				Description: "",
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.ProjectsService.CreateProject(map[string]interface{}{
		"name":         "TestProject",
		"description":  "Test project",
	}, map[string]string{})
	if err != nil {
		log.Fatalf("CreateInventory err: %s", err)
	}

	if !reflect.DeepEqual(result, expectCreateProjectResponse) {
		log.Fatalf("CreateProject resp not as expected, expected: %v, resp result: %v", expectCreateProjectResponse, result)
	}

	log.Println("CreateProject passed!")
}
