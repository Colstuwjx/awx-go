 # Ping API

## Usage

> List Inventories

```
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, _, err := awx.InventoriesService.ListInventories(map[string]string{
        "name": "Demo Inventory",
    })
    if err != nil {
        log.Fatalf("List Inventories err: %s", err)
    }

    log.Println("List Inventories: ", result)
}
```


