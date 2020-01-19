go-backlog
====

go library for backlog api.  
https://developer.nulab-inc.com/ja/docs/backlog/

## Installation
```bash
go get -u github.com/ksmt88/go-backlog
```

## Getting Started
```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ksmt88/go-backlog"
)

func main() {
	config := backlog.Configure{
		SpaceId: "Enter SpaceId",
		ApiKey:  "Enter ApiKey",
        Domain:  backlog.DomainJp,
	}
	client, err := backlog.NewClient(config, http.DefaultClient)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	space, err := client.GetSpace()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(space)
}
```
