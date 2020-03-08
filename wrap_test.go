package httpwrap

import (
	"net/http"
	"reflect"
	"testing"
)

type testResponseWriter struct {
	http.ResponseWriter
}

type tresponseWriterFlusherHijackerPusher responseWriterFlusherHijackerPusher

func TestWrap(t *testing.T) {
	for n, test := range [...]struct {
		Input, Output http.ResponseWriter
		Overrides     []Override
	}{
		{
			tresponseWriterFlusherHijackerPusher{},
			tresponseWriterFlusherHijackerPusher{},
			[]Override{},
		},
		{
			tresponseWriterFlusherHijackerPusher{},
			responseWriterHijackerPusher{},
			[]Override{
				OverrideFlusher(nil),
			},
		},
		{
			tresponseWriterFlusherHijackerPusher{},
			responseWriterFlusherPusher{},
			[]Override{
				OverrideHijacker(nil),
			},
		},
		{
			tresponseWriterFlusherHijackerPusher{},
			responseWriterFlusherHijacker{},
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
			responseWriterFlusher{},
			[]Override{
				OverrideFlusher(responseWriterFlusherHijackerPusher{}),
			},
		},
		{
			testResponseWriter{},
			responseWriterHijacker{},
			[]Override{
				OverrideHijacker(responseWriterFlusherHijackerPusher{}),
			},
		},
		{
			testResponseWriter{},
			responseWriterPusher{},
			[]Override{
				OverridePusher(responseWriterFlusherHijackerPusher{}),
			},
		},
	} {
		out := Wrap(test.Input, test.Overrides...)
		if reflect.TypeOf(test.Output) != reflect.TypeOf(out) {
			t.Errorf("test %d: expecting %T, got %T", n+1, test.Output, out)
		}
	}
}
