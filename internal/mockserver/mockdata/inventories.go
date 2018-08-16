package mockdata

var (
	// MockedListInventoriesResponse mocked `ListInventories` endpoint response
	MockedListInventoriesResponse = []byte(`
{
    "count": 1,
    "next": null,
    "previous": null,
    "results": [
        {
            "id": 1,
            "type": "inventory",
            "url": "/api/v2/inventories/1/",
            "related": {
                "created_by": "/api/v2/users/1/",
                "modified_by": "/api/v2/users/1/",
                "job_templates": "/api/v2/inventories/1/job_templates/",
                "variable_data": "/api/v2/inventories/1/variable_data/",
                "root_groups": "/api/v2/inventories/1/root_groups/",
                "object_roles": "/api/v2/inventories/1/object_roles/",
                "ad_hoc_commands": "/api/v2/inventories/1/ad_hoc_commands/",
                "script": "/api/v2/inventories/1/script/",
                "tree": "/api/v2/inventories/1/tree/",
                "access_list": "/api/v2/inventories/1/access_list/",
                "activity_stream": "/api/v2/inventories/1/activity_stream/",
                "instance_groups": "/api/v2/inventories/1/instance_groups/",
                "hosts": "/api/v2/inventories/1/hosts/",
                "groups": "/api/v2/inventories/1/groups/",
                "copy": "/api/v2/inventories/1/copy/",
                "update_inventory_sources": "/api/v2/inventories/1/update_inventory_sources/",
                "inventory_sources": "/api/v2/inventories/1/inventory_sources/",
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
                    "first_name": "",
                    "last_name": ""
                },
                "modified_by": {
                    "id": 1,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "object_roles": {
                    "use_role": {
                        "id": 23,
                        "description": "Can use the inventory in a job template",
                        "name": "Use"
                    },
                    "admin_role": {
                        "id": 21,
                        "description": "Can manage all aspects of the inventory",
                        "name": "Admin"
                    },
                    "adhoc_role": {
                        "id": 20,
                        "description": "May run ad hoc commands on an inventory",
                        "name": "Ad Hoc"
                    },
                    "update_role": {
                        "id": 24,
                        "description": "May update project or inventory or group using the configured source update system",
                        "name": "Update"
                    },
                    "read_role": {
                        "id": 22,
                        "description": "May view settings for the inventory",
                        "name": "Read"
                    }
                },
                "user_capabilities": {
                    "edit": true,
                    "copy": true,
                    "adhoc": true,
                    "delete": true
                }
            },
            "created": "2018-05-21T01:34:35.657185Z",
            "modified": "2018-05-30T09:42:22.412749Z",
            "name": "Demo Inventory",
            "description": "",
            "organization": 1,
            "kind": "",
            "host_filter": null,
            "variables": "",
            "has_active_failures": false,
            "total_hosts": 2,
            "hosts_with_active_failures": 0,
            "total_groups": 0,
            "groups_with_active_failures": 0,
            "has_inventory_sources": false,
            "total_inventory_sources": 0,
            "inventory_sources_with_failures": 0,
            "insights_credential": null,
            "pending_deletion": false
        }
    ]
}`)

	// MockedCreateInventoryResponse mocked `CreateInventory` endpoint response
	MockedCreateInventoryResponse = []byte(`
{
    "id": 6,
    "type": "inventory",
    "url": "/api/v2/inventories/6/",
    "related": {
        "named_url": "/api/v2/inventories/TestInventory++Default/",
        "created_by": "/api/v2/users/1/",
        "modified_by": "/api/v2/users/1/",
        "job_templates": "/api/v2/inventories/6/job_templates/",
        "variable_data": "/api/v2/inventories/6/variable_data/",
        "root_groups": "/api/v2/inventories/6/root_groups/",
        "object_roles": "/api/v2/inventories/6/object_roles/",
        "ad_hoc_commands": "/api/v2/inventories/6/ad_hoc_commands/",
        "script": "/api/v2/inventories/6/script/",
        "tree": "/api/v2/inventories/6/tree/",
        "access_list": "/api/v2/inventories/6/access_list/",
        "activity_stream": "/api/v2/inventories/6/activity_stream/",
        "instance_groups": "/api/v2/inventories/6/instance_groups/",
        "hosts": "/api/v2/inventories/6/hosts/",
        "groups": "/api/v2/inventories/6/groups/",
        "copy": "/api/v2/inventories/6/copy/",
        "update_inventory_sources": "/api/v2/inventories/6/update_inventory_sources/",
        "inventory_sources": "/api/v2/inventories/6/inventory_sources/",
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
            "first_name": "",
            "last_name": ""
        },
        "modified_by": {
            "id": 1,
            "username": "admin",
            "first_name": "",
            "last_name": ""
        },
        "object_roles": {
            "use_role": {
                "id": 80,
                "description": "Can use the inventory in a job template",
                "name": "Use"
            },
            "admin_role": {
                "id": 78,
                "description": "Can manage all aspects of the inventory",
                "name": "Admin"
            },
            "adhoc_role": {
                "id": 77,
                "description": "May run ad hoc commands on an inventory",
                "name": "Ad Hoc"
            },
            "update_role": {
                "id": 81,
                "description": "May update project or inventory or group using the configured source update system",
                "name": "Update"
            },
            "read_role": {
                "id": 79,
                "description": "May view settings for the inventory",
                "name": "Read"
            }
        },
        "user_capabilities": {
            "edit": true,
            "copy": true,
            "adhoc": true,
            "delete": true
        }
    },
    "created": "2018-08-13T01:59:47.160127Z",
    "modified": "2018-08-13T01:59:47.160140Z",
    "name": "TestInventory",
    "description": "for testing CreateInventory api",
    "organization": 1,
    "kind": "",
    "host_filter": null,
    "variables": "",
    "has_active_failures": false,
    "total_hosts": 0,
    "hosts_with_active_failures": 0,
    "total_groups": 0,
    "groups_with_active_failures": 0,
    "has_inventory_sources": false,
    "total_inventory_sources": 0,
    "inventory_sources_with_failures": 0,
    "insights_credential": null,
    "pending_deletion": false
}`)

	// MockedUpdateInventoryResponse mocked `UpdateInventory` endpoint response
	MockedUpdateInventoryResponse = []byte(`
{
        "id": 6,
        "type": "inventory",
        "url": "/api/v2/inventories/6/",
        "related": {
            "named_url": "/api/v2/inventories/TestInventory-update1++Default/",
            "created_by": "/api/v2/users/1/",
            "modified_by": "/api/v2/users/1/",
            "job_templates": "/api/v2/inventories/6/job_templates/",
            "variable_data": "/api/v2/inventories/6/variable_data/",
            "root_groups": "/api/v2/inventories/6/root_groups/",
            "object_roles": "/api/v2/inventories/6/object_roles/",
            "ad_hoc_commands": "/api/v2/inventories/6/ad_hoc_commands/",
            "script": "/api/v2/inventories/6/script/",
            "tree": "/api/v2/inventories/6/tree/",
            "access_list": "/api/v2/inventories/6/access_list/",
            "activity_stream": "/api/v2/inventories/6/activity_stream/",
            "instance_groups": "/api/v2/inventories/6/instance_groups/",
            "hosts": "/api/v2/inventories/6/hosts/",
            "groups": "/api/v2/inventories/6/groups/",
            "copy": "/api/v2/inventories/6/copy/",
            "update_inventory_sources": "/api/v2/inventories/6/update_inventory_sources/",
            "inventory_sources": "/api/v2/inventories/6/inventory_sources/",
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
                "first_name": "",
                "last_name": ""
            },
            "modified_by": {
                "id": 1,
                "username": "admin",
                "first_name": "",
                "last_name": ""
            },
            "object_roles": {
                "use_role": {
                    "id": 80,
                    "description": "Can use the inventory in a job template",
                    "name": "Use"
                },
                "admin_role": {
                    "id": 78,
                    "description": "Can manage all aspects of the inventory",
                    "name": "Admin"
                },
                "adhoc_role": {
                    "id": 77,
                    "description": "May run ad hoc commands on an inventory",
                    "name": "Ad Hoc"
                },
                "update_role": {
                    "id": 81,
                    "description": "May update project or inventory or group using the configured source update system",
                    "name": "Update"
                },
                "read_role": {
                    "id": 79,
                    "description": "May view settings for the inventory",
                    "name": "Read"
                }
            },
            "user_capabilities": {
                "edit": true,
                "copy": true,
                "adhoc": true,
                "delete": true
            }
        },
        "created": "2018-08-13T01:59:47.160127Z",
        "modified": "2018-08-13T01:59:47.160140Z",
        "name": "TestInventory-update1",
        "description": "for testing UpdateInventory api",
        "organization": 1,
        "kind": "",
        "host_filter": null,
        "variables": "",
        "has_active_failures": false,
        "total_hosts": 0,
        "hosts_with_active_failures": 0,
        "total_groups": 0,
        "groups_with_active_failures": 0,
        "has_inventory_sources": false,
        "total_inventory_sources": 0,
        "inventory_sources_with_failures": 0,
        "insights_credential": null,
        "pending_deletion": false
    }`)

	// MockedDeleteInventoryResponse mocked `Delete` endpoint response
	MockedDeleteInventoryResponse = []byte(`{}`)

	// MockedGetInventoryByIDResponse mocked `GetInventoryByID` endpoint response
	MockedGetInventoryByIDResponse = []byte(`
        {
            "id": 1,
            "type": "inventory",
            "url": "/api/v2/inventories/1/",
            "related": {
                "created_by": "/api/v2/users/1/",
                "modified_by": "/api/v2/users/1/",
                "job_templates": "/api/v2/inventories/1/job_templates/",
                "variable_data": "/api/v2/inventories/1/variable_data/",
                "root_groups": "/api/v2/inventories/1/root_groups/",
                "object_roles": "/api/v2/inventories/1/object_roles/",
                "ad_hoc_commands": "/api/v2/inventories/1/ad_hoc_commands/",
                "script": "/api/v2/inventories/1/script/",
                "tree": "/api/v2/inventories/1/tree/",
                "access_list": "/api/v2/inventories/1/access_list/",
                "activity_stream": "/api/v2/inventories/1/activity_stream/",
                "instance_groups": "/api/v2/inventories/1/instance_groups/",
                "hosts": "/api/v2/inventories/1/hosts/",
                "groups": "/api/v2/inventories/1/groups/",
                "copy": "/api/v2/inventories/1/copy/",
                "update_inventory_sources": "/api/v2/inventories/1/update_inventory_sources/",
                "inventory_sources": "/api/v2/inventories/1/inventory_sources/",
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
                    "first_name": "",
                    "last_name": ""
                },
                "modified_by": {
                    "id": 1,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "object_roles": {
                    "use_role": {
                        "id": 80,
                        "description": "Can use the inventory in a job template",
                        "name": "Use"
                    },
                    "admin_role": {
                        "id": 78,
                        "description": "Can manage all aspects of the inventory",
                        "name": "Admin"
                    },
                    "adhoc_role": {
                        "id": 77,
                        "description": "May run ad hoc commands on an inventory",
                        "name": "Ad Hoc"
                    },
                    "update_role": {
                        "id": 81,
                        "description": "May update project or inventory or group using the configured source update system",
                        "name": "Update"
                    },
                    "read_role": {
                        "id": 79,
                        "description": "May view settings for the inventory",
                        "name": "Read"
                    }
                },
                "user_capabilities": {
                    "edit": true,
                    "copy": true,
                    "adhoc": true,
                    "delete": true
                }
            },
            "created": "2018-05-21T01:34:35.657185Z",
            "modified": "2018-05-30T09:42:22.412749Z",
            "name": "Demo Inventory",
            "description": "",
            "organization": 1,
            "kind": "",
            "host_filter": null,
            "variables": "",
            "has_active_failures": false,
            "total_hosts": 2,
            "hosts_with_active_failures": 0,
            "total_groups": 0,
            "groups_with_active_failures": 0,
            "has_inventory_sources": false,
            "total_inventory_sources": 0,
            "inventory_sources_with_failures": 0,
            "insights_credential": null,
            "pending_deletion": false
        }`)
)
