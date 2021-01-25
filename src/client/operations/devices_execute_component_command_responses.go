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

// DevicesExecuteComponentCommandReader is a Reader for the DevicesExecuteComponentCommand structure.
type DevicesExecuteComponentCommandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevicesExecuteComponentCommandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDevicesExecuteComponentCommandCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDevicesExecuteComponentCommandCreated creates a DevicesExecuteComponentCommandCreated with default headers values
func NewDevicesExecuteComponentCommandCreated() *DevicesExecuteComponentCommandCreated {
	return &DevicesExecuteComponentCommandCreated{}
}

/*DevicesExecuteComponentCommandCreated handles this case with default header values.

Success
*/
type DevicesExecuteComponentCommandCreated struct {
	Payload *models.DeviceCommand
}

func (o *DevicesExecuteComponentCommandCreated) Error() string {
	return fmt.Sprintf("[POST /devices/{device_id}/components/{component_name}/commands/{command_name}][%d] devicesExecuteComponentCommandCreated  %+v", 201, o.Payload)
}

func (o *DevicesExecuteComponentCommandCreated) GetPayload() *models.DeviceCommand {
	return o.Payload
}

func (o *DevicesExecuteComponentCommandCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceCommand)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}