package mockdata

var (
	// MockedListGroupsResponse mocked `ListGroups` endpoint response
	MockedListGroupsResponse = []byte(`
    {
        "count": 12,
        "next": null,
        "previous": null,
        "results": [
            {
                "id": 21,
                "type": "group",
                "url": "/api/v2/groups/21/",
                "related": {
                    "created_by": "/api/v2/users/11/",
                    "job_host_summaries": "/api/v2/groups/21/job_host_summaries/",
                    "variable_data": "/api/v2/groups/21/variable_data/",
                    "job_events": "/api/v2/groups/21/job_events/",
                    "potential_children": "/api/v2/groups/21/potential_children/",
                    "ad_hoc_commands": "/api/v2/groups/21/ad_hoc_commands/",
                    "all_hosts": "/api/v2/groups/21/all_hosts/",
                    "activity_stream": "/api/v2/groups/21/activity_stream/",
                    "hosts": "/api/v2/groups/21/hosts/",
                    "children": "/api/v2/groups/21/children/",
                    "inventory_sources": "/api/v2/groups/21/inventory_sources/",
                    "inventory": "/api/v2/inventories/9/"
                },
                "summary_fields": {
                    "inventory": {
                        "id": 9,
                        "name": "test",
                        "description": "",
                        "has_active_failures": false,
                        "total_hosts": 8,
                        "hosts_with_active_failures": 0,
                        "total_groups": 4,
                        "groups_with_active_failures": 0,
                        "has_inventory_sources": false,
                        "total_inventory_sources": 0,
                        "inventory_sources_with_failures": 0,
                        "organization_id": 16,
                        "kind": ""
                    },
                    "created_by": {
                        "id": 11,
                        "username": "demouser",
                        "first_name": "Demo",
                        "last_name": "User"
                    },
                    "user_capabilities": {
                        "edit": true,
                        "copy": true,
                        "delete": true
                    }
                },
                "created": "2018-07-17T13:27:46.686176Z",
                "modified": "2018-07-17T13:28:07.127040Z",
                "name": "Demo Group",
                "description": "",
                "inventory": 9,
                "variables": "",
                "has_active_failures": false,
                "total_hosts": 3,
                "hosts_with_active_failures": 0,
                "total_groups": 0,
                "groups_with_active_failures": 0,
                "has_inventory_sources": false
            }
        ]
    }`)

	// MockedCreateGroupResponse mocked `CreateGroup` endpoint response
	MockedCreateGroupResponse = []byte(`
    {
        "id": 21,
        "type": "group",
        "url": "/api/v2/groups/21/",
        "related": {
            "created_by": "/api/v2/users/11/",
            "job_host_summaries": "/api/v2/groups/21/job_host_summaries/",
            "variable_data": "/api/v2/groups/21/variable_data/",
            "job_events": "/api/v2/groups/21/job_events/",
            "potential_children": "/api/v2/groups/21/potential_children/",
            "ad_hoc_commands": "/api/v2/groups/21/ad_hoc_commands/",
            "all_hosts": "/api/v2/groups/21/all_hosts/",
            "activity_stream": "/api/v2/groups/21/activity_stream/",
            "hosts": "/api/v2/groups/21/hosts/",
            "children": "/api/v2/groups/21/children/",
            "inventory_sources": "/api/v2/groups/21/inventory_sources/",
            "inventory": "/api/v2/inventories/9/"
        },
        "summary_fields": {
            "inventory": {
                "id": 9,
                "name": "test",
                "description": "",
                "has_active_failures": false,
                "total_hosts": 8,
                "hosts_with_active_failures": 0,
                "total_groups": 4,
                "groups_with_active_failures": 0,
                "has_inventory_sources": false,
                "total_inventory_sources": 0,
                "inventory_sources_with_failures": 0,
                "organization_id": 16,
                "kind": ""
            },
            "created_by": {
                "id": 11,
                "username": "demouser",
                "first_name": "Demo",
                "last_name": "User"
            },
            "user_capabilities": {
                "edit": true,
                "copy": true,
                "delete": true
            }
        },
        "created": "2018-07-17T13:27:46.686176Z",
        "modified": "2018-07-17T13:28:07.127040Z",
        "name": "Demo Group",
        "description": "",
        "inventory": 9,
        "variables": "",
        "has_active_failures": false,
        "total_hosts": 3,
        "hosts_with_active_failures": 0,
        "total_groups": 0,
        "groups_with_active_failures": 0,
        "has_inventory_sources": false
    }`)
)
