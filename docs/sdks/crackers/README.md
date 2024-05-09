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
	"github.com/unclesp1d3r/cipherswarm-agent-sdk/models/components"
	cipherswarmagentsdk "github.com/unclesp1d3r/cipherswarm-agent-sdk"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdk.New(
        cipherswarmagentsdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    var operatingSystem *string = cipherswarmagentsdk.String("<value>")

    var version *string = cipherswarmagentsdk.String("<value>")
    
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `operatingSystem`                                     | **string*                                             | :heavy_minus_sign:                                    | operating_system                                      |
| `version`                                             | **string*                                             | :heavy_minus_sign:                                    | version                                               |


### Response

**[*operations.CheckForCrackerUpdateResponse](../../models/operations/checkforcrackerupdateresponse.md), error**
| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 400                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |
