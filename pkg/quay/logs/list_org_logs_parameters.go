// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListOrgLogsParams creates a new ListOrgLogsParams object
// with the default values initialized.
func NewListOrgLogsParams() *ListOrgLogsParams {
	var ()
	return &ListOrgLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOrgLogsParamsWithTimeout creates a new ListOrgLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOrgLogsParamsWithTimeout(timeout time.Duration) *ListOrgLogsParams {
	var ()
	return &ListOrgLogsParams{

		timeout: timeout,
	}
}

// NewListOrgLogsParamsWithContext creates a new ListOrgLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOrgLogsParamsWithContext(ctx context.Context) *ListOrgLogsParams {
	var ()
	return &ListOrgLogsParams{

		Context: ctx,
	}
}

// NewListOrgLogsParamsWithHTTPClient creates a new ListOrgLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOrgLogsParamsWithHTTPClient(client *http.Client) *ListOrgLogsParams {
	var ()
	return &ListOrgLogsParams{
		HTTPClient: client,
	}
}

/*ListOrgLogsParams contains all the parameters to send to the API endpoint
for the list org logs operation typically these are written to a http.Request
*/
type ListOrgLogsParams struct {

	/*Endtime
	  Latest time to which to get logs. (%m/%d/%Y %Z)

	*/
	Endtime *string
	/*NextPage
	  The page token for the next page

	*/
	NextPage *string
	/*Orgname
	  The name of the organization

	*/
	Orgname string
	/*Page
	  The page number for the logs

	*/
	Page *int64
	/*Performer
	  Username for which to filter logs.

	*/
	Performer *string
	/*Starttime
	  Earliest time from which to get logs. (%m/%d/%Y %Z)

	*/
	Starttime *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list org logs params
func (o *ListOrgLogsParams) WithTimeout(timeout time.Duration) *ListOrgLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list org logs params
func (o *ListOrgLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list org logs params
func (o *ListOrgLogsParams) WithContext(ctx context.Context) *ListOrgLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list org logs params
func (o *ListOrgLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list org logs params
func (o *ListOrgLogsParams) WithHTTPClient(client *http.Client) *ListOrgLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list org logs params
func (o *ListOrgLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndtime adds the endtime to the list org logs params
func (o *ListOrgLogsParams) WithEndtime(endtime *string) *ListOrgLogsParams {
	o.SetEndtime(endtime)
	return o
}

// SetEndtime adds the endtime to the list org logs params
func (o *ListOrgLogsParams) SetEndtime(endtime *string) {
	o.Endtime = endtime
}

// WithNextPage adds the nextPage to the list org logs params
func (o *ListOrgLogsParams) WithNextPage(nextPage *string) *ListOrgLogsParams {
	o.SetNextPage(nextPage)
	return o
}

// SetNextPage adds the nextPage to the list org logs params
func (o *ListOrgLogsParams) SetNextPage(nextPage *string) {
	o.NextPage = nextPage
}

// WithOrgname adds the orgname to the list org logs params
func (o *ListOrgLogsParams) WithOrgname(orgname string) *ListOrgLogsParams {
	o.SetOrgname(orgname)
	return o
}

// SetOrgname adds the orgname to the list org logs params
func (o *ListOrgLogsParams) SetOrgname(orgname string) {
	o.Orgname = orgname
}

// WithPage adds the page to the list org logs params
func (o *ListOrgLogsParams) WithPage(page *int64) *ListOrgLogsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list org logs params
func (o *ListOrgLogsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerformer adds the performer to the list org logs params
func (o *ListOrgLogsParams) WithPerformer(performer *string) *ListOrgLogsParams {
	o.SetPerformer(performer)
	return o
}

// SetPerformer adds the performer to the list org logs params
func (o *ListOrgLogsParams) SetPerformer(performer *string) {
	o.Performer = performer
}

// WithStarttime adds the starttime to the list org logs params
func (o *ListOrgLogsParams) WithStarttime(starttime *string) *ListOrgLogsParams {
	o.SetStarttime(starttime)
	return o
}

// SetStarttime adds the starttime to the list org logs params
func (o *ListOrgLogsParams) SetStarttime(starttime *string) {
	o.Starttime = starttime
}

// WriteToRequest writes these params to a swagger request
func (o *ListOrgLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Endtime != nil {

		// query param endtime
		var qrEndtime string
		if o.Endtime != nil {
			qrEndtime = *o.Endtime
		}
		qEndtime := qrEndtime
		if qEndtime != "" {
			if err := r.SetQueryParam("endtime", qEndtime); err != nil {
				return err
			}
		}

	}

	if o.NextPage != nil {

		// query param next_page
		var qrNextPage string
		if o.NextPage != nil {
			qrNextPage = *o.NextPage
		}
		qNextPage := qrNextPage
		if qNextPage != "" {
			if err := r.SetQueryParam("next_page", qNextPage); err != nil {
				return err
			}
		}

	}

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Performer != nil {

		// query param performer
		var qrPerformer string
		if o.Performer != nil {
			qrPerformer = *o.Performer
		}
		qPerformer := qrPerformer
		if qPerformer != "" {
			if err := r.SetQueryParam("performer", qPerformer); err != nil {
				return err
			}
		}

	}

	if o.Starttime != nil {

		// query param starttime
		var qrStarttime string
		if o.Starttime != nil {
			qrStarttime = *o.Starttime
		}
		qStarttime := qrStarttime
		if qStarttime != "" {
			if err := r.SetQueryParam("starttime", qStarttime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
