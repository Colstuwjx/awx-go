package mockdata

var (
	// MockedListUsersResponse mocked `ListUsers` endpoint response
	MockedListUsersResponse = []byte(`
    {
        "count": 1,
        "next": null,
        "previous": null,
        "results": [
            {
                "id": 1,
                "type": "user",
                "url": "/api/v2/users/1/",
                "related": {
                    "admin_of_organizations": "/api/v2/users/1/admin_of_organizations/",
                    "organizations": "/api/v2/users/1/organizations/",
                    "roles": "/api/v2/users/1/roles/",
                    "access_list": "/api/v2/users/1/access_list/",
                    "teams": "/api/v2/users/1/teams/",
                    "credentials": "/api/v2/users/1/credentials/",
                    "activity_stream": "/api/v2/users/1/activity_stream/",
                    "projects": "/api/v2/users/1/projects/"
                },
                "summary_fields": {
                    "user_capabilities": {
                        "edit": true,
                        "delete": true
                    }
                },
                "created": "2017-09-06T02:55:30.492Z",
                "username": "admin",
                "first_name": "",
                "last_name": "",
                "email": "admin@example.com",
                "is_superuser": true,
                "is_system_auditor": false,
                "ldap_dn": "",
                "external_account": null
            }
        ]
    }`)

	// MockedCreateUserResponse mocked `CreateUser` endpoint response
	MockedCreateUserResponse = []byte(`
    {
        "id": 1,
        "type": "user",
        "url": "/api/v2/users/1/",
        "related": {
            "admin_of_organizations": "/api/v2/users/1/admin_of_organizations/",
            "organizations": "/api/v2/users/1/organizations/",
            "roles": "/api/v2/users/1/roles/",
            "access_list": "/api/v2/users/1/access_list/",
            "teams": "/api/v2/users/1/teams/",
            "credentials": "/api/v2/users/1/credentials/",
            "activity_stream": "/api/v2/users/1/activity_stream/",
            "projects": "/api/v2/users/1/projects/"
        },
        "summary_fields": {
            "user_capabilities": {
                "edit": true,
                "delete": true
            }
        },
        "created": "2017-09-06T02:55:30.492Z",
        "username": "admin",
        "first_name": "",
        "last_name": "",
        "email": "admin@example.com",
        "is_superuser": true,
        "is_system_auditor": false,
        "ldap_dn": "",
        "external_account": null
    }`)

	// MockedUpdateUserResponse mocked `UpdateUser` endpoint response
	MockedUpdateUserResponse = []byte(`
        {
            "id": 1,
            "type": "user",
            "url": "/api/v2/users/1/",
            "related": {
                "admin_of_organizations": "/api/v2/users/1/admin_of_organizations/",
                "organizations": "/api/v2/users/1/organizations/",
                "roles": "/api/v2/users/1/roles/",
                "access_list": "/api/v2/users/1/access_list/",
                "teams": "/api/v2/users/1/teams/",
                "credentials": "/api/v2/users/1/credentials/",
                "activity_stream": "/api/v2/users/1/activity_stream/",
                "projects": "/api/v2/users/1/projects/"
            },
            "summary_fields": {
                "user_capabilities": {
                    "edit": true,
                    "delete": true
                }
            },
            "created": "2017-09-06T02:55:30.492Z",
            "username": "admin",
            "first_name": "",
            "last_name": "",
            "email": "admin@example.net",
            "is_superuser": true,
            "is_system_auditor": false,
            "ldap_dn": "",
            "external_account": null
        }`)

	// MockedDeleteUserResponse mocked `DeleteUser` endpoint response
	MockedDeleteUserResponse = []byte(`{ }`)
)
