// File automatically generated with ./wrap.sh - DO NOT EDIT

// Package httpwrap wraps an http.ResponseWriter to override some method(s)
// while maintaining other possible interface implementations
package httpwrap

import (
	"io"
	"net/http"
)

// Headers is an interface for the Header method of the ResponseWriter interface
type Headers interface {
	Header() http.Header
}

// HeaderWriter is an interface for the WriteHeader method of the ResponseWriter
// interface
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
	http.CloseNotifier
	http.Flusher
	http.Hijacker
	http.Pusher
	StringWriter
}

// Wrap wraps the given ResponseWriter and overrides the methods requested.
// When using OverrideWriter make sure to use OverrideStringWriter, even if only
// with a nil value to disable it
func Wrap(w http.ResponseWriter, overrides ...Override) http.ResponseWriter {
	if len(overrides) == 0 {
		return w
	}
	var t types
	switch wt := w.(type) {
	case responseWriterCloseNotifier:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
	case responseWriterFlusher:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
	case responseWriterCloseNotifierFlusher:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Flusher = wt.Flusher
	case responseWriterHijacker:
		w = wt.ResponseWriter
		t.Hijacker = wt.Hijacker
	case responseWriterCloseNotifierHijacker:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Hijacker = wt.Hijacker
	case responseWriterFlusherHijacker:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.Hijacker = wt.Hijacker
	case responseWriterCloseNotifierFlusherHijacker:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Flusher = wt.Flusher
		t.Hijacker = wt.Hijacker
	case responseWriterPusher:
		w = wt.ResponseWriter
		t.Pusher = wt.Pusher
	case responseWriterCloseNotifierPusher:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Pusher = wt.Pusher
	case responseWriterFlusherPusher:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.Pusher = wt.Pusher
	case responseWriterCloseNotifierFlusherPusher:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Flusher = wt.Flusher
		t.Pusher = wt.Pusher
	case responseWriterHijackerPusher:
		w = wt.ResponseWriter
		t.Hijacker = wt.Hijacker
		t.Pusher = wt.Pusher
	case responseWriterCloseNotifierHijackerPusher:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Hijacker = wt.Hijacker
		t.Pusher = wt.Pusher
	case responseWriterFlusherHijackerPusher:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.Hijacker = wt.Hijacker
		t.Pusher = wt.Pusher
	case responseWriterCloseNotifierFlusherHijackerPusher:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Flusher = wt.Flusher
		t.Hijacker = wt.Hijacker
		t.Pusher = wt.Pusher
	case responseWriterStringWriter:
		w = wt.ResponseWriter
		t.StringWriter = wt.StringWriter
	case responseWriterCloseNotifierStringWriter:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.StringWriter = wt.StringWriter
	case responseWriterFlusherStringWriter:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.StringWriter = wt.StringWriter
	case responseWriterCloseNotifierFlusherStringWriter:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Flusher = wt.Flusher
		t.StringWriter = wt.StringWriter
	case responseWriterHijackerStringWriter:
		w = wt.ResponseWriter
		t.Hijacker = wt.Hijacker
		t.StringWriter = wt.StringWriter
	case responseWriterCloseNotifierHijackerStringWriter:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Hijacker = wt.Hijacker
		t.StringWriter = wt.StringWriter
	case responseWriterFlusherHijackerStringWriter:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.Hijacker = wt.Hijacker
		t.StringWriter = wt.StringWriter
	case responseWriterCloseNotifierFlusherHijackerStringWriter:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Flusher = wt.Flusher
		t.Hijacker = wt.Hijacker
		t.StringWriter = wt.StringWriter
	case responseWriterPusherStringWriter:
		w = wt.ResponseWriter
		t.Pusher = wt.Pusher
		t.StringWriter = wt.StringWriter
	case responseWriterCloseNotifierPusherStringWriter:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Pusher = wt.Pusher
		t.StringWriter = wt.StringWriter
	case responseWriterFlusherPusherStringWriter:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.Pusher = wt.Pusher
		t.StringWriter = wt.StringWriter
	case responseWriterCloseNotifierFlusherPusherStringWriter:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Flusher = wt.Flusher
		t.Pusher = wt.Pusher
		t.StringWriter = wt.StringWriter
	case responseWriterHijackerPusherStringWriter:
		w = wt.ResponseWriter
		t.Hijacker = wt.Hijacker
		t.Pusher = wt.Pusher
		t.StringWriter = wt.StringWriter
	case responseWriterCloseNotifierHijackerPusherStringWriter:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Hijacker = wt.Hijacker
		t.Pusher = wt.Pusher
		t.StringWriter = wt.StringWriter
	case responseWriterFlusherHijackerPusherStringWriter:
		w = wt.ResponseWriter
		t.Flusher = wt.Flusher
		t.Hijacker = wt.Hijacker
		t.Pusher = wt.Pusher
		t.StringWriter = wt.StringWriter
	case responseWriterCloseNotifierFlusherHijackerPusherStringWriter:
		w = wt.ResponseWriter
		t.CloseNotifier = wt.CloseNotifier
		t.Flusher = wt.Flusher
		t.Hijacker = wt.Hijacker
		t.Pusher = wt.Pusher
		t.StringWriter = wt.StringWriter
	default:
		t.CloseNotifier, _ = w.(http.CloseNotifier)
		t.Flusher, _ = w.(http.Flusher)
		t.Hijacker, _ = w.(http.Hijacker)
		t.Pusher, _ = w.(http.Pusher)
		t.StringWriter, _ = w.(StringWriter)
	}
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
	var bf uint64
	if t.CloseNotifier != nil {
		bf |= 1
	}
	if t.Flusher != nil {
		bf |= 2
	}
	if t.Hijacker != nil {
		bf |= 4
	}
	if t.Pusher != nil {
		bf |= 8
	}
	if t.StringWriter != nil {
		bf |= 16
	}
	if t.responseWriterOverride || bf == 0 {
		w = t.responseWriter
	}
	switch bf {
	case 1:
		return responseWriterCloseNotifier{
			w,
			t.CloseNotifier,
		}
	case 2:
		return responseWriterFlusher{
			w,
			t.Flusher,
		}
	case 3:
		return responseWriterCloseNotifierFlusher{
			w,
			t.CloseNotifier,
			t.Flusher,
		}
	case 4:
		return responseWriterHijacker{
			w,
			t.Hijacker,
		}
	case 5:
		return responseWriterCloseNotifierHijacker{
			w,
			t.CloseNotifier,
			t.Hijacker,
		}
	case 6:
		return responseWriterFlusherHijacker{
			w,
			t.Flusher,
			t.Hijacker,
		}
	case 7:
		return responseWriterCloseNotifierFlusherHijacker{
			w,
			t.CloseNotifier,
			t.Flusher,
			t.Hijacker,
		}
	case 8:
		return responseWriterPusher{
			w,
			t.Pusher,
		}
	case 9:
		return responseWriterCloseNotifierPusher{
			w,
			t.CloseNotifier,
			t.Pusher,
		}
	case 10:
		return responseWriterFlusherPusher{
			w,
			t.Flusher,
			t.Pusher,
		}
	case 11:
		return responseWriterCloseNotifierFlusherPusher{
			w,
			t.CloseNotifier,
			t.Flusher,
			t.Pusher,
		}
	case 12:
		return responseWriterHijackerPusher{
			w,
			t.Hijacker,
			t.Pusher,
		}
	case 13:
		return responseWriterCloseNotifierHijackerPusher{
			w,
			t.CloseNotifier,
			t.Hijacker,
			t.Pusher,
		}
	case 14:
		return responseWriterFlusherHijackerPusher{
			w,
			t.Flusher,
			t.Hijacker,
			t.Pusher,
		}
	case 15:
		return responseWriterCloseNotifierFlusherHijackerPusher{
			w,
			t.CloseNotifier,
			t.Flusher,
			t.Hijacker,
			t.Pusher,
		}
	case 16:
		return responseWriterStringWriter{
			w,
			t.StringWriter,
		}
	case 17:
		return responseWriterCloseNotifierStringWriter{
			w,
			t.CloseNotifier,
			t.StringWriter,
		}
	case 18:
		return responseWriterFlusherStringWriter{
			w,
			t.Flusher,
			t.StringWriter,
		}
	case 19:
		return responseWriterCloseNotifierFlusherStringWriter{
			w,
			t.CloseNotifier,
			t.Flusher,
			t.StringWriter,
		}
	case 20:
		return responseWriterHijackerStringWriter{
			w,
			t.Hijacker,
			t.StringWriter,
		}
	case 21:
		return responseWriterCloseNotifierHijackerStringWriter{
			w,
			t.CloseNotifier,
			t.Hijacker,
			t.StringWriter,
		}
	case 22:
		return responseWriterFlusherHijackerStringWriter{
			w,
			t.Flusher,
			t.Hijacker,
			t.StringWriter,
		}
	case 23:
		return responseWriterCloseNotifierFlusherHijackerStringWriter{
			w,
			t.CloseNotifier,
			t.Flusher,
			t.Hijacker,
			t.StringWriter,
		}
	case 24:
		return responseWriterPusherStringWriter{
			w,
			t.Pusher,
			t.StringWriter,
		}
	case 25:
		return responseWriterCloseNotifierPusherStringWriter{
			w,
			t.CloseNotifier,
			t.Pusher,
			t.StringWriter,
		}
	case 26:
		return responseWriterFlusherPusherStringWriter{
			w,
			t.Flusher,
			t.Pusher,
			t.StringWriter,
		}
	case 27:
		return responseWriterCloseNotifierFlusherPusherStringWriter{
			w,
			t.CloseNotifier,
			t.Flusher,
			t.Pusher,
			t.StringWriter,
		}
	case 28:
		return responseWriterHijackerPusherStringWriter{
			w,
			t.Hijacker,
			t.Pusher,
			t.StringWriter,
		}
	case 29:
		return responseWriterCloseNotifierHijackerPusherStringWriter{
			w,
			t.CloseNotifier,
			t.Hijacker,
			t.Pusher,
			t.StringWriter,
		}
	case 30:
		return responseWriterFlusherHijackerPusherStringWriter{
			w,
			t.Flusher,
			t.Hijacker,
			t.Pusher,
			t.StringWriter,
		}
	case 31:
		return responseWriterCloseNotifierFlusherHijackerPusherStringWriter{
			w,
			t.CloseNotifier,
			t.Flusher,
			t.Hijacker,
			t.Pusher,
			t.StringWriter,
		}
	}
	return w
}

type responseWriterCloseNotifier struct {
	http.ResponseWriter
	http.CloseNotifier
}

type responseWriterFlusher struct {
	http.ResponseWriter
	http.Flusher
}

type responseWriterCloseNotifierFlusher struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Flusher
}

type responseWriterHijacker struct {
	http.ResponseWriter
	http.Hijacker
}

type responseWriterCloseNotifierHijacker struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Hijacker
}

type responseWriterFlusherHijacker struct {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
}

type responseWriterCloseNotifierFlusherHijacker struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Flusher
	http.Hijacker
}

type responseWriterPusher struct {
	http.ResponseWriter
	http.Pusher
}

type responseWriterCloseNotifierPusher struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Pusher
}

type responseWriterFlusherPusher struct {
	http.ResponseWriter
	http.Flusher
	http.Pusher
}

type responseWriterCloseNotifierFlusherPusher struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Flusher
	http.Pusher
}

type responseWriterHijackerPusher struct {
	http.ResponseWriter
	http.Hijacker
	http.Pusher
}

type responseWriterCloseNotifierHijackerPusher struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Hijacker
	http.Pusher
}

type responseWriterFlusherHijackerPusher struct {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
	http.Pusher
}

type responseWriterCloseNotifierFlusherHijackerPusher struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Flusher
	http.Hijacker
	http.Pusher
}

type responseWriterStringWriter struct {
	http.ResponseWriter
	StringWriter
}

type responseWriterCloseNotifierStringWriter struct {
	http.ResponseWriter
	http.CloseNotifier
	StringWriter
}

type responseWriterFlusherStringWriter struct {
	http.ResponseWriter
	http.Flusher
	StringWriter
}

type responseWriterCloseNotifierFlusherStringWriter struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Flusher
	StringWriter
}

type responseWriterHijackerStringWriter struct {
	http.ResponseWriter
	http.Hijacker
	StringWriter
}

type responseWriterCloseNotifierHijackerStringWriter struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Hijacker
	StringWriter
}

type responseWriterFlusherHijackerStringWriter struct {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
	StringWriter
}

type responseWriterCloseNotifierFlusherHijackerStringWriter struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Flusher
	http.Hijacker
	StringWriter
}

type responseWriterPusherStringWriter struct {
	http.ResponseWriter
	http.Pusher
	StringWriter
}

type responseWriterCloseNotifierPusherStringWriter struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Pusher
	StringWriter
}

type responseWriterFlusherPusherStringWriter struct {
	http.ResponseWriter
	http.Flusher
	http.Pusher
	StringWriter
}

type responseWriterCloseNotifierFlusherPusherStringWriter struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Flusher
	http.Pusher
	StringWriter
}

type responseWriterHijackerPusherStringWriter struct {
	http.ResponseWriter
	http.Hijacker
	http.Pusher
	StringWriter
}

type responseWriterCloseNotifierHijackerPusherStringWriter struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Hijacker
	http.Pusher
	StringWriter
}

type responseWriterFlusherHijackerPusherStringWriter struct {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
	http.Pusher
	StringWriter
}

type responseWriterCloseNotifierFlusherHijackerPusherStringWriter struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Flusher
	http.Hijacker
	http.Pusher
	StringWriter
}
