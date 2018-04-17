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
	"github.com/seibert-media/k8s-ingress/domain"
	"github.com/seibert-media/k8s-ingress/ingress"
	"github.com/seibert-media/k8s-ingress/mocks"
)

var (
	versionInfo = flag.Bool("version", true, "show version info")
	urlPtr      = flag.String("url", "", "url to api")
	//stagingPtr  = flag.Bool("staging", false, "staging status")
	//dbg         = flag.Bool("debug", false, "enable debug mode")
	//sentryDsn   = flag.String("sentryDsn", "", "sentry dsn key")
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
	ingressSyncer := &ingress.Syncer{
		Applier: &mocks.DomainApplier{},
		Fetcher: &domain.Fetcher{
			URL:    *urlPtr,
			Client: http.DefaultClient,
		},
	}
	return ingressSyncer.Sync()
}
