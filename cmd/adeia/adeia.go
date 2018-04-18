// Copyright 2018 The adeia authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	flag "github.com/bborbe/flagenv"
	"github.com/golang/glog"
	"github.com/kolide/kit/version"
	"github.com/seibert-media/adeia"
	"github.com/seibert-media/adeia/domain"
	"github.com/seibert-media/adeia/ingress"
	k8s_homedir "k8s.io/client-go/util/homedir"
)

var (
	versionPtr     = flag.Bool("version", false, "show version info")
	urlPtr         = flag.String("url", "", "url to api")
	ingressNamePtr = flag.String("ingress-name", "", "name for ingress")
	serviceNamePtr = flag.String("service-name", "", "service name for ingress http-rule")
	servicePortPtr = flag.String("service-port", "", "port for ingress http-rule")
	namespacePtr   = flag.String("namespace", "", "k8s namespace to deploy ingresses")
	dryRunPtr      = flag.Bool("dry-run", false, "perform a trial run with no changes made and print ingress")
	kubeconfigPtr  *string
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	if home := k8s_homedir.HomeDir(); home != "" {
		kubeconfigPtr = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfigPtr = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if *versionPtr {
		fmt.Printf("-- //S/M adeia --\n")
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
	if len(*ingressNamePtr) == 0 {
		return errors.New("parameter name missing")
	}
	if len(*servicePortPtr) == 0 {
		return errors.New("parameter service-port missing")
	}
	if len(*namespacePtr) == 0 {
		return errors.New("parameter namespace missing")
	}
	if len(*kubeconfigPtr) == 0 {
		return errors.New("parameter kubeconfig missing")
	}

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
		ingressSyncer.Applier = &ingress.PrintApplier{
			Out: os.Stdout,
		}
	}
	return ingressSyncer.Sync()
}
