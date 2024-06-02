// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package cipherswarmagentsdkgo

import (
	"context"
	"fmt"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/internal/hooks"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/internal/utils"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// The production server
	"https://{defaultHost}",
	// The insecure server
	"http://{hostAddress}:{hostPort}",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	ServerDefaults    []map[string]string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
	Hooks             *hooks.Hooks
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], c.ServerDefaults[c.ServerIndex]
}

// CipherSwarmAgentSDK - CypherSwarm Agent API: The CypherSwarm Agent API is used to allow agents to connect to the CypherSwarm server.
type CipherSwarmAgentSDK struct {
	// Agents API
	Agents *Agents
	// Attacks API
	Attacks *Attacks
	// Crackers API
	Crackers *Crackers
	// Tasks API
	Tasks *Tasks
	// Client API
	Client *Client

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CipherSwarmAgentSDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CipherSwarmAgentSDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CipherSwarmAgentSDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CipherSwarmAgentSDK) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithDefaultHost allows setting the defaultHost variable for url substitution
func WithDefaultHost(defaultHost string) SDKOption {
	return func(sdk *CipherSwarmAgentSDK) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["defaultHost"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["defaultHost"] = fmt.Sprintf("%v", defaultHost)
		}
	}
}

// WithHostAddress allows setting the hostAddress variable for url substitution
func WithHostAddress(hostAddress string) SDKOption {
	return func(sdk *CipherSwarmAgentSDK) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["hostAddress"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["hostAddress"] = fmt.Sprintf("%v", hostAddress)
		}
	}
}

// WithHostPort allows setting the hostPort variable for url substitution
func WithHostPort(hostPort string) SDKOption {
	return func(sdk *CipherSwarmAgentSDK) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["hostPort"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["hostPort"] = fmt.Sprintf("%v", hostPort)
		}
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CipherSwarmAgentSDK) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(bearerAuth string) SDKOption {
	return func(sdk *CipherSwarmAgentSDK) {
		security := components.Security{BearerAuth: bearerAuth}
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *CipherSwarmAgentSDK) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *CipherSwarmAgentSDK) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CipherSwarmAgentSDK {
	sdk := &CipherSwarmAgentSDK{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.2",
			SDKVersion:        "0.3.0",
			GenVersion:        "2.338.7",
			UserAgent:         "speakeasy-sdk/go 0.3.0 2.338.7 1.2 github.com/unclesp1d3r/cipherswarm-agent-sdk-go",
			ServerDefaults: []map[string]string{
				{
					"defaultHost": "www.example.com",
				},
				{
					"hostAddress": "localhost",
					"hostPort":    "8080",
				},
			},
			Hooks: hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Agents = newAgents(sdk.sdkConfiguration)

	sdk.Attacks = newAttacks(sdk.sdkConfiguration)

	sdk.Crackers = newCrackers(sdk.sdkConfiguration)

	sdk.Tasks = newTasks(sdk.sdkConfiguration)

	sdk.Client = newClient(sdk.sdkConfiguration)

	return sdk
}
