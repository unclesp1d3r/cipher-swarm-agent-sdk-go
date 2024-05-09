# Client
(*Client*)

## Overview

Client API

### Available Operations

* [Configuration](#configuration) - Get Agent Configuration
* [Authenticate](#authenticate) - Authenticate Client

## Configuration

Returns the configuration for the agent.

### Example Usage

```go
package main

import(
	"github.com/unclesp1d3r/cipherswarm-agent-sdk/models/components"
	cipherswarmagentsdk "github.com/unclesp1d3r/cipherswarm-agent-sdk"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdk.New(
        cipherswarmagentsdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    
    ctx := context.Background()
    res, err := s.Client.Configuration(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ConfigurationResponse](../../models/operations/configurationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Authenticate

Authenticates the client. This is used to verify that the client is able to connect to the server.

### Example Usage

```go
package main

import(
	"github.com/unclesp1d3r/cipherswarm-agent-sdk/models/components"
	cipherswarmagentsdk "github.com/unclesp1d3r/cipherswarm-agent-sdk"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdk.New(
        cipherswarmagentsdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    
    ctx := context.Background()
    res, err := s.Client.Authenticate(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthenticationResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.AuthenticateResponse](../../models/operations/authenticateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
