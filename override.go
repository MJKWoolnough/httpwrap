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

// Writer is used to set an override for io.Writer
func Writer(t io.Writer) override {
	return writer{t}
}

type header struct {
	Header
}

func (i header) Set(t *types) {
	t.responseWriterOverride = true
	t.Header = i.Header
}

// Header is used to set an override for Header
func Header(t Header) override {
	return header{t}
}

type headerwriter struct {
	HeaderWriter
}

func (i headerwriter) Set(t *types) {
	t.responseWriterOverride = true
	t.HeaderWriter = i.HeaderWriter
}

// HeaderWriter is used to set an override for HeaderWriter
func HeaderWriter(t HeaderWriter) override {
	return headerwriter{t}
}

type closenotifier struct {
	http.CloseNotifier
}

func (i closenotifier) Set(t *types) {
	t.CloseNotifier = i.CloseNotifier
}

// CloseNotifier is used to set an override for http.CloseNotifier
func CloseNotifier(t http.CloseNotifier) override {
	return closenotifier{t}
}

type flusher struct {
	http.Flusher
}

func (i flusher) Set(t *types) {
	t.Flusher = i.Flusher
}

// Flusher is used to set an override for http.Flusher
func Flusher(t http.Flusher) override {
	return flusher{t}
}

type hijacker struct {
	http.Hijacker
}

func (i hijacker) Set(t *types) {
	t.Hijacker = i.Hijacker
}

// Hijacker is used to set an override for http.Hijacker
func Hijacker(t http.Hijacker) override {
	return hijacker{t}
}

type pusher struct {
	http.Pusher
}

func (i pusher) Set(t *types) {
	t.Pusher = i.Pusher
}

// Pusher is used to set an override for http.Pusher
func Pusher(t http.Pusher) override {
	return pusher{t}
}
