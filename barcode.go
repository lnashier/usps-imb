package usps

// #include <stdlib.h>
// #include "lib/usps4cb.h"
import "C"
import (
	// Import C libs
	_ "github.com/lnashier/usps-imb/lib"
	"unsafe"
)

// IMb provides USPS 4-State Customer Barcode (abbreviated 4CB, 4-CB, or USPS4CB).
func IMb(track, route string) (string, int) {
	trackCStr := C.CString(track)
	routeCStr := C.CString(route)
	barCStr := C.CString(string(make([]byte, 65)))
	defer C.free(unsafe.Pointer(trackCStr))
	defer C.free(unsafe.Pointer(routeCStr))
	defer C.free(unsafe.Pointer(barCStr))
	statusCode := C.USPS4CB(trackCStr, routeCStr, barCStr)
	return C.GoString(barCStr), int(statusCode)
}
