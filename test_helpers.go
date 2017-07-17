package gogccjit

// typedef int (*intFunc) (int n);
//
// int bridge_int_func(intFunc f, int n)
// {
//   return f(n);
// }
import "C"
import "unsafe"

func test_call(f unsafe.Pointer, arg int) int {
	return int(C.bridge_int_func(C.intFunc(f), C.int(arg)))
}
