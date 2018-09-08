package mockdata

var (
	// MockedListProjectsResponse mocked `ListProjects` endpoint response
	MockedListProjectsResponse = []byte(`
    {
        "count": 1,
        "next": null,
        "previous": null,
        "results": [{
            "id": 4,
            "type": "project",
            "url": "/api/v2/projects/4/",
            "related": {
                "created_by": "/api/v2/users/1/",
                "modified_by": "/api/v2/users/1/",
                "object_roles": "/api/v2/projects/4/object_roles/",
                "copy": "/api/v2/projects/4/copy/",
                "access_list": "/api/v2/projects/4/access_list/",
                "schedules": "/api/v2/projects/4/schedules/",
                "organization": "/api/v2/organizations/1/"
            },
            "summary_fields": {
                "organization": {
                    "id": 1,
                    "name": "Default",
                    "description": ""
                },
                "created_by": {
                    "id": 1,
                    "username": "admin",
                    "first_name": "admin",
                    "last_name": "admin"
                },
                "modified_by": {
                    "id": 1,
                    "username": "admin",
                    "first_name": "admin",
                    "last_name": "admin"
                },
                "object_roles": {
                    "admin_role": {
                        "id": 14,
                        "description": "Can manage all aspects of the project",
                        "name": "Admin"
                    },
                    "use_role": {
                        "id": 16,
                        "description": "Can manage all aspects of the project",
                        "name": "Use"
                    },
                    "update_role": {
                        "id": 17,
                        "description": "May update project or inventory or group using the configured source update system",
                        "name": "Update"
                    },
                    "read_role": {
                        "id": 15,
                        "description": "May view settings for the project",
                        "name": "Read"
                    }
                },
                "user_capabilities": {
                    "edit": true,
                    "start": true,
                    "copy": true,
                    "schedule": true,
                    "delete": true
                }
            },
            "created": "2018-06-28T16:31:15.809617Z",
            "modified": "2018-06-28T16:31:15.923732Z",
            "name": "Demo Project",
            "description": "",
            "local_path": "",
            "scm_type": "git",
            "scm_url": "https://github.com/ansible/ansible-tower-samples",
            "scm_clean": false,
            "scm_delete_on_update": false,
            "status": "never updated",
            "scm_delete_on_next_update": false,
            "scm_update_on_launch": true,
            "scm_update_cache_timeout": 0,
            "last_update_failed": false
        }]
    }`)

	// MockedCreateProjectResponse mocked `CreateProject` endpoint response
	MockedCreateProjectResponse = []byte(`
    {
        "id": 4,
        "type": "project",
        "url": "/api/v2/projects/4/",
        "related": {
            "created_by": "/api/v2/users/1/",
            "modified_by": "/api/v2/users/1/",
            "object_roles": "/api/v2/projects/4/object_roles/",
            "copy": "/api/v2/projects/4/copy/",
            "access_list": "/api/v2/projects/4/access_list/",
            "schedules": "/api/v2/projects/4/schedules/",
            "organization": "/api/v2/organizations/1/"
        },
        "summary_fields": {
            "organization": {
                "id": 1,
                "name": "Default",
                "description": ""
            },
            "created_by": {
                "id": 1,
                "username": "admin",
                "first_name": "admin",
                "last_name": "admin"
            },
            "modified_by": {
                "id": 1,
                "username": "admin",
                "first_name": "admin",
                "last_name": "admin"
            },
            "object_roles": {
                "admin_role": {
                    "id": 14,
                    "description": "Can manage all aspects of the project",
                    "name": "Admin"
                },
                "use_role": {
                    "id": 16,
                    "description": "Can manage all aspects of the project",
                    "name": "Use"
                },
                "update_role": {
                    "id": 17,
                    "description": "May update project or inventory or group using the configured source update system",
                    "name": "Update"
                },
                "read_role": {
                    "id": 15,
                    "description": "May view settings for the project",
                    "name": "Read"
                }
            },
            "user_capabilities": {
                "edit": true,
                "start": true,
                "copy": true,
                "schedule": true,
                "delete": true
            }
        },
        "created": "2018-06-28T16:31:15.809617Z",
        "modified": "2018-06-28T16:31:15.923732Z",
        "name": "Demo Project",
        "description": "",
        "local_path": "",
        "scm_type": "git",
        "scm_url": "https://github.com/ansible/ansible-tower-samples",
        "scm_clean": false,
        "scm_delete_on_update": false,
        "status": "never updated",
        "scm_delete_on_next_update": false,
        "scm_update_on_launch": true,
        "scm_update_cache_timeout": 0,
        "last_update_failed": false
    }`)
	// MockedUpdateProjectResponse mocked `UpdateProject` endpoint response
	MockedUpdateProjectResponse = []byte(`
    {
        "id": 4,
        "type": "project",
        "url": "/api/v2/projects/4/",
        "related": {
            "created_by": "/api/v2/users/1/",
            "modified_by": "/api/v2/users/1/",
            "object_roles": "/api/v2/projects/4/object_roles/",
            "copy": "/api/v2/projects/4/copy/",
            "access_list": "/api/v2/projects/4/access_list/",
            "schedules": "/api/v2/projects/4/schedules/",
            "organization": "/api/v2/organizations/1/"
        },
        "summary_fields": {
            "organization": {
                "id": 1,
                "name": "Default",
                "description": ""
            },
            "created_by": {
                "id": 1,
                "username": "admin",
                "first_name": "admin",
                "last_name": "admin"
            },
            "modified_by": {
                "id": 1,
                "username": "admin",
                "first_name": "admin",
                "last_name": "admin"
            },
            "object_roles": {
                "admin_role": {
                    "id": 14,
                    "description": "Can manage all aspects of the project",
                    "name": "Admin"
                },
                "use_role": {
                    "id": 16,
                    "description": "Can manage all aspects of the project",
                    "name": "Use"
                },
                "update_role": {
                    "id": 17,
                    "description": "May update project or inventory or group using the configured source update system",
                    "name": "Update"
                },
                "read_role": {
                    "id": 15,
                    "description": "May view settings for the project",
                    "name": "Read"
                }
            },
            "user_capabilities": {
                "edit": true,
                "start": true,
                "copy": true,
                "schedule": true,
                "delete": true
            }
        },
        "created": "2018-06-28T16:31:15.809617Z",
        "modified": "2018-06-28T16:31:15.923732Z",
        "name": "Demo Project",
        "description": "Test Update",
        "local_path": "",
        "scm_type": "git",
        "scm_url": "https://github.com/ansible/ansible-tower-samples",
        "scm_clean": false,
        "scm_delete_on_update": false,
        "status": "never updated",
        "scm_delete_on_next_update": false,
        "scm_update_on_launch": true,
        "scm_update_cache_timeout": 0,
        "last_update_failed": false
    }`)

	// MockedDeleteProjectResponse mocked `DeleteProject` endpoint response
	MockedDeleteProjectResponse = []byte(`{}`)
)
