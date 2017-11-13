// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/dlespiau/quay-exporter/pkg/models"
)

// RequestRepoBuildReader is a Reader for the RequestRepoBuild structure.
type RequestRepoBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestRepoBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewRequestRepoBuildCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRequestRepoBuildBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRequestRepoBuildUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRequestRepoBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRequestRepoBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRequestRepoBuildCreated creates a RequestRepoBuildCreated with default headers values
func NewRequestRepoBuildCreated() *RequestRepoBuildCreated {
	return &RequestRepoBuildCreated{}
}

/*RequestRepoBuildCreated handles this case with default header values.

Successful creation
*/
type RequestRepoBuildCreated struct {
}

func (o *RequestRepoBuildCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/build/][%d] requestRepoBuildCreated ", 201)
}

func (o *RequestRepoBuildCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestRepoBuildBadRequest creates a RequestRepoBuildBadRequest with default headers values
func NewRequestRepoBuildBadRequest() *RequestRepoBuildBadRequest {
	return &RequestRepoBuildBadRequest{}
}

/*RequestRepoBuildBadRequest handles this case with default header values.

Bad Request
*/
type RequestRepoBuildBadRequest struct {
	Payload *models.APIError
}

func (o *RequestRepoBuildBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/build/][%d] requestRepoBuildBadRequest  %+v", 400, o.Payload)
}

func (o *RequestRepoBuildBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestRepoBuildUnauthorized creates a RequestRepoBuildUnauthorized with default headers values
func NewRequestRepoBuildUnauthorized() *RequestRepoBuildUnauthorized {
	return &RequestRepoBuildUnauthorized{}
}

/*RequestRepoBuildUnauthorized handles this case with default header values.

Session required
*/
type RequestRepoBuildUnauthorized struct {
	Payload *models.APIError
}

func (o *RequestRepoBuildUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/build/][%d] requestRepoBuildUnauthorized  %+v", 401, o.Payload)
}

func (o *RequestRepoBuildUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestRepoBuildForbidden creates a RequestRepoBuildForbidden with default headers values
func NewRequestRepoBuildForbidden() *RequestRepoBuildForbidden {
	return &RequestRepoBuildForbidden{}
}

/*RequestRepoBuildForbidden handles this case with default header values.

Unauthorized access
*/
type RequestRepoBuildForbidden struct {
	Payload *models.APIError
}

func (o *RequestRepoBuildForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/build/][%d] requestRepoBuildForbidden  %+v", 403, o.Payload)
}

func (o *RequestRepoBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestRepoBuildNotFound creates a RequestRepoBuildNotFound with default headers values
func NewRequestRepoBuildNotFound() *RequestRepoBuildNotFound {
	return &RequestRepoBuildNotFound{}
}

/*RequestRepoBuildNotFound handles this case with default header values.

Not found
*/
type RequestRepoBuildNotFound struct {
	Payload *models.APIError
}

func (o *RequestRepoBuildNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/build/][%d] requestRepoBuildNotFound  %+v", 404, o.Payload)
}

func (o *RequestRepoBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
