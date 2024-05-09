# Agents
(*Agents*)

## Overview

Agents API

### Available Operations

* [ShowAgent](#showagent) - Gets an instance of an agent
* [UpdateAgent](#updateagent) - Updates the agent
* [HeartbeatAgent](#heartbeatagent) - Send a heartbeat for an agent
* [LastBenchmarkAgent](#lastbenchmarkagent) - last_benchmark agent
* [SubmitBenchmarkAgent](#submitbenchmarkagent) - submit_benchmark agent

## ShowAgent

Returns an agent

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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *int64*                                               | :heavy_check_mark:                                    | id                                                    |


### Response

**[*operations.ShowAgentResponse](../../models/operations/showagentresponse.md), error**
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
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    var id int64 = 828119

    var agentUpdate *components.AgentUpdate = &components.AgentUpdate{
        ID: 182255,
        Name: "<value>",
        ClientSignature: "<value>",
        OperatingSystem: "<value>",
        Devices: []string{
            "<value>",
        },
    }
    
    ctx := context.Background()
    res, err := s.Agents.UpdateAgent(ctx, id, agentUpdate)
    if err != nil {
        log.Fatal(err)
    }
    if res.Agent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `id`                                                              | *int64*                                                           | :heavy_check_mark:                                                | id                                                                |
| `agentUpdate`                                                     | [*components.AgentUpdate](../../models/components/agentupdate.md) | :heavy_minus_sign:                                                | N/A                                                               |


### Response

**[*operations.UpdateAgentResponse](../../models/operations/updateagentresponse.md), error**
| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |

## HeartbeatAgent

Send a heartbeat for an agent to keep it alive.

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

    var id int64 = 397037
    
    ctx := context.Background()
    res, err := s.Agents.HeartbeatAgent(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *int64*                                               | :heavy_check_mark:                                    | id                                                    |


### Response

**[*operations.HeartbeatAgentResponse](../../models/operations/heartbeatagentresponse.md), error**
| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |

## LastBenchmarkAgent

Returns the last benchmark date for an agent

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

    var id int64 = 714635
    
    ctx := context.Background()
    res, err := s.Agents.LastBenchmarkAgent(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentLastBenchmark != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *int64*                                               | :heavy_check_mark:                                    | id                                                    |


### Response

**[*operations.LastBenchmarkAgentResponse](../../models/operations/lastbenchmarkagentresponse.md), error**
| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |

## SubmitBenchmarkAgent

Submit a benchmark for an agent

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

    var id int64 = 946448

    var requestBody []components.HashcatBenchmark = []components.HashcatBenchmark{
        components.HashcatBenchmark{
            HashType: 268990,
            Runtime: 493062,
            HashSpeed: 1512.35,
            Device: 572344,
        },
    }
    
    ctx := context.Background()
    res, err := s.Agents.SubmitBenchmarkAgent(ctx, id, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `id`                                                                         | *int64*                                                                      | :heavy_check_mark:                                                           | id                                                                           |
| `requestBody`                                                                | [][components.HashcatBenchmark](../../models/components/hashcatbenchmark.md) | :heavy_minus_sign:                                                           | N/A                                                                          |


### Response

**[*operations.SubmitBenchmarkAgentResponse](../../models/operations/submitbenchmarkagentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
