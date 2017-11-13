// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddLabel Adds a label to a manifest
// swagger:model AddLabel
type AddLabel struct {

	// The key for the label
	// Required: true
	Key *string `json:"key"`

	// The media type for this label
	// Required: true
	MediaType *string `json:"media_type"`

	// The value for the label
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this add label
func (m *AddLabel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddLabel) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

var addLabelTypeMediaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["text/plain","application/json"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addLabelTypeMediaTypePropEnum = append(addLabelTypeMediaTypePropEnum, v)
	}
}

const (
	// AddLabelMediaTypeTextPlain captures enum value "text/plain"
	AddLabelMediaTypeTextPlain string = "text/plain"
	// AddLabelMediaTypeApplicationJSON captures enum value "application/json"
	AddLabelMediaTypeApplicationJSON string = "application/json"
)

// prop value enum
func (m *AddLabel) validateMediaTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addLabelTypeMediaTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AddLabel) validateMediaType(formats strfmt.Registry) error {

	if err := validate.Required("media_type", "body", m.MediaType); err != nil {
		return err
	}

	// value enum
	if err := m.validateMediaTypeEnum("media_type", "body", *m.MediaType); err != nil {
		return err
	}

	return nil
}

func (m *AddLabel) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddLabel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddLabel) UnmarshalBinary(b []byte) error {
	var res AddLabel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
