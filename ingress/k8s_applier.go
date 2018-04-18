// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ingress

import (
	"github.com/golang/glog"
	"github.com/pkg/errors"
	k8s_v1beta1 "k8s.io/api/extensions/v1beta1"
	k8s_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s_kubernetes "k8s.io/client-go/kubernetes"
	k8s_clientcmd "k8s.io/client-go/tools/clientcmd"
)

// K8sApplier add Ingress to Client.
type K8sApplier struct {
	Kubeconfig string
	Namespace  string
}

// Apply a list of domains.
func (a *K8sApplier) Apply(ingress *k8s_v1beta1.Ingress) error {
	config, err := k8s_clientcmd.BuildConfigFromFlags("", a.Kubeconfig)
	if err != nil {
		return errors.Wrap(err, "build k8s config from flags failed")
	}
	clientset, err := k8s_kubernetes.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "create k8s config failed")
	}
	_, err = clientset.ExtensionsV1beta1().Ingresses(a.Namespace).Get(ingress.Name, k8s_v1.GetOptions{})
	if err != nil {
		_, err = clientset.ExtensionsV1beta1().Ingresses(a.Namespace).Create(ingress)
		if err != nil {
			return errors.Wrap(err, "create Ingress failed")
		}
		glog.V(0).Infof("ingress %s created successful", ingress.Name)
		return nil
	}
	_, err = clientset.ExtensionsV1beta1().Ingresses(a.Namespace).Update(ingress)
	if err != nil {
		return errors.Wrap(err, "update Ingress failed")
	}
	glog.V(0).Infof("ingress %s updated successful", ingress.Name)
	return nil
}
