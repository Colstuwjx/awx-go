# User API

## Usage

> List Users

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, _, err := awx.UserService.ListUsers(map[string]string{
        "name": "Demo User",
    })
    if err != nil {
        log.Fatalf("List Users err: %s", err)
    }

    log.Println("List User: ", result)
}
```

> Create User

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.UserService.CreateUser(map[string]interface{}{
       "username":     "test",
       "description":  "for testing CreateUser api",
    }, map[string]string{})

    if err != nil {
        log.Fatalf("Create User err: %s", err)
    }

    log.Printf("User created. Username: %s", result.User.Username)
}
```

> Update User

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.UserService.UpdateUser(1, map[string]interface{}{
       "description":  "for testing Update api",
    }, map[string]string{})

    if err != nil {
        log.Fatalf("Update User err: %s", err)
    }

    log.Printf("Update finised. Description: %s", result.User.Description)
}
```

> Delete User

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.UserService.DeleteUser(1)

    if err != nil {
        log.Fatalf("Delete user err: %s", err)
    }

    log.Printf("User Deleted")
}
```

> Grant User Role

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.UserService.GrantRole(1, 24)

    if err != nil {
        log.Fatalf("Grant user role err: %s", err)
    }

    log.Printf("User role granted")
}
```

> Revoke User Role

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.UserService.RevokeRole(1, 24)

    if err != nil {
        log.Fatalf("Revoke user role err: %s", err)
    }

    log.Printf("User role revoked")
}
```