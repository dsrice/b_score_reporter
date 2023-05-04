package jaeger

import (
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

func CreateChileSpan(c echo.Context, name string) opentracing.Span {
	if opentracing.SpanFromContext(c.Request().Context()) == nil {
		return opentracing.StartSpan("start")
	}

	sp := jaegertracing.CreateChildSpan(c, name)

	return sp
}
