package gogccjit

// #cgo LDFLAGS: -lgccjit
// #include <libgccjit.h>
import "C"
import "runtime"

type contextCont struct {
	ctxt *C.struct_gcc_jit_context
}

type Context struct {
	*contextCont
}

func newContext(ctxt *C.struct_gcc_jit_context) Context {
	cont := &contextCont{ctxt}
	c := Context{cont}
	runtime.SetFinalizer(cont, func(c *contextCont) {
		C.gcc_jit_context_release(c.ctxt)
	})

	return c
}

func NewContext() Context {
	return newContext(C.gcc_jit_context_acquire())
}

func (parent Context) NewChildContext() Context {
	return newContext(C.gcc_jit_context_new_child_context(parent.ctxt))
}

type BoolOption int

const (
	BoolOptionDebugInfo         BoolOption = C.GCC_JIT_BOOL_OPTION_DEBUGINFO
	BoolOptionInitialTree       BoolOption = C.GCC_JIT_BOOL_OPTION_DUMP_INITIAL_TREE
	BoolOptionInitialGimple     BoolOption = C.GCC_JIT_BOOL_OPTION_DUMP_INITIAL_GIMPLE
	BoolOptionDumpGeneratedCode BoolOption = C.GCC_JIT_BOOL_OPTION_DUMP_GENERATED_CODE
	BoolOptionDumpSummary       BoolOption = C.GCC_JIT_BOOL_OPTION_DUMP_SUMMARY
	BoolOptionDumpEverything    BoolOption = C.GCC_JIT_BOOL_OPTION_DUMP_EVERYTHING
	BoolOptionSelfcheckGC       BoolOption = C.GCC_JIT_BOOL_OPTION_SELFCHECK_GC
	BoolOptionKeepIntermediates BoolOption = C.GCC_JIT_BOOL_OPTION_KEEP_INTERMEDIATES
)

func (c Context) SetBoolOption(opt BoolOption, value bool) {
	vInt := 0
	if value {
		vInt = 1
	}
	C.gcc_jit_context_set_bool_option(c.ctxt, C.enum_gcc_jit_bool_option(opt), C.int(vInt))
}

type StrOption int

const StrOptionProgName = C.GCC_JIT_STR_OPTION_PROGNAME

func (c Context) SetStrOption(opt StrOption, value string) {
	C.gcc_jit_context_set_str_option(c.ctxt, C.enum_gcc_jit_str_option(opt), C.CString(value))
}

type IntOption int

const IntOptionOptimizationLevel = C.GCC_JIT_INT_OPTION_OPTIMIZATION_LEVEL

func (c Context) SetIntOption(opt IntOption, value int) {
	C.gcc_jit_context_set_int_option(c.ctxt, C.enum_gcc_jit_int_option(opt), C.int(value))
}

func (c Context) AddCommandLineOption(optname string) {
	C.gcc_jit_context_add_command_line_option(c.ctxt, C.CString(optname))
}
