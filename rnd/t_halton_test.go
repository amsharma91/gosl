// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rnd

import (
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/plt"
)

func Test_halton01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("halton01. Halton Points")

	dim := 2
	npts := 100
	P := HaltonPoints(dim, npts)
	X := P[0]
	Y := P[1]

	if chk.Verbose {
		plt.SetForEps(1, 400)
		plt.Plot(X, Y, "'r.', clip_on=0")
		plt.Equal()
		plt.SaveD("/tmp/gosl", "halton01.eps")
	}
}
