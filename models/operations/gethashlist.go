// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"io"
	"net/http"
)

type GetHashListRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetHashListRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetHashListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	Stream io.ReadCloser
}

func (o *GetHashListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHashListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHashListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetHashListResponse) GetStream() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Stream
}