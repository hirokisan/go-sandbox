package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("sample", func() { // API defines the microservice endpoint and
	Title("The virtual wine sample")    // other global properties. There should be one
	Description("A simple goa service") // and exactly one API definition appearing in
	Scheme("http")                      // the design.
	Host("localhost:8080")
})

var _ = Resource("sample", func() { // Resources group related API endpoints
	BasePath("/sample")       // together. They map to REST resources for REST
	DefaultMedia(SampleMedia) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get sample by id") // with its path, parameters (both path
		Routing(GET("/:bottleID"))      // parameters and querystring values) and payload
		Params(func() {                 // (shape of the request body).
			Param("bottleID", Integer, "sample ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})

var SampleMedia = MediaType("application/vnd.sample+json", func() {
	Description("A sample of wine")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique sample ID")
		Attribute("href", String, "API href for making requests on the sample")
		Attribute("name", String, "Name of wine")
		Required("id", "href", "name")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("name")
	})
})
