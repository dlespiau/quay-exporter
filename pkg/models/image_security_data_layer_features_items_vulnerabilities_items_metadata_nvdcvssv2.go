// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadataNVDCVSSv2 image security data layer features items vulnerabilities items metadata n v d c v s sv2
// swagger:model imageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadataNVDCVSSv2
type ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadataNVDCVSSv2 struct {

	// score
	Score float64 `json:"Score,omitempty"`

	// vectors
	Vectors string `json:"Vectors,omitempty"`
}

// Validate validates this image security data layer features items vulnerabilities items metadata n v d c v s sv2
func (m *ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadataNVDCVSSv2) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadataNVDCVSSv2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadataNVDCVSSv2) UnmarshalBinary(b []byte) error {
	var res ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItemsMetadataNVDCVSSv2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
