# Project API

## Usage

> Project Updates Cancel

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    err := awx.ProjectUpdatesService.ProjectUpdateCancel(4)

    if err != nil {
        log.Fatalf("Cancel Update Projects err: %s", err)
    }

    log.Printf("Update Project cancelled.")
}
```

> Project Updates Get Update

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    err := awx.ProjectUpdatesService.ProjectUpdateGet(4)

    if err != nil {
        log.Fatalf("Get Update Projects err: %s", err)
    }

    log.Printf("Get Project done.")
}
```