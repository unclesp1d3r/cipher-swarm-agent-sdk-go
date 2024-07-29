<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"log"
	"os"
)

func main() {
	s := cipherswarmagentsdkgo.New(
		cipherswarmagentsdkgo.WithSecurity(os.Getenv("BEARER_AUTH")),
	)
	var id int64 = 135003
	ctx := context.Background()
	res, err := s.Agents.GetAgent(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	if res.Agent != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->