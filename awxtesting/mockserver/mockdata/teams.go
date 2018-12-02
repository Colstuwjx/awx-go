package mockdata

var (
	// MockedListTeamResponse mocked `ListTeams` endpoint response
	MockedListTeamResponse = []byte(`
    {
        "count": 1,
        "next": null,
        "previous": null,
        "results": [
            {
                "id": 1,
                "type": "team",
                "url": "/api/v2/teams/1/",
                "related": {
                    "created_by": "/api/v2/users/4/",
                    "modified_by": "/api/v2/users/4/",
                    "users": "/api/v2/teams/1/users/",
                    "roles": "/api/v2/teams/1/roles/",
                    "object_roles": "/api/v2/teams/1/object_roles/",
                    "credentials": "/api/v2/teams/1/credentials/",
                    "projects": "/api/v2/teams/1/projects/",
                    "activity_stream": "/api/v2/teams/1/activity_stream/",
                    "access_list": "/api/v2/teams/1/access_list/",
                    "organization": "/api/v2/organizations/1/"
                },
                "summary_fields": {
                    "organization": {
                        "id": 1,
                        "name": "Default",
                        "description": ""
                    },
                    "created_by": {
                        "id": 4,
                        "username": "admin",
                        "first_name": "",
                        "last_name": ""
                    },
                    "modified_by": {
                        "id": 4,
                        "username": "admin",
                        "first_name": "",
                        "last_name": ""
                    },
                    "object_roles": {
                        "admin_role": {
                            "id": 30,
                            "description": "Can manage all aspects of the team",
                            "name": "Admin"
                        },
                        "member_role": {
                            "id": 29,
                            "description": "User is a member of the team",
                            "name": "Member"
                        },
                        "read_role": {
                            "id": 31,
                            "description": "May view settings for the team",
                            "name": "Read"
                        }
                    },
                    "user_capabilities": {
                        "edit": true,
                        "delete": true
                    }
                },
                "created": "2018-12-01T13:33:30.692904Z",
                "modified": "2018-12-01T13:33:30.692904Z",
                "name": "test-team",
                "description": "",
                "organization": 1
            }
        ]
    }`)
	//MockedCreateTeamResponse mocked "CreateTeam" api endpoint
	MockedCreateTeamResponse = []byte(`
                {
                    "id": 1,
                    "type": "team",
                    "url": "/api/v2/teams/1/",
                    "related": {
                        "created_by": "/api/v2/users/4/",
                        "modified_by": "/api/v2/users/4/",
                        "users": "/api/v2/teams/1/users/",
                        "roles": "/api/v2/teams/1/roles/",
                        "object_roles": "/api/v2/teams/1/object_roles/",
                        "credentials": "/api/v2/teams/1/credentials/",
                        "projects": "/api/v2/teams/1/projects/",
                        "activity_stream": "/api/v2/teams/1/activity_stream/",
                        "access_list": "/api/v2/teams/1/access_list/",
                        "organization": "/api/v2/organizations/1/"
                    },
                    "summary_fields": {
                        "organization": {
                            "id": 1,
                            "name": "Default",
                            "description": ""
                        },
                        "created_by": {
                            "id": 4,
                            "username": "admin",
                            "first_name": "",
                            "last_name": ""
                        },
                        "modified_by": {
                            "id": 4,
                            "username": "admin",
                            "first_name": "",
                            "last_name": ""
                        },
                        "object_roles": {
                            "admin_role": {
                                "id": 30,
                                "description": "Can manage all aspects of the team",
                                "name": "Admin"
                            },
                            "member_role": {
                                "id": 29,
                                "description": "User is a member of the team",
                                "name": "Member"
                            },
                            "read_role": {
                                "id": 31,
                                "description": "May view settings for the team",
                                "name": "Read"
                            }
                        },
                        "user_capabilities": {
                            "edit": true,
                            "delete": true
                        }
                    },
                    "created": "2018-12-01T13:33:30.692904Z",
                    "modified": "2018-12-01T13:33:30.692904Z",
                    "name": "test-team",
                    "description": "",
                    "organization": 1
        }`)
	//MockedUpdateTeamResponse mocked "UpdateTeam" api endpoint
	MockedUpdateTeamResponse = []byte(`
        {
            "id": 1,
            "type": "team",
            "url": "/api/v2/teams/1/",
            "related": {
                "created_by": "/api/v2/users/4/",
                "modified_by": "/api/v2/users/4/",
                "users": "/api/v2/teams/1/users/",
                "roles": "/api/v2/teams/1/roles/",
                "object_roles": "/api/v2/teams/1/object_roles/",
                "credentials": "/api/v2/teams/1/credentials/",
                "projects": "/api/v2/teams/1/projects/",
                "activity_stream": "/api/v2/teams/1/activity_stream/",
                "access_list": "/api/v2/teams/1/access_list/",
                "organization": "/api/v2/organizations/1/"
            },
            "summary_fields": {
                "organization": {
                    "id": 1,
                    "name": "Default",
                    "description": ""
                },
                "created_by": {
                    "id": 4,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "modified_by": {
                    "id": 4,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "object_roles": {
                    "admin_role": {
                        "id": 30,
                        "description": "Can manage all aspects of the team",
                        "name": "Admin"
                    },
                    "member_role": {
                        "id": 29,
                        "description": "User is a member of the team",
                        "name": "Member"
                    },
                    "read_role": {
                        "id": 31,
                        "description": "May view settings for the team",
                        "name": "Read"
                    }
                },
                "user_capabilities": {
                    "edit": true,
                    "delete": true
                }
            },
            "created": "2018-12-01T13:33:30.692904Z",
            "modified": "2018-12-01T13:33:30.692904Z",
            "name": "test-team",
            "description": "Update test-team",
            "organization": 1
}`)
	// MockedDeleteTeamResponse mocked `DeleteTeam` endpoint response
	MockedDeleteTeamResponse = []byte(`{ }`)
	// MockedTeamGrantRoleResponse mocked `Team Grant Role` endpoint response
	MockedTeamGrantRoleResponse = []byte(`{ }`)
	// MockedTeamRevokeRoleResponse mocked `Team Revoke Role` endpoint response
	MockedTeamRevokeRoleResponse = []byte(`{ }`)
)
