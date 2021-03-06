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
	"github.com/go-openapi/strfmt"
)

// NewContinuousDataExportsListParams creates a new ContinuousDataExportsListParams object
// with the default values initialized.
func NewContinuousDataExportsListParams() *ContinuousDataExportsListParams {

	return &ContinuousDataExportsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContinuousDataExportsListParamsWithTimeout creates a new ContinuousDataExportsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContinuousDataExportsListParamsWithTimeout(timeout time.Duration) *ContinuousDataExportsListParams {

	return &ContinuousDataExportsListParams{

		timeout: timeout,
	}
}

// NewContinuousDataExportsListParamsWithContext creates a new ContinuousDataExportsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewContinuousDataExportsListParamsWithContext(ctx context.Context) *ContinuousDataExportsListParams {

	return &ContinuousDataExportsListParams{

		Context: ctx,
	}
}

// NewContinuousDataExportsListParamsWithHTTPClient creates a new ContinuousDataExportsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContinuousDataExportsListParamsWithHTTPClient(client *http.Client) *ContinuousDataExportsListParams {

	return &ContinuousDataExportsListParams{
		HTTPClient: client,
	}
}

/*ContinuousDataExportsListParams contains all the parameters to send to the API endpoint
for the continuous data exports list operation typically these are written to a http.Request
*/
type ContinuousDataExportsListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the continuous data exports list params
func (o *ContinuousDataExportsListParams) WithTimeout(timeout time.Duration) *ContinuousDataExportsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the continuous data exports list params
func (o *ContinuousDataExportsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the continuous data exports list params
func (o *ContinuousDataExportsListParams) WithContext(ctx context.Context) *ContinuousDataExportsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the continuous data exports list params
func (o *ContinuousDataExportsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the continuous data exports list params
func (o *ContinuousDataExportsListParams) WithHTTPClient(client *http.Client) *ContinuousDataExportsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the continuous data exports list params
func (o *ContinuousDataExportsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ContinuousDataExportsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
