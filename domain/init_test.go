// Copyright 2018 The adeia authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDomain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domain Suite")
}
