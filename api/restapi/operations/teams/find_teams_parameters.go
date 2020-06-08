// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewFindTeamsParams creates a new FindTeamsParams object
// no default values defined in spec.
func NewFindTeamsParams() FindTeamsParams {

	return FindTeamsParams{}
}

// FindTeamsParams contains all the bound params for the find teams operation
// typically these are obtained from a http.Request
//
// swagger:parameters findTeams
type FindTeamsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*pass an optional search string for looking up teams
	  Required: true
	  In: path
	*/
	SeasonCode string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFindTeamsParams() beforehand.
func (o *FindTeamsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rSeasonCode, rhkSeasonCode, _ := route.Params.GetOK("seasonCode")
	if err := o.bindSeasonCode(rSeasonCode, rhkSeasonCode, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSeasonCode binds and validates parameter SeasonCode from path.
func (o *FindTeamsParams) bindSeasonCode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.SeasonCode = raw

	return nil
}
