// Package econtext provides Echo integration with golang.org/x/net/context.
package econtext

import (
	"github.com/labstack/echo"
	"golang.org/x/net/context"
)

type ctx struct {
	c *echo.Context
	context.Context
}

func (c ctx) Value(key interface{}) interface{} {
	if key == ckey {
		return c.c
	}
	if v := c.c.Get(key.(string)); v != nil {
		return v
	}
	return c.Context.Value(key)
}
