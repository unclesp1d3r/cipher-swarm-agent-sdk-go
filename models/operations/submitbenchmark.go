// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"net/http"
)

type SubmitBenchmarkRequest struct {
	// id
	ID          int64                         `pathParam:"style=simple,explode=false,name=id"`
	RequestBody []components.HashcatBenchmark `request:"mediaType=application/json"`
}

func (o *SubmitBenchmarkRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *SubmitBenchmarkRequest) GetRequestBody() []components.HashcatBenchmark {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type SubmitBenchmarkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SubmitBenchmarkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SubmitBenchmarkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SubmitBenchmarkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}