/*
Copyright 2022 The Kubermatic Kubernetes Platform contributors.

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

package scenarios

import (
	"context"
	"fmt"

	kubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/cmd/conformance-tester/pkg/types"
	"k8c.io/kubermatic/v2/pkg/machine/provider"
	clusterv1alpha1 "k8c.io/machine-controller/sdk/apis/cluster/v1alpha1"
	"k8c.io/machine-controller/sdk/providerconfig"

	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	nutanixCPUs     = 2
	nutanixMemoryMB = 4096
	nutanixDiskSize = 40
)

type nutanixScenario struct {
	baseScenario
}

func (s *nutanixScenario) compatibleOperatingSystems() sets.Set[providerconfig.OperatingSystem] {
	return sets.New[providerconfig.OperatingSystem](
		providerconfig.OperatingSystemUbuntu,
		providerconfig.OperatingSystemRHEL,
		providerconfig.OperatingSystemFlatcar,
		providerconfig.OperatingSystemRockyLinux,
	)
}

func (s *nutanixScenario) IsValid() error {
	if err := s.baseScenario.IsValid(); err != nil {
		return err
	}

	if compat := s.compatibleOperatingSystems(); !compat.Has(s.operatingSystem) {
		return fmt.Errorf("provider supports only %v", sets.List(compat))
	}

	return nil
}

func (s *nutanixScenario) Cluster(secrets types.Secrets) *kubermaticv1.ClusterSpec {
	return &kubermaticv1.ClusterSpec{
		Cloud: kubermaticv1.CloudSpec{
			DatacenterName: secrets.Nutanix.KKPDatacenter,
			Nutanix: &kubermaticv1.NutanixCloudSpec{
				Username: secrets.Nutanix.Username,
				Password: secrets.Nutanix.Password,
				CSI: &kubermaticv1.NutanixCSIConfig{
					Endpoint: secrets.Nutanix.CSIEndpoint,
					Password: secrets.Nutanix.CSIPassword,
					Username: secrets.Nutanix.CSIUsername,
				},
				ProxyURL:    secrets.Nutanix.ProxyURL,
				ClusterName: secrets.Nutanix.ClusterName,
				ProjectName: secrets.Nutanix.ProjectName,
			},
		},
		Version: s.clusterVersion,
	}
}

func (s *nutanixScenario) MachineDeployments(_ context.Context, num int, secrets types.Secrets, cluster *kubermaticv1.Cluster, sshPubKeys []string) ([]clusterv1alpha1.MachineDeployment, error) {
	cloudProviderSpec := provider.NewNutanixConfig().
		WithSubnetName(secrets.Nutanix.SubnetName).
		WithCPUs(nutanixCPUs).
		WithMemoryMB(nutanixMemoryMB).
		WithDiskSize(nutanixDiskSize).
		Build()

	md, err := s.createMachineDeployment(cluster, num, cloudProviderSpec, sshPubKeys, secrets)
	if err != nil {
		return nil, err
	}

	return []clusterv1alpha1.MachineDeployment{md}, nil
}
