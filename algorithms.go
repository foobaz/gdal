package gdal

/*
#include "go_gdal.h"
#include "gdal_version.h"

#cgo linux  CFLAGS: -I/usr/include/gdal
#cgo linux  LDFLAGS: -lgdal
#cgo darwin pkg-config: gdal
#cgo windows LDFLAGS: -lgdal.dll
*/
import "C"
import (
	"fmt"
	"unsafe"
)

var _ = fmt.Println

/* --------------------------------------------- */
/* Misc functions                                */
/* --------------------------------------------- */

// Compute optimal PCT for RGB image
func ComputeMedianCutPCT(
	red, green, blue RasterBand,
	colors int,
	ct ColorTable,
	progress ProgressFunc,
	data interface{},
) int {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	err := C.GDALComputeMedianCutPCT(
		red.cval,
		green.cval,
		blue.cval,
		nil,
		C.int(colors),
		ct.cval,
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return int(err)
}

// 24bit to 8bit conversion with dithering
func DitherRGB2PCT(
	red, green, blue, target RasterBand,
	ct ColorTable,
	progress ProgressFunc,
	data interface{},
) int {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	err := C.GDALDitherRGB2PCT(
		red.cval,
		green.cval,
		blue.cval,
		target.cval,
		ct.cval,
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return int(err)
}

// Compute checksum for image region
func (rb RasterBand) Checksum(xOff, yOff, xSize, ySize int) int {
	sum := C.GDALChecksumImage(rb.cval, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize))
	return int(sum)
}

// Compute the proximity of all pixels in the image to a set of pixels in the source image
func (src RasterBand) ComputeProximity(
	dest RasterBand,
	options []string,
	progress ProgressFunc,
	data interface{},
) error {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	err := C.GDALComputeProximity(
		src.cval,
		dest.cval,
		(**C.char)(unsafe.Pointer(&opts[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return err
}

// Fill selected raster regions by interpolation from the edges
func (src RasterBand) FillNoData(
	mask RasterBand,
	distance float64,
	iterations int,
	options []string,
	progress ProgressFunc,
	data interface{},
) error {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	err := C.GDALFillNodata(
		src.cval,
		mask.cval,
		C.double(distance),
		0,
		C.int(iterations),
		(**C.char)(unsafe.Pointer(&opts[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return error(err)
}

// Create polygon coverage from raster data using an integer buffer
func (src RasterBand) Polygonize(
	mask RasterBand,
	layer Layer,
	fieldIndex int,
	options []string,
	progress ProgressFunc,
	data interface{},
) error {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	err := C.GDALPolygonize(
		src.cval,
		mask.cval,
		layer.cval,
		C.int(fieldIndex),
		(**C.char)(unsafe.Pointer(&opts[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return error(err)
}

// Create polygon coverage from raster data using a floating point buffer
func (src RasterBand) FPolygonize(
	mask RasterBand,
	layer Layer,
	fieldIndex int,
	options []string,
	progress ProgressFunc,
	data interface{},
) error {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	err := C.GDALFPolygonize(
		src.cval,
		mask.cval,
		layer.cval,
		C.int(fieldIndex),
		(**C.char)(unsafe.Pointer(&opts[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return error(err)
}

// Removes small raster polygons
func (src RasterBand) SieveFilter(
	mask, dest RasterBand,
	threshold, connectedness int,
	options []string,
	progress ProgressFunc,
	data interface{},
) error {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	err := C.GDALSieveFilter(
		src.cval,
		mask.cval,
		dest.cval,
		C.int(threshold),
		C.int(connectedness),
		(**C.char)(unsafe.Pointer(&opts[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return error(err)
}

/* --------------------------------------------- */
/* Warp functions                                */
/* --------------------------------------------- */

type WarpOptions *C.GDALWarpOptions 

// Reproject image
func (src Dataset) ReprojectImage(
	srcProjWKT string,
	dst Dataset,
	dstProjWKT string,
	resampleAlg ResampleAlg,
	memLimit, maxError float64,
	progress ProgressFunc,
	data interface{},
	options WarpOptions,
) error {
	var pf C.GDALProgressFunc
	var pa unsafe.Pointer
	if progress != nil {
		pf = C.goGDALProgressFuncProxyB()
		pa = unsafe.Pointer(&goGDALProgressFuncProxyArgs{progress, data})
	}
	
	var c_srcWKT, c_dstWKT *C.char
	if srcProjWKT != "" {
		c_srcWKT = C.CString(srcProjWKT)
		defer C.free(unsafe.Pointer(c_srcWKT))
	}
	if dstProjWKT != "" {
		c_dstWKT = C.CString(dstProjWKT)
		defer C.free(unsafe.Pointer(c_dstWKT))
	}

	err := C.GDALReprojectImage(
		src.cval,
		c_srcWKT,
		dst.cval,
		c_dstWKT,
		C.GDALResampleAlg(resampleAlg),
		C.double(memLimit),
		C.double(maxError),
		pf, 
		pa,
		options,
	)
	if err != 0 {
		return error(err)
	}
	return nil
}

//Unimplemented: CreateGenImgProjTransformer
//Unimplemented: CreateGenImgProjTransformer2
//Unimplemented: CreateGenImgProjTransformer3
//Unimplemented: SetGenImgProjTransformerDstGeoTransform
//Unimplemented: DestroyGenImgProjTransformer
//Unimplemented: GenImgProjTransform

//Unimplemented: CreateReprojectionTransformer
//Unimplemented: DestroyReprojection
//Unimplemented: ReprojectionTransform
//Unimplemented: CreateGCPTransformer
//Unimplemented: CreateGCPRefineTransformer
//Unimplemented: DestroyGCPTransformer
//Unimplemented: GCPTransform

//Unimplemented: CreateTPSTransformer
//Unimplemented: DestroyTPSTransformer
//Unimplemented: TPSTransform

//Unimplemented: CreateRPCTransformer
//Unimplemented: DestroyRPCTransformer
//Unimplemented: RPCTransform

//Unimplemented: CreateGeoLocTransformer
//Unimplemented: DestroyGeoLocTransformer
//Unimplemented: GeoLocTransform

//Unimplemented: CreateApproxTransformer
//Unimplemented: DestroyApproxTransformer
//Unimplemented: ApproxTransform

//Unimplemented: SimpleImageWarp
//Unimplemented: SuggestedWarpOutput
//Unimplemented: SuggsetedWarpOutput2
//Unimplemented: SerializeTransformer
//Unimplemented: DeserializeTransformer

//Unimplemented: TransformGeolocations

/* --------------------------------------------- */
/* Contour line functions                        */
/* --------------------------------------------- */

//Unimplemented: CreateContourGenerator
//Unimplemented: FeedLine
//Unimplemented: Destroy
//Unimplemented: ContourWriter
//Unimplemented: ContourGenerate

/* --------------------------------------------- */
/* Rasterizer functions                          */
/* --------------------------------------------- */

// Burn geometries into raster
//Unimplmemented: RasterizeGeometries

// Burn geometries from the specified list of layers into the raster
//Unimplemented: RasterizeLayers

// Burn geometries from the specified list of layers into the raster
//Unimplemented: RasterizeLayersBuf

/* --------------------------------------------- */
/* Gridding functions                            */
/* --------------------------------------------- */

//Unimplemented: CreateGrid
//Unimplemented: ComputeMatchingPoints
