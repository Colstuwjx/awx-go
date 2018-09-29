# Project API

## Usage

> List Projects

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, _, err := awx.ProjectService.ListProjects(map[string]string{
        "name": "Demo Project",
    })
    if err != nil {
        log.Fatalf("List Projects err: %s", err)
    }

    log.Println("List Projects: ", result)
}
```

> Create Project

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.ProjectService.CreateProject(map[string]interface{}{
       "name":         "TestProject",
       "description":  "for testing CreateProject api",
       "organization": 1,
       "kind":         "",
       "host_filter":  "",
       "variables":    "",
    }, map[string]string{})

    if err != nil {
        log.Fatalf("Create Projects err: %s", err)
    }

    log.Printf("Project created. Project ID: %d", result.Project.ID)
}
```

> Update Project

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.ProjectService.UpdateProject(4, map[string]interface{}{
       "description":  "Update Example",
    }, map[string]string{})

    if err != nil {
        log.Fatalf("Update Projects err: %s", err)
    }

    log.Printf("Project Updated. Project ID: %d", result.ID)
}
```

> Delete Project

```go
package main
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.ProjectService.DeleteProject(4)

    if err != nil {
        log.Fatalf("Delete Projects err: %s", err)
    }

    log.Printf("Project Deleted. Project ID: %d", result.ID)
}
```
