package gdal

/*
#include <cpl_vsi.h>
//#include "go_gdal.h"
//#include "gdal_version.h"

#cgo linux  CFLAGS: -I/usr/include/gdal
#cgo linux  LDFLAGS: -lgdal
#cgo darwin pkg-config: gdal
#cgo windows LDFLAGS: -lgdal.dll
*/
import "C"
import (
	"unsafe"
)

type VSILFile *C.VSILFILE

func VSIFileFromMemBuffer(pszFilename string, pabyData []byte, takeOwnership bool) VSILFile {
	cFilename := C.CString(pszFilename)
	defer C.free(unsafe.Pointer(cFilename))

	var bTakeOwnership C.int
	if takeOwnership {
		bTakeOwnership = C.TRUE
	} else {
		bTakeOwnership = C.FALSE
	}
	f := C.VSIFileFromMemBuffer(cFilename, (*C.GByte)(&pabyData[0]), C.vsi_l_offset(len(pabyData)), bTakeOwnership)
	return VSILFile(f)
}
