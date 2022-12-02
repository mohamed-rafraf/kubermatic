/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apiserver

import (
	"k8c.io/kubermatic/v2/pkg/resources"
	"k8c.io/kubermatic/v2/pkg/resources/certificates"
	"k8c.io/kubermatic/v2/pkg/resources/certificates/triple"
	"k8c.io/reconciler/pkg/reconciling"
)

type etcdClientCertificateReconcilerData interface {
	GetRootCA() (*triple.KeyPair, error)
}

// EtcdClientCertificateReconciler returns a function to create/update the secret with the client certificate for authenticating against etcd.
func EtcdClientCertificateReconciler(data etcdClientCertificateReconcilerData) reconciling.NamedSecretReconcilerFactory {
	return certificates.GetClientCertificateReconciler(
		resources.ApiserverEtcdClientCertificateSecretName,
		"apiserver",
		nil,
		resources.ApiserverEtcdClientCertificateCertSecretKey,
		resources.ApiserverEtcdClientCertificateKeySecretKey,
		data.GetRootCA)
}
