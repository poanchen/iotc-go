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

// NewJobsGetDevicesParams creates a new JobsGetDevicesParams object
// with the default values initialized.
func NewJobsGetDevicesParams() *JobsGetDevicesParams {
	var ()
	return &JobsGetDevicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobsGetDevicesParamsWithTimeout creates a new JobsGetDevicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobsGetDevicesParamsWithTimeout(timeout time.Duration) *JobsGetDevicesParams {
	var ()
	return &JobsGetDevicesParams{

		timeout: timeout,
	}
}

// NewJobsGetDevicesParamsWithContext creates a new JobsGetDevicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobsGetDevicesParamsWithContext(ctx context.Context) *JobsGetDevicesParams {
	var ()
	return &JobsGetDevicesParams{

		Context: ctx,
	}
}

// NewJobsGetDevicesParamsWithHTTPClient creates a new JobsGetDevicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobsGetDevicesParamsWithHTTPClient(client *http.Client) *JobsGetDevicesParams {
	var ()
	return &JobsGetDevicesParams{
		HTTPClient: client,
	}
}

/*JobsGetDevicesParams contains all the parameters to send to the API endpoint
for the jobs get devices operation typically these are written to a http.Request
*/
type JobsGetDevicesParams struct {

	/*JobID
	  Unique ID of the job.

	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the jobs get devices params
func (o *JobsGetDevicesParams) WithTimeout(timeout time.Duration) *JobsGetDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the jobs get devices params
func (o *JobsGetDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the jobs get devices params
func (o *JobsGetDevicesParams) WithContext(ctx context.Context) *JobsGetDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the jobs get devices params
func (o *JobsGetDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the jobs get devices params
func (o *JobsGetDevicesParams) WithHTTPClient(client *http.Client) *JobsGetDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the jobs get devices params
func (o *JobsGetDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the jobs get devices params
func (o *JobsGetDevicesParams) WithJobID(jobID string) *JobsGetDevicesParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the jobs get devices params
func (o *JobsGetDevicesParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *JobsGetDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param job_id
	if err := r.SetPathParam("job_id", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}