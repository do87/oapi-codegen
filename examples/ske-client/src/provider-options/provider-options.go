// Package provideroptions provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version (devel) DO NOT EDIT.
package provideroptions

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	skeclient "github.com/do87/oapi-codegen/examples/ske-client"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for CRIName.
const (
	CONTAINERD CRIName = "containerd"
	DOCKER     CRIName = "docker"
)

// Defines values for RuntimeErrorCode.
const (
	SKE_API_SERVER_ERROR      RuntimeErrorCode = "SKE_API_SERVER_ERROR"
	SKE_CONFIGURATION_PROBLEM RuntimeErrorCode = "SKE_CONFIGURATION_PROBLEM"
	SKE_INFRA_ERROR           RuntimeErrorCode = "SKE_INFRA_ERROR"
	SKE_QUOTA_EXCEEDED        RuntimeErrorCode = "SKE_QUOTA_EXCEEDED"
	SKE_RATE_LIMITS           RuntimeErrorCode = "SKE_RATE_LIMITS"
	SKE_REMAINING_RESOURCES   RuntimeErrorCode = "SKE_REMAINING_RESOURCES"
	SKE_TMP_AUTH_ERROR        RuntimeErrorCode = "SKE_TMP_AUTH_ERROR"
	SKE_UNREADY_NODES         RuntimeErrorCode = "SKE_UNREADY_NODES"
	SKE_UNSPECIFIED           RuntimeErrorCode = "SKE_UNSPECIFIED"
)

// AvailabilityZone defines model for AvailabilityZone.
type AvailabilityZone struct {
	Name *string `json:"name,omitempty"`
}

// CRI defines model for CRI.
type CRI struct {
	Name *CRIName `json:"name,omitempty"`
}

// CRIName defines model for CRI.Name.
type CRIName string

// KubernetesVersion defines model for KubernetesVersion.
type KubernetesVersion struct {
	ExpirationDate *time.Time         `json:"expirationDate,omitempty"`
	FeatureGates   *map[string]string `json:"featureGates,omitempty"`
	State          *string            `json:"state,omitempty"`
	Version        *string            `json:"version,omitempty"`
}

// MachineImage defines model for MachineImage.
type MachineImage struct {
	Name     *string                `json:"name,omitempty"`
	Versions *[]MachineImageVersion `json:"versions,omitempty"`
}

// MachineImageVersion defines model for MachineImageVersion.
type MachineImageVersion struct {
	CRI            *[]CRI     `json:"cri,omitempty"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	State          *string    `json:"state,omitempty"`
	Version        *string    `json:"version,omitempty"`
}

// MachineType defines model for MachineType.
type MachineType struct {
	CPU    *int32  `json:"cpu,omitempty"`
	Memory *int32  `json:"memory,omitempty"`
	Name   *string `json:"name,omitempty"`
}

// ProviderOptions defines model for ProviderOptions.
type ProviderOptions struct {
	AvailabilityZones  *[]AvailabilityZone  `json:"availabilityZones,omitempty"`
	KubernetesVersions *[]KubernetesVersion `json:"kubernetesVersions,omitempty"`
	MachineImages      *[]MachineImage      `json:"machineImages,omitempty"`
	MachineTypes       *[]MachineType       `json:"machineTypes,omitempty"`
	VolumeTypes        *[]VolumeType        `json:"volumeTypes,omitempty"`
}

// RuntimeError defines model for RuntimeError.
type RuntimeError struct {
	// Code - Code:    "SKE_UNSPECIFIED"
	//   Message: "An error occurred. Please open a support ticket if this error persists."
	// - Code:    "SKE_TMP_AUTH_ERROR"
	//   Message: "Authentication failed. This is a temporary error. Please wait while the system recovers."
	// - Code:    "SKE_QUOTA_EXCEEDED"
	//   Message: "Your project's resource quotas are exhausted. Please make sure your quota is sufficient for the ordered cluster."
	// - Code:    "SKE_RATE_LIMITS"
	//   Message: "While provisioning your cluster, request rate limits where incurred. Please wait while the system recovers."
	// - Code:    "SKE_INFRA_ERROR"
	//   Message: "An error occurred with the underlying infrastructure. Please open a support ticket if this error persists."
	// - Code:    "SKE_REMAINING_RESOURCES"
	//   Message: "There are remaining Kubernetes resources in your cluster that prevent deletion. Please make sure to remove them."
	// - Code:    "SKE_CONFIGURATION_PROBLEM"
	//   Message: "A configuration error occurred. Please open a support ticket if this error persists."
	// - Code:    "SKE_UNREADY_NODES"
	//   Message: "Not all worker nodes are ready. Please open a support ticket if this error persists."
	// - Code:    "SKE_API_SERVER_ERROR"
	//   Message: "The Kubernetes API server is not reporting readiness. Please open a support ticket if this error persists."
	Code    *RuntimeErrorCode `json:"code,omitempty"`
	Details *string           `json:"details,omitempty"`
	Message *string           `json:"message,omitempty"`
}

// RuntimeErrorCode - Code:    "SKE_UNSPECIFIED"
//
//		Message: "An error occurred. Please open a support ticket if this error persists."
//	  - Code:    "SKE_TMP_AUTH_ERROR"
//	    Message: "Authentication failed. This is a temporary error. Please wait while the system recovers."
//	  - Code:    "SKE_QUOTA_EXCEEDED"
//	    Message: "Your project's resource quotas are exhausted. Please make sure your quota is sufficient for the ordered cluster."
//	  - Code:    "SKE_RATE_LIMITS"
//	    Message: "While provisioning your cluster, request rate limits where incurred. Please wait while the system recovers."
//	  - Code:    "SKE_INFRA_ERROR"
//	    Message: "An error occurred with the underlying infrastructure. Please open a support ticket if this error persists."
//	  - Code:    "SKE_REMAINING_RESOURCES"
//	    Message: "There are remaining Kubernetes resources in your cluster that prevent deletion. Please make sure to remove them."
//	  - Code:    "SKE_CONFIGURATION_PROBLEM"
//	    Message: "A configuration error occurred. Please open a support ticket if this error persists."
//	  - Code:    "SKE_UNREADY_NODES"
//	    Message: "Not all worker nodes are ready. Please open a support ticket if this error persists."
//	  - Code:    "SKE_API_SERVER_ERROR"
//	    Message: "The Kubernetes API server is not reporting readiness. Please open a support ticket if this error persists."
type RuntimeErrorCode string

// VolumeType defines model for VolumeType.
type VolumeType struct {
	Name *string `json:"name,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client skeclient.HttpRequestDoer
}

// Creates a new Client, with reasonable defaults
func NewClient(server string, httpClient skeclient.HttpRequestDoer) *Client {
	// create a client with sane default values
	client := Client{
		Server: server,
		Client: httpClient,
	}
	return &client
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetProviderOptions request
	GetProviderOptions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetProviderOptions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetProviderOptionsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetProviderOptionsRequest generates requests for GetProviderOptions
func NewGetProviderOptionsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/provider-options")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, httpClient skeclient.HttpRequestDoer) *ClientWithResponses {
	return &ClientWithResponses{NewClient(server, httpClient)}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetProviderOptions request
	GetProviderOptionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetProviderOptionsResponse, error)
}

type GetProviderOptionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ProviderOptions
	JSONDefault  *RuntimeError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r GetProviderOptionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetProviderOptionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetProviderOptionsWithResponse request returning *GetProviderOptionsResponse
func (c *ClientWithResponses) GetProviderOptionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetProviderOptionsResponse, error) {
	rsp, err := c.GetProviderOptions(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetProviderOptionsResponse(rsp)
}

// ParseGetProviderOptionsResponse parses an HTTP response from a GetProviderOptionsWithResponse call
func (c *ClientWithResponses) ParseGetProviderOptionsResponse(rsp *http.Response) (*GetProviderOptionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetProviderOptionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = nil

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ProviderOptions
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RuntimeError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
