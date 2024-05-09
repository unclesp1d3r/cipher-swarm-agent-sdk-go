# Attacks
(*Attacks*)

## Overview

Attacks API

### Available Operations

* [ShowAttack](#showattack) - show attack
* [HashListAttack](#hashlistattack) - Get the hash list

## ShowAttack

Returns an attack by id. This is used to get the details of an attack.

### Example Usage

```go
package main

import(
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    var id int64 = 69912
    
    ctx := context.Background()
    res, err := s.Attacks.ShowAttack(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *int64*                                                  | :heavy_check_mark:                                       | id                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*components.Attack](../../models/components/attack.md), error**
| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401,404               | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |

## HashListAttack

Returns the hash list for an attack.

### Example Usage

```go
package main

import(
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    var id int64 = 295812
    
    ctx := context.Background()
    res, err := s.Attacks.HashListAttack(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *int64*                                                  | :heavy_check_mark:                                       | id                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[io.ReadCloser](../../.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
