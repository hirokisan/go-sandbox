// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "jwt": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/hirokisan/go-sandbox/go-api/design/jwt
// --out=$(GOPATH)/src/github.com/hirokisan/go-sandbox/go-api/jwt
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// SecureBasicContext provides the basic secure action context.
type SecureBasicContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewSecureBasicContext parses the incoming request URL and body, performs validations and creates the
// context used by the basic controller secure action.
func NewSecureBasicContext(ctx context.Context, r *http.Request, service *goa.Service) (*SecureBasicContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := SecureBasicContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SecureBasicContext) OK(r *Success) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.examples.security.success")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *SecureBasicContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// UnsecureBasicContext provides the basic unsecure action context.
type UnsecureBasicContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewUnsecureBasicContext parses the incoming request URL and body, performs validations and creates the
// context used by the basic controller unsecure action.
func NewUnsecureBasicContext(ctx context.Context, r *http.Request, service *goa.Service) (*UnsecureBasicContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UnsecureBasicContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UnsecureBasicContext) OK(r *Success) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.examples.security.success")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// SecureJWTContext provides the jwt secure action context.
type SecureJWTContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Fail *bool
}

// NewSecureJWTContext parses the incoming request URL and body, performs validations and creates the
// context used by the jwt controller secure action.
func NewSecureJWTContext(ctx context.Context, r *http.Request, service *goa.Service) (*SecureJWTContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := SecureJWTContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramFail := req.Params["fail"]
	if len(paramFail) > 0 {
		rawFail := paramFail[0]
		if fail, err2 := strconv.ParseBool(rawFail); err2 == nil {
			tmp1 := &fail
			rctx.Fail = tmp1
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("fail", rawFail, "boolean"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SecureJWTContext) OK(r *Success) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.examples.security.success")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *SecureJWTContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// SigninJWTContext provides the jwt signin action context.
type SigninJWTContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewSigninJWTContext parses the incoming request URL and body, performs validations and creates the
// context used by the jwt controller signin action.
func NewSigninJWTContext(ctx context.Context, r *http.Request, service *goa.Service) (*SigninJWTContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := SigninJWTContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *SigninJWTContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *SigninJWTContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// UnsecureJWTContext provides the jwt unsecure action context.
type UnsecureJWTContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewUnsecureJWTContext parses the incoming request URL and body, performs validations and creates the
// context used by the jwt controller unsecure action.
func NewUnsecureJWTContext(ctx context.Context, r *http.Request, service *goa.Service) (*UnsecureJWTContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UnsecureJWTContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UnsecureJWTContext) OK(r *Success) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.examples.security.success")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
