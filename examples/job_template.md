# Job Template API

## Usage

> List Job Templates

```
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, _, err := awx.JobTemplateService.ListJobTemplates(map[string]string{})
    if err != nil {
        log.Fatalf("List Job Templates err: %s", err)
    }

    log.Println("List Job Templates: ", result)
}
```

> Launch Job Template

```
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    var (
        yourJobTemplateId = 1
        yourInventoryId = 1
    )

    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.JobTemplateService.Launch(yourJobTemplateId, map[string]interface{}{
        "inventory": yourInventoryId,
    }, map[string]string{})
    if err != nil {
        log.Fatalf("Lauch err: %s", err)
    }

    log.Println("Launch Job Template: ", result)
}
```
