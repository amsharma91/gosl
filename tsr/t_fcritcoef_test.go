// Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tsr

import (
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/utl"
)

func Test_fcritcoef01(tst *testing.T) {

	prevTs := verbose()
	defer func() {
		verbose() = prevTs
		if err := recover(); err != nil {
			tst.Error("[1;31mSome error has happened:[0m\n", err)
		}
	}()

	//verbose() = false
	chk.PrintTitle("fcritcoef01")

	φ := 30.0
	Ma, _ := Mmatch(0, φ, "cmp")
	Mb := Phi2M(φ, "cam")
	φb := M2Phi(Ma, "cam")
	Mc := Phi2M(φ, "oct")
	φc := M2Phi(Mc, "oct")
	Md := SmpCalcμ(φ, 0)
	Me := NewSmpCalcμ(φ, -1.0, 0.0, 1.0, 1e-3)
	Mf := NewSmpCalcμ(φ, 1.0, 0.0, 1.0, 1e-3)
	Mg := NewSmpCalcμ(φ, 1.0, 0.0, 10.0, 1e-7)
	io.Pforan("Ma (cam) = %v\n", Ma)
	io.Pforan("Mb (cam) = %v\n", Mb)
	io.Pforan("Mc (oct) = %v\n", Mc)
	io.Pforan("Md (oct) = %v\n", Md)
	io.Pforan("Me (oct) = %v\n", Me)
	io.Pforan("Mf (oct) = %v\n", Mf)
	io.Pforan("Mg (oct) = %v\n", Mg)
	utl.CheckScalar(tst, "Ma-Mb", 1e-17, Ma, Mb)
	utl.CheckScalar(tst, "φ-φb", 1e-14, φ, φb)
	utl.CheckScalar(tst, "φ-φc", 1e-14, φ, φc)
	utl.CheckScalar(tst, "Mc-Md", 1e-17, Mc, Md)
	utl.CheckScalar(tst, "Mc-Me", 1e-15, Mc, Me)
	utl.CheckScalar(tst, "Mc-Mf", 1e-15, Mc, Mf)
	utl.CheckScalar(tst, "Mc-Mg", 1e-15, Mc, Mg)
}
