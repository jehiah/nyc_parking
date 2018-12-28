// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAtTimeByCurbIDParams creates a new GetAtTimeByCurbIDParams object
// with the default values initialized.
func NewGetAtTimeByCurbIDParams() *GetAtTimeByCurbIDParams {
	var ()
	return &GetAtTimeByCurbIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAtTimeByCurbIDParamsWithTimeout creates a new GetAtTimeByCurbIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAtTimeByCurbIDParamsWithTimeout(timeout time.Duration) *GetAtTimeByCurbIDParams {
	var ()
	return &GetAtTimeByCurbIDParams{

		timeout: timeout,
	}
}

// NewGetAtTimeByCurbIDParamsWithContext creates a new GetAtTimeByCurbIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAtTimeByCurbIDParamsWithContext(ctx context.Context) *GetAtTimeByCurbIDParams {
	var ()
	return &GetAtTimeByCurbIDParams{

		Context: ctx,
	}
}

// NewGetAtTimeByCurbIDParamsWithHTTPClient creates a new GetAtTimeByCurbIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAtTimeByCurbIDParamsWithHTTPClient(client *http.Client) *GetAtTimeByCurbIDParams {
	var ()
	return &GetAtTimeByCurbIDParams{
		HTTPClient: client,
	}
}

/*GetAtTimeByCurbIDParams contains all the parameters to send to the API endpoint
for the get at time by curb id operation typically these are written to a http.Request
*/
type GetAtTimeByCurbIDParams struct {

	/*AccessKey
	  The API access key for the request.

	*/
	AccessKey string
	/*ID
	  The ID of the curb to return.

	*/
	ID string
	/*Time
	  The time (in ISO-8601 format) to find the rules for. If not provided, we will use the current time.


	*/
	Time *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get at time by curb id params
func (o *GetAtTimeByCurbIDParams) WithTimeout(timeout time.Duration) *GetAtTimeByCurbIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get at time by curb id params
func (o *GetAtTimeByCurbIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get at time by curb id params
func (o *GetAtTimeByCurbIDParams) WithContext(ctx context.Context) *GetAtTimeByCurbIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get at time by curb id params
func (o *GetAtTimeByCurbIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get at time by curb id params
func (o *GetAtTimeByCurbIDParams) WithHTTPClient(client *http.Client) *GetAtTimeByCurbIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get at time by curb id params
func (o *GetAtTimeByCurbIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKey adds the accessKey to the get at time by curb id params
func (o *GetAtTimeByCurbIDParams) WithAccessKey(accessKey string) *GetAtTimeByCurbIDParams {
	o.SetAccessKey(accessKey)
	return o
}

// SetAccessKey adds the accessKey to the get at time by curb id params
func (o *GetAtTimeByCurbIDParams) SetAccessKey(accessKey string) {
	o.AccessKey = accessKey
}

// WithID adds the id to the get at time by curb id params
func (o *GetAtTimeByCurbIDParams) WithID(id string) *GetAtTimeByCurbIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get at time by curb id params
func (o *GetAtTimeByCurbIDParams) SetID(id string) {
	o.ID = id
}

// WithTime adds the time to the get at time by curb id params
func (o *GetAtTimeByCurbIDParams) WithTime(time *strfmt.DateTime) *GetAtTimeByCurbIDParams {
	o.SetTime(time)
	return o
}

// SetTime adds the time to the get at time by curb id params
func (o *GetAtTimeByCurbIDParams) SetTime(time *strfmt.DateTime) {
	o.Time = time
}

// WriteToRequest writes these params to a swagger request
func (o *GetAtTimeByCurbIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param access_key
	qrAccessKey := o.AccessKey
	qAccessKey := qrAccessKey
	if qAccessKey != "" {
		if err := r.SetQueryParam("access_key", qAccessKey); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Time != nil {

		// query param time
		var qrTime strfmt.DateTime
		if o.Time != nil {
			qrTime = *o.Time
		}
		qTime := qrTime.String()
		if qTime != "" {
			if err := r.SetQueryParam("time", qTime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}