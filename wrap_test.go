package httpwrap

import (
	"net/http"
	"reflect"
	"testing"
)

type testResponseWriter struct {
	http.ResponseWriter
}

type tresponseWriterCloseNotifierFlusherHijackerPusher responseWriterCloseNotifierFlusherHijackerPusher

func TestWrap(t *testing.T) {
	for n, test := range [...]struct {
		Input, Output http.ResponseWriter
		Overrides     []Override
	}{
		{
			tresponseWriterCloseNotifierFlusherHijackerPusher{},
			tresponseWriterCloseNotifierFlusherHijackerPusher{},
			[]Override{},
		},
		{
			tresponseWriterCloseNotifierFlusherHijackerPusher{},
			responseWriterFlusherHijackerPusher{},
			[]Override{
				OverrideCloseNotifier(nil),
			},
		},
		{
			tresponseWriterCloseNotifierFlusherHijackerPusher{},
			responseWriterCloseNotifierHijackerPusher{},
			[]Override{
				OverrideFlusher(nil),
			},
		},
		{
			tresponseWriterCloseNotifierFlusherHijackerPusher{},
			responseWriterCloseNotifierFlusherPusher{},
			[]Override{
				OverrideHijacker(nil),
			},
		},
		{
			tresponseWriterCloseNotifierFlusherHijackerPusher{},
			responseWriterCloseNotifierFlusherHijacker{},
			[]Override{
				OverridePusher(nil),
			},
		},
		{
			testResponseWriter{},
			testResponseWriter{},
			[]Override{},
		},
		{
			testResponseWriter{},
			responseWriterCloseNotifier{},
			[]Override{
				OverrideCloseNotifier(responseWriterCloseNotifierFlusherHijackerPusher{}),
			},
		},
		{
			testResponseWriter{},
			responseWriterFlusher{},
			[]Override{
				OverrideFlusher(responseWriterCloseNotifierFlusherHijackerPusher{}),
			},
		},
		{
			testResponseWriter{},
			responseWriterHijacker{},
			[]Override{
				OverrideHijacker(responseWriterCloseNotifierFlusherHijackerPusher{}),
			},
		},
		{
			testResponseWriter{},
			responseWriterPusher{},
			[]Override{
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
