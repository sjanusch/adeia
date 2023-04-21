// Copyright 2018 The adeia authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

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
var duration = 3 * time.Second

var _ = BeforeSuite(func() {
	var err error
	pathToServerBinary, err = gexec.Build("github.com/seibert-media/adeia/cmd/adeia")
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
		"ingress-name": "test-name",
		"namespace":    "test-namespace",
		"service-name": "test-service",
		"service-port": "8080",
		"dry-run":      "true",
		"kubeconfig":   "~/.kube/config",
	}
})

var _ = Describe("the adeia", func() {
	var err error
	Describe("when asked for version", func() {
		It("prints version string", func() {
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, "-version"), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(duration * time.Second)
			Expect(serverSession.ExitCode()).To(Equal(0))
			Expect(serverSession.Out).To(gbytes.Say(`-- //S/M adeia --
unknown - version unknown
  branch: 	unknown
  revision: 	unknown
  build date: 	unknown
  build user: 	unknown
  go version: 	go1.20.1
`))
		})
	})
	Describe("when validating parameters", func() {
		It("returns with exitcode != 0 if no parameters have been given", func() {
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(duration * time.Second)
			Expect(serverSession.ExitCode()).NotTo(Equal(0))
		})
		It("returns with exitcode 0 if called with valid args", func() {
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(duration * time.Second)
			Expect(serverSession.ExitCode()).To(Equal(0))
		})
		It("returns error when service-name arg is missing", func() {
			delete(validargs, "service-name")
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(duration * time.Second)
			Expect(serverSession.ExitCode()).NotTo(Equal(0))
			Expect(serverSession.Err).To(gbytes.Say("parameter service-name missing"))
		})
		It("returns error when name arg is missing", func() {
			delete(validargs, "ingress-name")
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(duration * time.Second)
			Expect(serverSession.ExitCode()).NotTo(Equal(0))
			Expect(serverSession.Err).To(gbytes.Say("parameter name missing"))
		})
		It("returns error when service-port arg is missing", func() {
			delete(validargs, "service-port")
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(duration * time.Second)
			Expect(serverSession.ExitCode()).NotTo(Equal(0))
			Expect(serverSession.Err).To(gbytes.Say("parameter service-port missing"))
		})
		It("returns error when namespace arg is missing", func() {
			delete(validargs, "namespace")
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(duration * time.Second)
			Expect(serverSession.ExitCode()).NotTo(Equal(0))
			Expect(serverSession.Err).To(gbytes.Say("parameter namespace missing"))
		})
	})
	Describe("when called with valid input", func() {
		It("call given url", func() {
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(duration * time.Second)
			Expect(serverSession.ExitCode()).To(Equal(0))
			Expect(len(server.ReceivedRequests())).To(Equal(1))
		})
	})
	Describe("called with dry run", func() {
		It("writes Ingress object to stdout", func() {
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(duration * time.Second)
			Expect(serverSession.ExitCode()).To(Equal(0))
			Expect(serverSession.Out).To(gbytes.Say(`apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    traefik.ingress.kubernetes.io/router.tls.certresolver: default
  creationTimestamp: null
  name: test-name
  namespace: test-namespace
spec:
  ingressClassName: traefik2
  rules:
  - host: a.example.com
    http:
      paths:
      - backend:
          service:
            name: test-service
            port:
              name: a.example.com-test-service
              number: 8080
        path: /
        pathType: Prefix
  - host: b.example.com
    http:
      paths:
      - backend:
          service:
            name: test-service
            port:
              name: b.example.com-test-service
              number: 8080
        path: /
        pathType: Prefix
status:
  loadBalancer: {}`))
		})

		It("call with dry run and different Ingress specs", func() {
			validargs["ingress-name"] = "superingress"
			validargs["service-port"] = "8080"
			validargs["service-name"] = "superservicename"
			validargs["namespace"] = "superspace"
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, validargs.list()...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(duration * time.Second)
			Expect(serverSession.ExitCode()).To(Equal(0))
			Expect(serverSession.Out).To(gbytes.Say(`apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    traefik.ingress.kubernetes.io/router.tls.certresolver: default
  creationTimestamp: null
  name: superingress
  namespace: superspace
spec:
  ingressClassName: traefik2
  rules:
  - host: a.example.com
    http:
      paths:
      - backend:
          service:
            name: superservicename
            port:
              name: a.example.com-superservicename
              number: 8080
        path: /
        pathType: Prefix
  - host: b.example.com
    http:
      paths:
      - backend:
          service:
            name: superservicename
            port:
              name: b.example.com-superservicename
              number: 8080
        path: /
        pathType: Prefix
status:
  loadBalancer: {}`))
		})
	})

	Describe("when given parameters via environment", func() {
		Describe("when no arguments are given via command line", func() {
			BeforeEach(func() {
				validargs = nil
			})
			It("uses version environment variable", func() {
				cmd := exec.Command(pathToServerBinary, validargs.list()...)
				cmd.Env = []string{"VERSION=true"}
				serverSession, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
				Expect(err).To(BeNil())
				serverSession.Wait(duration * time.Second)
				Expect(serverSession.ExitCode()).To(Equal(0))
			})
		})
		Describe("when version is set via command line", func() {
			BeforeEach(func() {
				validargs = map[string]string{
					"version": "true",
				}
			})
			It("uses command line argument value prioritized over environment", func() {
				cmd := exec.Command(pathToServerBinary, validargs.list()...)
				cmd.Env = []string{"VERSION=false"}
				serverSession, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
				Expect(err).To(BeNil())
				serverSession.Wait(duration * time.Second)
				Expect(serverSession.ExitCode()).To(Equal(0))
			})
		})
	})
})

func TestSystem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Test Suite")
}
