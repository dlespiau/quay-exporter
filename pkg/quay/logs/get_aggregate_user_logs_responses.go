// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/dlespiau/quay-exporter/pkg/models"
)

// GetAggregateUserLogsReader is a Reader for the GetAggregateUserLogs structure.
type GetAggregateUserLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAggregateUserLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAggregateUserLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAggregateUserLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAggregateUserLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAggregateUserLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAggregateUserLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAggregateUserLogsOK creates a GetAggregateUserLogsOK with default headers values
func NewGetAggregateUserLogsOK() *GetAggregateUserLogsOK {
	return &GetAggregateUserLogsOK{}
}

/*GetAggregateUserLogsOK handles this case with default header values.

Successful invocation
*/
type GetAggregateUserLogsOK struct {
}

func (o *GetAggregateUserLogsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/aggregatelogs][%d] getAggregateUserLogsOK ", 200)
}

func (o *GetAggregateUserLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAggregateUserLogsBadRequest creates a GetAggregateUserLogsBadRequest with default headers values
func NewGetAggregateUserLogsBadRequest() *GetAggregateUserLogsBadRequest {
	return &GetAggregateUserLogsBadRequest{}
}

/*GetAggregateUserLogsBadRequest handles this case with default header values.

Bad Request
*/
type GetAggregateUserLogsBadRequest struct {
	Payload *models.APIError
}

func (o *GetAggregateUserLogsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/aggregatelogs][%d] getAggregateUserLogsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAggregateUserLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAggregateUserLogsUnauthorized creates a GetAggregateUserLogsUnauthorized with default headers values
func NewGetAggregateUserLogsUnauthorized() *GetAggregateUserLogsUnauthorized {
	return &GetAggregateUserLogsUnauthorized{}
}

/*GetAggregateUserLogsUnauthorized handles this case with default header values.

Session required
*/
type GetAggregateUserLogsUnauthorized struct {
	Payload *models.APIError
}

func (o *GetAggregateUserLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/aggregatelogs][%d] getAggregateUserLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAggregateUserLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAggregateUserLogsForbidden creates a GetAggregateUserLogsForbidden with default headers values
func NewGetAggregateUserLogsForbidden() *GetAggregateUserLogsForbidden {
	return &GetAggregateUserLogsForbidden{}
}

/*GetAggregateUserLogsForbidden handles this case with default header values.

Unauthorized access
*/
type GetAggregateUserLogsForbidden struct {
	Payload *models.APIError
}

func (o *GetAggregateUserLogsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/aggregatelogs][%d] getAggregateUserLogsForbidden  %+v", 403, o.Payload)
}

func (o *GetAggregateUserLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAggregateUserLogsNotFound creates a GetAggregateUserLogsNotFound with default headers values
func NewGetAggregateUserLogsNotFound() *GetAggregateUserLogsNotFound {
	return &GetAggregateUserLogsNotFound{}
}

/*GetAggregateUserLogsNotFound handles this case with default header values.

Not found
*/
type GetAggregateUserLogsNotFound struct {
	Payload *models.APIError
}

func (o *GetAggregateUserLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/aggregatelogs][%d] getAggregateUserLogsNotFound  %+v", 404, o.Payload)
}

func (o *GetAggregateUserLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
