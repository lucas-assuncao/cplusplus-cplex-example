// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY LUASSUNCAO.

// WARNING: This file has automatically been generated on Thu, 05 May 2022 12:28:54 -03.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package model

/*
#cgo LDFLAGS: -L/Applications/CPLEX_Studio1210/cplex/lib/x86-64_osx/static_pic -L/Applications/CPLEX_Studio1210/concert/lib/x86-64_osx/static_pic -L/Users/luassuncao/c++-projects/cplusplus-cplex-example/c-cplex-lib/bin -lilocplex -lconcert -lcplex -lm -lpthread -lmodel
#cgo CFLAGS: -I/Users/luassuncao/c++-projects/cplusplus-cplex-example/c-cplex-lib/src -Isrc -m64 -O3 -fPIC -fexceptions -DNDEBUG -DIL_STD -DLONG_MAX=0x7FFFFFFFL -I/Applications/CPLEX_Studio1210/cplex/include -I/Applications/CPLEX_Studio1210/concert/include
#include "model.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}
