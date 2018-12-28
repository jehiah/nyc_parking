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

// NewGetByCurbIDParams creates a new GetByCurbIDParams object
// with the default values initialized.
func NewGetByCurbIDParams() *GetByCurbIDParams {
	var ()
	return &GetByCurbIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetByCurbIDParamsWithTimeout creates a new GetByCurbIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetByCurbIDParamsWithTimeout(timeout time.Duration) *GetByCurbIDParams {
	var ()
	return &GetByCurbIDParams{

		timeout: timeout,
	}
}

// NewGetByCurbIDParamsWithContext creates a new GetByCurbIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetByCurbIDParamsWithContext(ctx context.Context) *GetByCurbIDParams {
	var ()
	return &GetByCurbIDParams{

		Context: ctx,
	}
}

// NewGetByCurbIDParamsWithHTTPClient creates a new GetByCurbIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetByCurbIDParamsWithHTTPClient(client *http.Client) *GetByCurbIDParams {
	var ()
	return &GetByCurbIDParams{
		HTTPClient: client,
	}
}

/*GetByCurbIDParams contains all the parameters to send to the API endpoint
for the get by curb id operation typically these are written to a http.Request
*/
type GetByCurbIDParams struct {

	/*AccessKey
	  The API access key for the request.

	*/
	AccessKey string
	/*ID
	  The ID of the curb to return.

	*/
	ID string
	/*TempRulesWindowEnd
	  We will return any temporary rules we find that are in effect at or before this time, in ISO-8601 format.
	Defaults to 7 days after `temp_rules_window_start`.


	*/
	TempRulesWindowEnd *strfmt.DateTime
	/*TempRulesWindowStart
	  We will return any temporary rules we find that are in effect at or after this time, in ISO-8601 format.
	Defaults to the current time.


	*/
	TempRulesWindowStart *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get by curb id params
func (o *GetByCurbIDParams) WithTimeout(timeout time.Duration) *GetByCurbIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get by curb id params
func (o *GetByCurbIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get by curb id params
func (o *GetByCurbIDParams) WithContext(ctx context.Context) *GetByCurbIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get by curb id params
func (o *GetByCurbIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get by curb id params
func (o *GetByCurbIDParams) WithHTTPClient(client *http.Client) *GetByCurbIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get by curb id params
func (o *GetByCurbIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKey adds the accessKey to the get by curb id params
func (o *GetByCurbIDParams) WithAccessKey(accessKey string) *GetByCurbIDParams {
	o.SetAccessKey(accessKey)
	return o
}

// SetAccessKey adds the accessKey to the get by curb id params
func (o *GetByCurbIDParams) SetAccessKey(accessKey string) {
	o.AccessKey = accessKey
}

// WithID adds the id to the get by curb id params
func (o *GetByCurbIDParams) WithID(id string) *GetByCurbIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get by curb id params
func (o *GetByCurbIDParams) SetID(id string) {
	o.ID = id
}

// WithTempRulesWindowEnd adds the tempRulesWindowEnd to the get by curb id params
func (o *GetByCurbIDParams) WithTempRulesWindowEnd(tempRulesWindowEnd *strfmt.DateTime) *GetByCurbIDParams {
	o.SetTempRulesWindowEnd(tempRulesWindowEnd)
	return o
}

// SetTempRulesWindowEnd adds the tempRulesWindowEnd to the get by curb id params
func (o *GetByCurbIDParams) SetTempRulesWindowEnd(tempRulesWindowEnd *strfmt.DateTime) {
	o.TempRulesWindowEnd = tempRulesWindowEnd
}

// WithTempRulesWindowStart adds the tempRulesWindowStart to the get by curb id params
func (o *GetByCurbIDParams) WithTempRulesWindowStart(tempRulesWindowStart *strfmt.DateTime) *GetByCurbIDParams {
	o.SetTempRulesWindowStart(tempRulesWindowStart)
	return o
}

// SetTempRulesWindowStart adds the tempRulesWindowStart to the get by curb id params
func (o *GetByCurbIDParams) SetTempRulesWindowStart(tempRulesWindowStart *strfmt.DateTime) {
	o.TempRulesWindowStart = tempRulesWindowStart
}

// WriteToRequest writes these params to a swagger request
func (o *GetByCurbIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.TempRulesWindowEnd != nil {

		// query param temp_rules_window_end
		var qrTempRulesWindowEnd strfmt.DateTime
		if o.TempRulesWindowEnd != nil {
			qrTempRulesWindowEnd = *o.TempRulesWindowEnd
		}
		qTempRulesWindowEnd := qrTempRulesWindowEnd.String()
		if qTempRulesWindowEnd != "" {
			if err := r.SetQueryParam("temp_rules_window_end", qTempRulesWindowEnd); err != nil {
				return err
			}
		}

	}

	if o.TempRulesWindowStart != nil {

		// query param temp_rules_window_start
		var qrTempRulesWindowStart strfmt.DateTime
		if o.TempRulesWindowStart != nil {
			qrTempRulesWindowStart = *o.TempRulesWindowStart
		}
		qTempRulesWindowStart := qrTempRulesWindowStart.String()
		if qTempRulesWindowStart != "" {
			if err := r.SetQueryParam("temp_rules_window_start", qTempRulesWindowStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}