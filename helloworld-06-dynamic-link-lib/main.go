package main

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib -lhelloworld

#include "helloworld.h"
*/
import "C"

func main() {
	C.Says(C.CString("Hello, world!\n"))
}
