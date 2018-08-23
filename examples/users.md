 # User API

## Usage

> List Users

```
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
```
import (
    "log" 
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.UserService.CreateUser(map[string]interface{}{
       "username":     "test",
       "description":  "for testing CreateUser api",
    }, map[string]string{})

    if err != nil {
        log.Fatalf("Create User err: %s", err)
    }

    log.Printf("User created. User ID: %d", result.User.Username")
}
```