// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "cellar": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/hirokisan/go-sandbox/go-api/design/api1
// --out=$(GOPATH)/src/github.com/hirokisan/go-sandbox/go-api/api1
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// A bottle of wine (default view)
//
// Identifier: application/vnd.bottle+json; view=default
type Bottle struct {
	// API href for making requests on the bottle
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique bottle ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// Name of wine
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the Bottle media type instance.
func (mt *Bottle) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// DecodeBottle decodes the Bottle instance encoded in resp body.
func (c *Client) DecodeBottle(resp *http.Response) (*Bottle, error) {
	var decoded Bottle
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
