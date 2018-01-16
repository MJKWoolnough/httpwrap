# httpwrap
--
    import "github.com/MJKWoolnough/httpwrap"

Package httpwrap wraps an http.ResponseWriter to override some method(s) while
maintaining other possible interface implementations

## Usage

#### func  OverrideCloseNotifier

```go
func OverrideCloseNotifier(t http.CloseNotifier) override
```
OverrideCloseNotifier is used to set an override for http.CloseNotifier

#### func  OverrideFlusher

```go
func OverrideFlusher(t http.Flusher) override
```
OverrideFlusher is used to set an override for http.Flusher

#### func  OverrideHeaderWriter

```go
func OverrideHeaderWriter(t HeaderWriter) override
```
OverrideHeaderWriter is used to set an override for HeaderWriter

#### func  OverrideHeaders

```go
func OverrideHeaders(t Headers) override
```
OverrideHeaders is used to set an override for Headers

#### func  OverrideHijacker

```go
func OverrideHijacker(t http.Hijacker) override
```
OverrideHijacker is used to set an override for http.Hijacker

#### func  OverridePusher

```go
func OverridePusher(t http.Pusher) override
```
OverridePusher is used to set an override for http.Pusher

#### func  OverrideWriter

```go
func OverrideWriter(t io.Writer) override
```
OverrideWriter is used to set an override for io.Writer

#### func  Wrap

```go
func Wrap(w http.ResponseWriter, overrides ...override) http.ResponseWriter
```
Wrap wraps the given ResponseWriter and overrides the methods requested.

#### type HeaderWriter

```go
type HeaderWriter interface {
	WriteHeader(int)
}
```

HeaderWriter is an interface for the WriteHeader method of the ResponseWriter
interface

#### type Headers

```go
type Headers interface {
	Header() http.Header
}
```

Headers is an interface for the Header method of the ResponseWriter interface
