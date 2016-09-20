// Copyright 2015 Dorival Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package msh

func init() {
	Functions["tri3"] = Tri3
	Functions["tri6"] = Tri6
	Functions["tri10"] = Tri10
	Functions["tri15"] = Tri15
}

// Tri3 calculates the shape functions (S) and derivatives of shape functions (dSdR) of tri3
// elements at {r,s,t} natural coordinates. The derivatives are calculated only if derivs==true.
func Tri3(S []float64, dSdR [][]float64, R []float64, derivs bool) {
	/*      s
	        |
	        2, (0,1)
	        | ',
	        |   ',
	        |     ',
	        |       ',
	        |         ',
	        |           ',
	        |             ',
	        |               ',
	        | (0,0)           ', (1,0)
	        0-------------------1 ---- r
	*/
	r, s := R[0], R[1]
	S[0] = 1.0 - r - s
	S[1] = r
	S[2] = s

	if !derivs {
		return
	}

	dSdR[0][0] = -1.0
	dSdR[1][0] = 1.0
	dSdR[2][0] = 0.0

	dSdR[0][1] = -1.0
	dSdR[1][1] = 0.0
	dSdR[2][1] = 1.0
}

// Tri6 calculates the shape functions (S) and derivatives of shape functions (dSdR) of tri6
// elements at {r,s,t} natural coordinates. The derivatives are calculated only if derivs==true.
func Tri6(S []float64, dSdR [][]float64, R []float64, derivs bool) {
	/*      s
	        |
	        2, (0,1)
	        | ',
	        |   ',
	        |     ',
	        |       ',
	        5         '4
	        |           ',
	        |             ',
	        |               ',
	        | (0,0)           ', (1,0)
	        0---------3---------1 ---- r
	*/
	r, s := R[0], R[1]
	S[0] = 1.0 - (r+s)*(3.0-2.0*(r+s))
	S[1] = r * (2.0*r - 1.0)
	S[2] = s * (2.0*s - 1.0)
	S[3] = 4.0 * r * (1.0 - (r + s))
	S[4] = 4.0 * r * s
	S[5] = 4.0 * s * (1.0 - (r + s))

	if !derivs {
		return
	}

	dSdR[0][0] = -3.0 + 4.0*(r+s)
	dSdR[1][0] = 4.0*r - 1.0
	dSdR[2][0] = 0.0
	dSdR[3][0] = 4.0 - 8.0*r - 4.0*s
	dSdR[4][0] = 4.0 * s
	dSdR[5][0] = -4.0 * s

	dSdR[0][1] = -3.0 + 4.0*(r+s)
	dSdR[1][1] = 0.0
	dSdR[2][1] = 4.0*s - 1.0
	dSdR[3][1] = -4.0 * r
	dSdR[4][1] = 4.0 * r
	dSdR[5][1] = 4.0 - 4.0*r - 8.0*s
}

// Tri10 calculates the shape functions (S) and derivatives of shape functions (dSdR) of tri10
// elements at {r,s,t} natural coordinates. The derivatives are calculated only if derivs==true.
func Tri10(S []float64, dSdR [][]float64, R []float64, derivs bool) {
	/*
	   s
	   |
	   2, (0,1)
	   | ',
	   |   ',
	   5     '7
	   |       ',
	   |         ',
	   8      9    '4
	   |             ',
	   | (0,0)         ', (1,0)
	   0-----3-----6-----1 ---- r
	*/

	r, s := R[0], R[1]
	z := 1.0 - r - s
	t1 := s * (3.0*s - 1.0)
	t2 := z * (3.0*z - 1.0)
	t3 := r * (3.0*r - 1.0)

	S[0] = 0.5 * t2 * (3.0*z - 2.0)
	S[1] = 0.5 * t3 * (3.0*r - 2.0)
	S[2] = 0.5 * t1 * (3.0*s - 2.0)
	S[3] = 4.5 * r * t2
	S[4] = 4.5 * s * t3
	S[5] = 4.5 * z * t1
	S[6] = 4.5 * z * t3
	S[7] = 4.5 * r * t1
	S[8] = 4.5 * s * t2
	S[9] = 27.0 * s * z * r

	if !derivs {
		return
	}

	q0 := 4.5 * (6.0*z - 1.0)
	q1 := 4.5 * s * (3.0*s - 1.0)
	q2 := 4.5 * z * (3.0*z - 1.0)
	q3 := 4.5 * r * (3.0*r - 1.0)
	q4 := 4.5 * (6.0*s - 1.0)
	q5 := 4.5 * (6.0*r - 1.0)
	q6 := q0 * s
	q7 := q0 * r
	q8 := -0.5 * (27.0*z*z - 18.0*z + 2.0)
	q9 := 0.5 * (27.0*s*s - 18.0*s + 2.0)
	q10 := 0.5 * (27.0*r*r - 18.0*r + 2.0)

	dSdR[0][0] = q8
	dSdR[1][0] = q10
	dSdR[2][0] = 0.0
	dSdR[3][0] = q2 - q7
	dSdR[4][0] = s * q5
	dSdR[5][0] = -q1
	dSdR[6][0] = z*q5 - q3
	dSdR[7][0] = q1
	dSdR[8][0] = -q6
	dSdR[9][0] = 27.0 * s * (z - r)

	dSdR[0][1] = q8
	dSdR[1][1] = 0.0
	dSdR[2][1] = q9
	dSdR[3][1] = -q7
	dSdR[4][1] = q3
	dSdR[5][1] = z*q4 - q1
	dSdR[6][1] = -q3
	dSdR[7][1] = r * q4
	dSdR[8][1] = q2 - q6
	dSdR[9][1] = 27.0 * r * (z - s)
}

// Tri15 calculates the shape functions (S) and derivatives of shape functions (dSdR) of tri15
// elements at {r,s,t} natural coordinates. The derivatives are calculated only if derivs==true.
func Tri15(S []float64, dSdR [][]float64, R []float64, derivs bool) {
	/*      s
	           ^
	           |
	         2
	           @,(0,1)
	           | ',
	           |   ', 9
	        10 @     @,
	           |  14   ',   4
	         5 @    @     @
	           |           ',  8
	        11 @  12@   @    '@
	           |       13      ',
	           |(0,0)            ', (1,0)
	           @----@----@----@----@  --> r
	         0      6    3    7     1
	*/
	r, s := R[0], R[1]
	pt1 := 128.0 / 3.0
	pt2 := 32.0 / 3.0
	cc := 1.0 - r - s
	t1 := r - 0.25
	t2 := r - 0.5
	t3 := r - 0.75
	t4 := s - 0.25
	t5 := s - 0.5
	t6 := s - 0.75
	t7 := cc - 0.25
	t8 := cc - 0.5
	t9 := cc - 0.75
	S[0] = pt2 * cc * t7 * t8 * t9
	S[1] = pt2 * r * t1 * t2 * t3
	S[2] = pt2 * s * t4 * t5 * t6
	S[3] = 64.0 * cc * r * t1 * t7
	S[4] = 64.0 * r * s * t1 * t4
	S[5] = 64.0 * s * cc * t4 * t7
	S[6] = pt1 * cc * r * t7 * t8
	S[7] = pt1 * cc * r * t1 * t2
	S[8] = pt1 * r * s * t1 * t2
	S[9] = pt1 * r * s * t4 * t5
	S[10] = pt1 * s * cc * t4 * t5
	S[11] = pt1 * s * cc * t7 * t8
	S[12] = 128.0 * r * s * cc * t7
	S[13] = 128.0 * r * s * t1 * cc
	S[14] = 128.0 * r * s * cc * t4

	if !derivs {
		return
	}

	dSdR[0][0] = -pt2 * (t8*t9*(t7+cc) + cc*t7*(t8+t9))
	dSdR[1][0] = pt2 * (t2*t3*(t1+r) + r*t1*(t3+t2))
	dSdR[2][0] = 0.0
	dSdR[3][0] = 64.0 * (cc*t7*(t1+r) - r*t1*(t7+cc))
	dSdR[4][0] = 64.0 * s * t4 * (t1 + r)
	dSdR[5][0] = -64.0 * s * t4 * (t7 + cc)
	dSdR[6][0] = pt1 * (cc*t7*t8 - r*(t8*(t7+cc)+cc*t7))
	dSdR[7][0] = pt1 * (cc*(t2*(t1+r)+r*t1) - r*t1*t2)
	dSdR[8][0] = pt1 * s * (t2*(t1+r) + r*t1)
	dSdR[9][0] = pt1 * s * t4 * t5
	dSdR[10][0] = -pt1 * s * t4 * t5
	dSdR[11][0] = -pt1 * s * (t8*(t7+cc) + cc*t7)
	dSdR[12][0] = 128.0 * s * (cc*t7 - r*(t7+cc))
	dSdR[13][0] = 128.0 * s * (cc*(t1+r) - r*t1)
	dSdR[14][0] = 128.0 * s * t4 * (cc - r)

	dSdR[0][1] = -pt2 * (t8*t9*(t7+cc) + cc*t7*(t8+t9))
	dSdR[1][1] = 0.0
	dSdR[2][1] = pt2 * (t5*t6*(t4+s) + s*t4*(t6+t5))
	dSdR[3][1] = -64.0 * r * t1 * (t7 + cc)
	dSdR[4][1] = 64.0 * r * t1 * (t4 + s)
	dSdR[5][1] = 64.0 * (cc*t7*(t4+s) - s*t4*(t7+cc))
	dSdR[6][1] = -pt1 * r * (t8*(t7+cc) + cc*t7)
	dSdR[7][1] = -pt1 * r * t1 * t2
	dSdR[8][1] = pt1 * r * t1 * t2
	dSdR[9][1] = pt1 * r * (t5*(t4+s) + s*t4)
	dSdR[10][1] = pt1 * ((cc * (t5*(t4+s) + s*t4)) - s*t4*t5)
	dSdR[11][1] = pt1 * (cc*t7*t8 - s*(t8*(t7+cc)+cc*t7))
	dSdR[12][1] = 128.0 * r * (cc*t7 - s*(cc+t7))
	dSdR[13][1] = 128.0 * r * t1 * (cc - s)
	dSdR[14][1] = 128.0 * r * (cc*(t4+s) - s*t4)
}
