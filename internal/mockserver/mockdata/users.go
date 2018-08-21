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
                    "credentials": "/api/v2/users/1/credentials/",
                    "permissions": "/api/v2/users/1/permissions/",
                    "activity_stream": "/api/v2/users/1/activity_stream/",
                    "projects": "/api/v2/users/1/projects/",
                    "teams": "/api/v2/users/1/teams/"
                },
                "created": "2015-08-25T01:00:07.127Z",
                "username": "test",
                "first_name": "",
                "last_name": "",
                "email": "admin@example.com",
                "is_superuser": true,
                "ldap_dn": ""
            }
        ]
    }`)

	// MockedCreateUserResponse mocked `CreateUser` endpoint response
	MockedCreateUserResponse = []byte(`
    {
        "id": 1,
        "type": "user",
        "url": "/api/v1/users/1/",
        "related": {
            "admin_of_organizations": "/api/v1/users/1/admin_of_organizations/",
            "organizations": "/api/v1/users/1/organizations/",
            "credentials": "/api/v1/users/1/credentials/",
            "permissions": "/api/v1/users/1/permissions/",
            "activity_stream": "/api/v1/users/1/activity_stream/",
            "projects": "/api/v1/users/1/projects/",
            "teams": "/api/v1/users/1/teams/"
        },
        "created": "2015-08-25T01:00:07.127Z",
        "username": "test",
        "first_name": "",
        "last_name": "",
        "email": "admin@example.com",
        "is_superuser": true,
        "ldap_dn": ""
    }`)
)
