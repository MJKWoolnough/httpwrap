// File automatically generated with ./wrap.sh - DO NOT EDIT.

// Package httpwrap wraps an http.ResponseWriter to override some method(s)
// while maintaining other possible interface implementations.
package httpwrap // import "vimagination.zapto.org/httpwrap"

import (
	"io"
	"net/http"
)

// Headers is an interface for the Header method of the ResponseWriter interface.
type Headers interface {
	Header() http.Header
}

// HeaderWriter is an interface for the WriteHeader method of the ResponseWriter
// interface.
type HeaderWriter interface {
	WriteHeader(int)
}

type responseWriter struct {
	io.Writer
	Headers
	HeaderWriter
}

type types struct {
	responseWriterOverride bool
	responseWriter
	http.Flusher
	http.Hijacker
	http.Pusher
	io.StringWriter
}
func createTypes(w http.ResponseWriter) (http.ResponseWriter, types) {
	var t types

	switch wt := w.(type) {
	case responseWriterFlusher:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
	case responseWriterHijacker:
		w = wt.ResponseWriter
		t.Hijacker = wt.Hijacker
	case responseWriterFlusherHijacker:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.Hijacker = wt.Hijacker
	case responseWriterPusher:
		w = wt.ResponseWriter
		t.Pusher = wt.Pusher
	case responseWriterFlusherPusher:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.Pusher = wt.Pusher
	case responseWriterHijackerPusher:
		w = wt.ResponseWriter
		t.Hijacker = wt.Hijacker
		t.Pusher = wt.Pusher
	case responseWriterFlusherHijackerPusher:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.Hijacker = wt.Hijacker
		t.Pusher = wt.Pusher
	case responseWriterStringWriter:
		w = wt.ResponseWriter
		t.StringWriter = wt.StringWriter
	case responseWriterFlusherStringWriter:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.StringWriter = wt.StringWriter
	case responseWriterHijackerStringWriter:
		w = wt.ResponseWriter
		t.Hijacker = wt.Hijacker
		t.StringWriter = wt.StringWriter
	case responseWriterFlusherHijackerStringWriter:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.Hijacker = wt.Hijacker
		t.StringWriter = wt.StringWriter
	case responseWriterPusherStringWriter:
		w = wt.ResponseWriter
		t.Pusher = wt.Pusher
		t.StringWriter = wt.StringWriter
	case responseWriterFlusherPusherStringWriter:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.Pusher = wt.Pusher
		t.StringWriter = wt.StringWriter
	case responseWriterHijackerPusherStringWriter:
		w = wt.ResponseWriter
		t.Hijacker = wt.Hijacker
		t.Pusher = wt.Pusher
		t.StringWriter = wt.StringWriter
	case responseWriterFlusherHijackerPusherStringWriter:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.Hijacker = wt.Hijacker
		t.Pusher = wt.Pusher
		t.StringWriter = wt.StringWriter
	default:
		t.Flusher, _ = w.(http.Flusher)
		t.Hijacker, _ = w.(http.Hijacker)
		t.Pusher, _ = w.(http.Pusher)
		t.StringWriter, _ = w.(io.StringWriter)
	}

	return w, t
}

func createBitMask(t types) uint64 {
	var bf uint64

	if t.Flusher != nil {
		bf |= 1
	}

	if t.Hijacker != nil {
		bf |= 2
	}

	if t.Pusher != nil {
		bf |= 4
	}

	if t.StringWriter != nil {
		bf |= 8
	}

	return bf
}

func createWrapper(w http.ResponseWriter, t types, bf uint64) http.ResponseWriter {
	switch bf {
	case 1:
		return responseWriterFlusher{
			w,
			t.Flusher,
		}
	case 2:
		return responseWriterHijacker{
			w,
			t.Hijacker,
		}
	case 3:
		return responseWriterFlusherHijacker{
			w,
			t.Flusher,
			t.Hijacker,
		}
	case 4:
		return responseWriterPusher{
			w,
			t.Pusher,
		}
	case 5:
		return responseWriterFlusherPusher{
			w,
			t.Flusher,
			t.Pusher,
		}
	case 6:
		return responseWriterHijackerPusher{
			w,
			t.Hijacker,
			t.Pusher,
		}
	case 7:
		return responseWriterFlusherHijackerPusher{
			w,
			t.Flusher,
			t.Hijacker,
			t.Pusher,
		}
	case 8:
		return responseWriterStringWriter{
			w,
			t.StringWriter,
		}
	case 9:
		return responseWriterFlusherStringWriter{
			w,
			t.Flusher,
			t.StringWriter,
		}
	case 10:
		return responseWriterHijackerStringWriter{
			w,
			t.Hijacker,
			t.StringWriter,
		}
	case 11:
		return responseWriterFlusherHijackerStringWriter{
			w,
			t.Flusher,
			t.Hijacker,
			t.StringWriter,
		}
	case 12:
		return responseWriterPusherStringWriter{
			w,
			t.Pusher,
			t.StringWriter,
		}
	case 13:
		return responseWriterFlusherPusherStringWriter{
			w,
			t.Flusher,
			t.Pusher,
			t.StringWriter,
		}
	case 14:
		return responseWriterHijackerPusherStringWriter{
			w,
			t.Hijacker,
			t.Pusher,
			t.StringWriter,
		}
	case 15:
		return responseWriterFlusherHijackerPusherStringWriter{
			w,
			t.Flusher,
			t.Hijacker,
			t.Pusher,
			t.StringWriter,
		}
	}

	return w
}

// Wrap wraps the given ResponseWriter and overrides the methods requested.
// When using OverrideWriter make sure to use OverrideStringWriter, even if only
// with a nil value to disable it.
func Wrap(w http.ResponseWriter, overrides ...Override) http.ResponseWriter {
	if len(overrides) == 0 {
		return w
	}

	w, t := createTypes(w)

	if rw, ok := w.(responseWriter); ok {
		t.responseWriter = rw
	} else {
		t.Writer = w
		t.Headers = w
		t.HeaderWriter = w
	}

	for _, o := range overrides {
		o.Set(&t)
	}

	bf := createBitMask(t)

	if t.responseWriterOverride || bf == 0 {
		w = t.responseWriter
	}

	return createWrapper(w, t, bf)
}

type responseWriterFlusher struct {
	http.ResponseWriter
	http.Flusher
}

type responseWriterHijacker struct {
	http.ResponseWriter
	http.Hijacker
}

type responseWriterFlusherHijacker struct {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
}

type responseWriterPusher struct {
	http.ResponseWriter
	http.Pusher
}

type responseWriterFlusherPusher struct {
	http.ResponseWriter
	http.Flusher
	http.Pusher
}

type responseWriterHijackerPusher struct {
	http.ResponseWriter
	http.Hijacker
	http.Pusher
}

type responseWriterFlusherHijackerPusher struct {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
	http.Pusher
}

type responseWriterStringWriter struct {
	http.ResponseWriter
	io.StringWriter
}

type responseWriterFlusherStringWriter struct {
	http.ResponseWriter
	http.Flusher
	io.StringWriter
}

type responseWriterHijackerStringWriter struct {
	http.ResponseWriter
	http.Hijacker
	io.StringWriter
}

type responseWriterFlusherHijackerStringWriter struct {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
	io.StringWriter
}

type responseWriterPusherStringWriter struct {
	http.ResponseWriter
	http.Pusher
	io.StringWriter
}

type responseWriterFlusherPusherStringWriter struct {
	http.ResponseWriter
	http.Flusher
	http.Pusher
	io.StringWriter
}

type responseWriterHijackerPusherStringWriter struct {
	http.ResponseWriter
	http.Hijacker
	http.Pusher
	io.StringWriter
}

type responseWriterFlusherHijackerPusherStringWriter struct {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
	http.Pusher
	io.StringWriter
}
