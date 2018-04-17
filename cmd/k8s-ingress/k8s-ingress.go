// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime"

	"net/http"

	"github.com/golang/glog"
	"github.com/kolide/kit/version"
	"github.com/seibert-media/k8s-ingress/pkg"
	"github.com/seibert-media/k8s-ingress/pkg/domain"
	"github.com/seibert-media/k8s-ingress/pkg/ingress"
)

var (
	versionPtr     = flag.Bool("version", false, "show version info")
	urlPtr         = flag.String("url", "", "url to api")
	namePtr        = flag.String("name", "", "name for ingress")
	serviceNamePtr = flag.String("service-name", "", "service name for ingress http-rule")
	serverPortPtr  = flag.String("server-port", "", "port for ingress http-rule")
	namespacePtr   = flag.String("namespace", "", "k8s namespace to deploy ingresses")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if *versionPtr {
		fmt.Printf("-- //S/M k8s-ingress --\n")
		version.PrintFull()
		os.Exit(0)
	}

	if err := do(); err != nil {
		glog.Error(err)
		os.Exit(1)
	}
}

func do() error {
	if len(*urlPtr) == 0 {
		return errors.New("parameter url missing")
	}
	if len(*serviceNamePtr) == 0 {
		return errors.New("parameter service-name missing")
	}
	if len(*namePtr) == 0 {
		return errors.New("parameter name missing")
	}
	if len(*serverPortPtr) == 0 {
		return errors.New("parameter server-port missing")
	}
	if len(*namespacePtr) == 0 {
		return errors.New("parameter namespace missing")
	}
	ingressSyncer := &pkg.Syncer{
		Applier: &ingress.PrintApplier{},
		Creator: &ingress.Creator{},
		Fetcher: &domain.Fetcher{
			URL:    *urlPtr,
			Client: http.DefaultClient,
		},
	}
	return ingressSyncer.Sync()
}
