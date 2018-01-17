package httpwrap

import (
	"net/http"
	"reflect"
	"testing"
)

type testResponseWriter struct {
	http.ResponseWriter
}

func TestWrap(t *testing.T) {
	for n, test := range [...]struct {
		Input, Output http.ResponseWriter
		Overrides     []override
	}{
		{
			responseWriterCloseNotifierFlusherHijackerPusher{},
			responseWriterCloseNotifierFlusherHijackerPusher{},
			[]override{},
		},
		{
			responseWriterCloseNotifierFlusherHijackerPusher{},
			responseWriterFlusherHijackerPusher{},
			[]override{
				OverrideCloseNotifier(nil),
			},
		},
		{
			responseWriterCloseNotifierFlusherHijackerPusher{},
			responseWriterCloseNotifierHijackerPusher{},
			[]override{
				OverrideFlusher(nil),
			},
		},
		{
			responseWriterCloseNotifierFlusherHijackerPusher{},
			responseWriterCloseNotifierFlusherPusher{},
			[]override{
				OverrideHijacker(nil),
			},
		},
		{
			responseWriterCloseNotifierFlusherHijackerPusher{},
			responseWriterCloseNotifierFlusherHijacker{},
			[]override{
				OverridePusher(nil),
			},
		},
		{
			testResponseWriter{},
			testResponseWriter{},
			[]override{},
		},
		{
			testResponseWriter{},
			responseWriterCloseNotifier{},
			[]override{
				OverrideCloseNotifier(responseWriterCloseNotifierFlusherHijackerPusher{}),
			},
		},
		{
			testResponseWriter{},
			responseWriterFlusher{},
			[]override{
				OverrideFlusher(responseWriterCloseNotifierFlusherHijackerPusher{}),
			},
		},
		{
			testResponseWriter{},
			responseWriterHijacker{},
			[]override{
				OverrideHijacker(responseWriterCloseNotifierFlusherHijackerPusher{}),
			},
		},
		{
			testResponseWriter{},
			responseWriterPusher{},
			[]override{
				OverridePusher(responseWriterCloseNotifierFlusherHijackerPusher{}),
			},
		},
	} {
		out := Wrap(test.Input, test.Overrides...)
		if reflect.TypeOf(test.Output) != reflect.TypeOf(out) {
			t.Errorf("test %d: expecting %T, got %T", n+1, test.Output, out)
		}
	}
}
