// File automatically generated with ./override.sh - DO NOT EDIT.

package httpwrap

import (
	"io"
	"net/http"
)

// Override is an interface for overriding interfaces on a http.ResponseWriter.
type Override interface {
	Set(*types)
}

type writer struct {
	io.Writer
}

func (i writer) Set(t *types) {
	t.responseWriterOverride = true
	t.Writer = i.Writer
}

// OverrideWriter is used to set an override for io.Writer.
func OverrideWriter(t io.Writer) Override {
	return writer{t}
}

type headers struct {
	Headers
}

func (i headers) Set(t *types) {
	t.responseWriterOverride = true
	t.Headers = i.Headers
}

// OverrideHeaders is used to set an override for Headers.
func OverrideHeaders(t Headers) Override {
	return headers{t}
}

type headerwriter struct {
	HeaderWriter
}

func (i headerwriter) Set(t *types) {
	t.responseWriterOverride = true
	t.HeaderWriter = i.HeaderWriter
}

// OverrideHeaderWriter is used to set an override for HeaderWriter.
func OverrideHeaderWriter(t HeaderWriter) Override {
	return headerwriter{t}
}

type flusher struct {
	http.Flusher
}

func (i flusher) Set(t *types) {
	t.Flusher = i.Flusher
}

// OverrideFlusher is used to set an override for http.Flusher.
func OverrideFlusher(t http.Flusher) Override {
	return flusher{t}
}

type hijacker struct {
	http.Hijacker
}

func (i hijacker) Set(t *types) {
	t.Hijacker = i.Hijacker
}

// OverrideHijacker is used to set an override for http.Hijacker.
func OverrideHijacker(t http.Hijacker) Override {
	return hijacker{t}
}

type pusher struct {
	http.Pusher
}

func (i pusher) Set(t *types) {
	t.Pusher = i.Pusher
}

// OverridePusher is used to set an override for http.Pusher.
func OverridePusher(t http.Pusher) Override {
	return pusher{t}
}

type stringwriter struct {
	io.StringWriter
}

func (i stringwriter) Set(t *types) {
	t.StringWriter = i.StringWriter
}

// OverrideStringWriter is used to set an override for io.StringWriter.
func OverrideStringWriter(t io.StringWriter) Override {
	return stringwriter{t}
}
