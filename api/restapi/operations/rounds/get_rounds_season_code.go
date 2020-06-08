// Code generated by go-swagger; DO NOT EDIT.

package rounds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRoundsSeasonCodeHandlerFunc turns a function with the right signature into a get rounds season code handler
type GetRoundsSeasonCodeHandlerFunc func(GetRoundsSeasonCodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRoundsSeasonCodeHandlerFunc) Handle(params GetRoundsSeasonCodeParams) middleware.Responder {
	return fn(params)
}

// GetRoundsSeasonCodeHandler interface for that can handle valid get rounds season code params
type GetRoundsSeasonCodeHandler interface {
	Handle(GetRoundsSeasonCodeParams) middleware.Responder
}

// NewGetRoundsSeasonCode creates a new http.Handler for the get rounds season code operation
func NewGetRoundsSeasonCode(ctx *middleware.Context, handler GetRoundsSeasonCodeHandler) *GetRoundsSeasonCode {
	return &GetRoundsSeasonCode{Context: ctx, Handler: handler}
}

/*GetRoundsSeasonCode swagger:route GET /rounds/{seasonCode} rounds getRoundsSeasonCode

Teams when the season

Rounds when {seasonCode}


*/
type GetRoundsSeasonCode struct {
	Context *middleware.Context
	Handler GetRoundsSeasonCodeHandler
}

func (o *GetRoundsSeasonCode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRoundsSeasonCodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}