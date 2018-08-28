# Job API

## Usage

> Get Job

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    var (
        yourJobId = 301
    )

    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.JobService.GetJob(yourJobId, map[string]string{})
    if err != nil {
        log.Fatalf("Get Job err: %s", err)
    }

    log.Println("Get job: ", result)
}
```

> Cancel Job

```
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    var (
        yourJobId = 301
    )

    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.JobService.CancelJob(yourJobId, map[string]interface{}{}, map[string]string{})
    if err != nil {
        log.Fatalf("Cancel Job err: %s", err)
    }

    log.Println("Cancel job: ", result)
}
```

> Relaunch Job

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    var (
        yourJobId = 301
    )

    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.JobService.RelaunchJob(yourJobId, map[string]interface{}{"hosts": "all"}, map[string]string{})
    if err != nil {
        log.Fatalf("Relaunch Job err: %s", err)
    }

    log.Println("Relaunch job: ", result)
}
```

> Get Host Summaries

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    var (
        yourJobId = 301
    )

    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, _, err := awx.JobService.GetHostSummaries(yourJobId, map[string]string{})
    if err != nil {
        log.Fatalf("Get Host Summaries err: %s", err)
    }

    log.Println("Get Host Summaries: ", result)
}
```

> Get Job Events

```
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    var (
        yourJobId = 301
    )

    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, _, err := awx.JobService.GetJobEvents(yourJobId, map[string]string{
        "order_by":  "start_line",
        "page_size": "1000000",
    })
    if err != nil {
        log.Fatalf("Get Job Events err: %s", err)
    }

    log.Println("Get Job Events: ", result)
}
```
