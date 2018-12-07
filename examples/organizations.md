# Organizations API

## Usage

> List Organizations

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, _, err := awx.OrganizationService.ListOrganizations(map[string]string{
    "name": "test-organization",
    })
    if err != nil {
        log.Fatalf("List Organizations err: %s", err)
    }

    log.Println("List Organization: ", result)
}
```

> Create Organization

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.OrganizationService.CreateOrganization(map[string]interface{}{
        "name":         "test-organization",
        "description": "test organization",
        }, map[string]string{})
    if err != nil {
        log.Fatalf("Create Organization err: %s", err)
    }

    log.Printf("Organization created")
}
```

> Update Organization

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.OrganizationService.UpdateOrganization(4, map[string]interface{}{
        "description":  "Update test-organization",
        }, map[string]string{})

    if err != nil {
        log.Fatalf("Update Organization err: %s", err)
    }

    log.Printf("Update finised.")
}
```

> Delete Organization

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.OrganizationService.DeleteOrganization(1)

    if err != nil {
        log.Fatalf("Delete Organization err: %s", err)
    }

    log.Printf("Organization Deleted")
}
```

> Grant Organization Role

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    err := awx.OrganizationService.GrantRole(1, 170)

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
    err := awx.OrganizationService.GrantRole(1, 170)

    if err != nil {
        log.Fatalf("Revoke user role err: %s", err)
    }

    log.Printf("User role revoked")
}
```
