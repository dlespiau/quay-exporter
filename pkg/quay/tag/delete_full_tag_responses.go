// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/dlespiau/quay-exporter/pkg/models"
)

// DeleteFullTagReader is a Reader for the DeleteFullTag structure.
type DeleteFullTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFullTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteFullTagNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteFullTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteFullTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteFullTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteFullTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFullTagNoContent creates a DeleteFullTagNoContent with default headers values
func NewDeleteFullTagNoContent() *DeleteFullTagNoContent {
	return &DeleteFullTagNoContent{}
}

/*DeleteFullTagNoContent handles this case with default header values.

Deleted
*/
type DeleteFullTagNoContent struct {
}

func (o *DeleteFullTagNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/tag/{tag}][%d] deleteFullTagNoContent ", 204)
}

func (o *DeleteFullTagNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFullTagBadRequest creates a DeleteFullTagBadRequest with default headers values
func NewDeleteFullTagBadRequest() *DeleteFullTagBadRequest {
	return &DeleteFullTagBadRequest{}
}

/*DeleteFullTagBadRequest handles this case with default header values.

Bad Request
*/
type DeleteFullTagBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteFullTagBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/tag/{tag}][%d] deleteFullTagBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFullTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFullTagUnauthorized creates a DeleteFullTagUnauthorized with default headers values
func NewDeleteFullTagUnauthorized() *DeleteFullTagUnauthorized {
	return &DeleteFullTagUnauthorized{}
}

/*DeleteFullTagUnauthorized handles this case with default header values.

Session required
*/
type DeleteFullTagUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteFullTagUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/tag/{tag}][%d] deleteFullTagUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFullTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFullTagForbidden creates a DeleteFullTagForbidden with default headers values
func NewDeleteFullTagForbidden() *DeleteFullTagForbidden {
	return &DeleteFullTagForbidden{}
}

/*DeleteFullTagForbidden handles this case with default header values.

Unauthorized access
*/
type DeleteFullTagForbidden struct {
	Payload *models.APIError
}

func (o *DeleteFullTagForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/tag/{tag}][%d] deleteFullTagForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFullTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFullTagNotFound creates a DeleteFullTagNotFound with default headers values
func NewDeleteFullTagNotFound() *DeleteFullTagNotFound {
	return &DeleteFullTagNotFound{}
}

/*DeleteFullTagNotFound handles this case with default header values.

Not found
*/
type DeleteFullTagNotFound struct {
	Payload *models.APIError
}

func (o *DeleteFullTagNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/tag/{tag}][%d] deleteFullTagNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFullTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}