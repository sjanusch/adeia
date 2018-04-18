// Copyright 2018 The adeia authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	flag "github.com/bborbe/flagenv"
	"github.com/golang/glog"
	"github.com/kolide/kit/version"
	"github.com/pkg/errors"
	"github.com/seibert-media/adeia"
	"github.com/seibert-media/adeia/domain"
	"github.com/seibert-media/adeia/ingress"
)

var (
	versionPtr     = flag.Bool("version", false, "show version info")
	urlPtr         = flag.String("url", "", "url to api")
	ingressNamePtr = flag.String("ingress-name", "", "name for ingress")
	serviceNamePtr = flag.String("service-name", "", "service name for Ingress http-rule")
	servicePortPtr = flag.String("service-port", "", "port for Ingress http-rule")
	namespacePtr   = flag.String("namespace", "", "k8s namespace to deploy ingresses")
	dryRunPtr      = flag.Bool("dry-run", false, "perform a trial run with no changes made and print ingress")
	kubeconfigPtr  = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if *versionPtr {
		fmt.Printf("-- //S/M adeia --\n")
		version.PrintFull()
		os.Exit(0)
	}

	if err := do(); err != nil {
		type stackTracer interface {
			StackTrace() errors.StackTrace
		}
		cause, ok := err.(stackTracer)
		if ok {
			glog.V(1).Info(cause.StackTrace())
		}
		glog.Fatal(err)
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
	if len(*ingressNamePtr) == 0 {
		return errors.New("parameter name missing")
	}
	if len(*servicePortPtr) == 0 {
		return errors.New("parameter service-port missing")
	}
	if len(*namespacePtr) == 0 {
		return errors.New("parameter namespace missing")
	}
	glog.V(1).Infof("arg dry-run: %b", *dryRunPtr)
	glog.V(1).Infof("arg kubeconfig: %s", *kubeconfigPtr)
	glog.V(1).Infof("arg url: %s", *urlPtr)
	glog.V(1).Infof("arg namespace: %s", *namespacePtr)
	glog.V(1).Infof("arg ingress-name: %s", *ingressNamePtr)
	glog.V(1).Infof("arg service-name: %s", *serviceNamePtr)
	glog.V(1).Infof("arg service-port: %s", *servicePortPtr)

	ingressSyncer := &adeia.Syncer{
		Applier: &ingress.K8sApplier{
			Kubeconfig: *kubeconfigPtr,
			Namespace:  *namespacePtr,
		},
		Creator: &ingress.Creator{
			Ingressname: *ingressNamePtr,
			Serviceport: *servicePortPtr,
			Servicename: *serviceNamePtr,
			Namespace:   *namespacePtr,
		},
		Fetcher: &domain.Fetcher{
			URL:    *urlPtr,
			Client: http.DefaultClient,
		},
	}
	if *dryRunPtr {
		glog.V(2).Infof("dry-run=true => use print applier")
		ingressSyncer.Applier = &ingress.PrintApplier{
			Out: os.Stdout,
		}
	}
	return ingressSyncer.Sync()
}
