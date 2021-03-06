// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "jwt": jwt Resource Client
//
// Command:
// $ goagen
// --design=github.com/hirokisan/go-sandbox/go-api/design/jwt
// --out=$(GOPATH)/src/github.com/hirokisan/go-sandbox/go-api/jwt
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// SecureJWTPath computes a request path to the secure action of jwt.
func SecureJWTPath() string {

	return fmt.Sprintf("/api/v1/jwt")
}

// This action is secured with the jwt scheme
func (c *Client) SecureJWT(ctx context.Context, path string, fail *bool) (*http.Response, error) {
	req, err := c.NewSecureJWTRequest(ctx, path, fail)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSecureJWTRequest create the request corresponding to the secure action endpoint of the jwt resource.
func (c *Client) NewSecureJWTRequest(ctx context.Context, path string, fail *bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if fail != nil {
		tmp7 := strconv.FormatBool(*fail)
		values.Set("fail", tmp7)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// SigninJWTPath computes a request path to the signin action of jwt.
func SigninJWTPath() string {

	return fmt.Sprintf("/api/v1/jwt/signin")
}

// Creates a valid JWT
func (c *Client) SigninJWT(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSigninJWTRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSigninJWTRequest create the request corresponding to the signin action endpoint of the jwt resource.
func (c *Client) NewSigninJWTRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.SigninBasicAuthSigner != nil {
		if err := c.SigninBasicAuthSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// UnsecureJWTPath computes a request path to the unsecure action of jwt.
func UnsecureJWTPath() string {

	return fmt.Sprintf("/api/v1/jwt/unsecure")
}

// This action does not require auth
func (c *Client) UnsecureJWT(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewUnsecureJWTRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUnsecureJWTRequest create the request corresponding to the unsecure action endpoint of the jwt resource.
func (c *Client) NewUnsecureJWTRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
