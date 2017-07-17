package gogccjit

// #cgo LDFLAGS: -lgccjit
// #include <libgccjit.h>
import "C"
import "unsafe"
import "runtime"

type resultCont struct {
	result *C.gcc_jit_result
}

type Result struct {
	*resultCont
}

func newResult(result *C.gcc_jit_result) Result {
	cont := &resultCont{result}
	runtime.SetFinalizer(cont, func(r *resultCont) {
		C.gcc_jit_result_release(r.result)
	})

	return Result{cont}
}

func (c Context) Compile() Result {
	return newResult(C.gcc_jit_context_compile(c.ctxt))
}

func (res Result) GetCode(funcName string) unsafe.Pointer {
	return C.gcc_jit_result_get_code(res.result, C.CString(funcName))
}

func (res Result) GetGlobal(name string) unsafe.Pointer {
	return C.gcc_jit_result_get_global(res.result, C.CString(name))
}
