// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StreamManagerMetricsOutgoingGetReader is a Reader for the StreamManagerMetricsOutgoingGet structure.
type StreamManagerMetricsOutgoingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StreamManagerMetricsOutgoingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStreamManagerMetricsOutgoingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStreamManagerMetricsOutgoingGetOK creates a StreamManagerMetricsOutgoingGetOK with default headers values
func NewStreamManagerMetricsOutgoingGetOK() *StreamManagerMetricsOutgoingGetOK {
	return &StreamManagerMetricsOutgoingGetOK{}
}

/*StreamManagerMetricsOutgoingGetOK handles this case with default header values.

StreamManagerMetricsOutgoingGetOK stream manager metrics outgoing get o k
*/
type StreamManagerMetricsOutgoingGetOK struct {
	Payload int32
}

func (o *StreamManagerMetricsOutgoingGetOK) Error() string {
	return fmt.Sprintf("[GET /stream_manager/metrics/outgoing][%d] streamManagerMetricsOutgoingGetOK  %+v", 200, o.Payload)
}

func (o *StreamManagerMetricsOutgoingGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StreamManagerMetricsOutgoingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}