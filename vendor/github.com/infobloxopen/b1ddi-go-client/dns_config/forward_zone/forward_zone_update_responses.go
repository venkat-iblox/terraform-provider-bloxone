// Code generated by go-swagger; DO NOT EDIT.

package forward_zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
	b1cliruntime "github.com/infobloxopen/b1ddi-go-client/runtime"
)

// ForwardZoneUpdateReader is a Reader for the ForwardZoneUpdate structure.
type ForwardZoneUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForwardZoneUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewForwardZoneUpdateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewForwardZoneUpdateCreated creates a ForwardZoneUpdateCreated with default headers values
func NewForwardZoneUpdateCreated() *ForwardZoneUpdateCreated {
	return &ForwardZoneUpdateCreated{}
}

/* ForwardZoneUpdateCreated describes a response with status code 201, with default header values.

PATCH operation response
*/
type ForwardZoneUpdateCreated struct {
	Payload *models.ConfigUpdateForwardZoneResponse
}

func (o *ForwardZoneUpdateCreated) Error() string {
	return fmt.Sprintf("[PATCH /dns/forward_zone/{id}][%d] forwardZoneUpdateCreated  %+v", 201, o.Payload)
}
func (o *ForwardZoneUpdateCreated) GetPayload() *models.ConfigUpdateForwardZoneResponse {
	return o.Payload
}

func (o *ForwardZoneUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigUpdateForwardZoneResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
