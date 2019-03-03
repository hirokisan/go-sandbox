package main

import (
	"context"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/goadesign/goa"
	"github.com/hirokisan/go-sandbox/go-api/jwt/app"
	"github.com/hirokisan/go-sandbox/go-api/lib/util"
)

type User struct {
	Name string
	Pass string
}

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

			session, _ := mgo.Dial("mongodb://localhost/")
			defer session.Close()
			//u := &User{
			//	Name: un,
			//	Pass: ph,
			//}
			//err := session.DB("local").C("users").Insert(u)
			var u User
			if err := session.DB("local").C("users").Find(bson.M{"name": user}).One(&u); err != nil {
			}
		  goa.LogInfo(ctx, "user", u)

			err := util.PasswordVerify(u.Pass, pass)
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
