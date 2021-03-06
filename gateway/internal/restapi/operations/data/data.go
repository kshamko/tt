// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DataHandlerFunc turns a function with the right signature into a data handler
type DataHandlerFunc func(DataParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DataHandlerFunc) Handle(params DataParams) middleware.Responder {
	return fn(params)
}

// DataHandler interface for that can handle valid data params
type DataHandler interface {
	Handle(DataParams) middleware.Responder
}

// NewData creates a new http.Handler for the data operation
func NewData(ctx *middleware.Context, handler DataHandler) *Data {
	return &Data{Context: ctx, Handler: handler}
}

/* Data swagger:route GET /data/{id} Data data

Data data API

*/
type Data struct {
	Context *middleware.Context
	Handler DataHandler
}

func (o *Data) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDataParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
