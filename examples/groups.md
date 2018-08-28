# Group API

## Usage

> List Groups

```go
package main

import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, _, err := awx.GroupService.ListGroups(map[string]string{
        "name": "Demo Group",
    })
    if err != nil {
        log.Fatalf("List Groups err: %s", err)
    }

    log.Println("List Groups: ", result)
}
```

> Create Group

```go
package main

import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.GroupService.CreateGroup(map[string]interface{}{
       "name":     "test",
       "description":  "for testing CreateGroup api",
    }, map[string]string{})

    if err != nil {
        log.Fatalf("Create Group err: %s", err)
    }

    log.Printf("Group created. Group ID: %d", result.Group.ID")
}
```

> Update Group

```go
package main

import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.GroupService.UpdateGroup(21, map[string]interface{}{
        "description": "Add description here",
    }, map[string]string{})

    if err != nil {
        log.Fatalf("Update Group err: %s", err)
    }

    log.Printf("Updated Group ID: %d", 21)
}
```

> Delete Group

```go
package main

import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.GroupService.DeleteGroup(12)

    if err != nil {
        log.Fatalf("Delete Group err: %s", err)
    }

    log.Printf("Group deleted. Group ID: %d", 12)
}
```