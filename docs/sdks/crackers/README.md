# Crackers
(*Crackers*)

## Overview

Crackers API

### Available Operations

* [CheckForCrackerUpdate](#checkforcrackerupdate) - Check for Cracker Update

## CheckForCrackerUpdate

Check for a cracker update, based on the operating system and version.

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
    var operatingSystem *string = cipherswarmagentsdkgo.String("<value>")

    var version *string = cipherswarmagentsdkgo.String("<value>")
    ctx := context.Background()
    res, err := s.Crackers.CheckForCrackerUpdate(ctx, operatingSystem, version)
    if err != nil {
        log.Fatal(err)
    }
    if res.CrackerUpdate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `operatingSystem`                                        | **string*                                                | :heavy_minus_sign:                                       | operating_system                                         |
| `version`                                                | **string*                                                | :heavy_minus_sign:                                       | version                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.CheckForCrackerUpdateResponse](../../models/operations/checkforcrackerupdateresponse.md), error**
| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 400                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |
