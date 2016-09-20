// Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/cpmech/gosl/chk"
)

func Test_fileIO1(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fileIO1")

	fn := "test/dorival/file.sim"
	chk.String(tst, "file.sim", filepath.Base(fn))
	chk.String(tst, ".sim", filepath.Ext(fn))
	chk.String(tst, "file", FnKey(fn))
	chk.String(tst, ".sim", FnExt(fn))
	chk.String(tst, "test/dorival/file", PathKey(fn))

	gn := "test/dorival/file.h5"
	chk.String(tst, "file.h5", filepath.Base(gn))
	chk.String(tst, ".h5", filepath.Ext(gn))
	chk.String(tst, "file", FnKey(gn))
	chk.String(tst, ".h5", FnExt(gn))
	chk.String(tst, "test/dorival/file", PathKey(gn))

	Pf("\n")
	Pf("fn   = %s\n", fn)
	Pf("base = %s\n", filepath.Base(fn))
	Pf("ext  = %s\n", filepath.Ext(fn))
	Pf("fnk  = %s\n", FnKey(fn))
	Pf("\n")

	fn = "test/dorival/file"
	chk.String(tst, "file", filepath.Base(fn))
	chk.String(tst, "", filepath.Ext(fn))
	chk.String(tst, "file", FnKey(fn))
	chk.String(tst, "test/dorival/file", PathKey(fn))

	Pf("\n")
	Pf("fn   = %s\n", fn)
	Pf("base = %s\n", filepath.Base(fn))
	Pf("ext  = %s\n", filepath.Ext(fn))
	Pf("fnk  = %s\n", FnKey(fn))
	Pf("\n")

	fn = "test/dorival/file."
	chk.String(tst, "file.", filepath.Base(fn))
	chk.String(tst, ".", filepath.Ext(fn))
	chk.String(tst, "file", FnKey(fn))
	chk.String(tst, "test/dorival/file", PathKey(fn))

	Pf("\n")
	Pf("fn   = %s\n", fn)
	Pf("base = %s\n", filepath.Base(fn))
	Pf("ext  = %s\n", filepath.Ext(fn))
	Pf("fnk  = %s\n", FnKey(fn))
	Pf("\n")

	fn = "test/dorival/f.extension"
	chk.String(tst, "f.extension", filepath.Base(fn))
	chk.String(tst, ".extension", filepath.Ext(fn))
	chk.String(tst, "f", FnKey(fn))
	chk.String(tst, "test/dorival/f", PathKey(fn))

	Pf("\n")
	Pf("fn   = %s\n", fn)
	Pf("base = %s\n", filepath.Base(fn))
	Pf("ext  = %s\n", filepath.Ext(fn))
	Pf("fnk  = %s\n", FnKey(fn))
	Pf("pathkey = %s\n", PathKey(fn))
	Pf("\n")
}

func Test_fileIO2(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fileIO2")

	os.MkdirAll("/tmp/gosl", 0777)

	fn := "/tmp/gosl/gosl_t_01_fileio.res"
	var bout bytes.Buffer
	Ff(&bout, "just testing %g\n", 666.0)
	AppendToFile(fn, &bout)

	ReadLines(fn, func(idx int, line string) (stop bool) {
		if line != "just testing 666" {
			chk.Panic("read wrong line: '%v'", line)
		}
		return false
	})
}

func Test_fileIO3(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fileIO3")

	type Test struct {
		Id     int
		Cells  []int
		Types  []string
		Values []float64
	}
	t := Test{0, []int{7, 3, 5}, []string{"a", "x", "p", "y"}, []float64{666}}
	Pf("t = %v\n", t)

	b, err := json.Marshal(&t)
	if err != nil {
		chk.Panic("marshal failed for %+v", t)
	}
	WriteBytesToFileD("/tmp/gosl/", "gosl_jsontest.res", b)
	PfBlue("file written /tmp/gosl/gosl_jsontest.res\n")
}

func Test_fileIO4(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fileIO4")

	theline := "Hello World !!!"
	WriteFileSD("/tmp/gosl", "filestring.txt", theline)

	f, err := OpenFileR("/tmp/gosl/filestring.txt")
	if err != nil {
		chk.Panic("%v", err)
	}

	ReadLinesFile(f, func(idx int, line string) (stop bool) {
		Pforan("line = %v\n", line)
		chk.String(tst, line, theline)
		return
	})
}
