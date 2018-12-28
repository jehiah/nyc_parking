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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetByLocationParams creates a new GetByLocationParams object
// with the default values initialized.
func NewGetByLocationParams() *GetByLocationParams {
	var ()
	return &GetByLocationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetByLocationParamsWithTimeout creates a new GetByLocationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetByLocationParamsWithTimeout(timeout time.Duration) *GetByLocationParams {
	var ()
	return &GetByLocationParams{

		timeout: timeout,
	}
}

// NewGetByLocationParamsWithContext creates a new GetByLocationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetByLocationParamsWithContext(ctx context.Context) *GetByLocationParams {
	var ()
	return &GetByLocationParams{

		Context: ctx,
	}
}

// NewGetByLocationParamsWithHTTPClient creates a new GetByLocationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetByLocationParamsWithHTTPClient(client *http.Client) *GetByLocationParams {
	var ()
	return &GetByLocationParams{
		HTTPClient: client,
	}
}

/*GetByLocationParams contains all the parameters to send to the API endpoint
for the get by location operation typically these are written to a http.Request
*/
type GetByLocationParams struct {

	/*AccessKey
	  The API access key for the request.

	*/
	AccessKey string
	/*Latitude
	  Latitude to return results for.

	*/
	Latitude float64
	/*Longitude
	  Longitude to return results for.

	*/
	Longitude float64
	/*PermittedUse
	  If set, only return segments that permit this use at some point in time.

	*/
	PermittedUse *string
	/*PrimaryUse
	  If set, only return segments that have this primary use at some point in time.


	*/
	PrimaryUse *string
	/*RadiusKm
	  Distance, in kilometers, from (latitude, longitude) we will return
	results for.


	*/
	RadiusKm float64
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
	/*VehicleType
	  If `primary_use` or `permitted_use` is set AND `vehicle_type` is set, only return
	segments whose primary or permitted use, respectively, includes this vehicle type.
	`vehicle_type=all` will only return segments that permit every vehicle to engage
	in the use, while setting `primary_use` or `permitted_use` **without** setting
	`vehicle_type` returns segments that allow at least one kind of vehicle to engage
	in the use.


	*/
	VehicleType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get by location params
func (o *GetByLocationParams) WithTimeout(timeout time.Duration) *GetByLocationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get by location params
func (o *GetByLocationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get by location params
func (o *GetByLocationParams) WithContext(ctx context.Context) *GetByLocationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get by location params
func (o *GetByLocationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get by location params
func (o *GetByLocationParams) WithHTTPClient(client *http.Client) *GetByLocationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get by location params
func (o *GetByLocationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKey adds the accessKey to the get by location params
func (o *GetByLocationParams) WithAccessKey(accessKey string) *GetByLocationParams {
	o.SetAccessKey(accessKey)
	return o
}

// SetAccessKey adds the accessKey to the get by location params
func (o *GetByLocationParams) SetAccessKey(accessKey string) {
	o.AccessKey = accessKey
}

// WithLatitude adds the latitude to the get by location params
func (o *GetByLocationParams) WithLatitude(latitude float64) *GetByLocationParams {
	o.SetLatitude(latitude)
	return o
}

// SetLatitude adds the latitude to the get by location params
func (o *GetByLocationParams) SetLatitude(latitude float64) {
	o.Latitude = latitude
}

// WithLongitude adds the longitude to the get by location params
func (o *GetByLocationParams) WithLongitude(longitude float64) *GetByLocationParams {
	o.SetLongitude(longitude)
	return o
}

// SetLongitude adds the longitude to the get by location params
func (o *GetByLocationParams) SetLongitude(longitude float64) {
	o.Longitude = longitude
}

// WithPermittedUse adds the permittedUse to the get by location params
func (o *GetByLocationParams) WithPermittedUse(permittedUse *string) *GetByLocationParams {
	o.SetPermittedUse(permittedUse)
	return o
}

// SetPermittedUse adds the permittedUse to the get by location params
func (o *GetByLocationParams) SetPermittedUse(permittedUse *string) {
	o.PermittedUse = permittedUse
}

// WithPrimaryUse adds the primaryUse to the get by location params
func (o *GetByLocationParams) WithPrimaryUse(primaryUse *string) *GetByLocationParams {
	o.SetPrimaryUse(primaryUse)
	return o
}

// SetPrimaryUse adds the primaryUse to the get by location params
func (o *GetByLocationParams) SetPrimaryUse(primaryUse *string) {
	o.PrimaryUse = primaryUse
}

// WithRadiusKm adds the radiusKm to the get by location params
func (o *GetByLocationParams) WithRadiusKm(radiusKm float64) *GetByLocationParams {
	o.SetRadiusKm(radiusKm)
	return o
}

// SetRadiusKm adds the radiusKm to the get by location params
func (o *GetByLocationParams) SetRadiusKm(radiusKm float64) {
	o.RadiusKm = radiusKm
}

// WithTempRulesWindowEnd adds the tempRulesWindowEnd to the get by location params
func (o *GetByLocationParams) WithTempRulesWindowEnd(tempRulesWindowEnd *strfmt.DateTime) *GetByLocationParams {
	o.SetTempRulesWindowEnd(tempRulesWindowEnd)
	return o
}

// SetTempRulesWindowEnd adds the tempRulesWindowEnd to the get by location params
func (o *GetByLocationParams) SetTempRulesWindowEnd(tempRulesWindowEnd *strfmt.DateTime) {
	o.TempRulesWindowEnd = tempRulesWindowEnd
}

// WithTempRulesWindowStart adds the tempRulesWindowStart to the get by location params
func (o *GetByLocationParams) WithTempRulesWindowStart(tempRulesWindowStart *strfmt.DateTime) *GetByLocationParams {
	o.SetTempRulesWindowStart(tempRulesWindowStart)
	return o
}

// SetTempRulesWindowStart adds the tempRulesWindowStart to the get by location params
func (o *GetByLocationParams) SetTempRulesWindowStart(tempRulesWindowStart *strfmt.DateTime) {
	o.TempRulesWindowStart = tempRulesWindowStart
}

// WithVehicleType adds the vehicleType to the get by location params
func (o *GetByLocationParams) WithVehicleType(vehicleType *string) *GetByLocationParams {
	o.SetVehicleType(vehicleType)
	return o
}

// SetVehicleType adds the vehicleType to the get by location params
func (o *GetByLocationParams) SetVehicleType(vehicleType *string) {
	o.VehicleType = vehicleType
}

// WriteToRequest writes these params to a swagger request
func (o *GetByLocationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param latitude
	qrLatitude := o.Latitude
	qLatitude := swag.FormatFloat64(qrLatitude)
	if qLatitude != "" {
		if err := r.SetQueryParam("latitude", qLatitude); err != nil {
			return err
		}
	}

	// query param longitude
	qrLongitude := o.Longitude
	qLongitude := swag.FormatFloat64(qrLongitude)
	if qLongitude != "" {
		if err := r.SetQueryParam("longitude", qLongitude); err != nil {
			return err
		}
	}

	if o.PermittedUse != nil {

		// query param permitted_use
		var qrPermittedUse string
		if o.PermittedUse != nil {
			qrPermittedUse = *o.PermittedUse
		}
		qPermittedUse := qrPermittedUse
		if qPermittedUse != "" {
			if err := r.SetQueryParam("permitted_use", qPermittedUse); err != nil {
				return err
			}
		}

	}

	if o.PrimaryUse != nil {

		// query param primary_use
		var qrPrimaryUse string
		if o.PrimaryUse != nil {
			qrPrimaryUse = *o.PrimaryUse
		}
		qPrimaryUse := qrPrimaryUse
		if qPrimaryUse != "" {
			if err := r.SetQueryParam("primary_use", qPrimaryUse); err != nil {
				return err
			}
		}

	}

	// query param radius_km
	qrRadiusKm := o.RadiusKm
	qRadiusKm := swag.FormatFloat64(qrRadiusKm)
	if qRadiusKm != "" {
		if err := r.SetQueryParam("radius_km", qRadiusKm); err != nil {
			return err
		}
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

	if o.VehicleType != nil {

		// query param vehicle_type
		var qrVehicleType string
		if o.VehicleType != nil {
			qrVehicleType = *o.VehicleType
		}
		qVehicleType := qrVehicleType
		if qVehicleType != "" {
			if err := r.SetQueryParam("vehicle_type", qVehicleType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}