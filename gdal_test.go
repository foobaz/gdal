// Copyright 2011 go-gdal. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import "testing"

func TestTiffDriver(t *testing.T) {
	_, err := GetDriverByName("GTiff")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestShapeDriver(t *testing.T) {
	drv := OGRDriverByName("ESRI Shapefile")
	drv.TestCapability("NA")
}

func TestHistogram(t *testing.T) {
	drv, err := GetDriverByName("MEM")
	if err != nil {
		t.Errorf("%+v", err)
	}
	ds := drv.Create("/vsimem/tmp", 10, 10, 1, Byte, nil)
	defer ds.Close()
	band := ds.RasterBand(1)
	data := make([]uint8, 100)
	for i := uint8(0); i < 100; i++ {
		data[i] = i
	}
	err = band.IO(Write, 0, 0, 10, 10, data, 10, 10, 0, 0)
	if err != nil {
		t.Errorf("%+v", err)
	}
	band.FlushCache()
	hist, err := band.Histogram(0, 100, 10, 1, 0, DummyProgress, nil)
	if err != nil {
		t.Errorf("%+v", err)
	}
	for i := 0; i < 10; i++ {
		if hist[i] != 10 {
			t.Errorf("failed to compute histogram. got: %+v, expected: 10\n", hist[i])
		}
	}
}

func TestConfigOption(t *testing.T) {
	k, v := "GDAL_GO_TEST", "ON"
	SetConfigOption(k, v)
	value := GetConfigOption(k, "")
	if value != v {
		t.Errorf("Invalid value: %s\n", value)
	}
}
