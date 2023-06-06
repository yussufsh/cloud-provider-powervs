// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_placement_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudPlacementgroupsMembersPostParams creates a new PcloudPlacementgroupsMembersPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPlacementgroupsMembersPostParams() *PcloudPlacementgroupsMembersPostParams {
	return &PcloudPlacementgroupsMembersPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPlacementgroupsMembersPostParamsWithTimeout creates a new PcloudPlacementgroupsMembersPostParams object
// with the ability to set a timeout on a request.
func NewPcloudPlacementgroupsMembersPostParamsWithTimeout(timeout time.Duration) *PcloudPlacementgroupsMembersPostParams {
	return &PcloudPlacementgroupsMembersPostParams{
		timeout: timeout,
	}
}

// NewPcloudPlacementgroupsMembersPostParamsWithContext creates a new PcloudPlacementgroupsMembersPostParams object
// with the ability to set a context for a request.
func NewPcloudPlacementgroupsMembersPostParamsWithContext(ctx context.Context) *PcloudPlacementgroupsMembersPostParams {
	return &PcloudPlacementgroupsMembersPostParams{
		Context: ctx,
	}
}

// NewPcloudPlacementgroupsMembersPostParamsWithHTTPClient creates a new PcloudPlacementgroupsMembersPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPlacementgroupsMembersPostParamsWithHTTPClient(client *http.Client) *PcloudPlacementgroupsMembersPostParams {
	return &PcloudPlacementgroupsMembersPostParams{
		HTTPClient: client,
	}
}

/*
PcloudPlacementgroupsMembersPostParams contains all the parameters to send to the API endpoint

	for the pcloud placementgroups members post operation.

	Typically these are written to a http.Request.
*/
type PcloudPlacementgroupsMembersPostParams struct {

	/* Body.

	   Parameters for adding a server to a Server Placement Group
	*/
	Body *models.PlacementGroupServer

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* PlacementGroupID.

	   Placement Group ID
	*/
	PlacementGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud placementgroups members post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPlacementgroupsMembersPostParams) WithDefaults() *PcloudPlacementgroupsMembersPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud placementgroups members post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPlacementgroupsMembersPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud placementgroups members post params
func (o *PcloudPlacementgroupsMembersPostParams) WithTimeout(timeout time.Duration) *PcloudPlacementgroupsMembersPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud placementgroups members post params
func (o *PcloudPlacementgroupsMembersPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud placementgroups members post params
func (o *PcloudPlacementgroupsMembersPostParams) WithContext(ctx context.Context) *PcloudPlacementgroupsMembersPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud placementgroups members post params
func (o *PcloudPlacementgroupsMembersPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud placementgroups members post params
func (o *PcloudPlacementgroupsMembersPostParams) WithHTTPClient(client *http.Client) *PcloudPlacementgroupsMembersPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud placementgroups members post params
func (o *PcloudPlacementgroupsMembersPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud placementgroups members post params
func (o *PcloudPlacementgroupsMembersPostParams) WithBody(body *models.PlacementGroupServer) *PcloudPlacementgroupsMembersPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud placementgroups members post params
func (o *PcloudPlacementgroupsMembersPostParams) SetBody(body *models.PlacementGroupServer) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud placementgroups members post params
func (o *PcloudPlacementgroupsMembersPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPlacementgroupsMembersPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud placementgroups members post params
func (o *PcloudPlacementgroupsMembersPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPlacementGroupID adds the placementGroupID to the pcloud placementgroups members post params
func (o *PcloudPlacementgroupsMembersPostParams) WithPlacementGroupID(placementGroupID string) *PcloudPlacementgroupsMembersPostParams {
	o.SetPlacementGroupID(placementGroupID)
	return o
}

// SetPlacementGroupID adds the placementGroupId to the pcloud placementgroups members post params
func (o *PcloudPlacementgroupsMembersPostParams) SetPlacementGroupID(placementGroupID string) {
	o.PlacementGroupID = placementGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPlacementgroupsMembersPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param placement_group_id
	if err := r.SetPathParam("placement_group_id", o.PlacementGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
