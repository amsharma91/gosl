// Copyright 2015 Dorival Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package msh

import (
	"math"
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/num"
	"github.com/cpmech/gosl/utl"
)

func Test_shp01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("shp01")

	r := []float64{0, 0, 0}

	verb := true
	for name, _ := range Functions {

		io.Pfyel("--------------------------------- %-6s---------------------------------\n", name)

		// check S
		tol := 1e-17
		if name == "tri10" {
			tol = 1e-14
		}
		checkShape(tst, name, tol, verb)

		// check dSdR
		tol = 1e-14
		if name == "lin5" || name == "lin4" || name == "tri10" || name == "qua12" || name == "qua16" {
			tol = 1e-10
		}
		if name == "tri15" {
			tol = 1e-9
		}
		checkDerivs(tst, name, r, tol, verb)

		io.PfGreen("OK\n")
	}
}

// checkShape checks that shape functions result in 1.0 @ nodes
func checkShape(tst *testing.T, shape string, tol float64, verbose bool) {

	// information
	fcn := Functions[shape]
	ndim := GeomNdim[shape]
	nverts := NumVerts[shape]
	coords := NatCoords[shape]

	// allocate slices
	S := make([]float64, nverts)
	dSdR := utl.DblsAlloc(nverts, ndim)

	// loop over all vertices
	errS := 0.0
	r := []float64{0, 0, 0}
	for n := 0; n < nverts; n++ {

		// natural coordinates @ vertex
		for i := 0; i < ndim; i++ {
			r[i] = coords[i][n]
		}

		// compute function
		fcn(S, dSdR, r, false)

		// check
		if verbose {
			for _, val := range S {
				if math.Abs(val) < 1e-15 {
					val = 0
				}
				io.Pf("%3v", val)
			}
			io.Pf("\n")
		}
		for m := 0; m < nverts; m++ {
			if n == m {
				errS += math.Abs(S[m] - 1.0)
			} else {
				errS += math.Abs(S[m])
			}
		}
	}

	// error
	if errS > tol {
		tst.Errorf("%s failed with err = %g\n", shape, errS)
		return
	}
}

// checkDerivs checks dSdR derivatives of shape structures
func checkDerivs(tst *testing.T, shape string, r []float64, tol float64, verbose bool) {

	// information
	fcn := Functions[shape]
	ndim := GeomNdim[shape]
	nverts := NumVerts[shape]

	// allocate slices
	S := make([]float64, nverts)
	dSdR := utl.DblsAlloc(nverts, ndim)

	// auxiliary
	r_tmp := make([]float64, len(r))
	S_tmp := make([]float64, nverts)

	// analytical
	fcn(S, dSdR, r, true)

	// numerical
	for n := 0; n < nverts; n++ {
		for i := 0; i < ndim; i++ {
			dSndRi, _ := num.DerivCentral(func(t float64, args ...interface{}) (Sn float64) {
				copy(r_tmp, r)
				r_tmp[i] = t
				fcn(S_tmp, nil, r_tmp, false)
				Sn = S_tmp[n]
				return
			}, r[i], 1e-1)
			if verbose {
				io.Pfgrey2("  dS%ddR%d @ %5.2f = %v (num: %v)\n", n, i, r, dSdR[n][i], dSndRi)
			}
			if math.Abs(dSdR[n][i]-dSndRi) > tol {
				tst.Errorf("nurbs dS%ddR%d failed with err = %g\n", n, i, math.Abs(dSdR[n][i]-dSndRi))
				return
			}
		}
	}
}
