// Code generated by go-swagger; DO NOT EDIT.

package stage_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetProjectProjectNameStageStageNameResourceHandlerFunc turns a function with the right signature into a get project project name stage stage name resource handler
type GetProjectProjectNameStageStageNameResourceHandlerFunc func(GetProjectProjectNameStageStageNameResourceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProjectProjectNameStageStageNameResourceHandlerFunc) Handle(params GetProjectProjectNameStageStageNameResourceParams) middleware.Responder {
	return fn(params)
}

// GetProjectProjectNameStageStageNameResourceHandler interface for that can handle valid get project project name stage stage name resource params
type GetProjectProjectNameStageStageNameResourceHandler interface {
	Handle(GetProjectProjectNameStageStageNameResourceParams) middleware.Responder
}

// NewGetProjectProjectNameStageStageNameResource creates a new http.Handler for the get project project name stage stage name resource operation
func NewGetProjectProjectNameStageStageNameResource(ctx *middleware.Context, handler GetProjectProjectNameStageStageNameResourceHandler) *GetProjectProjectNameStageStageNameResource {
	return &GetProjectProjectNameStageStageNameResource{Context: ctx, Handler: handler}
}

/* GetProjectProjectNameStageStageNameResource swagger:route GET /project/{projectName}/stage/{stageName}/resource Stage Resource getProjectProjectNameStageStageNameResource

Get list of stage resources

*/
type GetProjectProjectNameStageStageNameResource struct {
	Context *middleware.Context
	Handler GetProjectProjectNameStageStageNameResourceHandler
}

func (o *GetProjectProjectNameStageStageNameResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetProjectProjectNameStageStageNameResourceParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
