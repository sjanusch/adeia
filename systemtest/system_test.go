// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system_test

import (
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var pathToServerBinary string
var serverSession *gexec.Session

var _ = BeforeSuite(func() {
	var err error
	pathToServerBinary, err = gexec.Build("github.com/seibert-media/k8s-ingress/cmd/k8s-ingress")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

var _ = AfterEach(func() {
	serverSession.Interrupt()
	Eventually(serverSession).Should(gexec.Exit())
})

var _ = Describe("the k8s-ingress", func() {
	It("return with exitcode != 0 without needed parameter", func() {
		var err error
		serverSession, err = gexec.Start(exec.Command(pathToServerBinary), GinkgoWriter, GinkgoWriter)
		Expect(err).To(BeNil())
		serverSession.Wait(time.Second)
		Expect(serverSession.ExitCode()).NotTo(Equal(0))
	})
	It("return with exitcode 0 if called with valid args", func() {
		var err error
		serverSession, err = gexec.Start(exec.Command(pathToServerBinary, "-logtostderr", "-v=0", "-url=http://localhost:8080"), GinkgoWriter, GinkgoWriter)
		Expect(err).To(BeNil())
		serverSession.Wait(time.Second)
		Expect(serverSession.ExitCode()).To(Equal(0))
	})
})

func TestSystem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Test Suite")
}
