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

> Create Job Template
```
import (
	"log"

    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
	
    result, err := awx.JobTemplateService.CreateJobTemplate(map[string]interface{}{
		"name":        "Example Create Job Template",
		"description": "Created from awx-go Example",
		"job_type":    "run",
		"inventory":   1,
		"project":     1,
		"playbook":    "playbook.yml",
		"verbosity":   0,
	}, map[string]string{})

	if err != nil {
		log.Fatalf("Create Inventories err: %s", err)
	}

	log.Printf("Inventory created. Inventory ID: %d", result)
}

```