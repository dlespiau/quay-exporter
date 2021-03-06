// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new build API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for build API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CancelRepoBuild Cancels a repository build.
*/
func (a *Client) CancelRepoBuild(params *CancelRepoBuildParams, authInfo runtime.ClientAuthInfoWriter) (*CancelRepoBuildNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelRepoBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cancelRepoBuild",
		Method:             "DELETE",
		PathPattern:        "/api/v1/repository/{repository}/build/{build_uuid}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelRepoBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CancelRepoBuildNoContent), nil

}

/*
GetRepoBuild Returns information about a build.
*/
func (a *Client) GetRepoBuild(params *GetRepoBuildParams, authInfo runtime.ClientAuthInfoWriter) (*GetRepoBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepoBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRepoBuild",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/build/{build_uuid}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRepoBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRepoBuildOK), nil

}

/*
GetRepoBuildLogs Return the build logs for the build specified by the build uuid.
*/
func (a *Client) GetRepoBuildLogs(params *GetRepoBuildLogsParams, authInfo runtime.ClientAuthInfoWriter) (*GetRepoBuildLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepoBuildLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRepoBuildLogs",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/build/{build_uuid}/logs",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRepoBuildLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRepoBuildLogsOK), nil

}

/*
GetRepoBuildStatus Return the status for the builds specified by the build uuids.
*/
func (a *Client) GetRepoBuildStatus(params *GetRepoBuildStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetRepoBuildStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepoBuildStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRepoBuildStatus",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/build/{build_uuid}/status",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRepoBuildStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRepoBuildStatusOK), nil

}

/*
GetRepoBuilds Get the list of repository builds.
*/
func (a *Client) GetRepoBuilds(params *GetRepoBuildsParams, authInfo runtime.ClientAuthInfoWriter) (*GetRepoBuildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepoBuildsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRepoBuilds",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/build/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRepoBuildsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRepoBuildsOK), nil

}

/*
RequestRepoBuild Request that a repository be built and pushed from the specified input.
*/
func (a *Client) RequestRepoBuild(params *RequestRepoBuildParams, authInfo runtime.ClientAuthInfoWriter) (*RequestRepoBuildCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestRepoBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "requestRepoBuild",
		Method:             "POST",
		PathPattern:        "/api/v1/repository/{repository}/build/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RequestRepoBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RequestRepoBuildCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
