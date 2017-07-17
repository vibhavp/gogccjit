package gogccjit

// #cgo LDFLAGS: -lgccjit
// #include <libgccjit.h>
import "C"
import "unsafe"

type RValue struct {
	rval *C.gcc_jit_rvalue
}

func (rval RValue) GetType() Type {
	return Type{C.gcc_jit_rvalue_get_type(rval.rval)}
}

func (rval RValue) AsObject() Object {
	return Object{C.gcc_jit_rvalue_as_object(rval.rval)}
}

func (c Context) NewRValueFromInt(numType Type, value int) RValue {
	return RValue{C.gcc_jit_context_new_rvalue_from_int(c.ctxt, numType.typ, C.int(value))}
}

func (c Context) NewRValueFromLong(numType Type, value int) RValue {
	return RValue{C.gcc_jit_context_new_rvalue_from_long(c.ctxt, numType.typ, C.long(value))}
}

func (c Context) Zero(numType Type) RValue {
	return RValue{C.gcc_jit_context_zero(c.ctxt, numType.typ)}
}

func (c Context) One(numType Type) RValue {
	return RValue{C.gcc_jit_context_one(c.ctxt, numType.typ)}
}

func (c Context) NewRValueFromDouble(numType Type, value float64) RValue {
	return RValue{C.gcc_jit_context_new_rvalue_from_double(c.ctxt, numType.typ, C.double(value))}
}

func (c Context) NewRValueFromPtr(ptrType Type, value unsafe.Pointer) RValue {
	return RValue{C.gcc_jit_context_new_rvalue_from_ptr(c.ctxt, ptrType.typ, value)}
}

func (c Context) Null(ptrType Type) RValue {
	return RValue{C.gcc_jit_context_null(c.ctxt, ptrType.typ)}
}

func (c Context) NewStringLiteral(value string) RValue {
	return RValue{C.gcc_jit_context_new_string_literal(c.ctxt, C.CString(value))}
}

type UnaryOp int

const (
	UnaryOpMinus         UnaryOp = C.GCC_JIT_UNARY_OP_MINUS
	UnaryOpBitwiseNegate UnaryOp = C.GCC_JIT_UNARY_OP_BITWISE_NEGATE
	UnaryOpLogicalNegate UnaryOp = C.GCC_JIT_UNARY_OP_LOGICAL_NEGATE
	UnaryOpAbs           UnaryOp = C.GCC_JIT_UNARY_OP_ABS
)

func (c Context) NewUnaryOp(loc Location, op UnaryOp, resType Type, rval RValue) RValue {
	return RValue{C.gcc_jit_context_new_unary_op(c.ctxt, loc.loc, C.enum_gcc_jit_unary_op(op), resType.typ, rval.rval)}
}

type BinaryOp int

const (
	BinaryOpPlus       BinaryOp = C.GCC_JIT_BINARY_OP_PLUS
	BinaryOpMinus      BinaryOp = C.GCC_JIT_BINARY_OP_MINUS
	BinaryOpMult       BinaryOp = C.GCC_JIT_BINARY_OP_MULT
	BinryOpDivide      BinaryOp = C.GCC_JIT_BINARY_OP_DIVIDE
	BinaryOpModulo     BinaryOp = C.GCC_JIT_BINARY_OP_MODULO
	BinaryOpBitwiseAnd BinaryOp = C.GCC_JIT_BINARY_OP_BITWISE_AND
	BinaryOpBitwiseXor BinaryOp = C.GCC_JIT_BINARY_OP_BITWISE_XOR
	BinaryOpBitwiseOr  BinaryOp = C.GCC_JIT_BINARY_OP_BITWISE_OR
	BinaryOpLogicalAnd BinaryOp = C.GCC_JIT_BINARY_OP_LOGICAL_AND
	BinaryOpLogicalOr  BinaryOp = C.GCC_JIT_BINARY_OP_LOGICAL_OR
	BinaryOpLShift     BinaryOp = C.GCC_JIT_BINARY_OP_LSHIFT
	BinaryOpRShift     BinaryOp = C.GCC_JIT_BINARY_OP_RSHIFT
)

func (c Context) NewBinaryOp(loc Location, op BinaryOp, resType Type, a RValue, b RValue) RValue {
	return RValue{C.gcc_jit_context_new_binary_op(c.ctxt, loc.loc, C.enum_gcc_jit_binary_op(op), resType.typ, a.rval, b.rval)}
}

type Lvalue struct {
	lval *C.gcc_jit_lvalue
}
