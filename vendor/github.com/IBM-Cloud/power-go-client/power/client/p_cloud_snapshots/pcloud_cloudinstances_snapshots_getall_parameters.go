// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_snapshots

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

// NewPcloudCloudinstancesSnapshotsGetallParams creates a new PcloudCloudinstancesSnapshotsGetallParams object
// with the default values initialized.
func NewPcloudCloudinstancesSnapshotsGetallParams() *PcloudCloudinstancesSnapshotsGetallParams {
	var ()
	return &PcloudCloudinstancesSnapshotsGetallParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesSnapshotsGetallParamsWithTimeout creates a new PcloudCloudinstancesSnapshotsGetallParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudCloudinstancesSnapshotsGetallParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesSnapshotsGetallParams {
	var ()
	return &PcloudCloudinstancesSnapshotsGetallParams{

		timeout: timeout,
	}
}

// NewPcloudCloudinstancesSnapshotsGetallParamsWithContext creates a new PcloudCloudinstancesSnapshotsGetallParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudCloudinstancesSnapshotsGetallParamsWithContext(ctx context.Context) *PcloudCloudinstancesSnapshotsGetallParams {
	var ()
	return &PcloudCloudinstancesSnapshotsGetallParams{

		Context: ctx,
	}
}

// NewPcloudCloudinstancesSnapshotsGetallParamsWithHTTPClient creates a new PcloudCloudinstancesSnapshotsGetallParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudCloudinstancesSnapshotsGetallParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesSnapshotsGetallParams {
	var ()
	return &PcloudCloudinstancesSnapshotsGetallParams{
		HTTPClient: client,
	}
}

/*PcloudCloudinstancesSnapshotsGetallParams contains all the parameters to send to the API endpoint
for the pcloud cloudinstances snapshots getall operation typically these are written to a http.Request
*/
type PcloudCloudinstancesSnapshotsGetallParams struct {

	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud cloudinstances snapshots getall params
func (o *PcloudCloudinstancesSnapshotsGetallParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesSnapshotsGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances snapshots getall params
func (o *PcloudCloudinstancesSnapshotsGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances snapshots getall params
func (o *PcloudCloudinstancesSnapshotsGetallParams) WithContext(ctx context.Context) *PcloudCloudinstancesSnapshotsGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances snapshots getall params
func (o *PcloudCloudinstancesSnapshotsGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances snapshots getall params
func (o *PcloudCloudinstancesSnapshotsGetallParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesSnapshotsGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances snapshots getall params
func (o *PcloudCloudinstancesSnapshotsGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances snapshots getall params
func (o *PcloudCloudinstancesSnapshotsGetallParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesSnapshotsGetallParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances snapshots getall params
func (o *PcloudCloudinstancesSnapshotsGetallParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesSnapshotsGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}