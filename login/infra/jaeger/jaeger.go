package jaeger

import (
	"fmt"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

func CreateChileSpan(c echo.Context, name string) opentracing.Span {
	defer func() opentracing.Span {
		err := recover()
		if err != nil {
			fmt.Println(err)
			return opentracing.StartSpan("test")
		}
		return nil
	}()

	sp := jaegertracing.CreateChildSpan(c, name)

	return sp
}
