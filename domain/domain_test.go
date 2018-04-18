// Copyright 2018 The adeia authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Domain", func() {
	var domain = Domain("foo")
	Describe("String", func() {
		It("returns content as string", func() {
			Expect(domain.String()).To(Equal("foo"))
		})
	})
})
