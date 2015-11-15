package ctxreq

import (
	"golang.org/x/net/context"
	"net/http"
)

// Unexported context key type.
type key int

// HTTP request context key.
const reqContextKey key = 0

// New context with HTTP request assigned.
func NewContext(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, reqContextKey, req)
}

// Retrieve HTTP request from context.
func FromContext(ctx context.Context) (*http.Request, bool) {
	req, ok := ctx.Value(reqContextKey).(*http.Request)
	return req, ok
}
