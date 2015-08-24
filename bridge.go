package econtext

import (
	"github.com/labstack/echo"
	"golang.org/x/net/context"
)

const (
	ckey       = "echo.Context"
	contextkey = "context.Context"
)

// FromContext extracts the bound golang.org/x/net/context.Context from a Echo
// context if one has been set, or nil if one is not available.
func FromContext(c echo.Context) context.Context {
	if ctx, ok := c.Get(ckey).(context.Context); ok {
		return ctx
	}
	return nil
}

// ToContext extracts the bound Echo context from a golang.org/x/net/context.Context
// if one has been set, or the empty Echo context if one is not available.
func ToContext(ctx context.Context) echo.Context {
	if c, ok := ctx.Value(contextkey).(*echo.Context); ok {
		return *c
	}
	return echo.Context{}
}

// Set makes a two-way binding between the given Echo request context and the
// given golang.org/x/net/context.Context. Returns the fresh context.Context that contains
// this binding. Using the ToC and From functions will allow you to convert
// between one and the other.
//
// Note that since context.Context's are immutable, you will have to call this
// function to "re-bind" the request's canonical context.Context if you ever
// decide to change it.
func Set(c *echo.Context, context context.Context) context.Context {
	ctx := ctx{c, context}
	c.Set(ckey, ctx)
	return ctx
}
