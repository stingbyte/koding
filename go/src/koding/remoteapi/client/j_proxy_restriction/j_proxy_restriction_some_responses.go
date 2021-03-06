package j_proxy_restriction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JProxyRestrictionSomeReader is a Reader for the JProxyRestrictionSome structure.
type JProxyRestrictionSomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JProxyRestrictionSomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJProxyRestrictionSomeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJProxyRestrictionSomeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJProxyRestrictionSomeOK creates a JProxyRestrictionSomeOK with default headers values
func NewJProxyRestrictionSomeOK() *JProxyRestrictionSomeOK {
	return &JProxyRestrictionSomeOK{}
}

/*JProxyRestrictionSomeOK handles this case with default header values.

Request processed successfully
*/
type JProxyRestrictionSomeOK struct {
	Payload *models.DefaultResponse
}

func (o *JProxyRestrictionSomeOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProxyRestriction.some][%d] jProxyRestrictionSomeOK  %+v", 200, o.Payload)
}

func (o *JProxyRestrictionSomeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJProxyRestrictionSomeUnauthorized creates a JProxyRestrictionSomeUnauthorized with default headers values
func NewJProxyRestrictionSomeUnauthorized() *JProxyRestrictionSomeUnauthorized {
	return &JProxyRestrictionSomeUnauthorized{}
}

/*JProxyRestrictionSomeUnauthorized handles this case with default header values.

Unauthorized request
*/
type JProxyRestrictionSomeUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JProxyRestrictionSomeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProxyRestriction.some][%d] jProxyRestrictionSomeUnauthorized  %+v", 401, o.Payload)
}

func (o *JProxyRestrictionSomeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
