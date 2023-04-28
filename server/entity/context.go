package entity

import "context"

type (
	ContextConfig struct {
	}
	Context struct {
		context.Context
		TraceID string
	}
)

func NewContext(traceID string) *Context {
	old := context.Background()
	ctx := &Context{
		Context: old,
		TraceID: traceID,
	}

	return ctx
}

func WrapContext(old context.Context, traceID string) *Context {
	ctx := &Context{
		Context: old,
		TraceID: traceID,
	}
	return ctx
}
