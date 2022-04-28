package gopack

//#cgo CFLAGS: -I./../../c-cplex-lib/src -Isrc -m64 -O3 -fPIC -fexceptions -DNDEBUG -DIL_STD -DLONG_MAX=0x7FFFFFFFL -I./../../CPLEX_Studio1210//cplex/include -I./../../CPLEX_Studio1210//concert/include
//#cgo LDFLAGS: -L./../../c-cplex-lib/bin/ -L./../../CPLEX_Studio1210//cplex/lib/x86-64_osx/static_pic -L./../../CPLEX_Studio1210/concert/lib/x86-64_osx/static_pic -lilocplex -lconcert -lcplex -lm -lpthread -lmodel
//#include "model.h"
import "C"
import "fmt"

func Run() {

	/*env := C.CPXENVptr(nil);
	  lp := C.CPXLPptr(nil)
	  status := C.int(0)

	  env = C.CPXopenCPLEX (&status)
	  if env == nil {
	          fmt.Println("env error")
	  }

	  // Create the problem.
	  lp = C.CPXcreateprob (env, &status, C.CString("transport"));

	  if lp == nil {
	          fmt.Println("lp error")
	  }*/

	success := C.run_model()

	fmt.Println(success)
}
