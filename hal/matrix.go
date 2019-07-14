package hal

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS:  -lmatrix_creator_hal
#include "matrix.h"
*/
import (
	"C"
)

import (
	"fmt"
)

func WrapTest() {
	fmt.Println("called from wrapper")
}

func BusInit() {
	var result C.bool = C.busInit()
	fmt.Println(bool(result))
}
