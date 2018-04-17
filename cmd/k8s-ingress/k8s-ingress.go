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

	"github.com/golang/glog"
	"github.com/kolide/kit/version"
	"github.com/seibert-media/k8s-ingress/applier"
	"github.com/seibert-media/k8s-ingress/converter"
	"github.com/seibert-media/k8s-ingress/domain"
	"github.com/seibert-media/k8s-ingress/ingress"
)

var (
	versionInfo = flag.Bool("version", false, "show version info")
	urlPtr      = flag.String("url", "", "url to api")
	serviceName = flag.String("service-name", "", "service name for ingress http-rule")
	name        = flag.String("name", "", "name for ingress")
	serverPort  = flag.String("server-port", "", "port for ingress http-rule")
	namespace   = flag.String("namespace", "", "k8s namespace to deploy ingresses")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if *versionInfo {
		fmt.Printf("-- //S/M k8s-ingress --\n")
		version.PrintFull()
	}

	if err := do(); err != nil {
		glog.Error(err)
		os.Exit(1)
	}

	fmt.Println("finished")
}

func do() error {
	if len(*urlPtr) == 0 {
		return errors.New("parameter url missing")
	}
	if len(*serviceName) == 0 {
		return errors.New("parameter service-name missing")
	}
	if len(*name) == 0 {
		return errors.New("parameter name missing")
	}
	if len(*serverPort) == 0 {
		return errors.New("parameter server-port missing")
	}
	if len(*namespace) == 0 {
		return errors.New("parameter namespace missing")
	}
	ingressSyncer := &ingress.Syncer{
		Applier:   &applier.Applier{
			Converter: &converter.Converter{},
		},
		Fetcher:   &domain.Fetcher{},
	}
	return ingressSyncer.Sync()
}
