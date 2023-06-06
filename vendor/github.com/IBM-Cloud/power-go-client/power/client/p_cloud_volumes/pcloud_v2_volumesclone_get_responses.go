// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudV2VolumescloneGetReader is a Reader for the PcloudV2VolumescloneGet structure.
type PcloudV2VolumescloneGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2VolumescloneGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudV2VolumescloneGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudV2VolumescloneGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudV2VolumescloneGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudV2VolumescloneGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudV2VolumescloneGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudV2VolumescloneGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudV2VolumescloneGetOK creates a PcloudV2VolumescloneGetOK with default headers values
func NewPcloudV2VolumescloneGetOK() *PcloudV2VolumescloneGetOK {
	return &PcloudV2VolumescloneGetOK{}
}

/*
PcloudV2VolumescloneGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudV2VolumescloneGetOK struct {
	Payload *models.VolumesCloneDetail
}

// IsSuccess returns true when this pcloud v2 volumesclone get o k response has a 2xx status code
func (o *PcloudV2VolumescloneGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud v2 volumesclone get o k response has a 3xx status code
func (o *PcloudV2VolumescloneGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone get o k response has a 4xx status code
func (o *PcloudV2VolumescloneGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 volumesclone get o k response has a 5xx status code
func (o *PcloudV2VolumescloneGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone get o k response a status code equal to that given
func (o *PcloudV2VolumescloneGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudV2VolumescloneGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetOK  %+v", 200, o.Payload)
}

func (o *PcloudV2VolumescloneGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetOK  %+v", 200, o.Payload)
}

func (o *PcloudV2VolumescloneGetOK) GetPayload() *models.VolumesCloneDetail {
	return o.Payload
}

func (o *PcloudV2VolumescloneGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumesCloneDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetBadRequest creates a PcloudV2VolumescloneGetBadRequest with default headers values
func NewPcloudV2VolumescloneGetBadRequest() *PcloudV2VolumescloneGetBadRequest {
	return &PcloudV2VolumescloneGetBadRequest{}
}

/*
PcloudV2VolumescloneGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudV2VolumescloneGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone get bad request response has a 2xx status code
func (o *PcloudV2VolumescloneGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone get bad request response has a 3xx status code
func (o *PcloudV2VolumescloneGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone get bad request response has a 4xx status code
func (o *PcloudV2VolumescloneGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone get bad request response has a 5xx status code
func (o *PcloudV2VolumescloneGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone get bad request response a status code equal to that given
func (o *PcloudV2VolumescloneGetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudV2VolumescloneGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudV2VolumescloneGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudV2VolumescloneGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetUnauthorized creates a PcloudV2VolumescloneGetUnauthorized with default headers values
func NewPcloudV2VolumescloneGetUnauthorized() *PcloudV2VolumescloneGetUnauthorized {
	return &PcloudV2VolumescloneGetUnauthorized{}
}

/*
PcloudV2VolumescloneGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudV2VolumescloneGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone get unauthorized response has a 2xx status code
func (o *PcloudV2VolumescloneGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone get unauthorized response has a 3xx status code
func (o *PcloudV2VolumescloneGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone get unauthorized response has a 4xx status code
func (o *PcloudV2VolumescloneGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone get unauthorized response has a 5xx status code
func (o *PcloudV2VolumescloneGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone get unauthorized response a status code equal to that given
func (o *PcloudV2VolumescloneGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudV2VolumescloneGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV2VolumescloneGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV2VolumescloneGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetForbidden creates a PcloudV2VolumescloneGetForbidden with default headers values
func NewPcloudV2VolumescloneGetForbidden() *PcloudV2VolumescloneGetForbidden {
	return &PcloudV2VolumescloneGetForbidden{}
}

/*
PcloudV2VolumescloneGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudV2VolumescloneGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone get forbidden response has a 2xx status code
func (o *PcloudV2VolumescloneGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone get forbidden response has a 3xx status code
func (o *PcloudV2VolumescloneGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone get forbidden response has a 4xx status code
func (o *PcloudV2VolumescloneGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone get forbidden response has a 5xx status code
func (o *PcloudV2VolumescloneGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone get forbidden response a status code equal to that given
func (o *PcloudV2VolumescloneGetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudV2VolumescloneGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudV2VolumescloneGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudV2VolumescloneGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetNotFound creates a PcloudV2VolumescloneGetNotFound with default headers values
func NewPcloudV2VolumescloneGetNotFound() *PcloudV2VolumescloneGetNotFound {
	return &PcloudV2VolumescloneGetNotFound{}
}

/*
PcloudV2VolumescloneGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudV2VolumescloneGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone get not found response has a 2xx status code
func (o *PcloudV2VolumescloneGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone get not found response has a 3xx status code
func (o *PcloudV2VolumescloneGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone get not found response has a 4xx status code
func (o *PcloudV2VolumescloneGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone get not found response has a 5xx status code
func (o *PcloudV2VolumescloneGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone get not found response a status code equal to that given
func (o *PcloudV2VolumescloneGetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudV2VolumescloneGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudV2VolumescloneGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudV2VolumescloneGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetInternalServerError creates a PcloudV2VolumescloneGetInternalServerError with default headers values
func NewPcloudV2VolumescloneGetInternalServerError() *PcloudV2VolumescloneGetInternalServerError {
	return &PcloudV2VolumescloneGetInternalServerError{}
}

/*
PcloudV2VolumescloneGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudV2VolumescloneGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone get internal server error response has a 2xx status code
func (o *PcloudV2VolumescloneGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone get internal server error response has a 3xx status code
func (o *PcloudV2VolumescloneGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone get internal server error response has a 4xx status code
func (o *PcloudV2VolumescloneGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 volumesclone get internal server error response has a 5xx status code
func (o *PcloudV2VolumescloneGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud v2 volumesclone get internal server error response a status code equal to that given
func (o *PcloudV2VolumescloneGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudV2VolumescloneGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2VolumescloneGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2VolumescloneGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
