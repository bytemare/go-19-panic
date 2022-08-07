// SPDX-License-Identifier: MIT
//
// Copyright (C) 2021 Daniel Bourdrez. All Rights Reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree or at
// https://spdx.org/licenses/MIT.html

package nist

/*
EcPoint generalizes the structures in filippo.io/nistec. There used to be things in there,
but the panic works without any of it.
*/
type EcPoint[point any] interface{}

type Interface interface {
	Set(element Interface)
}

// Point implements Interface.
type Point[P EcPoint[P]] struct {
	P P
}

func (e *Point[P]) Set(element Interface) {
	// The real code doesn't really do following, don't judge me.
	_ = element.(*Point[P])
}

// Group is a thing from which we can spawn Points.
type Group[Point EcPoint[Point]] struct {
	new func() Point
}

func (g Group[P]) NewPoint() Interface {
	p := g.new()
	return g.newPoint(p)
}

func (g Group[P]) newPoint(p P) *Point[P] {
	return &Point[P]{
		P: p,
	}
}
