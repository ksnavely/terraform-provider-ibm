// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewGetVpcsIDParams creates a new GetVpcsIDParams object
// with the default values initialized.
func NewGetVpcsIDParams() *GetVpcsIDParams {
	var ()
	return &GetVpcsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVpcsIDParamsWithTimeout creates a new GetVpcsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVpcsIDParamsWithTimeout(timeout time.Duration) *GetVpcsIDParams {
	var ()
	return &GetVpcsIDParams{

		timeout: timeout,
	}
}

// NewGetVpcsIDParamsWithContext creates a new GetVpcsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVpcsIDParamsWithContext(ctx context.Context) *GetVpcsIDParams {
	var ()
	return &GetVpcsIDParams{

		Context: ctx,
	}
}

// NewGetVpcsIDParamsWithHTTPClient creates a new GetVpcsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVpcsIDParamsWithHTTPClient(client *http.Client) *GetVpcsIDParams {
	var ()
	return &GetVpcsIDParams{
		HTTPClient: client,
	}
}

/*GetVpcsIDParams contains all the parameters to send to the API endpoint
for the get vpcs ID operation typically these are written to a http.Request
*/
type GetVpcsIDParams struct {

	/*ID
	  The VPC idenitifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vpcs ID params
func (o *GetVpcsIDParams) WithTimeout(timeout time.Duration) *GetVpcsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vpcs ID params
func (o *GetVpcsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vpcs ID params
func (o *GetVpcsIDParams) WithContext(ctx context.Context) *GetVpcsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vpcs ID params
func (o *GetVpcsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vpcs ID params
func (o *GetVpcsIDParams) WithHTTPClient(client *http.Client) *GetVpcsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vpcs ID params
func (o *GetVpcsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get vpcs ID params
func (o *GetVpcsIDParams) WithID(id string) *GetVpcsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vpcs ID params
func (o *GetVpcsIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get vpcs ID params
func (o *GetVpcsIDParams) WithVersion(version string) *GetVpcsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get vpcs ID params
func (o *GetVpcsIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetVpcsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}