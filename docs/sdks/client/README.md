# Client
(*Client*)

## Overview

Client API

### Available Operations

* [GetConfiguration](#getconfiguration) - Get Agent Configuration
* [Authenticate](#authenticate) - Authenticate Client

## GetConfiguration

Returns the configuration for the agent.

### Example Usage

```go
package main

import(
	"os"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )

    ctx := context.Background()
    res, err := s.Client.GetConfiguration(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetConfigurationResponse](../../models/operations/getconfigurationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Authenticate

Authenticates the client. This is used to verify that the client is able to connect to the server.

### Example Usage

```go
package main

import(
	"os"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity(os.Getenv("BEARER_AUTH")),
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.AuthenticateResponse](../../models/operations/authenticateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
