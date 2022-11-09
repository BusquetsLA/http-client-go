package httpgo

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers      http.Header // default headers from http pkg
	connTimeout  time.Duration
	resTimeout   time.Duration
	maxIdleConns int
	disTimeout   bool
}

type ClientBuilder interface {
	// methods for configuration, should be defined only when creating the client
	Build() Client
	SetHeaders(headers http.Header) ClientBuilder
	SetConnTimeout(timeout time.Duration) ClientBuilder
	SetResTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConns(maxConns int) ClientBuilder
	DisableTimeouts(disTimeouts bool) ClientBuilder
}

func New() ClientBuilder { // single http client being used every time for every request
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		headers:      c.headers,
		maxIdleConns: c.maxIdleConns,
		connTimeout:  c.connTimeout,
		resTimeout:   c.resTimeout,
		disTimeout:   c.disTimeout,
	}
	return &client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnTimeout(timeout time.Duration) ClientBuilder {
	c.connTimeout = timeout
	return c
}

func (c *clientBuilder) SetResTimeout(timeout time.Duration) ClientBuilder {
	c.resTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConns(maxConns int) ClientBuilder {
	c.maxIdleConns = maxConns
	return c
}

func (c *clientBuilder) DisableTimeouts(disTimeout bool) ClientBuilder {
	c.disTimeout = disTimeout
	return c
}
