# Tasks
(*Tasks*)

## Overview

Tasks API

### Available Operations

* [NewTask](#newtask) - Request a new task from server
* [ShowTask](#showtask) - Request the task information
* [SubmitCrack](#submitcrack) - Submit a cracked hash result for a task
* [SubmitStatus](#submitstatus) - Submit a status update for a task
* [AcceptTask](#accepttask) - Accept Task
* [ExhaustedTask](#exhaustedtask) - Notify of Exhausted Task
* [AbandonTask](#abandontask) - Abandon Task

## NewTask

Request a new task from the server, if available.

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
    res, err := s.Tasks.NewTask(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Task != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.NewTaskResponse](../../models/operations/newtaskresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ShowTask

Request the task information from the server.

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

    var id int64 = 576214
    
    ctx := context.Background()
    res, err := s.Tasks.ShowTask(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Task != nil {
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

**[*operations.ShowTaskResponse](../../models/operations/showtaskresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## SubmitCrack

Submit a cracked hash result for a task.

### Example Usage

```go
package main

import(
	"github.com/unclesp1d3r/cipherswarm-agent-sdk/models/components"
	cipherswarmagentsdk "github.com/unclesp1d3r/cipherswarm-agent-sdk"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk/types"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdk.New(
        cipherswarmagentsdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    var id int64 = 150559

    var hashcatResult *components.HashcatResult = &components.HashcatResult{
        Timestamp: types.MustTimeFromString("2022-06-14T17:51:53.220Z"),
        Hash: "<value>",
        PlainText: "<value>",
    }
    
    ctx := context.Background()
    res, err := s.Tasks.SubmitCrack(ctx, id, hashcatResult)
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


### Response

**[*operations.SubmitCrackResponse](../../models/operations/submitcrackresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## SubmitStatus

Submit a status update for a task. This includes the status of the current guess and the devices.

### Example Usage

```go
package main

import(
	"github.com/unclesp1d3r/cipherswarm-agent-sdk/models/components"
	cipherswarmagentsdk "github.com/unclesp1d3r/cipherswarm-agent-sdk"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk/types"
	"context"
	"log"
)

func main() {
    s := cipherswarmagentsdk.New(
        cipherswarmagentsdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    var id int64 = 584703

    taskStatus := components.TaskStatus{
        OriginalLine: "<value>",
        Time: types.MustTimeFromString("2023-12-03T10:27:22.986Z"),
        Session: "<value>",
        HashcatGuess: components.HashcatGuess{
            GuessBase: "<value>",
            GuessBaseCount: 602220,
            GuessBaseOffset: 22120,
            GuessBasePercentage: 9978.16,
            GuessMod: "<value>",
            GuessModCount: 948688,
            GuessModOffset: 68184,
            GuessModPercentage: 5432.29,
            GuessMode: 140481,
        },
        Status: 477845,
        Target: "<value>",
        Progress: []int64{
            448822,
        },
        RestorePoint: 466580,
        RecoveredHashes: []int64{
            213301,
        },
        RecoveredSalts: []int64{
            527166,
        },
        Rejected: 235279,
        DeviceStatuses: []components.DeviceStatus{
            components.DeviceStatus{
                DeviceID: 722842,
                DeviceName: "<value>",
                DeviceType: components.TheTypeOfTheDeviceCPU,
                Speed: 939700,
                Utilization: 725728,
                Temperature: 757147,
            },
        },
        TimeStart: types.MustTimeFromString("2022-03-04T10:33:44.215Z"),
        EstimatedStop: types.MustTimeFromString("2022-10-11T11:57:10.785Z"),
    }
    
    ctx := context.Background()
    res, err := s.Tasks.SubmitStatus(ctx, id, taskStatus)
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


### Response

**[*operations.SubmitStatusResponse](../../models/operations/submitstatusresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AcceptTask

Accept an offered task from the server.

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

    var id int64 = 262707
    
    ctx := context.Background()
    res, err := s.Tasks.AcceptTask(ctx, id)
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

**[*operations.AcceptTaskResponse](../../models/operations/accepttaskresponse.md), error**
| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 404,422               | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |

## ExhaustedTask

Notify the server that the task is exhausted. This will mark the task as completed.

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

    var id int64 = 461909
    
    ctx := context.Background()
    res, err := s.Tasks.ExhaustedTask(ctx, id)
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

**[*operations.ExhaustedTaskResponse](../../models/operations/exhaustedtaskresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AbandonTask

Abandon a task. This will mark the task as abandoned. Usually used when the client is unable to complete the task.

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

    var id int64 = 446205
    
    ctx := context.Background()
    res, err := s.Tasks.AbandonTask(ctx, id)
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

**[*operations.AbandonTaskResponse](../../models/operations/abandontaskresponse.md), error**
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.StateError | 422                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |
