package mockdata

var (
	// MockedListOrganizationResponse mocked `ListOrganizations` endpoint response
	MockedListOrganizationResponse = []byte(`
    {
        "count": 1,
        "next": null,
        "previous": null,
        "results": [
            {
                "id": 1,
                "type": "organization",
                "url": "/api/v2/organizations/1/",
                "related": {
                    "notification_templates_error": "/api/v2/organizations/1/notification_templates_error/",
                    "notification_templates_success": "/api/v2/organizations/1/notification_templates_success/",
                    "users": "/api/v2/organizations/1/users/",
                    "notification_templates_any": "/api/v2/organizations/1/notification_templates_any/",
                    "notification_templates": "/api/v2/organizations/1/notification_templates/",
                    "applications": "/api/v2/organizations/1/applications/",
                    "instance_groups": "/api/v2/organizations/1/instance_groups/",
                    "credentials": "/api/v2/organizations/1/credentials/",
                    "inventories": "/api/v2/organizations/1/inventories/",
                    "projects": "/api/v2/organizations/1/projects/",
                    "workflow_job_templates": "/api/v2/organizations/1/workflow_job_templates/",
                    "object_roles": "/api/v2/organizations/1/object_roles/",
                    "access_list": "/api/v2/organizations/1/access_list/",
                    "teams": "/api/v2/organizations/1/teams/",
                    "admins": "/api/v2/organizations/1/admins/",
                    "activity_stream": "/api/v2/organizations/1/activity_stream/"
                },
                "summary_fields": {
                    "object_roles": {
                        "admin_role": {
                            "id": 6,
                            "description": "Can manage all aspects of the organization",
                            "name": "Admin"
                        },
                        "read_role": {
                            "id": 3,
                            "description": "May view settings for the organization",
                            "name": "Read"
                        },
                        "member_role": {
                            "id": 4,
                            "description": "User is a member of the organization",
                            "name": "Member"
                        },
                        "execute_role": {
                            "id": 5,
                            "description": "May run any executable resources in the organization",
                            "name": "Execute"
                        },
                        "notification_admin_role": {
                            "id": 7,
                            "description": "Can manage all notifications of the organization",
                            "name": "Notification Admin"
                        },
                        "workflow_admin_role": {
                            "id": 8,
                            "description": "Can manage all workflows of the organization",
                            "name": "Workflow Admin"
                        },
                        "credential_admin_role": {
                            "id": 2,
                            "description": "Can manage all credentials of the organization",
                            "name": "Credential Admin"
                        },
                        "job_template_admin_role": {
                            "id": 9,
                            "description": "Can manage all job templates of the organization",
                            "name": "Job Template Admin"
                        },
                        "project_admin_role": {
                            "id": 10,
                            "description": "Can manage all projects of the organization",
                            "name": "Project Admin"
                        },
                        "auditor_role": {
                            "id": 11,
                            "description": "Can view all settings for the organization",
                            "name": "Auditor"
                        },
                        "inventory_admin_role": {
                            "id": 12,
                            "description": "Can manage all inventories of the organization",
                            "name": "Inventory Admin"
                        }
                    },
                    "user_capabilities": {
                        "edit": true,
                        "delete": true
                    },
                    "related_field_counts": {
                        "job_templates": 1,
                        "users": 1,
                        "teams": 4,
                        "admins": 0,
                        "inventories": 2,
                        "projects": 1
                    }
                },
                "created": "2018-12-01T12:10:00.152668Z",
                "modified": "2018-12-01T12:10:00.152668Z",
                "name": "Default",
                "description": "",
                "custom_virtualenv": null
            }
        ]
    }`)

	// MockedCreateOrganizationResponse mocked `CreateOrganization` endpoint response
	MockedCreateOrganizationResponse = []byte(`
     {
                    "id": 1,
                    "type": "organization",
                    "url": "/api/v2/organizations/1/",
                    "related": {
                        "notification_templates_error": "/api/v2/organizations/1/notification_templates_error/",
                        "notification_templates_success": "/api/v2/organizations/1/notification_templates_success/",
                        "users": "/api/v2/organizations/1/users/",
                        "notification_templates_any": "/api/v2/organizations/1/notification_templates_any/",
                        "notification_templates": "/api/v2/organizations/1/notification_templates/",
                        "applications": "/api/v2/organizations/1/applications/",
                        "instance_groups": "/api/v2/organizations/1/instance_groups/",
                        "credentials": "/api/v2/organizations/1/credentials/",
                        "inventories": "/api/v2/organizations/1/inventories/",
                        "projects": "/api/v2/organizations/1/projects/",
                        "workflow_job_templates": "/api/v2/organizations/1/workflow_job_templates/",
                        "object_roles": "/api/v2/organizations/1/object_roles/",
                        "access_list": "/api/v2/organizations/1/access_list/",
                        "teams": "/api/v2/organizations/1/teams/",
                        "admins": "/api/v2/organizations/1/admins/",
                        "activity_stream": "/api/v2/organizations/1/activity_stream/"
                    },
                    "summary_fields": {
                        "object_roles": {
                            "admin_role": {
                                "id": 6,
                                "description": "Can manage all aspects of the organization",
                                "name": "Admin"
                            },
                            "read_role": {
                                "id": 3,
                                "description": "May view settings for the organization",
                                "name": "Read"
                            },
                            "member_role": {
                                "id": 4,
                                "description": "User is a member of the organization",
                                "name": "Member"
                            },
                            "execute_role": {
                                "id": 5,
                                "description": "May run any executable resources in the organization",
                                "name": "Execute"
                            },
                            "notification_admin_role": {
                                "id": 7,
                                "description": "Can manage all notifications of the organization",
                                "name": "Notification Admin"
                            },
                            "workflow_admin_role": {
                                "id": 8,
                                "description": "Can manage all workflows of the organization",
                                "name": "Workflow Admin"
                            },
                            "credential_admin_role": {
                                "id": 2,
                                "description": "Can manage all credentials of the organization",
                                "name": "Credential Admin"
                            },
                            "job_template_admin_role": {
                                "id": 9,
                                "description": "Can manage all job templates of the organization",
                                "name": "Job Template Admin"
                            },
                            "project_admin_role": {
                                "id": 10,
                                "description": "Can manage all projects of the organization",
                                "name": "Project Admin"
                            },
                            "auditor_role": {
                                "id": 11,
                                "description": "Can view all settings for the organization",
                                "name": "Auditor"
                            },
                            "inventory_admin_role": {
                                "id": 12,
                                "description": "Can manage all inventories of the organization",
                                "name": "Inventory Admin"
                            }
                        },
                        "user_capabilities": {
                            "edit": true,
                            "delete": true
                        },
                        "related_field_counts": {
                            "job_templates": 1,
                            "users": 1,
                            "teams": 4,
                            "admins": 0,
                            "inventories": 2,
                            "projects": 1
                        }
                    },
                    "created": "2018-12-01T12:10:00.152668Z",
                    "modified": "2018-12-01T12:10:00.152668Z",
                    "name": "test-organization",
                    "description": "",
                    "custom_virtualenv": null
}`)

	// MockedUpdateOrganizationResponse mocked `UpdateOrganization` endpoint response
	MockedUpdateOrganizationResponse = []byte(`
        {
                       "id": 1,
                       "type": "organization",
                       "url": "/api/v2/organizations/1/",
                       "related": {
                           "notification_templates_error": "/api/v2/organizations/1/notification_templates_error/",
                           "notification_templates_success": "/api/v2/organizations/1/notification_templates_success/",
                           "users": "/api/v2/organizations/1/users/",
                           "notification_templates_any": "/api/v2/organizations/1/notification_templates_any/",
                           "notification_templates": "/api/v2/organizations/1/notification_templates/",
                           "applications": "/api/v2/organizations/1/applications/",
                           "instance_groups": "/api/v2/organizations/1/instance_groups/",
                           "credentials": "/api/v2/organizations/1/credentials/",
                           "inventories": "/api/v2/organizations/1/inventories/",
                           "projects": "/api/v2/organizations/1/projects/",
                           "workflow_job_templates": "/api/v2/organizations/1/workflow_job_templates/",
                           "object_roles": "/api/v2/organizations/1/object_roles/",
                           "access_list": "/api/v2/organizations/1/access_list/",
                           "teams": "/api/v2/organizations/1/teams/",
                           "admins": "/api/v2/organizations/1/admins/",
                           "activity_stream": "/api/v2/organizations/1/activity_stream/"
                       },
                       "summary_fields": {
                           "object_roles": {
                               "admin_role": {
                                   "id": 6,
                                   "description": "Can manage all aspects of the organization",
                                   "name": "Admin"
                               },
                               "read_role": {
                                   "id": 3,
                                   "description": "May view settings for the organization",
                                   "name": "Read"
                               },
                               "member_role": {
                                   "id": 4,
                                   "description": "User is a member of the organization",
                                   "name": "Member"
                               },
                               "execute_role": {
                                   "id": 5,
                                   "description": "May run any executable resources in the organization",
                                   "name": "Execute"
                               },
                               "notification_admin_role": {
                                   "id": 7,
                                   "description": "Can manage all notifications of the organization",
                                   "name": "Notification Admin"
                               },
                               "workflow_admin_role": {
                                   "id": 8,
                                   "description": "Can manage all workflows of the organization",
                                   "name": "Workflow Admin"
                               },
                               "credential_admin_role": {
                                   "id": 2,
                                   "description": "Can manage all credentials of the organization",
                                   "name": "Credential Admin"
                               },
                               "job_template_admin_role": {
                                   "id": 9,
                                   "description": "Can manage all job templates of the organization",
                                   "name": "Job Template Admin"
                               },
                               "project_admin_role": {
                                   "id": 10,
                                   "description": "Can manage all projects of the organization",
                                   "name": "Project Admin"
                               },
                               "auditor_role": {
                                   "id": 11,
                                   "description": "Can view all settings for the organization",
                                   "name": "Auditor"
                               },
                               "inventory_admin_role": {
                                   "id": 12,
                                   "description": "Can manage all inventories of the organization",
                                   "name": "Inventory Admin"
                               }
                           },
                           "user_capabilities": {
                               "edit": true,
                               "delete": true
                           },
                           "related_field_counts": {
                               "job_templates": 1,
                               "users": 1,
                               "teams": 4,
                               "admins": 0,
                               "inventories": 2,
                               "projects": 1
                           }
                       },
                       "created": "2018-12-01T12:10:00.152668Z",
                       "modified": "2018-12-01T12:10:00.152668Z",
                       "name": "test-organization",
                       "description": "Update test-organization",
                       "custom_virtualenv": null
   }`)
	// MockedDeleteOrganizationResponse mocked `DeleteOrganization` endpoint response
	MockedDeleteOrganizationResponse = []byte(`{ }`)
)
