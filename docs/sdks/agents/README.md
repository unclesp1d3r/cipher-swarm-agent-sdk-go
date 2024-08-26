# Agents
(*Agents*)

## Overview

Agents API

### Available Operations

* [GetAgent](#getagent) - Gets an instance of an agent
* [UpdateAgent](#updateagent) - Updates the agent
* [SendHeartbeat](#sendheartbeat) - Send a heartbeat for an agent
* [SubmitBenchmark](#submitbenchmark) - submit agent benchmarks
* [SubmitErrorAgent](#submiterroragent) - Submit an error for an agent
* [SetAgentShutdown](#setagentshutdown) - shutdown agent

## GetAgent

Returns an agent

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

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *int64*                                                  | :heavy_check_mark:                                       | id                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAgentResponse](../../models/operations/getagentresponse.md), error**

### Errors

| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |


## UpdateAgent

Updates an agent

### Example Usage

```go
package main

import(
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/operations"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var id int64 = 828119

    var requestBody *operations.UpdateAgentRequestBody = &operations.UpdateAgentRequestBody{
        ID: 96,
        Name: "effertz-frami",
        ClientSignature: "Plan 9, Plan 9 Fourth Edition",
        OperatingSystem: "linux",
        Devices: []string{
            "cpu",
            "GPU",
        },
    }
    ctx := context.Background()
    res, err := s.Agents.UpdateAgent(ctx, id, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Agent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |
| `id`                                                                                    | *int64*                                                                                 | :heavy_check_mark:                                                                      | id                                                                                      |
| `requestBody`                                                                           | [*operations.UpdateAgentRequestBody](../../models/operations/updateagentrequestbody.md) | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |

### Response

**[*operations.UpdateAgentResponse](../../models/operations/updateagentresponse.md), error**

### Errors

| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |


## SendHeartbeat

Send a heartbeat for an agent to keep it alive.

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
    var id int64 = 992386
    ctx := context.Background()
    res, err := s.Agents.SendHeartbeat(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.SendHeartbeatResponse](../../models/operations/sendheartbeatresponse.md), error**

### Errors

| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |


## SubmitBenchmark

Submit a benchmark for an agent

### Example Usage

```go
package main

import(
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/operations"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var id int64 = 303399

    requestBody := operations.SubmitBenchmarkRequestBody{
        HashcatBenchmarks: []components.HashcatBenchmark{
            components.HashcatBenchmark{
                HashType: 442220,
                Runtime: 8499,
                HashSpeed: 156.49,
                Device: 322052,
            },
        },
    }
    ctx := context.Background()
    res, err := s.Agents.SubmitBenchmark(ctx, id, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `id`                                                                                           | *int64*                                                                                        | :heavy_check_mark:                                                                             | id                                                                                             |
| `requestBody`                                                                                  | [operations.SubmitBenchmarkRequestBody](../../models/operations/submitbenchmarkrequestbody.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.SubmitBenchmarkResponse](../../models/operations/submitbenchmarkresponse.md), error**

### Errors

| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 400,401               | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |


## SubmitErrorAgent

Submit an error for an agent

### Example Usage

```go
package main

import(
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/operations"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var id int64 = 607526
    ctx := context.Background()
    res, err := s.Agents.SubmitErrorAgent(ctx, id, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `id`                                                                                              | *int64*                                                                                           | :heavy_check_mark:                                                                                | id                                                                                                |
| `requestBody`                                                                                     | [*operations.SubmitErrorAgentRequestBody](../../models/operations/submiterroragentrequestbody.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |

### Response

**[*operations.SubmitErrorAgentResponse](../../models/operations/submiterroragentresponse.md), error**

### Errors

| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401,404               | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |


## SetAgentShutdown

Marks the agent as shutdown and offline, freeing any assigned tasks.

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
    var id int64 = 811605
    ctx := context.Background()
    res, err := s.Agents.SetAgentShutdown(ctx, id)
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

**[*operations.SetAgentShutdownResponse](../../models/operations/setagentshutdownresponse.md), error**

### Errors

| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |
