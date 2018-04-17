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
	"fmt"
	"github.com/onsi/gomega/gbytes"
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

type args map[string]string

func (a args) list() []string {
	var result []string
	for k, v := range a {
		if len(v) == 0 {
			result = append(result, fmt.Sprintf("-%s", k))
		} else {
			result = append(result, fmt.Sprintf("-%s=%s", k, v))
		}
	}
	return result
}

var validargs args

var _ = BeforeEach(func() {
	validargs = map[string]string{
		"logtostderr":  "",
		"v":            "0",
		"url":          "http://localhost:8080",
		"service-name": "test-service",
		"name":         "test-name",
		"server-port": 	"8080",
	}
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
		serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
		Expect(err).To(BeNil())
		serverSession.Wait(time.Second)
		Expect(serverSession.ExitCode()).To(Equal(0))
	})
	It("return error when service-name arg is missing", func() {
		var err error
		delete(validargs, "service-name")
		serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
		Expect(err).To(BeNil())
		serverSession.Wait(time.Second)
		Expect(serverSession.ExitCode()).NotTo(Equal(0))
		Expect(serverSession.Err).To(gbytes.Say("parameter service-name missing"))
	})
	It("return error when name arg is missing", func() {
		var err error
		delete(validargs, "name")
		serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
		Expect(err).To(BeNil())
		serverSession.Wait(time.Second)
		Expect(serverSession.ExitCode()).NotTo(Equal(0))
		Expect(serverSession.Err).To(gbytes.Say("parameter name missing"))
	})
	It("return error when server-port arg is missing", func() {
		var err error
		delete(validargs, "server-port")
		serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
		Expect(err).To(BeNil())
		serverSession.Wait(time.Second)
		Expect(serverSession.ExitCode()).NotTo(Equal(0))
		Expect(serverSession.Err).To(gbytes.Say("parameter server-port missing"))
	})
})

func TestSystem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Test Suite")
}
