package econtext

import (
	"github.com/labstack/echo"
	"golang.org/x/net/context"
)

// Middleware is a Echo middleware that binds a new golang.org/x/net/context.Context to
// every request. This binding is two-way, and you can use the ToC and FromC
// functions to convert between one and the other.
//
// Note that since context.Context's are immutable, you will have to call Set to
// "re-bind" the request's canonical context.Context if you ever decide to
// change it, otherwise only the original context.Context (as set by this
// middleware) will be bound.
func Middleware() echo.HandlerFunc {
	return func(c *echo.Context) error {
		Set(c, context.Background())
		return nil
	}
}
