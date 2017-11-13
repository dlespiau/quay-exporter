// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/dlespiau/quay-exporter/pkg/models"
)

// NewCreateStarParams creates a new CreateStarParams object
// with the default values initialized.
func NewCreateStarParams() *CreateStarParams {
	var ()
	return &CreateStarParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateStarParamsWithTimeout creates a new CreateStarParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateStarParamsWithTimeout(timeout time.Duration) *CreateStarParams {
	var ()
	return &CreateStarParams{

		timeout: timeout,
	}
}

// NewCreateStarParamsWithContext creates a new CreateStarParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateStarParamsWithContext(ctx context.Context) *CreateStarParams {
	var ()
	return &CreateStarParams{

		Context: ctx,
	}
}

// NewCreateStarParamsWithHTTPClient creates a new CreateStarParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateStarParamsWithHTTPClient(client *http.Client) *CreateStarParams {
	var ()
	return &CreateStarParams{
		HTTPClient: client,
	}
}

/*CreateStarParams contains all the parameters to send to the API endpoint
for the create star operation typically these are written to a http.Request
*/
type CreateStarParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.NewStarredRepository

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create star params
func (o *CreateStarParams) WithTimeout(timeout time.Duration) *CreateStarParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create star params
func (o *CreateStarParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create star params
func (o *CreateStarParams) WithContext(ctx context.Context) *CreateStarParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create star params
func (o *CreateStarParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create star params
func (o *CreateStarParams) WithHTTPClient(client *http.Client) *CreateStarParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create star params
func (o *CreateStarParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create star params
func (o *CreateStarParams) WithBody(body *models.NewStarredRepository) *CreateStarParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create star params
func (o *CreateStarParams) SetBody(body *models.NewStarredRepository) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateStarParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
