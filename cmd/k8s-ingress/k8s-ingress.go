// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/golang/glog"
	"github.com/kolide/kit/version"
	"github.com/seibert-media/k8s-ingress/domain"
	"github.com/seibert-media/k8s-ingress/ingress"
	"github.com/seibert-media/k8s-ingress/mocks"
)

var (
	stagingPtr  = flag.Bool("staging", false, "staging status")
	versionInfo = flag.Bool("version", true, "show version info")
	dbg         = flag.Bool("debug", false, "enable debug mode")
	sentryDsn   = flag.String("sentryDsn", "", "sentry dsn key")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")

	if *versionInfo {
		fmt.Printf("-- //S/M k8s-ingress --\n")
		version.PrintFull()
	}

	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	ingressSyncer := &ingress.Syncer{
		Applier: &mocks.DomainApplier{},
		Fetcher: &domain.Fetcher{},
	}
	if err := ingressSyncer.Sync(); err != nil {
		glog.Exit(err)
	}

	fmt.Println("finished")
}
