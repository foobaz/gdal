package gdal

/*
#include "go_gdal.h"
#include "gdal_version.h"

#cgo linux  pkg-config: gdal
#cgo darwin pkg-config: gdal
#cgo windows LDFLAGS: -Lc:/gdal/release-1600-x64/lib -lgdal_i
#cgo windows CFLAGS: -IC:/gdal/release-1600-x64/include
*/
import "C"
import (
	"unsafe"
)

// Get a configuration option.
func GetConfigOption(key, def string) string {
	k := C.CString(key)
	d := C.CString(def)
	defer C.free(unsafe.Pointer(k))
	defer C.free(unsafe.Pointer(d))
	opt := C.CPLGetConfigOption(k, d)
	return C.GoString(opt)
}

// Set a configuration option.
func SetConfigOption(key, value string) {
	k := C.CString(key)
	v := C.CString(value)
	defer C.free(unsafe.Pointer(k))
	defer C.free(unsafe.Pointer(v))
	C.CPLSetConfigOption(k, v)
}
