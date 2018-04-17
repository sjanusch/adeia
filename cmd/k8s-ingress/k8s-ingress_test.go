// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
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
		"url":          server.URL(),
		"service-name": "test-service",
		"name":         "test-name",
		"server-port":  "8080",
		"namespace":    "test-namespace",
	}
})

var _ = Describe("the k8s-ingress", func() {
	var err error

	Describe("when asked for version", func() {
		It("prints version string", func() {
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, "-version"), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(time.Second)
			Expect(serverSession.ExitCode()).To(Equal(0))
			Expect(serverSession.Out).To(gbytes.Say(`-- //S/M k8s-ingress --
unknown - version unknown
  branch: 	unknown
  revision: 	unknown
  build date: 	unknown
  build user: 	unknown
  go version: 	unknown
`))
		})
	})
	Describe("when validating parameters", func() {
		It("returns with exitcode != 0 if no parameters have been given", func() {
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(time.Second)
			Expect(serverSession.ExitCode()).NotTo(Equal(0))
		})
		It("returns with exitcode 0 if called with valid args", func() {
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(time.Second)
			Expect(serverSession.ExitCode()).To(Equal(0))
		})
		It("returns error when service-name arg is missing", func() {
			delete(validargs, "service-name")
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(time.Second)
			Expect(serverSession.ExitCode()).NotTo(Equal(0))
			Expect(serverSession.Err).To(gbytes.Say("parameter service-name missing"))
		})
		It("returns error when name arg is missing", func() {
			delete(validargs, "name")
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(time.Second)
			Expect(serverSession.ExitCode()).NotTo(Equal(0))
			Expect(serverSession.Err).To(gbytes.Say("parameter name missing"))
		})
		It("returns error when server-port arg is missing", func() {
			delete(validargs, "server-port")
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(time.Second)
			Expect(serverSession.ExitCode()).NotTo(Equal(0))
			Expect(serverSession.Err).To(gbytes.Say("parameter server-port missing"))
		})
		It("returns error when server-port arg is missing", func() {
			delete(validargs, "namespace")
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(time.Second)
			Expect(serverSession.ExitCode()).NotTo(Equal(0))
			Expect(serverSession.Err).To(gbytes.Say("parameter namespace missing"))
		})
	})

	Describe("when called with valid input", func() {

		It("call given url", func() {
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(time.Second)
			Expect(serverSession.ExitCode()).To(Equal(0))
			Expect(len(server.ReceivedRequests())).To(Equal(1))
		})
	})

	It("writes ingress object to stdout", func() {
		var err error
		serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
		Expect(err).To(BeNil())
		serverSession.Wait(time.Second)
		Expect(serverSession.ExitCode()).To(Equal(0))
		Expect(serverSession.Out).To(gbytes.Say(`test`))
	})
})

func TestSystem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Test Suite")
}
