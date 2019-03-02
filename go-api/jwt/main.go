//go:generate goagen bootstrap -d github.com/hirokisan/go-sandbox/go-api/design/jwt

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/hirokisan/go-sandbox/go-api/jwt/app"
)

func main() {
	// Create service
	service := goa.New("")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// JWT
	jwtMiddleware, _ := NewJWTMiddleware()
	app.UseJWTMiddleware(service, jwtMiddleware)

	// Mount "jwt" controller
	c, _ := NewJWTController(service)
	app.MountJWTController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
