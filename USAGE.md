<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"log"
)

func main() {
	s := cipherswarmagentsdkgo.New(
		cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Agents.GetAgent(ctx, 135003)
	if err != nil {
		log.Fatal(err)
	}
	if res.Agent != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->