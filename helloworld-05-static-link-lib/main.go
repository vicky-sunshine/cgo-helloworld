package main

/*
#cgo CFLAGS: -I./helloworld
#cgo LDFLAGS: -L./helloworld -lhelloworld
#include "helloworld.h"
*/
import "C"

func main() {
	C.Says(C.CString("Hello, world!\n"))
}
