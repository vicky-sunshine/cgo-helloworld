package main

/*
#include "helloworld.h"
*/
import "C"

func main() {
	C.Says(C.CString("Hello, world!\n"))
}
