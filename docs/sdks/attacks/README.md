# Attacks
(*Attacks*)

## Overview

Attacks API

### Available Operations

* [GetAttack](#getattack) - show attack
* [GetHashList](#gethashlist) - Get the hash list

## GetAttack

Returns an attack by id. This is used to get the details of an attack.

### Example Usage

```go
package main

import(
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Attacks.GetAttack(ctx, 535539)
    if err != nil {
        log.Fatal(err)
    }
    if res.Attack != nil {
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

**[*operations.GetAttackResponse](../../models/operations/getattackresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401, 404              | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## GetHashList

Returns the hash list for an attack.

### Example Usage

```go
package main

import(
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Attacks.GetHashList(ctx, 904122)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
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

**[*operations.GetHashListResponse](../../models/operations/gethashlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |