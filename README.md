# CipherSwarm Agent SDK for Go

![Made With SpeakEasy](https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454)
![GitHub License](https://img.shields.io/github/license/unclesp1d3r/cipherswarm-agent-sdk-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/unclesp1d3r/cipherswarm-agent-sdk-go.svg)](https://pkg.go.dev/github.com/unclesp1d3r/cipherswarm-agent-sdk-go)

Welcome to the CipherSwarm Agent SDK for Go! This SDK provides the tools you need to write agents for the CipherSwarm platform. CipherSwarm is a powerful, distributed system for secure, scalable data processing and analysis.

## Features

- Agent Creation: Easily create and manage agents for the CipherSwarm platform.
- Secure Communication: Utilize built-in encryption and authentication mechanisms to ensure secure communication between agents and the CipherSwarm network.
- Scalability: Designed to handle high-throughput data processing tasks with ease.
- Extensibility: Flexible architecture allows for customization and extension of agent functionality.

## Getting Started

Prerequisites

- Go 1.16 or later
- A working installation of the CipherSwarm platform. Refer to the [CipherSwarm documentation](https://github.com/unclesp1d3r/CipherSwarm) for setup instructions.

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/unclesp1d3r/cipherswarm-agent-sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"log"
	"os"
)

func main() {
	s := cipherswarmagentsdkgo.New(
		cipherswarmagentsdkgo.WithSecurity(os.Getenv("BEARER_AUTH")),
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Agents](docs/sdks/agents/README.md)

* [GetAgent](docs/sdks/agents/README.md#getagent) - Gets an instance of an agent
* [UpdateAgent](docs/sdks/agents/README.md#updateagent) - Updates the agent
* [SendHeartbeat](docs/sdks/agents/README.md#sendheartbeat) - Send a heartbeat for an agent
* [SubmitBenchmark](docs/sdks/agents/README.md#submitbenchmark) - submit_benchmark agent
* [SubmitErrorAgent](docs/sdks/agents/README.md#submiterroragent) - Submit an error for an agent
* [SetAgentShutdown](docs/sdks/agents/README.md#setagentshutdown) - shutdown agent

### [Attacks](docs/sdks/attacks/README.md)

* [GetAttack](docs/sdks/attacks/README.md#getattack) - show attack
* [GetHashList](docs/sdks/attacks/README.md#gethashlist) - Get the hash list

### [Crackers](docs/sdks/crackers/README.md)

* [CheckForCrackerUpdate](docs/sdks/crackers/README.md#checkforcrackerupdate) - Check for Cracker Update

### [Tasks](docs/sdks/tasks/README.md)

* [GetNewTask](docs/sdks/tasks/README.md#getnewtask) - Request a new task from server
* [GetTask](docs/sdks/tasks/README.md#gettask) - Request the task information
* [SendCrack](docs/sdks/tasks/README.md#sendcrack) - Submit a cracked hash result for a task
* [SendStatus](docs/sdks/tasks/README.md#sendstatus) - Submit a status update for a task
* [SetTaskAccepted](docs/sdks/tasks/README.md#settaskaccepted) - Accept Task
* [SetTaskExhausted](docs/sdks/tasks/README.md#settaskexhausted) - Notify of Exhausted Task
* [SetTaskAbandoned](docs/sdks/tasks/README.md#settaskabandoned) - Abandon Task
* [GetTaskZaps](docs/sdks/tasks/README.md#gettaskzaps) - Get Completed Hashes

### [Client](docs/sdks/client/README.md)

* [GetConfiguration](docs/sdks/client/README.md#getconfiguration) - Get Agent Configuration
* [Authenticate](docs/sdks/client/README.md#authenticate) - Authenticate Client
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |

### Example

```go
package main

import (
	"context"
	"errors"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/sdkerrors"
	"log"
	"os"
)

func main() {
	s := cipherswarmagentsdkgo.New(
		cipherswarmagentsdkgo.WithSecurity(os.Getenv("BEARER_AUTH")),
	)
	var id int64 = 135003
	ctx := context.Background()
	res, err := s.Agents.GetAgent(ctx, id)
	if err != nil {

		var e *sdkerrors.ErrorObject
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://{defaultHost}` | `defaultHost` (default is `www.example.com`) |
| 1 | `http://{hostAddress}:{hostPort}` | `hostAddress` (default is `localhost`), `hostPort` (default is `8080`) |

#### Example

```go
package main

import (
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"log"
	"os"
)

func main() {
	s := cipherswarmagentsdkgo.New(
		cipherswarmagentsdkgo.WithServerIndex(1),
		cipherswarmagentsdkgo.WithSecurity(os.Getenv("BEARER_AUTH")),
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

#### Variables

Some of the server options above contain variables. If you want to set the values of those variables, the following options are provided for doing so:
 * `WithDefaultHost string`
 * `WithHostAddress string`
 * `WithHostPort string`

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"log"
	"os"
)

func main() {
	s := cipherswarmagentsdkgo.New(
		cipherswarmagentsdkgo.WithServerURL("https://{defaultHost}"),
		cipherswarmagentsdkgo.WithSecurity(os.Getenv("BEARER_AUTH")),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `BearerAuth` | http         | HTTP Bearer  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"log"
	"os"
)

func main() {
	s := cipherswarmagentsdkgo.New(
		cipherswarmagentsdkgo.WithSecurity(os.Getenv("BEARER_AUTH")),
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
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	s := cipherswarmagentsdkgo.New(
		cipherswarmagentsdkgo.WithSecurity(os.Getenv("BEARER_AUTH")),
	)
	var id int64 = 135003
	ctx := context.Background()
	res, err := s.Agents.GetAgent(ctx, id, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Agent != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/retry"
	"log"
	"os"
)

func main() {
	s := cipherswarmagentsdkgo.New(
		cipherswarmagentsdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		cipherswarmagentsdkgo.WithSecurity(os.Getenv("BEARER_AUTH")),
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
<!-- End Retries [retries] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update.

Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically from the CipherSwarm API specification. As a result, we are unable to accept pull requests for changes to the SDK itself. However, we are always looking for ways to improve the SDK and make it more useful for developers. If you have an idea for a new feature or improvement, please open an issue on the GitHub repository.

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
