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

// NewDevicesGetModuleTelemetryValueParams creates a new DevicesGetModuleTelemetryValueParams object
// with the default values initialized.
func NewDevicesGetModuleTelemetryValueParams() *DevicesGetModuleTelemetryValueParams {
	var ()
	return &DevicesGetModuleTelemetryValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDevicesGetModuleTelemetryValueParamsWithTimeout creates a new DevicesGetModuleTelemetryValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDevicesGetModuleTelemetryValueParamsWithTimeout(timeout time.Duration) *DevicesGetModuleTelemetryValueParams {
	var ()
	return &DevicesGetModuleTelemetryValueParams{

		timeout: timeout,
	}
}

// NewDevicesGetModuleTelemetryValueParamsWithContext creates a new DevicesGetModuleTelemetryValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewDevicesGetModuleTelemetryValueParamsWithContext(ctx context.Context) *DevicesGetModuleTelemetryValueParams {
	var ()
	return &DevicesGetModuleTelemetryValueParams{

		Context: ctx,
	}
}

// NewDevicesGetModuleTelemetryValueParamsWithHTTPClient creates a new DevicesGetModuleTelemetryValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDevicesGetModuleTelemetryValueParamsWithHTTPClient(client *http.Client) *DevicesGetModuleTelemetryValueParams {
	var ()
	return &DevicesGetModuleTelemetryValueParams{
		HTTPClient: client,
	}
}

/*DevicesGetModuleTelemetryValueParams contains all the parameters to send to the API endpoint
for the devices get module telemetry value operation typically these are written to a http.Request
*/
type DevicesGetModuleTelemetryValueParams struct {

	/*DeviceID
	  Unique ID of the device.

	*/
	DeviceID string
	/*ModuleName
	  Name of the device module.

	*/
	ModuleName string
	/*TelemetryName
	  Name of this device telemetry.

	*/
	TelemetryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the devices get module telemetry value params
func (o *DevicesGetModuleTelemetryValueParams) WithTimeout(timeout time.Duration) *DevicesGetModuleTelemetryValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the devices get module telemetry value params
func (o *DevicesGetModuleTelemetryValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the devices get module telemetry value params
func (o *DevicesGetModuleTelemetryValueParams) WithContext(ctx context.Context) *DevicesGetModuleTelemetryValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the devices get module telemetry value params
func (o *DevicesGetModuleTelemetryValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the devices get module telemetry value params
func (o *DevicesGetModuleTelemetryValueParams) WithHTTPClient(client *http.Client) *DevicesGetModuleTelemetryValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the devices get module telemetry value params
func (o *DevicesGetModuleTelemetryValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the devices get module telemetry value params
func (o *DevicesGetModuleTelemetryValueParams) WithDeviceID(deviceID string) *DevicesGetModuleTelemetryValueParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the devices get module telemetry value params
func (o *DevicesGetModuleTelemetryValueParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithModuleName adds the moduleName to the devices get module telemetry value params
func (o *DevicesGetModuleTelemetryValueParams) WithModuleName(moduleName string) *DevicesGetModuleTelemetryValueParams {
	o.SetModuleName(moduleName)
	return o
}

// SetModuleName adds the moduleName to the devices get module telemetry value params
func (o *DevicesGetModuleTelemetryValueParams) SetModuleName(moduleName string) {
	o.ModuleName = moduleName
}

// WithTelemetryName adds the telemetryName to the devices get module telemetry value params
func (o *DevicesGetModuleTelemetryValueParams) WithTelemetryName(telemetryName string) *DevicesGetModuleTelemetryValueParams {
	o.SetTelemetryName(telemetryName)
	return o
}

// SetTelemetryName adds the telemetryName to the devices get module telemetry value params
func (o *DevicesGetModuleTelemetryValueParams) SetTelemetryName(telemetryName string) {
	o.TelemetryName = telemetryName
}

// WriteToRequest writes these params to a swagger request
func (o *DevicesGetModuleTelemetryValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device_id
	if err := r.SetPathParam("device_id", o.DeviceID); err != nil {
		return err
	}

	// path param module_name
	if err := r.SetPathParam("module_name", o.ModuleName); err != nil {
		return err
	}

	// path param telemetry_name
	if err := r.SetPathParam("telemetry_name", o.TelemetryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
