package mockdata

var (
	// MockedListHostsResponse mocked `ListHosts` endpoint response
	MockedListHostsResponse = []byte(`
        {
            "count": 1,
            "next": null,
            "previous": null,
            "results": [
                {
                    "id": 1,
                    "type": "host",
                    "url": "/api/v2/hosts/1/",
                    "related": {
                        "created_by": "/api/v2/users/2/",
                        "modified_by": "/api/v2/users/2/",
                        "job_host_summaries": "/api/v2/hosts/1/job_host_summaries/",
                        "variable_data": "/api/v2/hosts/1/variable_data/",
                        "job_events": "/api/v2/hosts/1/job_events/",
                        "ad_hoc_commands": "/api/v2/hosts/1/ad_hoc_commands/",
                        "inventory_sources": "/api/v2/hosts/1/inventory_sources/",
                        "fact_versions": "/api/v2/hosts/1/fact_versions/",
                        "smart_inventories": "/api/v2/hosts/1/smart_inventories/",
                        "groups": "/api/v2/hosts/1/groups/",
                        "activity_stream": "/api/v2/hosts/1/activity_stream/",
                        "all_groups": "/api/v2/hosts/1/all_groups/",
                        "ad_hoc_command_events": "/api/v2/hosts/1/ad_hoc_command_events/",
                        "insights": "/api/v2/hosts/1/insights/",
                        "inventory": "/api/v2/inventories/1/",
                        "ansible_facts": "/api/v2/hosts/1/ansible_facts/"
                    },
                    "summary_fields": {
                        "inventory": {
                            "id": 1,
                            "name": "Demo Inventory",
                            "description": "",
                            "has_active_failures": false,
                            "total_hosts": 1,
                            "hosts_with_active_failures": 0,
                            "total_groups": 2,
                            "groups_with_active_failures": 0,
                            "has_inventory_sources": false,
                            "total_inventory_sources": 0,
                            "inventory_sources_with_failures": 0,
                            "organization_id": 1,
                            "kind": ""
                        },
                        "created_by": {
                            "id": 2,
                            "username": "admin",
                            "first_name": "",
                            "last_name": ""
                        },
                        "modified_by": {
                            "id": 2,
                            "username": "admin",
                            "first_name": "",
                            "last_name": ""
                        },
                        "user_capabilities": {
                            "edit": true,
                            "delete": true
                        },
                        "groups": {
                            "count": 2,
                            "results": [
                                {
                                    "id": 19,
                                    "name": "ciao"
                                },
                                {
                                    "id": 21,
                                    "name": "test"
                                }
                            ]
                        },
                        "recent_jobs": []
                    },
                    "created": "2018-08-27T13:47:11.145028Z",
                    "modified": "2018-08-27T13:47:11.145042Z",
                    "name": "localhost",
                    "description": "",
                    "inventory": 1,
                    "enabled": true,
                    "instance_id": "",
                    "variables": "ansible_connection: local",
                    "has_active_failures": false,
                    "has_inventory_sources": false,
                    "last_job": null,
                    "last_job_host_summary": null,
                    "insights_system_id": null,
                    "ansible_facts_modified": null
                }
            ]
        }`)

	// MockedCreateHostResponse mocked `CreateHost` endpoint response
	MockedCreateHostResponse = []byte(`
        {
            "id": 3,
            "type": "host",
            "url": "/api/v2/hosts/3/",
            "related": {
                "named_url": "/api/v2/hosts/test++Demo Inventory++Default/",
                "created_by": "/api/v2/users/2/",
                "modified_by": "/api/v2/users/2/",
                "job_host_summaries": "/api/v2/hosts/3/job_host_summaries/",
                "variable_data": "/api/v2/hosts/3/variable_data/",
                "job_events": "/api/v2/hosts/3/job_events/",
                "ad_hoc_commands": "/api/v2/hosts/3/ad_hoc_commands/",
                "inventory_sources": "/api/v2/hosts/3/inventory_sources/",
                "fact_versions": "/api/v2/hosts/3/fact_versions/",
                "smart_inventories": "/api/v2/hosts/3/smart_inventories/",
                "groups": "/api/v2/hosts/3/groups/",
                "activity_stream": "/api/v2/hosts/3/activity_stream/",
                "all_groups": "/api/v2/hosts/3/all_groups/",
                "ad_hoc_command_events": "/api/v2/hosts/3/ad_hoc_command_events/",
                "insights": "/api/v2/hosts/3/insights/",
                "inventory": "/api/v2/inventories/1/",
                "ansible_facts": "/api/v2/hosts/3/ansible_facts/"
            },
            "summary_fields": {
                "inventory": {
                    "id": 1,
                    "name": "Demo Inventory",
                    "description": "",
                    "has_active_failures": false,
                    "total_hosts": 1,
                    "hosts_with_active_failures": 0,
                    "total_groups": 3,
                    "groups_with_active_failures": 0,
                    "has_inventory_sources": false,
                    "total_inventory_sources": 0,
                    "inventory_sources_with_failures": 0,
                    "organization_id": 1,
                    "kind": ""
                },
                "created_by": {
                    "id": 2,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "modified_by": {
                    "id": 2,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "user_capabilities": {
                    "edit": true,
                    "delete": true
                },
                "groups": {
                    "count": 0,
                    "results": []
                },
                "recent_jobs": []
            },
            "created": "2018-09-01T11:18:16.456501Z",
            "modified": "2018-09-01T11:18:16.456512Z",
            "name": "test",
            "description": "test create host",
            "inventory": 1,
            "enabled": true,
            "instance_id": "",
            "variables": "ansible_host: localhost",
            "has_active_failures": false,
            "has_inventory_sources": false,
            "last_job": null,
            "last_job_host_summary": null,
            "insights_system_id": null,
            "ansible_facts_modified": null
        }`)

	// MockedUpdateHostResponse mocked `UpdateHost` endpoint response
	MockedUpdateHostResponse = []byte(`
        {
            "id": 3,
            "type": "host",
            "url": "/api/v2/hosts/3/",
            "related": {
                "named_url": "/api/v2/hosts/testUpdate++Demo Inventory++Default/",
                "created_by": "/api/v2/users/2/",
                "modified_by": "/api/v2/users/2/",
                "job_host_summaries": "/api/v2/hosts/3/job_host_summaries/",
                "variable_data": "/api/v2/hosts/3/variable_data/",
                "job_events": "/api/v2/hosts/3/job_events/",
                "ad_hoc_commands": "/api/v2/hosts/3/ad_hoc_commands/",
                "inventory_sources": "/api/v2/hosts/3/inventory_sources/",
                "fact_versions": "/api/v2/hosts/3/fact_versions/",
                "smart_inventories": "/api/v2/hosts/3/smart_inventories/",
                "groups": "/api/v2/hosts/3/groups/",
                "activity_stream": "/api/v2/hosts/3/activity_stream/",
                "all_groups": "/api/v2/hosts/3/all_groups/",
                "ad_hoc_command_events": "/api/v2/hosts/3/ad_hoc_command_events/",
                "insights": "/api/v2/hosts/3/insights/",
                "inventory": "/api/v2/inventories/1/",
                "ansible_facts": "/api/v2/hosts/3/ansible_facts/"
            },
            "summary_fields": {
                "inventory": {
                    "id": 1,
                    "name": "Demo Inventory",
                    "description": "",
                    "has_active_failures": false,
                    "total_hosts": 1,
                    "hosts_with_active_failures": 0,
                    "total_groups": 3,
                    "groups_with_active_failures": 0,
                    "has_inventory_sources": false,
                    "total_inventory_sources": 0,
                    "inventory_sources_with_failures": 0,
                    "organization_id": 1,
                    "kind": ""
                },
                "created_by": {
                    "id": 2,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "modified_by": {
                    "id": 2,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "user_capabilities": {
                    "edit": true,
                    "delete": true
                },
                "groups": {
                    "count": 0,
                    "results": []
                },
                "recent_jobs": []
            },
            "created": "2018-09-01T11:18:16.456501Z",
            "modified": "2018-09-01T11:18:16.456512Z",
            "name": "testUpdate",
            "description": "test create host",
            "inventory": 1,
            "enabled": true,
            "instance_id": "",
            "variables": "ansible_host: localhost",
            "has_active_failures": false,
            "has_inventory_sources": false,
            "last_job": null,
            "last_job_host_summary": null,
            "insights_system_id": null,
            "ansible_facts_modified": null
        }`)

	// MockedAssociateGroupResponse mocked `AssociateGroup` endpoint response
	MockedAssociateGroupResponse = []byte(`
        {
            "id": 3,
            "type": "host",
            "url": "/api/v2/hosts/3/",
            "related": {
                "named_url": "/api/v2/hosts/testUpdate++Demo Inventory++Default/",
                "created_by": "/api/v2/users/2/",
                "modified_by": "/api/v2/users/2/",
                "job_host_summaries": "/api/v2/hosts/3/job_host_summaries/",
                "variable_data": "/api/v2/hosts/3/variable_data/",
                "job_events": "/api/v2/hosts/3/job_events/",
                "ad_hoc_commands": "/api/v2/hosts/3/ad_hoc_commands/",
                "inventory_sources": "/api/v2/hosts/3/inventory_sources/",
                "fact_versions": "/api/v2/hosts/3/fact_versions/",
                "smart_inventories": "/api/v2/hosts/3/smart_inventories/",
                "groups": "/api/v2/hosts/3/groups/",
                "activity_stream": "/api/v2/hosts/3/activity_stream/",
                "all_groups": "/api/v2/hosts/3/all_groups/",
                "ad_hoc_command_events": "/api/v2/hosts/3/ad_hoc_command_events/",
                "insights": "/api/v2/hosts/3/insights/",
                "inventory": "/api/v2/inventories/1/",
                "ansible_facts": "/api/v2/hosts/3/ansible_facts/"
            },
            "summary_fields": {
                "inventory": {
                    "id": 1,
                    "name": "Demo Inventory",
                    "description": "",
                    "has_active_failures": false,
                    "total_hosts": 1,
                    "hosts_with_active_failures": 0,
                    "total_groups": 3,
                    "groups_with_active_failures": 0,
                    "has_inventory_sources": false,
                    "total_inventory_sources": 0,
                    "inventory_sources_with_failures": 0,
                    "organization_id": 1,
                    "kind": ""
                },
                "created_by": {
                    "id": 2,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "modified_by": {
                    "id": 2,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "user_capabilities": {
                    "edit": true,
                    "delete": true
                },
                "groups": {
                    "count": 1,
                    "results": [
                        { "id": 10,
                          "name": "testGroup"
                        }
                    ]
                },
                "recent_jobs": []
            },
            "created": "2018-09-01T11:18:16.456501Z",
            "modified": "2018-09-01T11:18:16.456512Z",
            "name": "testUpdate",
            "description": "test create host",
            "inventory": 1,
            "enabled": true,
            "instance_id": "",
            "variables": "ansible_host: localhost",
            "has_active_failures": false,
            "has_inventory_sources": false,
            "last_job": null,
            "last_job_host_summary": null,
            "insights_system_id": null,
            "ansible_facts_modified": null
        }`)

	// MockedDisAssociateGroupResponse mocked `DisAssociateGroup` endpoint response
	MockedDisAssociateGroupResponse = []byte(`
    {
        "id": 3,
        "type": "host",
        "url": "/api/v2/hosts/3/",
        "related": {
            "named_url": "/api/v2/hosts/testUpdate++Demo Inventory++Default/",
            "created_by": "/api/v2/users/2/",
            "modified_by": "/api/v2/users/2/",
            "job_host_summaries": "/api/v2/hosts/3/job_host_summaries/",
            "variable_data": "/api/v2/hosts/3/variable_data/",
            "job_events": "/api/v2/hosts/3/job_events/",
            "ad_hoc_commands": "/api/v2/hosts/3/ad_hoc_commands/",
            "inventory_sources": "/api/v2/hosts/3/inventory_sources/",
            "fact_versions": "/api/v2/hosts/3/fact_versions/",
            "smart_inventories": "/api/v2/hosts/3/smart_inventories/",
            "groups": "/api/v2/hosts/3/groups/",
            "activity_stream": "/api/v2/hosts/3/activity_stream/",
            "all_groups": "/api/v2/hosts/3/all_groups/",
            "ad_hoc_command_events": "/api/v2/hosts/3/ad_hoc_command_events/",
            "insights": "/api/v2/hosts/3/insights/",
            "inventory": "/api/v2/inventories/1/",
            "ansible_facts": "/api/v2/hosts/3/ansible_facts/"
        },
        "summary_fields": {
            "inventory": {
                "id": 1,
                "name": "Demo Inventory",
                "description": "",
                "has_active_failures": false,
                "total_hosts": 1,
                "hosts_with_active_failures": 0,
                "total_groups": 3,
                "groups_with_active_failures": 0,
                "has_inventory_sources": false,
                "total_inventory_sources": 0,
                "inventory_sources_with_failures": 0,
                "organization_id": 1,
                "kind": ""
            },
            "created_by": {
                "id": 2,
                "username": "admin",
                "first_name": "",
                "last_name": ""
            },
            "modified_by": {
                "id": 2,
                "username": "admin",
                "first_name": "",
                "last_name": ""
            },
            "user_capabilities": {
                "edit": true,
                "delete": true
            },
            "groups": {
                "count": 0,
                "results": []
            },
            "recent_jobs": []
        },
        "created": "2018-09-01T11:18:16.456501Z",
        "modified": "2018-09-01T11:18:16.456512Z",
        "name": "testUpdate",
        "description": "test create host",
        "inventory": 1,
        "enabled": true,
        "instance_id": "",
        "variables": "ansible_host: localhost",
        "has_active_failures": false,
        "has_inventory_sources": false,
        "last_job": null,
        "last_job_host_summary": null,
        "insights_system_id": null,
        "ansible_facts_modified": null
    }`)

	// MockedDeleteHostResponse mocked `DeleteHost` endpoint response
	MockedDeleteHostResponse = []byte(`{}`)
)
