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

// CompactionManagerMetricsTotalCompactionsCompletedGetReader is a Reader for the CompactionManagerMetricsTotalCompactionsCompletedGet structure.
type CompactionManagerMetricsTotalCompactionsCompletedGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompactionManagerMetricsTotalCompactionsCompletedGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCompactionManagerMetricsTotalCompactionsCompletedGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCompactionManagerMetricsTotalCompactionsCompletedGetOK creates a CompactionManagerMetricsTotalCompactionsCompletedGetOK with default headers values
func NewCompactionManagerMetricsTotalCompactionsCompletedGetOK() *CompactionManagerMetricsTotalCompactionsCompletedGetOK {
	return &CompactionManagerMetricsTotalCompactionsCompletedGetOK{}
}

/*CompactionManagerMetricsTotalCompactionsCompletedGetOK handles this case with default header values.

CompactionManagerMetricsTotalCompactionsCompletedGetOK compaction manager metrics total compactions completed get o k
*/
type CompactionManagerMetricsTotalCompactionsCompletedGetOK struct {
	Payload interface{}
}

func (o *CompactionManagerMetricsTotalCompactionsCompletedGetOK) Error() string {
	return fmt.Sprintf("[GET /compaction_manager/metrics/total_compactions_completed][%d] compactionManagerMetricsTotalCompactionsCompletedGetOK  %+v", 200, o.Payload)
}

func (o *CompactionManagerMetricsTotalCompactionsCompletedGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CompactionManagerMetricsTotalCompactionsCompletedGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}