package main

/*
#include <stdio.h>

static void Says(const char* s) {
    puts(s);
}
*/
import "C"

func main() {
	C.Says(C.CString("Hello, world!\n"))
}
