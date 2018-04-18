// Copyright 2018 The adeia authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

// Domain name
type Domain string

func (d Domain) String() string {
	return string(d)
}
