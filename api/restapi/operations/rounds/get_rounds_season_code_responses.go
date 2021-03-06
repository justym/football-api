// Code generated by go-swagger; DO NOT EDIT.

package rounds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/justym/football-api/api/models"
)

// GetRoundsSeasonCodeOKCode is the HTTP code returned for type GetRoundsSeasonCodeOK
const GetRoundsSeasonCodeOKCode int = 200

/*GetRoundsSeasonCodeOK success finding rounds in the seasons

swagger:response getRoundsSeasonCodeOK
*/
type GetRoundsSeasonCodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.SeasonRounds `json:"body,omitempty"`
}

// NewGetRoundsSeasonCodeOK creates GetRoundsSeasonCodeOK with default headers values
func NewGetRoundsSeasonCodeOK() *GetRoundsSeasonCodeOK {

	return &GetRoundsSeasonCodeOK{}
}

// WithPayload adds the payload to the get rounds season code o k response
func (o *GetRoundsSeasonCodeOK) WithPayload(payload *models.SeasonRounds) *GetRoundsSeasonCodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rounds season code o k response
func (o *GetRoundsSeasonCodeOK) SetPayload(payload *models.SeasonRounds) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRoundsSeasonCodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRoundsSeasonCodeBadRequestCode is the HTTP code returned for type GetRoundsSeasonCodeBadRequest
const GetRoundsSeasonCodeBadRequestCode int = 400

/*GetRoundsSeasonCodeBadRequest bad request

swagger:response getRoundsSeasonCodeBadRequest
*/
type GetRoundsSeasonCodeBadRequest struct {
}

// NewGetRoundsSeasonCodeBadRequest creates GetRoundsSeasonCodeBadRequest with default headers values
func NewGetRoundsSeasonCodeBadRequest() *GetRoundsSeasonCodeBadRequest {

	return &GetRoundsSeasonCodeBadRequest{}
}

// WriteResponse to the client
func (o *GetRoundsSeasonCodeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
