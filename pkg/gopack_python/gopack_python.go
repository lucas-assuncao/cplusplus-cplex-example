package main

// #cgo pkg-config: /usr/local/opt/python@3.7/lib/pkgconfig/python-3.7.pc
// #include <Python.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {

	pycodeGo := `
   print("Funfou")
 `

	defer C.Py_Finalize()
	C.Py_Initialize()
	pycodeC := C.CString(pycodeGo)
	defer C.free(unsafe.Pointer(pycodeC))
	C.PyRun_SimpleString(pycodeC)

	fmt.Println("uepa")

}
