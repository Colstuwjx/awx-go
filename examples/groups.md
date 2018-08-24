 # Group API

## Usage

> List Groups

```
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
```
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