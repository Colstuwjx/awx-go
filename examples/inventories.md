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


> Create Inventory
```
import (
    "log" 
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.InventoriesService.CreateInventory(map[string]interface{}{
       "name":         "TestInventory",
       "description":  "for testing CreateInventory api",
       "organization": 1,
       "kind":         "",
       "host_filter":  "",
       "variables":    "",
    }, map[string]string{})

    if err != nil {
        log.Fatalf("Create Inventories err: %s", err)
    }

    log.Println("Inventory created. Inventory ID: %s", result.Inventory.ID")
}
```

> Update Inventory

```
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.InventoriesService.UpdateInventory(map[string]interface{}{
       "name":         "TestInventory - 1",
       "description":  "for testing UpdateInventory api",
       "organization": 1,
       "kind":         "",
       "host_filter":  "",
       "variables":    "",
    }, nil, "1")

    if err != nil {
        log.Fatalf("Update Inventories err: %s", err)
    }

    log.Printf("Update result %v", result.Name)

}
```
> Delete Inventory
```
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.InventoriesService.Delete("5")
    if err != nil {
        log.Fatalf("Delete Inventories err: %s", err)
    }

    log.Println("Inventroy 5 Deleted")

}
```

> GetInventoryByName
```
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.InventoriesService.GetInventoryByName(map[string]string{"name": "Demo Inventory"})
    if err != nil {
        log.Fatalf("Get Inventory by Name err: %s", err)
    }

    log.Printf("Demo Inventory id %d\n", result.ID)

}
```