// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/stackpath/terraform-provider-stackpath/stackpath/internal/models"
)

// DeleteWorkloadReader is a Reader for the DeleteWorkload structure.
type DeleteWorkloadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteWorkloadNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteWorkloadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteWorkloadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteWorkloadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteWorkloadNoContent creates a DeleteWorkloadNoContent with default headers values
func NewDeleteWorkloadNoContent() *DeleteWorkloadNoContent {
	return &DeleteWorkloadNoContent{}
}

/*DeleteWorkloadNoContent handles this case with default header values.

No content
*/
type DeleteWorkloadNoContent struct {
}

func (o *DeleteWorkloadNoContent) Error() string {
	return fmt.Sprintf("[DELETE /workload/v1/stacks/{stack_id}/workloads/{workload_id}][%d] deleteWorkloadNoContent ", 204)
}

func (o *DeleteWorkloadNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkloadUnauthorized creates a DeleteWorkloadUnauthorized with default headers values
func NewDeleteWorkloadUnauthorized() *DeleteWorkloadUnauthorized {
	return &DeleteWorkloadUnauthorized{}
}

/*DeleteWorkloadUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type DeleteWorkloadUnauthorized struct {
	Payload *models.StackpathapiStatus
}

func (o *DeleteWorkloadUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /workload/v1/stacks/{stack_id}/workloads/{workload_id}][%d] deleteWorkloadUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWorkloadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkloadInternalServerError creates a DeleteWorkloadInternalServerError with default headers values
func NewDeleteWorkloadInternalServerError() *DeleteWorkloadInternalServerError {
	return &DeleteWorkloadInternalServerError{}
}

/*DeleteWorkloadInternalServerError handles this case with default header values.

Internal server error.
*/
type DeleteWorkloadInternalServerError struct {
	Payload *models.StackpathapiStatus
}

func (o *DeleteWorkloadInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /workload/v1/stacks/{stack_id}/workloads/{workload_id}][%d] deleteWorkloadInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWorkloadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkloadDefault creates a DeleteWorkloadDefault with default headers values
func NewDeleteWorkloadDefault(code int) *DeleteWorkloadDefault {
	return &DeleteWorkloadDefault{
		_statusCode: code,
	}
}

/*DeleteWorkloadDefault handles this case with default header values.

Default error structure.
*/
type DeleteWorkloadDefault struct {
	_statusCode int

	Payload *models.StackpathapiStatus
}

// Code gets the status code for the delete workload default response
func (o *DeleteWorkloadDefault) Code() int {
	return o._statusCode
}

func (o *DeleteWorkloadDefault) Error() string {
	return fmt.Sprintf("[DELETE /workload/v1/stacks/{stack_id}/workloads/{workload_id}][%d] DeleteWorkload default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteWorkloadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
