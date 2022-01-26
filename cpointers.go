package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"time"
	"unsafe"
)

func preserveCPointer(ptr *C.char){
	time.Sleep(time.Second * 5)
	C.free(unsafe.Pointer(ptr))
}
