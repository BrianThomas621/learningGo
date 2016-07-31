package main

/*
#include <stdio.h>
#include <stdlib.h>

typedef int (*intFunc) ();

int

bridge_int_func(intFunc f)
{
	return f();
}

int fortytwo()
{
	return 42;
}

void myprint(char* s) {
    printf("%s", s);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	f := C.intFunc(C.fortytwo)
	fmt.Println(int(C.bridge_int_func(f)))
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

// See https://github.com/golang/go/wiki/cgo
