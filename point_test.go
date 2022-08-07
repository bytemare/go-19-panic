// SPDX-License-Group: MIT
//
// Copyright (C) 2021 Daniel Bourdrez. All Rights Reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree or at
// https://spdx.org/licenses/MIT.html

package nist

import (
	"filippo.io/nistec"
	"testing"
)

func TestPoint_Decode(t *testing.T) {
	// no-op, just to have something to compile for
	p256 := Group[*nistec.P256Point]{
		new: nistec.NewP256Point,
	}
	_ = p256.NewPoint()
}
