Echo Context
=======

[![GoDoc](https://godoc.org/github.com/dre1080/econtext?status.svg)](https://godoc.org/github.com/dre1080/econtext)

Context provides a canonical way to use `golang.org/x/net`'s [`context`][context] package
with [Echo][echo]. It provides two-way bindings between `context.Context`
objects and Echo's `echo.Context`, giving you a convenient way to look up Echo `Context`
variables from the `context.Context` and allowing you to freely convert one
request context to the other.

[echo]: https://github.com/labstack/echo
[context]: http://godoc.org/golang.org/x/net/context
