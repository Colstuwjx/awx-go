package awx

import (
	"testing"
	"time"
)

func TestListOrganizations(t *testing.T) {
	var (
		expectListOrganizationsResponse = []*Organization{
			{
				ID:   1,
				Type: "organization",
				URL:  "/api/v2/organizations/1/",
				Related: &Related{
					NotificationTemplatesError:   "/api/v2/organizations/1/notification_templates_error/",
					NotificationTemplatesSuccess: "/api/v2/organizations/1/notification_templates_success/",
					Users: "/api/v2/organizations/1/users/",
					NotificationTemplatesAny: "/api/v2/organizations/1/notification_templates_any/",
					NotificationTemplates:    "/api/v2/organizations/1/notification_templates/",
					Applications:             "/api/v2/organizations/1/applications/",
					InstanceGroups:           "/api/v2/organizations/1/instance_groups/",
					Credentials:              "/api/v2/organizations/1/credentials/",
					Inventories:              "/api/v2/organizations/1/inventories/",
					Projects:                 "/api/v2/organizations/1/projects/",
					WorkflowJobTemplate:      "/api/v2/organizations/1/workflow_job_templates/",
					ObjectRoles:              "/api/v2/organizations/1/object_roles/",
					AccessList:               "/api/v2/organizations/1/access_list/",
					Teams:                    "/api/v2/organizations/1/teams/",
					Admins:                   "/api/v2/organizations/1/admins/",
					ActivityStream:           "/api/v2/organizations/1/activity_stream/",
				},
				SummaryFields: &Summary{
					ObjectRoles: &ObjectRoles{
						AdminRole: &ApplyRole{
							ID:          6,
							Description: "Can manage all aspects of the organization",
							Name:        "Admin",
						},
						ReadRole: &ApplyRole{
							ID:          3,
							Description: "May view settings for the organization",
							Name:        "Read",
						},
						MemberRole: &ApplyRole{
							ID:          4,
							Description: "User is a member of the organization",
							Name:        "Member",
						},
						ExecuteRole: &ApplyRole{
							ID:          5,
							Description: "May run any executable resources in the organization",
							Name:        "Execute",
						},
						NotificationAdminRole: &ApplyRole{
							ID:          7,
							Description: "Can manage all notifications of the organization",
							Name:        "Notification Admin",
						},
						WorkflowAdminRole: &ApplyRole{
							ID:          8,
							Description: "Can manage all workflows of the organization",
							Name:        "Workflow Admin",
						},
						CredentialAdminRole: &ApplyRole{
							ID:          2,
							Description: "Can manage all credentials of the organization",
							Name:        "Credential Admin",
						},
						JobTemplateAdminRole: &ApplyRole{
							ID:          9,
							Description: "Can manage all job templates of the organization",
							Name:        "Job Template Admin",
						},
						ProjectAdminRole: &ApplyRole{
							ID:          10,
							Description: "Can manage all projects of the organization",
							Name:        "Project Admin",
						},
						AuditorRole: &ApplyRole{
							ID:          11,
							Description: "Can view all settings for the organization",
							Name:        "Auditor",
						},
						InventoryAdminRole: &ApplyRole{
							ID:          12,
							Description: "Can manage all inventories of the organization",
							Name:        "Inventory Admin",
						},
					},
					UserCapabilities: &UserCapabilities{
						Edit:   true,
						Delete: true,
					},
					RelatedFieldCounts: &RelatedFieldCount{
						JobTemplates: 1,
						Users:        1,
						Teams:        4,
						Admins:       0,
						Inventories:  2,
						Projects:     1,
					},
				},
				Created: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-12-01T12:10:00.152668Z")
					return t
				}(),
				Modified: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-12-01T12:10:00.152668Z")
					return t
				}(),
				Name:             "Default",
				Description:      "",
				CustomVirtualEnv: "",
			},
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, _, err := awx.OrganizationService.ListOrganizations(map[string]string{
		"name": "Default",
	})

	if err != nil {
		t.Fatalf("ListOrganizations err: %s", err)
	} else {
		checkAPICallResult(t, expectListOrganizationsResponse, result)
		t.Log("ListOrganizations passed!")
	}
}

func TestCreateOrganization(t *testing.T) {
	var (
		expectCreateOrganizationResponse = &Organization{
			ID:   1,
			Type: "organization",
			URL:  "/api/v2/organizations/1/",
			Related: &Related{
				NotificationTemplatesError:   "/api/v2/organizations/1/notification_templates_error/",
				NotificationTemplatesSuccess: "/api/v2/organizations/1/notification_templates_success/",
				Users: "/api/v2/organizations/1/users/",
				NotificationTemplatesAny: "/api/v2/organizations/1/notification_templates_any/",
				NotificationTemplates:    "/api/v2/organizations/1/notification_templates/",
				Applications:             "/api/v2/organizations/1/applications/",
				InstanceGroups:           "/api/v2/organizations/1/instance_groups/",
				Credentials:              "/api/v2/organizations/1/credentials/",
				Inventories:              "/api/v2/organizations/1/inventories/",
				Projects:                 "/api/v2/organizations/1/projects/",
				WorkflowJobTemplate:      "/api/v2/organizations/1/workflow_job_templates/",
				ObjectRoles:              "/api/v2/organizations/1/object_roles/",
				AccessList:               "/api/v2/organizations/1/access_list/",
				Teams:                    "/api/v2/organizations/1/teams/",
				Admins:                   "/api/v2/organizations/1/admins/",
				ActivityStream:           "/api/v2/organizations/1/activity_stream/",
			},
			SummaryFields: &Summary{
				ObjectRoles: &ObjectRoles{
					AdminRole: &ApplyRole{
						ID:          6,
						Description: "Can manage all aspects of the organization",
						Name:        "Admin",
					},
					ReadRole: &ApplyRole{
						ID:          3,
						Description: "May view settings for the organization",
						Name:        "Read",
					},
					MemberRole: &ApplyRole{
						ID:          4,
						Description: "User is a member of the organization",
						Name:        "Member",
					},
					ExecuteRole: &ApplyRole{
						ID:          5,
						Description: "May run any executable resources in the organization",
						Name:        "Execute",
					},
					NotificationAdminRole: &ApplyRole{
						ID:          7,
						Description: "Can manage all notifications of the organization",
						Name:        "Notification Admin",
					},
					WorkflowAdminRole: &ApplyRole{
						ID:          8,
						Description: "Can manage all workflows of the organization",
						Name:        "Workflow Admin",
					},
					CredentialAdminRole: &ApplyRole{
						ID:          2,
						Description: "Can manage all credentials of the organization",
						Name:        "Credential Admin",
					},
					JobTemplateAdminRole: &ApplyRole{
						ID:          9,
						Description: "Can manage all job templates of the organization",
						Name:        "Job Template Admin",
					},
					ProjectAdminRole: &ApplyRole{
						ID:          10,
						Description: "Can manage all projects of the organization",
						Name:        "Project Admin",
					},
					AuditorRole: &ApplyRole{
						ID:          11,
						Description: "Can view all settings for the organization",
						Name:        "Auditor",
					},
					InventoryAdminRole: &ApplyRole{
						ID:          12,
						Description: "Can manage all inventories of the organization",
						Name:        "Inventory Admin",
					},
				},
				UserCapabilities: &UserCapabilities{
					Edit:   true,
					Delete: true,
				},
				RelatedFieldCounts: &RelatedFieldCount{
					JobTemplates: 1,
					Users:        1,
					Teams:        4,
					Admins:       0,
					Inventories:  2,
					Projects:     1,
				},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-12-01T12:10:00.152668Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-12-01T12:10:00.152668Z")
				return t
			}(),
			Name:             "test-organization",
			Description:      "",
			CustomVirtualEnv: "",
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.OrganizationService.CreateOrganization(map[string]interface{}{
		"name":              "test-organization",
		"description":       "",
		"custom_virtualenv": "",
	}, map[string]string{})
	if err != nil {
		t.Fatalf("CreateOrganization err: %s", err)
	} else {
		checkAPICallResult(t, expectCreateOrganizationResponse, result)
		t.Log("CreateOrganization passed!")
	}
}

func TestUpdateOrganization(t *testing.T) {
	var (
		expectUpdateOrganizationResponse = &Organization{
			ID:   1,
			Type: "organization",
			URL:  "/api/v2/organizations/1/",
			Related: &Related{
				NotificationTemplatesError:   "/api/v2/organizations/1/notification_templates_error/",
				NotificationTemplatesSuccess: "/api/v2/organizations/1/notification_templates_success/",
				Users: "/api/v2/organizations/1/users/",
				NotificationTemplatesAny: "/api/v2/organizations/1/notification_templates_any/",
				NotificationTemplates:    "/api/v2/organizations/1/notification_templates/",
				Applications:             "/api/v2/organizations/1/applications/",
				InstanceGroups:           "/api/v2/organizations/1/instance_groups/",
				Credentials:              "/api/v2/organizations/1/credentials/",
				Inventories:              "/api/v2/organizations/1/inventories/",
				Projects:                 "/api/v2/organizations/1/projects/",
				WorkflowJobTemplate:      "/api/v2/organizations/1/workflow_job_templates/",
				ObjectRoles:              "/api/v2/organizations/1/object_roles/",
				AccessList:               "/api/v2/organizations/1/access_list/",
				Teams:                    "/api/v2/organizations/1/teams/",
				Admins:                   "/api/v2/organizations/1/admins/",
				ActivityStream:           "/api/v2/organizations/1/activity_stream/",
			},
			SummaryFields: &Summary{
				ObjectRoles: &ObjectRoles{
					AdminRole: &ApplyRole{
						ID:          6,
						Description: "Can manage all aspects of the organization",
						Name:        "Admin",
					},
					ReadRole: &ApplyRole{
						ID:          3,
						Description: "May view settings for the organization",
						Name:        "Read",
					},
					MemberRole: &ApplyRole{
						ID:          4,
						Description: "User is a member of the organization",
						Name:        "Member",
					},
					ExecuteRole: &ApplyRole{
						ID:          5,
						Description: "May run any executable resources in the organization",
						Name:        "Execute",
					},
					NotificationAdminRole: &ApplyRole{
						ID:          7,
						Description: "Can manage all notifications of the organization",
						Name:        "Notification Admin",
					},
					WorkflowAdminRole: &ApplyRole{
						ID:          8,
						Description: "Can manage all workflows of the organization",
						Name:        "Workflow Admin",
					},
					CredentialAdminRole: &ApplyRole{
						ID:          2,
						Description: "Can manage all credentials of the organization",
						Name:        "Credential Admin",
					},
					JobTemplateAdminRole: &ApplyRole{
						ID:          9,
						Description: "Can manage all job templates of the organization",
						Name:        "Job Template Admin",
					},
					ProjectAdminRole: &ApplyRole{
						ID:          10,
						Description: "Can manage all projects of the organization",
						Name:        "Project Admin",
					},
					AuditorRole: &ApplyRole{
						ID:          11,
						Description: "Can view all settings for the organization",
						Name:        "Auditor",
					},
					InventoryAdminRole: &ApplyRole{
						ID:          12,
						Description: "Can manage all inventories of the organization",
						Name:        "Inventory Admin",
					},
				},
				UserCapabilities: &UserCapabilities{
					Edit:   true,
					Delete: true,
				},
				RelatedFieldCounts: &RelatedFieldCount{
					JobTemplates: 1,
					Users:        1,
					Teams:        4,
					Admins:       0,
					Inventories:  2,
					Projects:     1,
				},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-12-01T12:10:00.152668Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-12-01T12:10:00.152668Z")
				return t
			}(),
			Name:        "test-organization",
			Description: "Update test-organization",
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.OrganizationService.UpdateOrganization(1, map[string]interface{}{
		"name":        "test-organization",
		"description": "Update test-organization",
	}, map[string]string{})
	if err != nil {
		t.Fatalf("CreateOrganization err: %s", err)
	} else {
		checkAPICallResult(t, expectUpdateOrganizationResponse, result)
		t.Log("CreateOrganization passed!")
	}
}
func TestDeleteOrganization(t *testing.T) {
	var (
		expectDeleteOrganizationResponse = &Organization{}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.OrganizationService.DeleteOrganization(1)

	if err != nil {
		t.Fatalf("DeleteOrganization err: %s", err)
	} else {
		checkAPICallResult(t, expectDeleteOrganizationResponse, result)
		t.Log("DeleteOrganization passed!")
	}
}
