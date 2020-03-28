# Credentiales API

## Usage

> List Credentiales

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, _, err := awx.CredentialService.ListCredentiales(map[string]string{
        "name": "Demo Credential",
    })
    if err != nil {
        log.Fatalf("List Credentiales err: %s", err)
    }

    log.Println("List Credentiales: ", result)
}
```

> Create Credential

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.CredentialService.CreateCredential(map[string]interface{}{
        "name":            "Demo Credential",
        "organization":    1,
        "credential_type": 1,
    }, map[string]string{})

    if err != nil {
        log.Fatalf("Create Credentiales err: %s", err)
    }

    log.Printf("Credential created. Credential ID: %d", result.Credential.ID")
}
```

> Update Credential

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

fun main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.CredentialService.UpdateCredential(1, map[string]interface{}{
       "name":            "Demo Credential",
       "description":     "Demo credential",
       "organization":    1,
       "credential_type": 1,
       "inputs": map[string]interface{}{
           "username": "admin",
           },
    }, nil)

    if err != nil {
        log.Fatalf("Update Credentiales err: %s", err)
    }

    log.Printf("Update result %v", result.Name)

}
```

> Delete Credential

```go
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.CredentialService.DeleteCredential(5)
    if err != nil {
        log.Fatalf("Delete Credentiales err: %s", err)
    }

    log.Println("Credential 5 Deleted")

}
```