//go:build !ignore_autogenerated

/*
Copyright  The Kubermatic Kubernetes Platform contributors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppNamespaceSpec) DeepCopyInto(out *AppNamespaceSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppNamespaceSpec.
func (in *AppNamespaceSpec) DeepCopy() *AppNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(AppNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDefinition) DeepCopyInto(out *ApplicationDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDefinition.
func (in *ApplicationDefinition) DeepCopy() *ApplicationDefinition {
	if in == nil {
		return nil
	}
	out := new(ApplicationDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDefinitionList) DeepCopyInto(out *ApplicationDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDefinitionList.
func (in *ApplicationDefinitionList) DeepCopy() *ApplicationDefinitionList {
	if in == nil {
		return nil
	}
	out := new(ApplicationDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDefinitionSpec) DeepCopyInto(out *ApplicationDefinitionSpec) {
	*out = *in
	if in.DefaultValues != nil {
		in, out := &in.DefaultValues, &out.DefaultValues
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultNamespace != nil {
		in, out := &in.DefaultNamespace, &out.DefaultNamespace
		*out = new(AppNamespaceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultDeployOptions != nil {
		in, out := &in.DefaultDeployOptions, &out.DefaultDeployOptions
		*out = new(DeployOptions)
		(*in).DeepCopyInto(*out)
	}
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]ApplicationVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDefinitionSpec.
func (in *ApplicationDefinitionSpec) DeepCopy() *ApplicationDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInstallation) DeepCopyInto(out *ApplicationInstallation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInstallation.
func (in *ApplicationInstallation) DeepCopy() *ApplicationInstallation {
	if in == nil {
		return nil
	}
	out := new(ApplicationInstallation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationInstallation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInstallationCondition) DeepCopyInto(out *ApplicationInstallationCondition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInstallationCondition.
func (in *ApplicationInstallationCondition) DeepCopy() *ApplicationInstallationCondition {
	if in == nil {
		return nil
	}
	out := new(ApplicationInstallationCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInstallationList) DeepCopyInto(out *ApplicationInstallationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationInstallation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInstallationList.
func (in *ApplicationInstallationList) DeepCopy() *ApplicationInstallationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationInstallationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationInstallationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInstallationSpec) DeepCopyInto(out *ApplicationInstallationSpec) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(AppNamespaceSpec)
		(*in).DeepCopyInto(*out)
	}
	out.ApplicationRef = in.ApplicationRef
	in.Values.DeepCopyInto(&out.Values)
	out.ReconciliationInterval = in.ReconciliationInterval
	if in.DeployOptions != nil {
		in, out := &in.DeployOptions, &out.DeployOptions
		*out = new(DeployOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInstallationSpec.
func (in *ApplicationInstallationSpec) DeepCopy() *ApplicationInstallationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationInstallationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInstallationStatus) DeepCopyInto(out *ApplicationInstallationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(map[ApplicationInstallationConditionType]ApplicationInstallationCondition, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ApplicationVersion != nil {
		in, out := &in.ApplicationVersion, &out.ApplicationVersion
		*out = new(ApplicationVersion)
		(*in).DeepCopyInto(*out)
	}
	if in.HelmRelease != nil {
		in, out := &in.HelmRelease, &out.HelmRelease
		*out = new(HelmRelease)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInstallationStatus.
func (in *ApplicationInstallationStatus) DeepCopy() *ApplicationInstallationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationInstallationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationRef) DeepCopyInto(out *ApplicationRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationRef.
func (in *ApplicationRef) DeepCopy() *ApplicationRef {
	if in == nil {
		return nil
	}
	out := new(ApplicationRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSource) DeepCopyInto(out *ApplicationSource) {
	*out = *in
	if in.Helm != nil {
		in, out := &in.Helm, &out.Helm
		*out = new(HelmSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSource.
func (in *ApplicationSource) DeepCopy() *ApplicationSource {
	if in == nil {
		return nil
	}
	out := new(ApplicationSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationTemplate) DeepCopyInto(out *ApplicationTemplate) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	if in.DependencyCredentials != nil {
		in, out := &in.DependencyCredentials, &out.DependencyCredentials
		*out = new(DependencyCredentials)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationTemplate.
func (in *ApplicationTemplate) DeepCopy() *ApplicationTemplate {
	if in == nil {
		return nil
	}
	out := new(ApplicationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationVersion) DeepCopyInto(out *ApplicationVersion) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationVersion.
func (in *ApplicationVersion) DeepCopy() *ApplicationVersion {
	if in == nil {
		return nil
	}
	out := new(ApplicationVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultingSelector) DeepCopyInto(out *DefaultingSelector) {
	*out = *in
	if in.Datacenters != nil {
		in, out := &in.Datacenters, &out.Datacenters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultingSelector.
func (in *DefaultingSelector) DeepCopy() *DefaultingSelector {
	if in == nil {
		return nil
	}
	out := new(DefaultingSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependencyCredentials) DeepCopyInto(out *DependencyCredentials) {
	*out = *in
	if in.HelmCredentials != nil {
		in, out := &in.HelmCredentials, &out.HelmCredentials
		*out = new(HelmCredentials)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependencyCredentials.
func (in *DependencyCredentials) DeepCopy() *DependencyCredentials {
	if in == nil {
		return nil
	}
	out := new(DependencyCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployOptions) DeepCopyInto(out *DeployOptions) {
	*out = *in
	if in.Helm != nil {
		in, out := &in.Helm, &out.Helm
		*out = new(HelmDeployOptions)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployOptions.
func (in *DeployOptions) DeepCopy() *DeployOptions {
	if in == nil {
		return nil
	}
	out := new(DeployOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitCredentials) DeepCopyInto(out *GitCredentials) {
	*out = *in
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(corev1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(corev1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(corev1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SSHKey != nil {
		in, out := &in.SSHKey, &out.SSHKey
		*out = new(corev1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitCredentials.
func (in *GitCredentials) DeepCopy() *GitCredentials {
	if in == nil {
		return nil
	}
	out := new(GitCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitReference) DeepCopyInto(out *GitReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitReference.
func (in *GitReference) DeepCopy() *GitReference {
	if in == nil {
		return nil
	}
	out := new(GitReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitSource) DeepCopyInto(out *GitSource) {
	*out = *in
	out.Ref = in.Ref
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(GitCredentials)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitSource.
func (in *GitSource) DeepCopy() *GitSource {
	if in == nil {
		return nil
	}
	out := new(GitSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmCredentials) DeepCopyInto(out *HelmCredentials) {
	*out = *in
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(corev1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(corev1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.RegistryConfigFile != nil {
		in, out := &in.RegistryConfigFile, &out.RegistryConfigFile
		*out = new(corev1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmCredentials.
func (in *HelmCredentials) DeepCopy() *HelmCredentials {
	if in == nil {
		return nil
	}
	out := new(HelmCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmDeployOptions) DeepCopyInto(out *HelmDeployOptions) {
	*out = *in
	out.Timeout = in.Timeout
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmDeployOptions.
func (in *HelmDeployOptions) DeepCopy() *HelmDeployOptions {
	if in == nil {
		return nil
	}
	out := new(HelmDeployOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmRelease) DeepCopyInto(out *HelmRelease) {
	*out = *in
	if in.Info != nil {
		in, out := &in.Info, &out.Info
		*out = new(HelmReleaseInfo)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmRelease.
func (in *HelmRelease) DeepCopy() *HelmRelease {
	if in == nil {
		return nil
	}
	out := new(HelmRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseInfo) DeepCopyInto(out *HelmReleaseInfo) {
	*out = *in
	in.FirstDeployed.DeepCopyInto(&out.FirstDeployed)
	in.LastDeployed.DeepCopyInto(&out.LastDeployed)
	in.Deleted.DeepCopyInto(&out.Deleted)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseInfo.
func (in *HelmReleaseInfo) DeepCopy() *HelmReleaseInfo {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmSource) DeepCopyInto(out *HelmSource) {
	*out = *in
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(bool)
		**out = **in
	}
	if in.PlainHTTP != nil {
		in, out := &in.PlainHTTP, &out.PlainHTTP
		*out = new(bool)
		**out = **in
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(HelmCredentials)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmSource.
func (in *HelmSource) DeepCopy() *HelmSource {
	if in == nil {
		return nil
	}
	out := new(HelmSource)
	in.DeepCopyInto(out)
	return out
}
