// File automatically generated with ./types.sh - DO NOT EDIT

package httpwrap

import "net/http"

func wrap(w http.ResponseWriter, t types) http.ResponseWriter {
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
