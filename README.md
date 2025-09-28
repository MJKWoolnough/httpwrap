# httpwrap

[![CI](https://github.com/MJKWoolnough/httpwrap/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/httpwrap/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/httpwrap.svg)](https://pkg.go.dev/vimagination.zapto.org/httpwrap)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/httpwrap)](https://goreportcard.com/report/vimagination.zapto.org/httpwrap)

--
    import "vimagination.zapto.org/httpwrap"

Package httpwrap wraps an http.ResponseWriter to override some method(s) while maintaining other possible interface implementations.

## Highlights

 - Wrap `http.ResponseWriter` with custom, select overrides.
 - Preserves many, optional interfaces (`http.Flusher`, `http.Hijacker`, `http.Pusher`, etc.).
 - Auto-generated code to make it easy to add new overridable interfaces.

## Usage

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"vimagination.zapto.org/httpwrap"
)

type writeLogger struct {
	http.ResponseWriter
}

func (w writeLogger) Write(p []byte) (int, error) {
	fmt.Println("Writing Bytes:", len(p))

	return w.ResponseWriter.Write(p)
}

func (w writeLogger) WriteString(p string) (int, error) {
	fmt.Println("Writing String:", len(p))

	return io.WriteString(w.ResponseWriter, p)
}

func main() {
	w := httptest.NewRecorder()
	l := writeLogger{w}

	ww := httpwrap.Wrap(w, httpwrap.OverrideWriter(l), httpwrap.OverrideStringWriter(l))

	ww.Write([]byte("Some Data\n"))
	io.WriteString(ww, "Some More Data\n")
	fmt.Println(w.Body)

	// Output:
	// Writing Bytes: 10
	// Writing String: 15
	// Some Data
	// Some More Data
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/httpwrap
