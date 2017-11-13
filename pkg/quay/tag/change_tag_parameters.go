// Code generated by go-swagger; DO NOT EDIT.

package tag

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

// NewChangeTagParams creates a new ChangeTagParams object
// with the default values initialized.
func NewChangeTagParams() *ChangeTagParams {
	var ()
	return &ChangeTagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeTagParamsWithTimeout creates a new ChangeTagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeTagParamsWithTimeout(timeout time.Duration) *ChangeTagParams {
	var ()
	return &ChangeTagParams{

		timeout: timeout,
	}
}

// NewChangeTagParamsWithContext creates a new ChangeTagParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeTagParamsWithContext(ctx context.Context) *ChangeTagParams {
	var ()
	return &ChangeTagParams{

		Context: ctx,
	}
}

// NewChangeTagParamsWithHTTPClient creates a new ChangeTagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeTagParamsWithHTTPClient(client *http.Client) *ChangeTagParams {
	var ()
	return &ChangeTagParams{
		HTTPClient: client,
	}
}

/*ChangeTagParams contains all the parameters to send to the API endpoint
for the change tag operation typically these are written to a http.Request
*/
type ChangeTagParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.ChangeTag
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Tag
	  The name of the tag

	*/
	Tag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change tag params
func (o *ChangeTagParams) WithTimeout(timeout time.Duration) *ChangeTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change tag params
func (o *ChangeTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change tag params
func (o *ChangeTagParams) WithContext(ctx context.Context) *ChangeTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change tag params
func (o *ChangeTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change tag params
func (o *ChangeTagParams) WithHTTPClient(client *http.Client) *ChangeTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change tag params
func (o *ChangeTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change tag params
func (o *ChangeTagParams) WithBody(body *models.ChangeTag) *ChangeTagParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change tag params
func (o *ChangeTagParams) SetBody(body *models.ChangeTag) {
	o.Body = body
}

// WithRepository adds the repository to the change tag params
func (o *ChangeTagParams) WithRepository(repository string) *ChangeTagParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the change tag params
func (o *ChangeTagParams) SetRepository(repository string) {
	o.Repository = repository
}

// WithTag adds the tag to the change tag params
func (o *ChangeTagParams) WithTag(tag string) *ChangeTagParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the change tag params
func (o *ChangeTagParams) SetTag(tag string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
