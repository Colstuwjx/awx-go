# Job API

## Usage

```
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    job, _, err := awx.JobService.GetJob(map[string]string{})
    if err != nil {
        log.Fatalf("Get job err: %s", err)
    }

    log.Println("Get job: ", job)
}
```
