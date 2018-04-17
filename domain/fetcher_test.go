// Copyright 2018 The adeia Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seibert-media/adeia/domain"
	"github.com/seibert-media/adeia/mocks"
)

func TestDomain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domain Suite")
}

var _ = Describe("Fetcher", func() {
	var (
		domainFetcher *domain.Fetcher
		httpClient    *mocks.DomainClient
	)

	BeforeEach(func() {
		httpClient = &mocks.DomainClient{}
		httpClient.GetReturns(&http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(`["www.example.com"]`)),
		}, nil)
		domainFetcher = &domain.Fetcher{
			Client: httpClient,
			URL:    "http://server.com/domains",
		}
	})

	Describe("Fetch", func() {
		It("returns no error", func() {
			_, err := domainFetcher.Fetch()
			Expect(err).To(BeNil())
		})
		It("returns one domain", func() {
			list, _ := domainFetcher.Fetch()
			Expect(list).To(HaveLen(1))
		})
		It("does http call", func() {
			domainFetcher.Fetch()
			Expect(httpClient.GetCallCount()).To(Equal(1))
		})
		It("does not-nil request", func() {
			domainFetcher.Fetch()
			Expect(httpClient.GetArgsForCall(0)).NotTo(BeNil())
		})
		It("is using right api url", func() {
			domainFetcher.Fetch()
			Expect(httpClient.GetArgsForCall(0)).To(Equal("http://server.com/domains"))
		})
		It("is using different api url", func() {
			domainFetcher.URL = "http://server.de/domains"
			domainFetcher.Fetch()
			Expect(httpClient.GetArgsForCall(0)).To(Equal("http://server.de/domains"))
		})
		It("does error on empty url", func() {
			domainFetcher.URL = ""
			_, err := domainFetcher.Fetch()
			Expect(err).NotTo(BeNil())
		})
		It("does return error if http call fails", func() {
			domainFetcher.URL = "foo"
			httpClient.GetReturns(nil, errors.New("test"))
			_, err := domainFetcher.Fetch()
			Expect(err).NotTo(BeNil())
		})
		It("does return error if client returns empty response", func() {
			httpClient.GetReturns(nil, nil)
			_, err := domainFetcher.Fetch()
			Expect(err).NotTo(BeNil())
		})
		Describe("when json list contains example.com", func() {

			BeforeEach(func() {
				response := &http.Response{}
				response.Body = ioutil.NopCloser(bytes.NewBufferString(`["example.com"]`))
				httpClient.GetReturns(response, nil)
			})

			It("returns a list with example.com", func() {
				list, _ := domainFetcher.Fetch()
				Expect(list).To(HaveLen(1))
				Expect(list[0]).To(Equal(domain.Domain("example.com")))
			})
		})
		Describe("when json list contains two domains", func() {

			BeforeEach(func() {
				response := &http.Response{}
				response.Body = ioutil.NopCloser(bytes.NewBufferString(`["a.example.com","b.example.com"]`))
				httpClient.GetReturns(response, nil)
			})

			It("returns a list with example.com", func() {
				list, _ := domainFetcher.Fetch()
				Expect(list).To(HaveLen(2))
				Expect(list[0]).To(Equal(domain.Domain("a.example.com")))
				Expect(list[1]).To(Equal(domain.Domain("b.example.com")))
			})
		})
		Describe("when json is not parseable", func() {

			BeforeEach(func() {
				response := &http.Response{}
				response.Body = ioutil.NopCloser(bytes.NewBufferString(`foo bar`))
				httpClient.GetReturns(response, nil)
			})

			It("returns an error", func() {
				_, err := domainFetcher.Fetch()
				Expect(err).NotTo(BeNil())
			})

			It("returns an empty list", func() {
				list, _ := domainFetcher.Fetch()
				Expect(list).To(HaveLen(0))
			})
		})
	})
})
