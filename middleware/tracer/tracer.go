package tracer

type ITracer interface {
	AssociateTraceWithClient(traceId string, clientId string)
	// AppendTrace appends a trace to the trace log
	AppendTrace(traceId string, traceInfo string)
	//ClearTrace clears the trace log
	ClearTrace(traceId string)
}
