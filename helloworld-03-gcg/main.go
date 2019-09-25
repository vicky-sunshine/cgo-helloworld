package main

/*
void go_says(char* s);

static void Says(char* s) {
	go_says(s);
}
*/
import "C"
import (
	"fmt"
)

func main() {
	C.Says(C.CString("Hello, world!\n"))
}

//export go_says
func go_says(s *C.char) {
	fmt.Print(C.GoString(s))
}
