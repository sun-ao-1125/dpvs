// Code generated by go-swagger; DO NOT EDIT.

package virtualserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetVsParams creates a new GetVsParams object
// with the default values initialized.
func NewGetVsParams() GetVsParams {

	var (
		// initialize parameters with default values

		statsDefault = bool(false)
	)

	return GetVsParams{
		Stats: &statsDefault,
	}
}

// GetVsParams contains all the bound params for the get vs operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetVs
type GetVsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	  Default: false
	*/
	Stats *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetVsParams() beforehand.
func (o *GetVsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qStats, qhkStats, _ := qs.GetOK("stats")
	if err := o.bindStats(qStats, qhkStats, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindStats binds and validates parameter Stats from query.
func (o *GetVsParams) bindStats(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetVsParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("stats", "query", "bool", raw)
	}
	o.Stats = &value

	return nil
}
