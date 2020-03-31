# Teams API

## Usage

> List Teams

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, _, err := awx.TeamService.ListTeams(map[string]string{
    "name": "test-team",
    })
    if err != nil {
        log.Fatalf("List Teams err: %s", err)
    }

    log.Println("List Team: ", result)
}
```

> Create Team

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.TeamService.CreateTeam(map[string]interface{}{
        "name":         "test-team",
        "organization": 1,
        }, map[string]string{})
    if err != nil {
        log.Fatalf("Create Team err: %s", err)
    }

    log.Printf("Team created")
}
```

> Update Team

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.TeamService.UpdateTeam(4, map[string]interface{}{
        "name":         "test-team",
        "organization": 1,
        "description":  "Update test-team",
        }, map[string]string{})

    if err != nil {
        log.Fatalf("Update Team err: %s", err)
    }

    log.Printf("Update finised.")
}
```

> Delete Team

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.TeamService.DeleteTeam(1)

    if err != nil {
        log.Fatalf("Delete Team err: %s", err)
    }

    log.Printf("Team Deleted")
}
```

> Grant Team Role

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    err := awx.TeamService.GrantRole(1, 170)

    if err != nil {
        log.Fatalf("Grant user role err: %s", err)
    }

    log.Printf("User role granted")
}
```

> Revoke User Role

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    err := awx.TeamService.GrantRole(1, 170)

    if err != nil {
        log.Fatalf("Revoke user role err: %s", err)
    }

    log.Printf("User role revoked")
}
```
