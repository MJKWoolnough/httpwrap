// File automatically generated with ./override.sh - DO NOT EDIT

package httpwrap

import (
	"io"
	"net/http"
)

type override interface {
	Set(*types)
}

type writer struct {
	io.Writer
}

func (i writer) Set(t *types) {
	t.responseWriterOverride = true
	t.Writer = i.Writer
}

// OverrideWriter is used to set an override for io.Writer
func OverrideWriter(t io.Writer) override {
	return writer{t}
}

type headers struct {
	Headers
}

func (i headers) Set(t *types) {
	t.responseWriterOverride = true
	t.Headers = i.Headers
}

// OverrideHeaders is used to set an override for Headers
func OverrideHeaders(t Headers) override {
	return headers{t}
}

type headerwriter struct {
	HeaderWriter
}

func (i headerwriter) Set(t *types) {
	t.responseWriterOverride = true
	t.HeaderWriter = i.HeaderWriter
}

// OverrideHeaderWriter is used to set an override for HeaderWriter
func OverrideHeaderWriter(t HeaderWriter) override {
	return headerwriter{t}
}

type closenotifier struct {
	http.CloseNotifier
}

func (i closenotifier) Set(t *types) {
	t.CloseNotifier = i.CloseNotifier
}

// OverrideCloseNotifier is used to set an override for http.CloseNotifier
func OverrideCloseNotifier(t http.CloseNotifier) override {
	return closenotifier{t}
}

type flusher struct {
	http.Flusher
}

func (i flusher) Set(t *types) {
	t.Flusher = i.Flusher
}

// OverrideFlusher is used to set an override for http.Flusher
func OverrideFlusher(t http.Flusher) override {
	return flusher{t}
}

type hijacker struct {
	http.Hijacker
}

func (i hijacker) Set(t *types) {
	t.Hijacker = i.Hijacker
}

// OverrideHijacker is used to set an override for http.Hijacker
func OverrideHijacker(t http.Hijacker) override {
	return hijacker{t}
}

type pusher struct {
	http.Pusher
}

func (i pusher) Set(t *types) {
	t.Pusher = i.Pusher
}

// OverridePusher is used to set an override for http.Pusher
func OverridePusher(t http.Pusher) override {
	return pusher{t}
}
