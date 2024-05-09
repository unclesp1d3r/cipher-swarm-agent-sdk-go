<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	cipherswarmagentsdk "github.com/unclesp1d3r/cipherswarm-agent-sdk"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk/models/components"
	"log"
)

func main() {
	s := cipherswarmagentsdk.New(
		cipherswarmagentsdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	var id int64 = 969902

	ctx := context.Background()
	res, err := s.Agents.ShowAgent(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	if res.Agent != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->