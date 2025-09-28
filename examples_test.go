package httpwrap_test

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

func Example() {
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
