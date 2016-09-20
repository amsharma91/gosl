// Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package la

import (
	"time"

	"github.com/cpmech/gosl/utl"
)

type LinSol interface {
	InitR(tR *Triplet, symmetric, verbose, timing bool) error  // init for Real solution
	InitC(tC *TripletC, symmetric, verbose, timing bool) error // init for Complex solution
	Fact() error                                               // factorise
	SolveR(xR, bR []float64, sum_b_to_root bool) error         // solve Real problem. x = inv(A) * b
	SolveC(xR, xC, bR, bC []float64, sum_b_to_root bool) error // solve Complex problem x = inv(A) * b
	Clean()                                                    // deallocate memory
	SetOrdScal(ordering, scaling string) error                 // set ordering and scaling method
}

// lsAllocators is a "factory" for making linear solvers
var lsAllocators = map[string]func() LinSol{} // maps solver name to solver allocator

// GetSolver returns a linear solver by name. e.g. "umfpack" or "mumps"
func GetSolver(name string) LinSol {
	allocator, ok := lsAllocators[name]
	if !ok {
		utl.Panic(_linsol_err01, name)
	}
	return allocator()
}

// linSolData holds all data necessary to solve a sparse linear system like A.x = b
// Two direct solvers are used on the background: UMFPACK or MUMPS. The second one
// can be run in parallel via MPI. Both real and complex matrices are available
type linSolData struct {
	name  string    // solver name
	sym   bool      // is symmetric
	cmplx bool      // is complex
	verb  bool      // verbose call
	ton   bool      // timing is on
	tR    *Triplet  // triplet structure (real)
	tC    *TripletC // triplet structure (complex)
	tini  time.Time // initial time
}

// error messages
var (
	_linsol_err01 = "linsol.go: GetSolver: cannot find solver named %s in factory of linear solvers"
)
