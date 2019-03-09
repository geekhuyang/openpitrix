// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixJob openpitrix job
// swagger:model openpitrixJob
type OpenpitrixJob struct {

	// app id
	AppID string `json:"app_id,omitempty"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// directive
	Directive string `json:"directive,omitempty"`

	// error code
	ErrorCode int64 `json:"error_code,omitempty"`

	// executor
	Executor string `json:"executor,omitempty"`

	// job action
	JobAction string `json:"job_action,omitempty"`

	// job id
	JobID string `json:"job_id,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// owner path
	OwnerPath string `json:"owner_path,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// runtime id
	RuntimeID string `json:"runtime_id,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// task count
	TaskCount int64 `json:"task_count,omitempty"`

	// version id
	VersionID string `json:"version_id,omitempty"`
}

// Validate validates this openpitrix job
func (m *OpenpitrixJob) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixJob) UnmarshalBinary(b []byte) error {
	var res OpenpitrixJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
