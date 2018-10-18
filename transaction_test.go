package apm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.elastic.co/apm"
	"go.elastic.co/apm/transport/transporttest"
)

func TestStartTransactionTraceContextOptions(t *testing.T) {
	traceContext := startTransactionTraceContextOptions(t, false, false)
	assert.False(t, traceContext.Options.Requested())
	assert.False(t, traceContext.Options.MaybeRecorded())

	traceContext = startTransactionTraceContextOptions(t, false, true)
	assert.False(t, traceContext.Options.Requested())
	assert.False(t, traceContext.Options.MaybeRecorded())

	traceContext = startTransactionTraceContextOptions(t, true, false)
	assert.True(t, traceContext.Options.Requested())
	assert.True(t, traceContext.Options.MaybeRecorded())

	traceContext = startTransactionTraceContextOptions(t, true, true)
	assert.True(t, traceContext.Options.Requested())
	assert.True(t, traceContext.Options.MaybeRecorded())
}

func startTransactionTraceContextOptions(t *testing.T, requested, maybeRecorded bool) apm.TraceContext {
	tracer, _ := transporttest.NewRecorderTracer()
	defer tracer.Close()
	tracer.SetSampler(samplerFunc(func(apm.TraceContext) bool {
		panic("nope")
	}))

	opts := apm.TransactionOptions{
		TraceContext: apm.TraceContext{
			Trace: apm.TraceID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			Span:  apm.SpanID{0, 1, 2, 3, 4, 5, 6, 7},
		},
	}
	opts.TraceContext.Options = opts.TraceContext.Options.WithRequested(requested)
	opts.TraceContext.Options = opts.TraceContext.Options.WithMaybeRecorded(maybeRecorded)

	tx := tracer.StartTransactionOptions("name", "type", opts)
	result := tx.TraceContext()
	tx.Discard()
	return result
}

func TestStartTransactionInvalidTraceContext(t *testing.T) {
	startTransactionInvalidTraceContext(t, apm.TraceContext{
		// Trace is all zeroes, which is invalid.
		Span: apm.SpanID{0, 1, 2, 3, 4, 5, 6, 7},
	})
	startTransactionInvalidTraceContext(t, apm.TraceContext{
		Trace: apm.TraceID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		// Span is all zeroes, which is invalid.
	})
}

func startTransactionInvalidTraceContext(t *testing.T, traceContext apm.TraceContext) {
	tracer, _ := transporttest.NewRecorderTracer()
	defer tracer.Close()

	var samplerCalled bool
	tracer.SetSampler(samplerFunc(func(apm.TraceContext) bool {
		samplerCalled = true
		return true
	}))

	opts := apm.TransactionOptions{TraceContext: traceContext}
	tx := tracer.StartTransactionOptions("name", "type", opts)
	tx.Discard()
	assert.True(t, samplerCalled)
}

type samplerFunc func(apm.TraceContext) bool

func (f samplerFunc) Sample(t apm.TraceContext) bool {
	return f(t)
}
