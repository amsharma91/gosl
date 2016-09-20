// Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"math"
	"testing"

	"github.com/cpmech/gosl/la"
	"github.com/cpmech/gosl/mpi"
	"github.com/cpmech/gosl/num"
	"github.com/cpmech/gosl/utl"
)

func main() {

	mpi.Start(false)
	defer func() {
		if err := recover(); err != nil {
			utl.PfRed("Some error has happened: %v\n", err)
		}
		mpi.Stop(false)
	}()

	utl.Tsilent = false
	if mpi.Rank() == 0 {
		utl.TTitle("TestJacobian 01b (MPI)")
	}
	if mpi.Size() != 2 {
		utl.Pf("this tests needs MPI 2 processors\n")
		return
	}

	ffcn := func(fx, x []float64) {
		fx[0] = math.Pow(x[0], 3.0) + x[1] - 1.0
		fx[1] = -x[0] + math.Pow(x[1], 3.0) + 1.0
	}
	Jfcn := func(dfdx *la.Triplet, x []float64) {
		dfdx.Start()
		if false {
			if mpi.Rank() == 0 {
				dfdx.Put(0, 0, 3.0*x[0]*x[0])
				dfdx.Put(1, 0, -1.0)
			} else {
				dfdx.Put(0, 1, 1.0)
				dfdx.Put(1, 1, 3.0*x[1]*x[1])
			}
		} else {
			if mpi.Rank() == 0 {
				dfdx.Put(0, 0, 3.0*x[0]*x[0])
				dfdx.Put(0, 1, 1.0)
			} else {
				dfdx.Put(1, 0, -1.0)
				dfdx.Put(1, 1, 3.0*x[1]*x[1])
			}
		}
	}
	x := []float64{0.5, 0.5}
	var tst testing.T
	num.CompareJac(&tst, ffcn, Jfcn, x, 1e-8, true)
}
