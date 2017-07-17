package gogccjit

// #cgo LDFLAGS: -lgccjit
// #include <libgccjit.h>
import "C"

type Object struct {
	obj *C.gcc_jit_object
}

func (o Object) GetContext() Context {
	c := C.gcc_jit_object_get_context(o.obj)
	return newContext(c)
}

func (o Object) GetDebugString() string {
	return C.GoString(C.gcc_jit_object_get_debug_string(o.obj))
}
