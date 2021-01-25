// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"com.azure.iot/iotcentral/iotcgo/models"
)

// DevicesGetModulePropertiesReader is a Reader for the DevicesGetModuleProperties structure.
type DevicesGetModulePropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevicesGetModulePropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDevicesGetModulePropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDevicesGetModulePropertiesOK creates a DevicesGetModulePropertiesOK with default headers values
func NewDevicesGetModulePropertiesOK() *DevicesGetModulePropertiesOK {
	return &DevicesGetModulePropertiesOK{}
}

/*DevicesGetModulePropertiesOK handles this case with default header values.

Success
*/
type DevicesGetModulePropertiesOK struct {
	Payload models.DeviceProperties
}

func (o *DevicesGetModulePropertiesOK) Error() string {
	return fmt.Sprintf("[GET /devices/{device_id}/modules/{module_name}/properties][%d] devicesGetModulePropertiesOK  %+v", 200, o.Payload)
}

func (o *DevicesGetModulePropertiesOK) GetPayload() models.DeviceProperties {
	return o.Payload
}

func (o *DevicesGetModulePropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}