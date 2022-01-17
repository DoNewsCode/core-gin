package mw

import (
	"time"

	"github.com/DoNewsCode/core/srvhttp"
	"github.com/gin-gonic/gin"
)

// Metrics is a gin middleware that adds request histogram. Setting addPath
// to true will make histogram to use request path as a dimension. This is ok
// with few total number of paths, but incurs performance issue if the
// cardinality of request path is high.
func Metrics(rds *srvhttp.RequestDurationSeconds, addPath bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := rds
		if addPath {
			r = rds.Route(c.FullPath())
		}
		defer func(begin time.Time) {
			r.Status(c.Writer.Status()).Observe(time.Since(begin))
		}(time.Now())
		c.Next()
	}
}
