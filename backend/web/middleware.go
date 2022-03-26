package web

import (
	cloudpropagator "github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
	"github.com/gin-gonic/gin"
	"github.com/p1ass/midare/logging"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const userIDContextKey = "userID"

// AuthMiddleware get session id from cookie and set user id to context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := getUserIDFromCookie(c)
		if err != nil {
			sendError(err, c)
			c.Abort()
			return
		}
		c.Set(userIDContextKey, userID)
		c.Next()
	}
}

// TracerMiddleware extracts trace information from request and injects it to context
// https://izumisy.work/entry/2022/01/10/225539
func TracerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// まずはX-Cloud-Trace-Contextからの読み取りをトライ
		// ここでエラーを拾っても何もできないので意図的にエラーは無視する
		if sc, _ := cloudpropagator.New().SpanContextFromRequest(c.Request); sc.IsValid() {
			ctx = trace.ContextWithRemoteSpanContext(ctx, sc)
		} else {
			// X-Cloud-Trace-ContextからValidな値が取れない場合には
			// traceparentヘッダからのTraceID/SpanIDのパースを試してみる
			prop := propagation.TraceContext{}
			ctx = prop.Extract(ctx, propagation.HeaderCarrier(c.Request.Header))
		}

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// LoggerMiddleware injects logger to context
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		logger := logging.New()
		newCtx := logging.Inject(ctx, logger)

		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
