//go:generate goagen bootstrap -d github.com/hirokisan/go-sandbox/go-api/design/design2

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/hirokisan/go-sandbox/go-api/api2/app"
)

func main() {
	// Create service
	service := goa.New("sample")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "sample" controller
	c := NewSampleController(service)
	app.MountSampleController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
