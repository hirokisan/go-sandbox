// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "sample": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/hirokisan/go-sandbox/go-api/design/design2
// --out=$(GOPATH)/src/github.com/hirokisan/go-sandbox/go-api/api2
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// ShowSampleContext provides the sample show action context.
type ShowSampleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	BottleID int
}

// NewShowSampleContext parses the incoming request URL and body, performs validations and creates the
// context used by the sample controller show action.
func NewShowSampleContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowSampleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowSampleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramBottleID := req.Params["bottleID"]
	if len(paramBottleID) > 0 {
		rawBottleID := paramBottleID[0]
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowSampleContext) OK(r *GoaExampleSample) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.sample+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowSampleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
