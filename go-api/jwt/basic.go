package main

import (
	"context"
	"net/http"

	"github.com/goadesign/goa"
	"github.com/hirokisan/go-sandbox/go-api/jwt/app"
	"github.com/hirokisan/go-sandbox/go-api/lib/util"
)

// NewBasicAuthMiddleware creates a middleware that checks for the presence of a basic auth header
// and validates its content.
func NewBasicAuthMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log basic auth info
			user, pass, ok := req.BasicAuth()
			// A real app would do something more interesting here
			if !ok {
				goa.LogInfo(ctx, "failed basic auth")
				return ErrUnauthorized("missing auth")
			}

			//pass := "sample"
			ph := "$2a$10$z0GL0ObLuhKp8e1I9ZRKUOjB8S9qFlnI4kI7iXB4QcKJHnclPXNLq=MISSING"

			err := util.PasswordVerify(ph, pass)
			if err != nil {
				panic(err)
			}

			// Proceed
			goa.LogInfo(ctx, "basic", "user", user, "pass", pass)
			return h(ctx, rw, req)
		}
	}
}

// BasicController implements the BasicAuth resource.
type BasicController struct {
	*goa.Controller
}

// NewBasicController creates a BasicAuth controller.
func NewBasicController(service *goa.Service) *BasicController {
	return &BasicController{Controller: service.NewController("BasicController")}
}

// Secure runs the secure action.
func (c *BasicController) Secure(ctx *app.SecureBasicContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}

// Unsecure runs the unsecure action.
func (c *BasicController) Unsecure(ctx *app.UnsecureBasicContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}
