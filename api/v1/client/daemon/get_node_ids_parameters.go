// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

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
)

// NewGetNodeIdsParams creates a new GetNodeIdsParams object
// with the default values initialized.
func NewGetNodeIdsParams() *GetNodeIdsParams {

	return &GetNodeIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodeIdsParamsWithTimeout creates a new GetNodeIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodeIdsParamsWithTimeout(timeout time.Duration) *GetNodeIdsParams {

	return &GetNodeIdsParams{

		timeout: timeout,
	}
}

// NewGetNodeIdsParamsWithContext creates a new GetNodeIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNodeIdsParamsWithContext(ctx context.Context) *GetNodeIdsParams {

	return &GetNodeIdsParams{

		Context: ctx,
	}
}

// NewGetNodeIdsParamsWithHTTPClient creates a new GetNodeIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNodeIdsParamsWithHTTPClient(client *http.Client) *GetNodeIdsParams {

	return &GetNodeIdsParams{
		HTTPClient: client,
	}
}

/*GetNodeIdsParams contains all the parameters to send to the API endpoint
for the get node ids operation typically these are written to a http.Request
*/
type GetNodeIdsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get node ids params
func (o *GetNodeIdsParams) WithTimeout(timeout time.Duration) *GetNodeIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get node ids params
func (o *GetNodeIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get node ids params
func (o *GetNodeIdsParams) WithContext(ctx context.Context) *GetNodeIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get node ids params
func (o *GetNodeIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get node ids params
func (o *GetNodeIdsParams) WithHTTPClient(client *http.Client) *GetNodeIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get node ids params
func (o *GetNodeIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodeIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}