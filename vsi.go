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

// Open file.
//
// This function opens a file with the desired access. Large files (larger than 2GB) should be supported. Binary access is always implied and the "b" does not need to be included in the pszAccess string.
//
// On windows it is possible to define the configuration option GDAL_FILE_IS_UTF8 to have pszFilename treated as being in the local encoding instead of UTF-8, retoring the pre-1.8.0 behavior of VSIFOpenL().
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX fopen() function.
//
// Parameters:
//	filename 	the file to open. UTF-8 encoded.
//	access 	access requested (ie. "r", "r+", "w".
//
// Returns:
//	NULL on failure, or the file handle.
func VSIFOpenL(filename, access string) VSILFile {
	pszFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(pszFilename))

	pszAccess := C.CString(access)
	defer C.free(unsafe.Pointer(pszAccess))

	fp := C.VSIFOpenL(pszFilename, pszAccess)
	return VSILFile(fp)
}

// Close file.
//
// This function closes the indicated file.
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX fclose() function.
//
// Parameters:
//	fp 	file handle opened with VSIFOpenL().
//
// Returns:
//	True on success or false on failure.
func VSIFCloseL(fp VSILFile) int {
	result := C.VSIFCloseL((*C.VSILFILE)(fp))
	return int(result)
}

// Seek to requested offset.
//
// Seek to the desired offset (nOffset) in the indicated file.
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX fseek() call.
//
// Parameters:
//	fp 	file handle opened with VSIFOpenL().
//	offset 	offset in bytes.
//	whence 	one of SEEK_SET, SEEK_CUR or SEEK_END.
//
// Returns:
//	True on success or false on failure.
func VSIFSeekL(fp VSILFile, offset, whence int) bool {
	err := C.VSIFSeekL((*C.VSILFILE)(fp), C.vsi_l_offset(offset), C.int(whence))
	return err == 0
}

// Tell current file offset.
//
// Returns the current file read/write offset in bytes from the beginning of the file.
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX ftell() call.
//
// Parameters:
//	fp 	file handle opened with VSIFOpenL().
//
// Returns:
//	file offset in bytes.
func VSIFTellL(fp VSILFile) int {
	offset := C.VSIFTellL((*C.VSILFILE)(fp))
	return int(offset)
}

// Analog of the POSIX rewind() call.
func VSIRewindL(fp VSILFile) {
	C.VSIRewindL((*C.VSILFILE)(fp))
}

// Read bytes from file.
//
// Reads nCount objects of nSize bytes from the indicated file at the current offset into the indicated buffer.
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX fread() call.
//
// Parameters:
//	buffer 	the buffer into which the data should be read (at least nCount * nSize bytes in size.
//	size 	size of objects to read in bytes.
//	count 	number of objects to read.
//	fp 	file handle opened with VSIFOpenL().
//
// Returns:
//	number of objects successfully read.
func VSIFReadL(buffer []byte, size, count int, fp VSILFile) int {
	nRead := C.VSIFReadL(unsafe.Pointer(&buffer[0]), C.size_t(size), C.size_t(count), (*C.VSILFILE)(fp))
	return int(nRead)
}

// Write bytes to file.
//
// Writes nCount objects of nSize bytes to the indicated file at the current offset into the indicated buffer.
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX fwrite() call.
//
// Parameters:
//	buffer 	the buffer from which the data should be written (at least nCount * nSize bytes in size.
//	size 	size of objects to read in bytes.
//	count 	number of objects to read.
//	fp 	file handle opened with VSIFOpenL().
//
// Returns:
//	number of objects successfully written.
func VSIFWriteL(buffer []byte, size, count int, fp VSILFile) int {
	nWritten := C.VSIFWriteL(unsafe.Pointer(&buffer[0]), C.size_t(size), C.size_t(count), (*C.VSILFILE)(fp))
	return int(nWritten)
}

// Test for end of file.
//
// Returns true if an end-of-file condition occured during the previous read operation. The end-of-file flag is cleared by a successfull VSIFSeekL() call.
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX feof() call.
//
// Parameters:
//	fp 	file handle opened with VSIFOpenL().
//
// Returns:
//	True if at EOF else false.
func VSIFEofL(fp VSILFile) bool {
	bEnd := C.VSIFEofL((*C.VSILFILE)(fp))
	return bEnd != C.FALSE
}

// Truncate/expand the file to the specified size.
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX ftruncate() call.
//
// Parameters:
//	fp 	file handle opened with VSIFOpenL().
//	newSize 	new size in bytes.
//
// Returns:
//	True on success or false on error.
func VSIFTruncateL(fp VSILFile, newSize int) bool {
	err := C.VSIFTruncateL((*C.VSILFILE)(fp), C.vsi_l_offset(newSize))
	return err == 0
}

// Flush pending writes to disk.
//
// For files in write or update mode and on filesystem types where it is applicable, all pending output on the file is flushed to the physical disk.
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX fflush() call.
//
// Parameters:
//	fp 	file handle opened with VSIFOpenL().
//
// Returns:
//	True on success or false on error.
func VSIFFlushL(fp VSILFile) bool {
	err := C.VSIFFlushL((*C.VSILFILE)(fp))
	return err == 0
}

// Analog of the POSIX putc() call.
//
// Parameters:
//	fp 	file handle opened with VSIFOpenL().
//
// Returns:
//	True on success or false on error.
func VSIFPutcL(c int, fp VSILFile) bool {
	nWritten := C.VSIFPutcL(C.int(c), (*C.VSILFILE)(fp))
	return nWritten > 0
}

// Get filesystem object info.
//
// Fetches status information about a filesystem object (file, directory, etc). The returned information is placed in the VSIStatBufL structure. For portability only the st_size (size in bytes), and st_mode (file type). This method is similar to VSIStat(), but will work on large files on systems where this requires special calls.
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX stat() function.
//
// Parameters:
//	filename 	the path of the filesystem object to be queried. UTF-8 encoded.
//
// Returns:
//	statBuf 	the structure with information.
//	success 	True on success or false on an error.
func VSIStatL(filename string) (statBuf VSIStatBufL, success bool) {
	pszFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(pszFilename))

	err := C.VSIStatL(pszFilename, (*C.VSIStatBufL)(&statBuf))
	success = err == 0

	return statBuf, success
}

// Get filesystem object info.
//
// Fetches status information about a filesystem object (file, directory, etc). The returned information is placed in the VSIStatBufL structure. For portability only the st_size (size in bytes), and st_mode (file type). This method is similar to VSIStat(), but will work on large files on systems where this requires special calls.
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX stat() function, with an extra parameter to specify which information is needed, which offers a potential for speed optimizations on specialized and potentially slow virtual filesystem objects (/vsigzip/, /vsicurl/)
//
// Parameters:
//	filename 	the path of the filesystem object to be queried. UTF-8 encoded.
//	flags 	0 to get all information, or VSI_STAT_EXISTS_FLAG, VSI_STAT_NATURE_FLAG or VSI_STAT_SIZE_FLAG, or a combination of those to get partial info.
//
// Returns:
//	statBuf 	the structure with information.
//	success 	True on success or false on an error.
func VSIStatExL(filename string, flags int) (statBuf VSIStatBufL, success bool) {
	pszFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(pszFilename))

	err := C.VSIStatExL(pszFilename, (*C.VSIStatBufL)(&statBuf), C.int(flags))
	success = err == 0

	return statBuf, success
}

// Returns if the filenames of the filesystem are case sensitive.
//
// This method retrieves to which filesystem belongs the passed filename and return TRUE if the filenames of that filesystem are case sensitive.
//
// Currently, this will return FALSE only for Windows real filenames. Other VSI virtual filesystems are case sensitive.
//
// This methods avoid ugly ifndef WIN32 / endif code, that is wrong when dealing with virtual filenames.
//
// Parameters:
//	filename 	the path of the filesystem object to be tested. UTF-8 encoded.
//
// Returns:
//	True if the filenames of the filesystem are case sensitive.
func VSIIsCaseSensitiveFS(filename string) bool {
	pszFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(pszFilename))

	bSensitive := C.VSIIsCaseSensitiveFS(pszFilename)
	return (bSensitive != 0)
}

// Create a directory.
//
// Create a new directory with the indicated mode. The mode is ignored on some platforms. A reasonable default mode value would be 0666. This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX mkdir() function.
//
// Parameters:
//	pszPathname 	the path to the directory to create. UTF-8 encoded.
//	mode 	the permissions mode.
func VSIMkdir(path string, mode int) bool {
	pszPath := C.CString(path)
	defer C.free(unsafe.Pointer(pszPath))

	err := C.VSIMkdir(pszPath, C.long(mode))
	return err == 0
}

// Delete a directory.
//
// Deletes a directory object from the file system. On some systems the directory must be empty before it can be deleted.
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX rmdir() function.
//
// Parameters:
//	pszDirname 	the path of the directory to be deleted. UTF-8 encoded.
//
// Returns:
//	True on success or false on an error.
func VSIRmdir(path string) bool {
	pszPath := C.CString(path)
	defer C.free(unsafe.Pointer(pszPath))

	err := C.VSIRmdir(pszPath)
	return err == 0
}

// Delete a file.
//
// Deletes a file object from the file system.
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX unlink() function.
//
// Parameters:
//	pszFilename 	the path of the file to be deleted. UTF-8 encoded.
//
// Returns:
//	0 on success or -1 on an error.
func VSIUnlink(path string) bool {
	pszPath := C.CString(path)
	defer C.free(unsafe.Pointer(pszPath))

	err := C.VSIUnlink(pszPath)
	return err == 0
}

// Rename a file.
//
// Renames a file object in the file system. It should be possible to rename a file onto a new filesystem, but it is safest if this function is only used to rename files that remain in the same directory.
//
// This method goes through the VSIFileHandler virtualization and may work on unusual filesystems such as in memory.
//
// Analog of the POSIX rename() function.
//
// Parameters:
//	oldpath 	the name of the file to be renamed. UTF-8 encoded.
//	newpath 	the name the file should be given. UTF-8 encoded.
//
// Returns:
//	0 on success or -1 on an error.
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

// Install "memory" file system handler.
//
// A special file handler is installed that allows block of memory to be treated as files. All portions of the file system underneath the base path "/vsimem/" will be handled by this driver.
//
// Normal VSI*L functions can be used freely to create and destroy memory arrays treating them as if they were real file system objects. Some additional methods exist to efficient create memory file system objects without duplicating original copies of the data or to "steal" the block of memory associated with a memory file.
//
// At this time the memory handler does not properly handle directory semantics for the memory portion of the filesystem. The VSIReadDir() function is not supported though this will be corrected in the future.
//
// Calling this function repeatedly should do no harm, though it is not necessary. It is already called the first time a virtualizable file access function (ie. VSIFOpenL(), VSIMkDir(), etc) is called.
func VSIInstallMemFileHandler() {
	C.VSIInstallMemFileHandler()
}

func VSIInstallLargeFileHandler() {
	C.VSIInstallLargeFileHandler()
}

// Install /vsisubfile/ virtual file handler.
//
// This virtual file system handler allows access to subregions of files, treating them as a file on their own to the virtual file system functions (VSIFOpenL(), etc).
//
// A special form of the filename is used to indicate a subportion of another file:
//
// /vsisubfile/<offset>[_<size>],<filename>
//
// The size parameter is optional. Without it the remainder of the file from the start offset as treated as part of the subfile. Otherwise only <size> bytes from <offset> are treated as part of the subfile. The <filename> portion may be a relative or absolute path using normal rules. The <offset> and <size> values are in bytes.
//
// e.g. /vsisubfile/1000_3000,/data/abc.ntf /vsisubfile/5000,../xyz/raw.dat
//
// Unlike the /vsimem/ or conventional file system handlers, there is no meaningful support for filesystem operations for creating new files, traversing directories, and deleting files within the /vsisubfile/ area. Only the VSIStatL(), VSIFOpenL() and operations based on the file handle returned by VSIFOpenL() operate properly.
func VSIInstallSubFileHandler() {
	C.VSIInstallSubFileHandler()
}

// Install /vsisparse/ virtual file handler.
//
// The sparse virtual file handler allows a virtual file to be composed from chunks of data in other files, potentially with large spaces in the virtual file set to a constant value. This can make it possible to test some sorts of operations on what seems to be a large file with image data set to a constant value. It is also helpful when wanting to add test files to the test suite that are too large, but for which most of the data can be ignored. It could, in theory, also be used to treat several files on different file systems as one large virtual file.
func VSIInstallSparseFileHandler() {
	C.VSIInstallSparseFileHandler()
}

func VSICleanupFileManager() {
	C.VSICleanupFileManager()
}

// Create memory "file" from a buffer.
//
// A virtual memory file is created from the passed buffer with the indicated filename. Under normal conditions the filename would need to be absolute and within the /vsimem/ portion of the filesystem.
//
// The buffer remains the responsibility of the caller, and should not go out of scope as long as it might be accessed as a file. In no circumstances does this function take a copy of the data contents.
func VSIFileFromMemBuffer(filename string, data []byte) VSILFile {
	pszFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(pszFilename))

	fp := C.VSIFileFromMemBuffer(pszFilename, (*C.GByte)(&data[0]), C.vsi_l_offset(len(data)), C.FALSE)
	return VSILFile(fp)
}

// Fetch buffer underlying memory file.
//
// This function returns a pointer to the memory buffer underlying a virtual "in memory" file. If bUnlinkAndSeize is TRUE the filesystem object will be deleted, and ownership of the buffer will pass to the caller otherwise the underlying file will remain in existance.
//
// Parameters:
//	pszFilename 	the name of the file to grab the buffer of.
//	pnDataLength 	(file) length returned in this variable.
//	bUnlinkAndSeize 	TRUE to remove the file, or FALSE to leave unaltered.
//
// Returns:
//	pointer to memory buffer or NULL on failure.
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
