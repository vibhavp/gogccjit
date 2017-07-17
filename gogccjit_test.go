package gogccjit

import (
	"testing"
)

func TestNewContext(t *testing.T) {
	c := NewContext()
	c.NewChildContext()
	c.SetBoolOption(BoolOptionDebugInfo, true)
}

func TestCompile(t *testing.T) {
	ctxt := NewContext()
	intType := ctxt.GetType(TypeInt)
	paramN := ctxt.NewParam(Location{}, intType, "n")
	fn := ctxt.NewFunction(Location{}, FunctionKindExported, intType, "minus", []Param{paramN}, false)
	block := fn.NewBlock("")
	block.EndWithReturn(Location{}, ctxt.NewUnaryOp(Location{}, UnaryOpMinus, intType, paramN.AsRvalue()))
	res := ctxt.Compile()
	f := res.GetCode("minus")

	out := test_call(f, 1)
	if out != -1 {
		t.Fatalf("incorrect result from compiled function. got=%d, want=-1", out)
	}
}
