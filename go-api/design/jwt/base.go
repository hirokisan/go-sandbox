package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("jwt", func() { // API defines the microservice endpoint and
	Title("jwt")       // other global properties. There should be one
	Description("jwt") // and exactly one API definition appearing in
	Scheme("http")     // the design.
	Host("localhost:8080")
})
