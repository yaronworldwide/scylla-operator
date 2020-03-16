// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla_v2/models"
)

// FindConfigTrickleFsyncIntervalInKbReader is a Reader for the FindConfigTrickleFsyncIntervalInKb structure.
type FindConfigTrickleFsyncIntervalInKbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigTrickleFsyncIntervalInKbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigTrickleFsyncIntervalInKbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigTrickleFsyncIntervalInKbDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigTrickleFsyncIntervalInKbOK creates a FindConfigTrickleFsyncIntervalInKbOK with default headers values
func NewFindConfigTrickleFsyncIntervalInKbOK() *FindConfigTrickleFsyncIntervalInKbOK {
	return &FindConfigTrickleFsyncIntervalInKbOK{}
}

/*FindConfigTrickleFsyncIntervalInKbOK handles this case with default header values.

Config value
*/
type FindConfigTrickleFsyncIntervalInKbOK struct {
	Payload int64
}

func (o *FindConfigTrickleFsyncIntervalInKbOK) Error() string {
	return fmt.Sprintf("[GET /config/trickle_fsync_interval_in_kb][%d] findConfigTrickleFsyncIntervalInKbOK  %+v", 200, o.Payload)
}

func (o *FindConfigTrickleFsyncIntervalInKbOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigTrickleFsyncIntervalInKbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigTrickleFsyncIntervalInKbDefault creates a FindConfigTrickleFsyncIntervalInKbDefault with default headers values
func NewFindConfigTrickleFsyncIntervalInKbDefault(code int) *FindConfigTrickleFsyncIntervalInKbDefault {
	return &FindConfigTrickleFsyncIntervalInKbDefault{
		_statusCode: code,
	}
}

/*FindConfigTrickleFsyncIntervalInKbDefault handles this case with default header values.

unexpected error
*/
type FindConfigTrickleFsyncIntervalInKbDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config trickle fsync interval in kb default response
func (o *FindConfigTrickleFsyncIntervalInKbDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigTrickleFsyncIntervalInKbDefault) Error() string {
	return fmt.Sprintf("[GET /config/trickle_fsync_interval_in_kb][%d] find_config_trickle_fsync_interval_in_kb default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigTrickleFsyncIntervalInKbDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigTrickleFsyncIntervalInKbDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}