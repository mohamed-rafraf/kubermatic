// Code generated by go-swagger; DO NOT EDIT.

package vmwareclouddirector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListVMwareCloudDirectorTemplatesReader is a Reader for the ListVMwareCloudDirectorTemplates structure.
type ListVMwareCloudDirectorTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVMwareCloudDirectorTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVMwareCloudDirectorTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVMwareCloudDirectorTemplatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVMwareCloudDirectorTemplatesOK creates a ListVMwareCloudDirectorTemplatesOK with default headers values
func NewListVMwareCloudDirectorTemplatesOK() *ListVMwareCloudDirectorTemplatesOK {
	return &ListVMwareCloudDirectorTemplatesOK{}
}

/* ListVMwareCloudDirectorTemplatesOK describes a response with status code 200, with default header values.

VMwareCloudDirectorTemplateList
*/
type ListVMwareCloudDirectorTemplatesOK struct {
	Payload models.VMwareCloudDirectorTemplateList
}

func (o *ListVMwareCloudDirectorTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/vmwareclouddirector/{dc}/templates/{catalog_name}][%d] listVMwareCloudDirectorTemplatesOK  %+v", 200, o.Payload)
}
func (o *ListVMwareCloudDirectorTemplatesOK) GetPayload() models.VMwareCloudDirectorTemplateList {
	return o.Payload
}

func (o *ListVMwareCloudDirectorTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVMwareCloudDirectorTemplatesDefault creates a ListVMwareCloudDirectorTemplatesDefault with default headers values
func NewListVMwareCloudDirectorTemplatesDefault(code int) *ListVMwareCloudDirectorTemplatesDefault {
	return &ListVMwareCloudDirectorTemplatesDefault{
		_statusCode: code,
	}
}

/* ListVMwareCloudDirectorTemplatesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListVMwareCloudDirectorTemplatesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list v mware cloud director templates default response
func (o *ListVMwareCloudDirectorTemplatesDefault) Code() int {
	return o._statusCode
}

func (o *ListVMwareCloudDirectorTemplatesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/vmwareclouddirector/{dc}/templates/{catalog_name}][%d] listVMwareCloudDirectorTemplates default  %+v", o._statusCode, o.Payload)
}
func (o *ListVMwareCloudDirectorTemplatesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListVMwareCloudDirectorTemplatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}