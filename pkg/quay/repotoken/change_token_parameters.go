// Code generated by go-swagger; DO NOT EDIT.

package repotoken

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

// NewChangeTokenParams creates a new ChangeTokenParams object
// with the default values initialized.
func NewChangeTokenParams() *ChangeTokenParams {
	var ()
	return &ChangeTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeTokenParamsWithTimeout creates a new ChangeTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeTokenParamsWithTimeout(timeout time.Duration) *ChangeTokenParams {
	var ()
	return &ChangeTokenParams{

		timeout: timeout,
	}
}

// NewChangeTokenParamsWithContext creates a new ChangeTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeTokenParamsWithContext(ctx context.Context) *ChangeTokenParams {
	var ()
	return &ChangeTokenParams{

		Context: ctx,
	}
}

// NewChangeTokenParamsWithHTTPClient creates a new ChangeTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeTokenParamsWithHTTPClient(client *http.Client) *ChangeTokenParams {
	var ()
	return &ChangeTokenParams{
		HTTPClient: client,
	}
}

/*ChangeTokenParams contains all the parameters to send to the API endpoint
for the change token operation typically these are written to a http.Request
*/
type ChangeTokenParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.TokenPermission
	/*Code
	  The token code

	*/
	Code string
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change token params
func (o *ChangeTokenParams) WithTimeout(timeout time.Duration) *ChangeTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change token params
func (o *ChangeTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change token params
func (o *ChangeTokenParams) WithContext(ctx context.Context) *ChangeTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change token params
func (o *ChangeTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change token params
func (o *ChangeTokenParams) WithHTTPClient(client *http.Client) *ChangeTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change token params
func (o *ChangeTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change token params
func (o *ChangeTokenParams) WithBody(body *models.TokenPermission) *ChangeTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change token params
func (o *ChangeTokenParams) SetBody(body *models.TokenPermission) {
	o.Body = body
}

// WithCode adds the code to the change token params
func (o *ChangeTokenParams) WithCode(code string) *ChangeTokenParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the change token params
func (o *ChangeTokenParams) SetCode(code string) {
	o.Code = code
}

// WithRepository adds the repository to the change token params
func (o *ChangeTokenParams) WithRepository(repository string) *ChangeTokenParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the change token params
func (o *ChangeTokenParams) SetRepository(repository string) {
	o.Repository = repository
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param code
	if err := r.SetPathParam("code", o.Code); err != nil {
		return err
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
