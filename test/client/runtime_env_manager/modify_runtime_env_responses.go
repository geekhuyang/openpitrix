// Code generated by go-swagger; DO NOT EDIT.

package runtime_env_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// ModifyRuntimeEnvReader is a Reader for the ModifyRuntimeEnv structure.
type ModifyRuntimeEnvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyRuntimeEnvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyRuntimeEnvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyRuntimeEnvOK creates a ModifyRuntimeEnvOK with default headers values
func NewModifyRuntimeEnvOK() *ModifyRuntimeEnvOK {
	return &ModifyRuntimeEnvOK{}
}

/*ModifyRuntimeEnvOK handles this case with default header values.

ModifyRuntimeEnvOK modify runtime env o k
*/
type ModifyRuntimeEnvOK struct {
	Payload *models.OpenpitrixModifyRuntimeEnvResponse
}

func (o *ModifyRuntimeEnvOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/runtime_envs][%d] modifyRuntimeEnvOK  %+v", 200, o.Payload)
}

func (o *ModifyRuntimeEnvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixModifyRuntimeEnvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}