package tracer

import "time"

type Span struct {
	// TraceId is the unique identifier for the trace
	TraceId string

	Tags map[string]string

	StartTime time.Time

	EndTime time.Time

	ChildSpans []*Span // ChildSpans are the spans that are children of this span
}

func (span *Span) Finish() {
	span.EndTime = time.Now()

	//TODO: Log the span
}

func (span *Span) AddChildSpan(childSpan *Span) {
	span.ChildSpans = append(span.ChildSpans, childSpan)
}

func NewSpan(traceId string) *Span {
	return &Span{
		TraceId:   traceId,
		StartTime: time.Now(),
	}
}
