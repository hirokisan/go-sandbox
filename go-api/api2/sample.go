package main

import (
	"github.com/goadesign/goa"
	"github.com/hirokisan/go-sandbox/go-api/api2/app"
)

// SampleController implements the sample resource.
type SampleController struct {
	*goa.Controller
}

// NewSampleController creates a sample controller.
func NewSampleController(service *goa.Service) *SampleController {
	return &SampleController{Controller: service.NewController("SampleController")}
}

// Show runs the show action.
func (c *SampleController) Show(ctx *app.ShowSampleContext) error {
	// SampleController_Show: start_implement

	// Put your logic here

	res := &app.GoaExampleSample{}
	return ctx.OK(res)
	// SampleController_Show: end_implement
}
