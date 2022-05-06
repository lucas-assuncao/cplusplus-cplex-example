package gopack_c

//#cgo CFLAGS: -I./../../cplusplus-cplex-lib/src -Isrc -m64 -O3 -fPIC -fexceptions -DNDEBUG -DIL_STD -DLONG_MAX=0x7FFFFFFFL -I./../../CPLEX_Studio1210//cplex/include -I./../../CPLEX_Studio1210//concert/include
//#cgo LDFLAGS: -L./../../cplusplus-cplex-lib/bin/ -L./../../CPLEX_Studio1210//cplex/lib/x86-64_osx/static_pic -L./../../CPLEX_Studio1210/concert/lib/x86-64_osx/static_pic -lilocplex -lconcert -lcplex -lm -lpthread -lmodelcplusplus
//#include "model.h"
import "C"
import "fmt"

func RunC() {
	success := C.run_model()

	fmt.Println(success)
}
