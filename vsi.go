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
type VSIStatBufL C.VSIStatBufL

func VSIStatL(filename string) (VSIStatBufL, bool) {
	pszFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(pszFilename))

	var stat VSIStatBufL
	err := C.VSIStatL(pszFilename, (*C.VSIStatBufL)(&stat))
	success := err == 0

	return stat, success
}

func VSIStatExL(filename string, flags int) (VSIStatBufL, bool) {
	pszFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(pszFilename))

	var stat VSIStatBufL
	err := C.VSIStatExL(pszFilename, (*C.VSIStatBufL)(&stat), C.int(flags))
	success := err == 0

	return stat, success
}

func VSIIsCaseSensitiveFS(filename string) bool {
	pszFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(pszFilename))

	bSensitive := C.VSIIsCaseSensitiveFS(pszFilename)
	return (bSensitive != 0)
}

func VSIMkdir(path string, mode int) bool {
	pszPath := C.CString(path)
	defer C.free(unsafe.Pointer(pszPath))

	err := C.VSIMkdir(pszPath, C.long(mode))
	return err == 0
}

func VSIRmdir(path string) bool {
	pszPath := C.CString(path)
	defer C.free(unsafe.Pointer(pszPath))

	err := C.VSIRmdir(pszPath)
	return err == 0
}

func VSIUnlink(path string) bool {
	pszPath := C.CString(path)
	defer C.free(unsafe.Pointer(pszPath))

	err := C.VSIUnlink(pszPath)
	return err == 0
}

func VSIRename(oldPath, newPath string) bool {
	pszOldPath := C.CString(oldPath)
	defer C.free(unsafe.Pointer(pszOldPath))

	pszNewPath := C.CString(newPath)
	defer C.free(unsafe.Pointer(pszNewPath))

	err := C.VSIRename(pszOldPath, pszNewPath)
	return err == 0
}

func VSIStrerror(errno int) string {
	pszErr := C.VSIStrerror(C.int(errno))
	return C.GoString(pszErr)
}

func VSIInstallMemFileHandler() {
	C.VSIInstallMemFileHandler()
}

func VSIInstallLargeFileHandler() {
	C.VSIInstallLargeFileHandler()
}

func VSIInstallSubFileHandler() {
	C.VSIInstallSubFileHandler()
}

func VSIInstallSparseFileHandler() {
	C.VSIInstallSparseFileHandler()
}

func VSICleanupFileManager() {
	C.VSICleanupFileManager()
}

func VSIFileFromMemBuffer(filename string, data []byte, takeOwnership bool) VSILFile {
	pszFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(pszFilename))

	var bTakeOwnership C.int
	if takeOwnership {
		bTakeOwnership = C.TRUE
	} else {
		bTakeOwnership = C.FALSE
	}
	f := C.VSIFileFromMemBuffer(pszFilename, (*C.GByte)(&data[0]), C.vsi_l_offset(len(data)), bTakeOwnership)
	return VSILFile(f)
}

func VSIGetMemFileBuffer(filename string, unlinkAndSeize bool) []byte {
	pszFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(pszFilename))

	var bUnlinkAndSeize C.int
	if unlinkAndSeize {
		bUnlinkAndSeize = C.TRUE
	} else {
		bUnlinkAndSeize = C.FALSE
	}

	var pnDataLength C.vsi_l_offset
	pabyData := C.VSIGetMemFileBuffer(pszFilename, &pnDataLength, bUnlinkAndSeize)
	if pabyData == nil {
		return nil
	}

	data := C.GoBytes(unsafe.Pointer(pabyData), C.int(pnDataLength))
	if unlinkAndSeize {
		C.VSIFree(unsafe.Pointer(pabyData))
	}

	return data
}
