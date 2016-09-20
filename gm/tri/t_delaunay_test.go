// Copyright 2015 Dorival Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tri

import (
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/plt"
)

func Test_delaunay01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("delaunay01")

	// points
	X := []float64{0, 1, 1, 0, 0.5}
	Y := []float64{0, 0, 1, 1, 0.5}

	// generate
	V, C, err := Delaunay(X, Y, chk.Verbose)
	if err != nil {
		tst.Errorf("%v\n", err)
		return
	}

	// check
	xout := make([]float64, len(V))
	yout := make([]float64, len(V))
	for i, v := range V {
		io.Pforan("vert %2d : coords = %v\n", i, v)
		xout[i] = v[0]
		yout[i] = v[1]
	}
	chk.Vector(tst, "X", 1e-15, xout, X)
	chk.Vector(tst, "Y", 1e-15, yout, Y)
	for i, c := range C {
		io.Pforan("cell %2d : verts = %v\n", i, c)
	}
	chk.Ints(tst, "verts of cell 0", C[0], []int{3, 0, 4})
	chk.Ints(tst, "verts of cell 1", C[1], []int{4, 1, 2})
	chk.Ints(tst, "verts of cell 2", C[2], []int{1, 4, 0})
	chk.Ints(tst, "verts of cell 3", C[3], []int{4, 2, 3})

	// plot
	if chk.Verbose {
		plt.SetForPng(1, 300, 150)
		Draw(V, C, nil)
		plt.Equal()
		plt.AxisRange(-0.1, 1.1, -0.1, 1.1)
		plt.Gll("x", "y", "")
		plt.SaveD("/tmp/gosl/tri", "delaunay01.png")
	}
}
