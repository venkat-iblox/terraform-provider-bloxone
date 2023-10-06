// Code generated by go-swagger; DO NOT EDIT.

package global

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

// GlobalUpdateReader is a Reader for the GlobalUpdate structure.
type GlobalUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GlobalUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewGlobalUpdateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewGlobalUpdateCreated creates a GlobalUpdateCreated with default headers values
func NewGlobalUpdateCreated() *GlobalUpdateCreated {
	return &GlobalUpdateCreated{}
}

/* GlobalUpdateCreated describes a response with status code 201, with default header values.

PATCH operation response
*/
type GlobalUpdateCreated struct {
	Payload *models.ConfigUpdateGlobalResponse
}

func (o *GlobalUpdateCreated) Error() string {
	return fmt.Sprintf("[PATCH /dns/global][%d] globalUpdateCreated  %+v", 201, o.Payload)
}
func (o *GlobalUpdateCreated) GetPayload() *models.ConfigUpdateGlobalResponse {
	return o.Payload
}

func (o *GlobalUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigUpdateGlobalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
