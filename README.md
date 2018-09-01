# awx-go

[![Build Status](https://travis-ci.org/Colstuwjx/awx-go.svg?branch=master)](https://travis-ci.org/Colstuwjx/awx-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Colstuwjx/awx-go)](https://goreportcard.com/report/github.com/Colstuwjx/awx-go)
[![codecov](https://codecov.io/gh/Colstuwjx/awx-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Colstuwjx/awx-go)

AWX SDK for the Go programming language.

![AWX-GO-ROBOT](images/awx-go-robot.png)

## Installing

If you are using Go 1.5 with the GO15VENDOREXPERIMENT=1 vendoring flag, or 1.6 and higher you can use the following command to retrieve the SDK. The SDK will be included.

```
go get -u github.com/Colstuwjx/awx-go
```

## Example

We can simply import awx-go and call its services, such as PingService:

```
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.PingService.Ping()
    if err != nil {
        log.Fatalf("Ping awx err: %s", err)
    }

    log.Println("Ping awx: ", result)
}
```

More examples could be found at [here](https://github.com/Colstuwjx/awx-go/tree/master/examples).

## Roadmap

awx-go is still in developing, and it's roadmap could be found at [here](https://github.com/Colstuwjx/awx-go/blob/master/ROADMAP.md).

## Contribute

There are many ways to contribute to awx-go.

* Submit bugs via [Github issues](https://github.com/Colstuwjx/awx-go/issues);
* Take [pull request](https://github.com/Colstuwjx/awx-go/pulls) for fixes or features;
* Mail [me](mailto:wjx_colstu@hotmail.com)
