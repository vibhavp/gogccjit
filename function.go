package gogccjit

// #cgo LDFLAGS: -lgccjit
// #include <libgccjit.h>
import "C"
import "unsafe"

type Param struct {
	param *C.gcc_jit_param
}

func (c Context) NewParam(loc Location, paramType Type, name string) Param {
	return Param{C.gcc_jit_context_new_param(c.ctxt, loc.loc, paramType.typ, C.CString(name))}
}

func (p Param) AsLvalue() Lvalue {
	return Lvalue{C.gcc_jit_param_as_lvalue(p.param)}
}

func (p Param) AsRvalue() RValue {
	return RValue{C.gcc_jit_param_as_rvalue(p.param)}
}

func (p Param) AsObject() Object {
	return Object{C.gcc_jit_param_as_object(p.param)}
}

type Function struct {
	fn *C.gcc_jit_function
}

type FunctionKind int

const (
	FunctionKindExported     FunctionKind = C.GCC_JIT_FUNCTION_EXPORTED
	FunctionKindInternal     FunctionKind = C.GCC_JIT_FUNCTION_INTERNAL
	FunctionKindImported     FunctionKind = C.GCC_JIT_FUNCTION_IMPORTED
	FunctionKindAlwaysInline FunctionKind = C.GCC_JIT_FUNCTION_ALWAYS_INLINE
)

func (c Context) NewFunction(loc Location, kind FunctionKind, returnType Type, name string, params []Param, isVariadic bool) Function {
	var cParams []*C.gcc_jit_param
	for _, param := range params {
		cParams = append(cParams, param.param)
	}

	varInt := 0
	if isVariadic {
		varInt = 1
	}

	return Function{
		fn: C.gcc_jit_context_new_function(c.ctxt, loc.loc, C.enum_gcc_jit_function_kind(kind), returnType.typ, C.CString(name), C.int(len(params)), (**C.gcc_jit_param)(unsafe.Pointer(&cParams[0])), C.int(varInt)),
	}
}

func (c Context) GetBuiltinFunction(name string) Function {
	return Function{C.gcc_jit_context_get_builtin_function(c.ctxt, C.CString(name))}
}

func (fn Function) AsObject() Object {
	return Object{C.gcc_jit_function_as_object(fn.fn)}
}

func (fn Function) GetParam(index int) Param {
	return Param{C.gcc_jit_function_get_param(fn.fn, C.int(index))}
}

func (fn Function) DumpToDot(path string) {
	C.gcc_jit_function_dump_to_dot(fn.fn, C.CString(path))
}

func (fn Function) NewLocal(loc Location, localType Type, name string) Lvalue {
	return Lvalue{C.gcc_jit_function_new_local(fn.fn, loc.loc, localType.typ, C.CString(name))}
}

type Block struct {
	block *C.gcc_jit_block
}

func (fn Function) NewBlock(name string) Block {
	return Block{C.gcc_jit_function_new_block(fn.fn, C.CString(name))}
}

func (block Block) AsObject() Object {
	return Object{C.gcc_jit_block_as_object(block.block)}
}

func (block Block) GetFunction() Function {
	return Function{C.gcc_jit_block_get_function(block.block)}
}

func (block Block) AddEval(loc Location, rvalue RValue) {
	C.gcc_jit_block_add_eval(block.block, loc.loc, rvalue.rval)
}

func (block Block) AddAssignment(loc Location, lvalue Lvalue, rvalue RValue) {
	C.gcc_jit_block_add_assignment(block.block, loc.loc, lvalue.lval, rvalue.rval)
}

func (block Block) AddAssignmentOp(loc Location, lvalue Lvalue, op BinaryOp, rvalue RValue) {
	C.gcc_jit_block_add_assignment_op(block.block, loc.loc, lvalue.lval, C.enum_gcc_jit_binary_op(op), rvalue.rval)
}

func (block Block) AddComment(loc Location, text string) {
	C.gcc_jit_block_add_comment(block.block, loc.loc, C.CString(text))
}

func (block Block) EndWithConditional(loc Location, boolVal RValue, onTrue Block, onFalse Block) {
	C.gcc_jit_block_end_with_conditional(block.block, loc.loc, boolVal.rval, onTrue.block, onFalse.block)
}

func (block Block) EndWithJump(loc Location, target Block) {
	C.gcc_jit_block_end_with_jump(block.block, loc.loc, target.block)
}

func (block Block) EndWithReturn(loc Location, rvalue RValue) {
	C.gcc_jit_block_end_with_return(block.block, loc.loc, rvalue.rval)
}

func (block Block) EndWithVoidReturn(loc Location) {
	C.gcc_jit_block_end_with_void_return(block.block, loc.loc)
}
