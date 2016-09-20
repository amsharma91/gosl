// Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rnd

import (
	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/utl"
)

// TextHist prints a text histogram
//  Input:
//   labels -- labels
//   counts -- frequencies
func TextHist(labels []string, counts []int, barlen int) string {

	// check
	chk.IntAssert(len(labels), len(counts))
	if len(counts) < 2 {
		return "counts slice is too short\n"
	}

	// scale
	fmax := counts[0]
	lmax := 0
	Lmax := 0
	for i, f := range counts {
		fmax = imax(fmax, f)
		lmax = imax(lmax, len(labels[i]))
		Lmax = imax(Lmax, len(io.Sf("%d", f)))
	}
	if fmax < 1 {
		return io.Sf("max frequency is too small: fmax=%d\n", fmax)
	}
	scale := float64(barlen) / float64(fmax)

	// print
	sz := io.Sf("%d", lmax+1)
	Sz := io.Sf("%d", Lmax+1)
	l := ""
	total := 0
	for i, f := range counts {
		l += io.Sf("%"+sz+"s | %"+Sz+"d ", labels[i], f)
		n := int(float64(f) * scale)
		for j := 0; j < n; j++ {
			l += "#"
		}
		l += "\n"
		total += f
	}
	l += io.Sf("%"+sz+"s   %"+Sz+"d\n", "", total)
	return l
}

// BuildTextHist builds a text histogram
//  Input:
//   xmin      -- station xmin
//   xmax      -- station xmax
//   nstations -- number of stations
//   values    -- values to be counted
//   numfmt    -- number format
//   barlen    -- max length of bar
func BuildTextHist(xmin, xmax float64, nstations int, values []float64, numfmt string, barlen int) string {
	hist := Histogram{Stations: utl.LinSpace(xmin, xmax, nstations)}
	hist.Count(values, true)
	return TextHist(hist.GenLabels(numfmt), hist.Counts, 60)
}

// Histogram holds data for computing/plotting histograms
//
//  bin[i] corresponds to station[i] <= x < station[i+1]
//
//       [ bin[0] )[ bin[1] )[ bin[2] )[ bin[3] )[ bin[4] )
//    ---|---------|---------|---------|---------|---------|---  x
//     s[0]      s[1]      s[2]      s[3]      s[4]      s[5]
//
type Histogram struct {
	Stations []float64 // stations
	Counts   []int     // counts
}

// FindBin finds where x falls in
// returns -1 if x is outside the range
func (o Histogram) FindBin(x float64) int {

	// check
	if len(o.Stations) < 2 {
		chk.Panic("Histogram must have at least 2 stations")
	}
	if x < o.Stations[0] {
		return -1
	}
	if x >= o.Stations[len(o.Stations)-1] {
		return -1
	}

	// perform binary search
	upper := len(o.Stations)
	lower := 0
	mid := 0
	for upper-lower > 1 {
		mid = (upper + lower) / 2
		if x >= o.Stations[mid] {
			lower = mid
		} else {
			upper = mid
		}
	}
	return lower
}

// Count counts how many items fall within each bin
func (o *Histogram) Count(vals []float64, clear bool) {

	// check
	if len(o.Stations) < 2 {
		chk.Panic("Histogram must have at least 2 stations")
	}

	// allocate/clear counts
	nbins := len(o.Stations) - 1
	if len(o.Counts) != nbins {
		o.Counts = make([]int, nbins)
	} else if clear {
		for i := 0; i < nbins; i++ {
			o.Counts[i] = 0
		}
	}

	// add entries to bins
	for _, x := range vals {
		idx := o.FindBin(x)
		if idx >= 0 {
			o.Counts[idx] += 1
		}
	}
}

// GenLabels generate nice labels identifying bins
func (o Histogram) GenLabels(numfmt string) (labels []string) {
	if len(o.Stations) < 2 {
		chk.Panic("Histogram must have at least 2 stations")
	}
	nbins := len(o.Stations) - 1
	labels = make([]string, nbins)
	for i := 0; i < nbins; i++ {
		labels[i] = io.Sf("["+numfmt+","+numfmt+")", o.Stations[i], o.Stations[i+1])
	}
	return
}

// IntHistogram holds data for computing/plotting histograms with integers
//
//  bin[i] corresponds to station[i] <= x < station[i+1]
//
//       [ bin[0] )[ bin[1] )[ bin[2] )[ bin[3] )[ bin[4] )
//    ---|---------|---------|---------|---------|---------|---  x
//     s[0]      s[1]      s[2]      s[3]      s[4]      s[5]
//
type IntHistogram struct {
	Stations []int // stations
	Counts   []int // counts
}

// FindBin finds where x falls in
// returns -1 if x is outside the range
func (o IntHistogram) FindBin(x int) int {

	// check
	if len(o.Stations) < 2 {
		chk.Panic("IntHistogram must have at least 2 stations")
	}
	if x < o.Stations[0] {
		return -1
	}
	if x >= o.Stations[len(o.Stations)-1] {
		return -1
	}

	// perform binary search
	upper := len(o.Stations)
	lower := 0
	mid := 0
	for upper-lower > 1 {
		mid = (upper + lower) / 2
		if x >= o.Stations[mid] {
			lower = mid
		} else {
			upper = mid
		}
	}
	return lower
}

// Count counts how many items fall within each bin
func (o *IntHistogram) Count(vals []int, clear bool) {

	// check
	if len(o.Stations) < 2 {
		chk.Panic("IntHistogram must have at least 2 stations")
	}

	// allocate/clear counts
	nbins := len(o.Stations) - 1
	if len(o.Counts) != nbins {
		o.Counts = make([]int, nbins)
	} else if clear {
		for i := 0; i < nbins; i++ {
			o.Counts[i] = 0
		}
	}

	// add entries to bins
	for _, x := range vals {
		idx := o.FindBin(x)
		if idx >= 0 {
			o.Counts[idx] += 1
		}
	}
}

// GenLabels generate nice labels identifying bins
func (o IntHistogram) GenLabels(numfmt string) (labels []string) {
	if len(o.Stations) < 2 {
		chk.Panic("IntHistogram must have at least 2 stations")
	}
	nbins := len(o.Stations) - 1
	labels = make([]string, nbins)
	for i := 0; i < nbins; i++ {
		labels[i] = io.Sf("["+numfmt+","+numfmt+")", o.Stations[i], o.Stations[i+1])
	}
	return
}
