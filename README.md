# httpwrap
--
    import "github.com/MJKWoolnough/httpwrap"

Package httpwrap wraps an http.ResponseWriter to override some method(s) while
maintaining other possible interface implementations

## Usage

#### func  Wrap

```go
func Wrap(w http.ResponseWriter, overrides ...Override) http.ResponseWriter
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

#### type Override

```go
type Override interface {
	Set(*types)
}
```

Override is an interface for overriding interfaces on a http.ResponseWriter

#### func  OverrideCloseNotifier

```go
func OverrideCloseNotifier(t http.CloseNotifier) Override
```
OverrideCloseNotifier is used to set an override for http.CloseNotifier

#### func  OverrideFlusher

```go
func OverrideFlusher(t http.Flusher) Override
```
OverrideFlusher is used to set an override for http.Flusher

#### func  OverrideHeaderWriter

```go
func OverrideHeaderWriter(t HeaderWriter) Override
```
OverrideHeaderWriter is used to set an override for HeaderWriter

#### func  OverrideHeaders

```go
func OverrideHeaders(t Headers) Override
```
OverrideHeaders is used to set an override for Headers

#### func  OverrideHijacker

```go
func OverrideHijacker(t http.Hijacker) Override
```
OverrideHijacker is used to set an override for http.Hijacker

#### func  OverridePusher

```go
func OverridePusher(t http.Pusher) Override
```
OverridePusher is used to set an override for http.Pusher

#### func  OverrideStringWriter

```go
func OverrideStringWriter(t StringWriter) Override
```
OverrideStringWriter is used to set an override for StringWriter

#### func  OverrideWriter

```go
func OverrideWriter(t io.Writer) Override
```
OverrideWriter is used to set an override for io.Writer

#### type StringWriter

```go
type StringWriter interface {
	WriteString(string) (int, error)
}
```

StringWriter is an interface for writing strings
