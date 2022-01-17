package mw

import (
	"github.com/DoNewsCode/core/key"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func ExampleWithTrace() {
	g := gin.New()
	g.Use(Trace(opentracing.GlobalTracer(), key.New("module", "foo")))
	g.Handle("GET", "/", func(context *gin.Context) {
		// Do stuff
	})
}
