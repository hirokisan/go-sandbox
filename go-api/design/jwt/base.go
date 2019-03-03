package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("jwt", func() {
	Title("jwt")
	Description("jwt")
	Version("v1")
	Scheme("http")
	BasePath("/api/v1")
	Consumes("application/json")
	Produces("application/json")
	Host("localhost:8080")
})
