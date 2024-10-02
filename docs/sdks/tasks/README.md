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

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

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

    ctx := context.Background()
    res, err := s.Tasks.GetTask(ctx, 771489)
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

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401, 404              | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## SendCrack

Submit a cracked hash result for a task.

### Example Usage

```go
package main

import(
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"context"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/types"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Tasks.SendCrack(ctx, 302642, &components.HashcatResult{
        Timestamp: types.MustTimeFromString("2024-10-01T21:16:43.171-04:00"),
        Hash: "dummy_hash_2",
        PlainText: "dummy_plain",
    })
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

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 404                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## SendStatus

Submit a status update for a task. This includes the status of the current guess and the devices.

### Example Usage

```go
package main

import(
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"context"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/types"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"log"
)

func main() {
    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Tasks.SendStatus(ctx, 144718, components.TaskStatus{
        OriginalLine: "<value>",
        Time: types.MustTimeFromString("2022-08-12T20:48:19.251Z"),
        Session: "<value>",
        HashcatGuess: components.HashcatGuess{
            GuessBase: "<value>",
            GuessBaseCount: 593946,
            GuessBaseOffset: 380021,
            GuessBasePercentage: 4693.54,
            GuessMod: "<value>",
            GuessModCount: 292965,
            GuessModOffset: 508837,
            GuessModPercentage: 2392.76,
            GuessMode: 114928,
        },
        Status: 402894,
        Target: "<value>",
        Progress: []int64{
            319182,
            596493,
            642941,
        },
        RestorePoint: 336085,
        RecoveredHashes: []int64{
            896201,
            370747,
            65865,
        },
        RecoveredSalts: []int64{
            572686,
            925372,
            270430,
        },
        Rejected: 672868,
        DeviceStatuses: []components.DeviceStatus{
            components.DeviceStatus{
                DeviceID: 383051,
                DeviceName: "<value>",
                DeviceType: components.DeviceTypeGpu,
                Speed: 182572,
                Utilization: 740400,
                Temperature: 334336,
            },
            components.DeviceStatus{
                DeviceID: 330568,
                DeviceName: "<value>",
                DeviceType: components.DeviceTypeCPU,
                Speed: 364941,
                Utilization: 953679,
                Temperature: 240269,
            },
            components.DeviceStatus{
                DeviceID: 556137,
                DeviceName: "<value>",
                DeviceType: components.DeviceTypeCPU,
                Speed: 923394,
                Utilization: 990009,
                Temperature: 440098,
            },
        },
        TimeStart: types.MustTimeFromString("2022-12-21T12:08:07.379Z"),
        EstimatedStop: types.MustTimeFromString("2022-10-27T18:57:43.860Z"),
    })
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

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401, 422              | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

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

    ctx := context.Background()
    res, err := s.Tasks.SetTaskAccepted(ctx, 893037)
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

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 404, 422              | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

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

    ctx := context.Background()
    res, err := s.Tasks.SetTaskExhausted(ctx, 700537)
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

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401, 404              | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

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

    ctx := context.Background()
    res, err := s.Tasks.SetTaskAbandoned(ctx, 885883)
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

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorObject                  | 401, 404                               | application/json                       |
| sdkerrors.SetTaskAbandonedResponseBody | 422                                    | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

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

    ctx := context.Background()
    res, err := s.Tasks.GetTaskZaps(ctx, 174947)
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |