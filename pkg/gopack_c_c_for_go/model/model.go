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

// RUN_MODEL function as declared in src/model.h:24
func RUN_MODEL() int32 {
	__ret := C.run_model()
	__v := (int32)(__ret)
	return __v
}
