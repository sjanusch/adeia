// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system_test

import (
	"net/http"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/ghttp"
)

var pathToServerBinary string
var serverSession *gexec.Session
var server *ghttp.Server

var _ = BeforeSuite(func() {
	var err error
	pathToServerBinary, err = gexec.Build("github.com/seibert-media/k8s-ingress/cmd/k8s-ingress")
	Expect(err).NotTo(HaveOccurred())
})

var _ = BeforeEach(func() {
	server = ghttp.NewServer()
	server.RouteToHandler(http.MethodGet, "/", ghttp.RespondWithJSONEncoded(http.StatusOK, []string{"a.example.com", "b.example.com"}))
})

var _ = AfterEach(func() {
	serverSession.Interrupt()
	Eventually(serverSession).Should(gexec.Exit())
	server.Close()
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

var _ = Describe("the k8s-ingress", func() {
	var err error
	It("return with exitcode != 0 without needed parameter", func() {
		serverSession, err = gexec.Start(exec.Command(pathToServerBinary), GinkgoWriter, GinkgoWriter)
		Expect(err).To(BeNil())
		serverSession.Wait(time.Second)
		Expect(serverSession.ExitCode()).NotTo(Equal(0))
	})
	It("return with exitcode 0 if called with valid args", func() {
		serverSession, err = gexec.Start(exec.Command(pathToServerBinary, "-logtostderr", "-v=0", "-url="+server.URL()), GinkgoWriter, GinkgoWriter)
		Expect(err).To(BeNil())
		serverSession.Wait(time.Second)
		Expect(serverSession.ExitCode()).To(Equal(0))
	})
	It("call given url", func() {
		serverSession, err = gexec.Start(exec.Command(pathToServerBinary, "-logtostderr", "-v=0", "-url="+server.URL()), GinkgoWriter, GinkgoWriter)
		Expect(err).To(BeNil())
		serverSession.Wait(time.Second)
		Expect(serverSession.ExitCode()).To(Equal(0))
		Expect(len(server.ReceivedRequests())).To(Equal(1))
	})
})

func TestSystem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Test Suite")
}
