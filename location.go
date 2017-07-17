package gogccjit

// #cgo LDFLAGS: -lgccjit
// #include <libgccjit.h>
import "C"

type Location struct {
	loc *C.gcc_jit_location
}

func (c Context) NewLocation(filename string, line, column int) Location {
	return Location{
		loc: C.gcc_jit_context_new_location(c.ctxt, C.CString(filename), C.int(line), C.int(column)),
	}
}
