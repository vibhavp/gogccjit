package gogccjit

// #cgo LDFLAGS: -lgccjit
// #include <libgccjit.h>
import "C"
import "unsafe"

type Type struct {
	typ *C.gcc_jit_type
}

type JITType int

const (
	TypeVoid              JITType = C.GCC_JIT_TYPE_VOID
	TypeVoidPtr           JITType = C.GCC_JIT_TYPE_VOID
	TypeBool              JITType = C.GCC_JIT_TYPE_VOID
	TypeChar              JITType = C.GCC_JIT_TYPE_VOID
	TypeSignedChar        JITType = C.GCC_JIT_TYPE_SIGNED_CHAR
	TypeUnsignedChar      JITType = C.GCC_JIT_TYPE_UNSIGNED_CHAR
	TypeShort             JITType = C.GCC_JIT_TYPE_SHORT
	TypeUnsignedShort     JITType = C.GCC_JIT_TYPE_UNSIGNED_SHORT
	TypeInt               JITType = C.GCC_JIT_TYPE_INT
	TypeUnsignedInt       JITType = C.GCC_JIT_TYPE_UNSIGNED_INT
	TypeLong              JITType = C.GCC_JIT_TYPE_LONG
	TypeUnsignedLong      JITType = C.GCC_JIT_TYPE_UNSIGNED_LONG
	TypeLongLong          JITType = C.GCC_JIT_TYPE_LONG_LONG
	TypeUnsignedLongLong  JITType = C.GCC_JIT_TYPE_UNSIGNED_LONG_LONG
	TypeFloat             JITType = C.GCC_JIT_TYPE_FLOAT
	TypeDouble            JITType = C.GCC_JIT_TYPE_DOUBLE
	TypeLongDouble        JITType = C.GCC_JIT_TYPE_LONG_DOUBLE
	TypeConstCharPtr      JITType = C.GCC_JIT_TYPE_CONST_CHAR_PTR
	TypeSizeT             JITType = C.GCC_JIT_TYPE_SIZE_T
	TypeFilePtr           JITType = C.GCC_JIT_TYPE_FILE_PTR
	TypeComplexFloat      JITType = C.GCC_JIT_TYPE_COMPLEX_FLOAT
	TypeComplexDouble     JITType = C.GCC_JIT_TYPE_COMPLEX_DOUBLE
	TypeComplexLongDouble JITType = C.GCC_JIT_TYPE_COMPLEX_LONG_DOUBLE
)

func (c Context) GetType(typ JITType) Type {
	cTyp := C.gcc_jit_context_get_type(c.ctxt, C.enum_gcc_jit_types(typ))
	return Type{cTyp}
}

func (c Context) GetIntType(numBytes int, isSigned bool) Type {
	sig := 0
	if isSigned {
		sig = 1
	}

	return Type{C.gcc_jit_context_get_int_type(c.ctxt, C.int(numBytes), C.int(sig))}
}

func GetPointer(typ Type) Type {
	cTyp := C.gcc_jit_type_get_pointer(typ.typ)
	return Type{cTyp}
}

func GetConst(typ Type) Type {
	cTyp := C.gcc_jit_type_get_const(typ.typ)
	return Type{cTyp}
}

func GetVolatile(typ Type) Type {
	cTyp := C.gcc_jit_type_get_volatile(typ.typ)
	return Type{cTyp}
}

func (c Context) NewArrayType(loc Location, elemType Type, numElements int) Type {
	cTyp := C.gcc_jit_context_new_array_type(c.ctxt, loc.loc, elemType.typ, C.int(numElements))
	return Type{cTyp}
}

// func (c Context) GetAligned(typ Type, byteAlignment uintptr) Type {
// 	var cTyp C.gcc_jit_type = C.gcc_jit_type_get_aligned(typ.typ, C.size_t(byteAlignment))
// 	return Type{cTyp}
// }

type Struct struct {
	str *C.gcc_jit_struct
}

type Field struct {
	field *C.gcc_jit_field
}

func (ctxt Context) NewField(loc Location, typ Type, name string) Field {
	cField := C.gcc_jit_context_new_field(ctxt.ctxt, loc.loc, typ.typ, C.CString(name))
	return Field{cField}
}

func (field *Field) AsObject() Object {
	cObj := C.gcc_jit_field_as_object(field.field)
	return Object{cObj}
}

func (c Context) NewStructType(loc Location, name string, fields []Field) Struct {
	var cFields []*C.gcc_jit_field
	for _, field := range fields {
		cFields = append(cFields, field.field)
	}

	cStr := C.gcc_jit_context_new_struct_type(c.ctxt, loc.loc, C.CString(name), C.int(len(fields)), (**C.gcc_jit_field)(unsafe.Pointer(&cFields[0])))
	return Struct{cStr}
}

func (c Context) NewOpaqueStruct(loc Location, name string) Struct {
	cStr := C.gcc_jit_context_new_opaque_struct(c.ctxt, loc.loc, C.CString(name))
	return Struct{cStr}
}

func (str Struct) AsType() Type {
	return Type{C.gcc_jit_struct_as_type(str.str)}
}

func (str Struct) SetFields(loc Location, fields []Field) {
	var cFields []*C.gcc_jit_field
	for _, field := range fields {
		cFields = append(cFields, field.field)
	}

	C.gcc_jit_struct_set_fields(str.str, loc.loc, C.int(len(fields)), (**C.gcc_jit_field)(unsafe.Pointer(&cFields[0])))
}

func (c Context) NewUnionType(loc Location, name string, fields []Field) Type {
	var cFields []*C.gcc_jit_field
	for _, field := range fields {
		cFields = append(cFields, field.field)
	}

	cTyp := C.gcc_jit_context_new_union_type(c.ctxt, loc.loc, C.CString(name), C.int(len(fields)), (**C.gcc_jit_field)(unsafe.Pointer(&cFields[0])))

	return Type{cTyp}
}
