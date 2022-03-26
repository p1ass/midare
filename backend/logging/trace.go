package logging

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type Tracer struct {
	SpanID  string
	TraceID string
}

func ExtractTracer(ctx context.Context) Tracer {
	sc := trace.SpanFromContext(ctx).SpanContext()

	return Tracer{
		SpanID:  sc.SpanID().String(),
		TraceID: sc.TraceID().String(),
	}
}
