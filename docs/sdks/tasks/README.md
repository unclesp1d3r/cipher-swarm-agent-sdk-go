# Tasks
(*Tasks*)

## Overview

Tasks API

### Available Operations

* [GetNewTask](#getnewtask) - Request a new task from server
* [GetTask](#gettask) - Request the task information
* [SendCrack](#sendcrack) - Submit a cracked hash result for a task
* [SendStatus](#sendstatus) - Submit a status update for a task
* [SetTaskAccepted](#settaskaccepted) - Accept Task
* [SetTaskExhausted](#settaskexhausted) - Notify of Exhausted Task
* [SetTaskAbandoned](#settaskabandoned) - Abandon Task
* [GetTaskZaps](#gettaskzaps) - Get Completed Hashes

## GetNewTask

Request a new task from the server, if available.

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
    res, err := s.Tasks.GetNewTask(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Task != nil {
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

**[*operations.GetNewTaskResponse](../../models/operations/getnewtaskresponse.md), error**

### Errors

| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |


## GetTask

Request the task information from the server.

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
    var id int64 = 771489
    ctx := context.Background()
    res, err := s.Tasks.GetTask(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Task != nil {
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

**[*operations.GetTaskResponse](../../models/operations/gettaskresponse.md), error**

### Errors

| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401,404               | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |


## SendCrack

Submit a cracked hash result for a task.

### Example Usage

```go
package main

import(
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/types"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var id int64 = 302642

    var hashcatResult *components.HashcatResult = &components.HashcatResult{
        Timestamp: types.MustTimeFromString("2024-08-25T21:36:04.134-04:00"),
        Hash: "dummy_hash_2",
        PlainText: "dummy_plain",
    }
    ctx := context.Background()
    res, err := s.Tasks.SendCrack(ctx, id, hashcatResult)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `id`                                                                  | *int64*                                                               | :heavy_check_mark:                                                    | id                                                                    |
| `hashcatResult`                                                       | [*components.HashcatResult](../../models/components/hashcatresult.md) | :heavy_minus_sign:                                                    | N/A                                                                   |
| `opts`                                                                | [][operations.Option](../../models/operations/option.md)              | :heavy_minus_sign:                                                    | The options for this request.                                         |

### Response

**[*operations.SendCrackResponse](../../models/operations/sendcrackresponse.md), error**

### Errors

| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 404                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |


## SendStatus

Submit a status update for a task. This includes the status of the current guess and the devices.

### Example Usage

```go
package main

import(
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/types"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var id int64 = 204258

    taskStatus := components.TaskStatus{
        OriginalLine: "<value>",
        Time: types.MustTimeFromString("2023-10-13T23:08:24.639Z"),
        Session: "<value>",
        HashcatGuess: components.HashcatGuess{
            GuessBase: "<value>",
            GuessBaseCount: 380021,
            GuessBaseOffset: 469354,
            GuessBasePercentage: 2929.65,
            GuessMod: "<value>",
            GuessModCount: 508837,
            GuessModOffset: 239276,
            GuessModPercentage: 1149.28,
            GuessMode: 402894,
        },
        Status: 889036,
        Target: "<value>",
        Progress: []int64{
            319182,
        },
        RestorePoint: 596493,
        RecoveredHashes: []int64{
            642941,
        },
        RecoveredSalts: []int64{
            336085,
        },
        Rejected: 906004,
        DeviceStatuses: []components.DeviceStatus{
            components.DeviceStatus{
                DeviceID: 896201,
                DeviceName: "<value>",
                DeviceType: components.DeviceTypeCPU,
                Speed: 65865,
                Utilization: 921724,
                Temperature: 572686,
            },
        },
        TimeStart: types.MustTimeFromString("2024-10-11T04:57:47.416Z"),
        EstimatedStop: types.MustTimeFromString("2022-10-24T09:23:01.835Z"),
    }
    ctx := context.Background()
    res, err := s.Tasks.SendStatus(ctx, id, taskStatus)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `id`                                                           | *int64*                                                        | :heavy_check_mark:                                             | id                                                             |
| `taskStatus`                                                   | [components.TaskStatus](../../models/components/taskstatus.md) | :heavy_check_mark:                                             | status                                                         |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.SendStatusResponse](../../models/operations/sendstatusresponse.md), error**

### Errors

| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401,404,422           | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |


## SetTaskAccepted

Accept an offered task from the server.

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
    var id int64 = 893037
    ctx := context.Background()
    res, err := s.Tasks.SetTaskAccepted(ctx, id)
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

**[*operations.SetTaskAcceptedResponse](../../models/operations/settaskacceptedresponse.md), error**

### Errors

| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 404,422               | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |


## SetTaskExhausted

Notify the server that the task is exhausted. This will mark the task as completed.

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
    var id int64 = 700537
    ctx := context.Background()
    res, err := s.Tasks.SetTaskExhausted(ctx, id)
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

**[*operations.SetTaskExhaustedResponse](../../models/operations/settaskexhaustedresponse.md), error**

### Errors

| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401,404               | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |


## SetTaskAbandoned

Abandon a task. This will mark the task as abandoned. Usually used when the client is unable to complete the task.

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
    var id int64 = 885883
    ctx := context.Background()
    res, err := s.Tasks.SetTaskAbandoned(ctx, id)
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

**[*operations.SetTaskAbandonedResponse](../../models/operations/settaskabandonedresponse.md), error**

### Errors

| Error Object                           | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorObject                  | 401,404                                | application/json                       |
| sdkerrors.SetTaskAbandonedResponseBody | 422                                    | application/json                       |
| sdkerrors.SDKError                     | 4xx-5xx                                | */*                                    |


## GetTaskZaps

Gets the completed hashes for a task. This is a text file that should be added to the monitored directory to remove the hashes from the list during runtime.

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
    var id int64 = 174947
    ctx := context.Background()
    res, err := s.Tasks.GetTaskZaps(ctx, id)
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

**[*operations.GetTaskZapsResponse](../../models/operations/gettaskzapsresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
