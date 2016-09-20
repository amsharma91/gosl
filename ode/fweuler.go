// Copyright 2015 Dorival Pedroso and Raul Durand. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

func fweuler_accept(o *ODE, y []float64) {
}

// forward-Euler
func fweuler_step(o *ODE, y []float64, x float64, args ...interface{}) (rerr float64, err error) {
	o.nfeval += 1
	err = o.fcn(o.f[0], x, y, args...)
	if err != nil {
		return
	}
	for i := 0; i < o.ndim; i++ {
		y[i] += o.h * o.f[0][i]
	}
	return 1e+20, err // must not be used with automatic substepping
}
