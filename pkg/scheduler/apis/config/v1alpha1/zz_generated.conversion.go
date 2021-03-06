// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.
Copyright 2020 Authors of Arktos - file modified.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
	v1alpha1 "k8s.io/kube-scheduler/config/v1alpha1"
	config "k8s.io/kubernetes/pkg/scheduler/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha1.KubeSchedulerConfiguration)(nil), (*config.KubeSchedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration(a.(*v1alpha1.KubeSchedulerConfiguration), b.(*config.KubeSchedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeSchedulerConfiguration)(nil), (*v1alpha1.KubeSchedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(a.(*config.KubeSchedulerConfiguration), b.(*v1alpha1.KubeSchedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.KubeSchedulerLeaderElectionConfiguration)(nil), (*config.KubeSchedulerLeaderElectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(a.(*v1alpha1.KubeSchedulerLeaderElectionConfiguration), b.(*config.KubeSchedulerLeaderElectionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeSchedulerLeaderElectionConfiguration)(nil), (*v1alpha1.KubeSchedulerLeaderElectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(a.(*config.KubeSchedulerLeaderElectionConfiguration), b.(*v1alpha1.KubeSchedulerLeaderElectionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.Plugin)(nil), (*config.Plugin)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Plugin_To_config_Plugin(a.(*v1alpha1.Plugin), b.(*config.Plugin), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Plugin)(nil), (*v1alpha1.Plugin)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Plugin_To_v1alpha1_Plugin(a.(*config.Plugin), b.(*v1alpha1.Plugin), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.PluginConfig)(nil), (*config.PluginConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PluginConfig_To_config_PluginConfig(a.(*v1alpha1.PluginConfig), b.(*config.PluginConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PluginConfig)(nil), (*v1alpha1.PluginConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PluginConfig_To_v1alpha1_PluginConfig(a.(*config.PluginConfig), b.(*v1alpha1.PluginConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.PluginSet)(nil), (*config.PluginSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PluginSet_To_config_PluginSet(a.(*v1alpha1.PluginSet), b.(*config.PluginSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PluginSet)(nil), (*v1alpha1.PluginSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PluginSet_To_v1alpha1_PluginSet(a.(*config.PluginSet), b.(*v1alpha1.PluginSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.Plugins)(nil), (*config.Plugins)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Plugins_To_config_Plugins(a.(*v1alpha1.Plugins), b.(*config.Plugins), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Plugins)(nil), (*v1alpha1.Plugins)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Plugins_To_v1alpha1_Plugins(a.(*config.Plugins), b.(*v1alpha1.Plugins), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.SchedulerAlgorithmSource)(nil), (*config.SchedulerAlgorithmSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(a.(*v1alpha1.SchedulerAlgorithmSource), b.(*config.SchedulerAlgorithmSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SchedulerAlgorithmSource)(nil), (*v1alpha1.SchedulerAlgorithmSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(a.(*config.SchedulerAlgorithmSource), b.(*v1alpha1.SchedulerAlgorithmSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.SchedulerPolicyConfigMapSource)(nil), (*config.SchedulerPolicyConfigMapSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource(a.(*v1alpha1.SchedulerPolicyConfigMapSource), b.(*config.SchedulerPolicyConfigMapSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SchedulerPolicyConfigMapSource)(nil), (*v1alpha1.SchedulerPolicyConfigMapSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(a.(*config.SchedulerPolicyConfigMapSource), b.(*v1alpha1.SchedulerPolicyConfigMapSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.SchedulerPolicyFileSource)(nil), (*config.SchedulerPolicyFileSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource(a.(*v1alpha1.SchedulerPolicyFileSource), b.(*config.SchedulerPolicyFileSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SchedulerPolicyFileSource)(nil), (*v1alpha1.SchedulerPolicyFileSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(a.(*config.SchedulerPolicyFileSource), b.(*v1alpha1.SchedulerPolicyFileSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.SchedulerPolicySource)(nil), (*config.SchedulerPolicySource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SchedulerPolicySource_To_config_SchedulerPolicySource(a.(*v1alpha1.SchedulerPolicySource), b.(*config.SchedulerPolicySource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SchedulerPolicySource)(nil), (*v1alpha1.SchedulerPolicySource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(a.(*config.SchedulerPolicySource), b.(*v1alpha1.SchedulerPolicySource), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration(in *v1alpha1.KubeSchedulerConfiguration, out *config.KubeSchedulerConfiguration, s conversion.Scope) error {
	out.SchedulerName = in.SchedulerName
	if err := Convert_v1alpha1_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(&in.AlgorithmSource, &out.AlgorithmSource, s); err != nil {
		return err
	}
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	if err := Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := configv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.HealthzBindAddress = in.HealthzBindAddress
	out.MetricsBindAddress = in.MetricsBindAddress
	if err := configv1alpha1.Convert_v1alpha1_DebuggingConfiguration_To_config_DebuggingConfiguration(&in.DebuggingConfiguration, &out.DebuggingConfiguration, s); err != nil {
		return err
	}
	out.DisablePreemption = in.DisablePreemption
	out.PercentageOfNodesToScore = in.PercentageOfNodesToScore
	out.BindTimeoutSeconds = (*int64)(unsafe.Pointer(in.BindTimeoutSeconds))
	out.Plugins = (*config.Plugins)(unsafe.Pointer(in.Plugins))
	out.PluginConfig = *(*[]config.PluginConfig)(unsafe.Pointer(&in.PluginConfig))
	if err := configv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ResourceProviderClientConnection, &out.ResourceProviderClientConnection, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration(in *v1alpha1.KubeSchedulerConfiguration, out *config.KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_config_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *config.KubeSchedulerConfiguration, out *v1alpha1.KubeSchedulerConfiguration, s conversion.Scope) error {
	out.SchedulerName = in.SchedulerName
	if err := Convert_config_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(&in.AlgorithmSource, &out.AlgorithmSource, s); err != nil {
		return err
	}
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	if err := Convert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := configv1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.HealthzBindAddress = in.HealthzBindAddress
	out.MetricsBindAddress = in.MetricsBindAddress
	if err := configv1alpha1.Convert_config_DebuggingConfiguration_To_v1alpha1_DebuggingConfiguration(&in.DebuggingConfiguration, &out.DebuggingConfiguration, s); err != nil {
		return err
	}
	out.DisablePreemption = in.DisablePreemption
	out.PercentageOfNodesToScore = in.PercentageOfNodesToScore
	out.BindTimeoutSeconds = (*int64)(unsafe.Pointer(in.BindTimeoutSeconds))
	out.Plugins = (*v1alpha1.Plugins)(unsafe.Pointer(in.Plugins))
	out.PluginConfig = *(*[]v1alpha1.PluginConfig)(unsafe.Pointer(&in.PluginConfig))
	if err := configv1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ResourceProviderClientConnection, &out.ResourceProviderClientConnection, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration is an autogenerated conversion function.
func Convert_config_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *config.KubeSchedulerConfiguration, out *v1alpha1.KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(in *v1alpha1.KubeSchedulerLeaderElectionConfiguration, out *config.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	if err := configv1alpha1.Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&in.LeaderElectionConfiguration, &out.LeaderElectionConfiguration, s); err != nil {
		return err
	}
	out.LockObjectNamespace = in.LockObjectNamespace
	out.LockObjectName = in.LockObjectName
	return nil
}

// Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(in *v1alpha1.KubeSchedulerLeaderElectionConfiguration, out *config.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(in, out, s)
}

func autoConvert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in *config.KubeSchedulerLeaderElectionConfiguration, out *v1alpha1.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	if err := configv1alpha1.Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElectionConfiguration, &out.LeaderElectionConfiguration, s); err != nil {
		return err
	}
	out.LockObjectNamespace = in.LockObjectNamespace
	out.LockObjectName = in.LockObjectName
	return nil
}

// Convert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration is an autogenerated conversion function.
func Convert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in *config.KubeSchedulerLeaderElectionConfiguration, out *v1alpha1.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_Plugin_To_config_Plugin(in *v1alpha1.Plugin, out *config.Plugin, s conversion.Scope) error {
	out.Name = in.Name
	out.Weight = in.Weight
	return nil
}

// Convert_v1alpha1_Plugin_To_config_Plugin is an autogenerated conversion function.
func Convert_v1alpha1_Plugin_To_config_Plugin(in *v1alpha1.Plugin, out *config.Plugin, s conversion.Scope) error {
	return autoConvert_v1alpha1_Plugin_To_config_Plugin(in, out, s)
}

func autoConvert_config_Plugin_To_v1alpha1_Plugin(in *config.Plugin, out *v1alpha1.Plugin, s conversion.Scope) error {
	out.Name = in.Name
	out.Weight = in.Weight
	return nil
}

// Convert_config_Plugin_To_v1alpha1_Plugin is an autogenerated conversion function.
func Convert_config_Plugin_To_v1alpha1_Plugin(in *config.Plugin, out *v1alpha1.Plugin, s conversion.Scope) error {
	return autoConvert_config_Plugin_To_v1alpha1_Plugin(in, out, s)
}

func autoConvert_v1alpha1_PluginConfig_To_config_PluginConfig(in *v1alpha1.PluginConfig, out *config.PluginConfig, s conversion.Scope) error {
	out.Name = in.Name
	out.Args = in.Args
	return nil
}

// Convert_v1alpha1_PluginConfig_To_config_PluginConfig is an autogenerated conversion function.
func Convert_v1alpha1_PluginConfig_To_config_PluginConfig(in *v1alpha1.PluginConfig, out *config.PluginConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_PluginConfig_To_config_PluginConfig(in, out, s)
}

func autoConvert_config_PluginConfig_To_v1alpha1_PluginConfig(in *config.PluginConfig, out *v1alpha1.PluginConfig, s conversion.Scope) error {
	out.Name = in.Name
	out.Args = in.Args
	return nil
}

// Convert_config_PluginConfig_To_v1alpha1_PluginConfig is an autogenerated conversion function.
func Convert_config_PluginConfig_To_v1alpha1_PluginConfig(in *config.PluginConfig, out *v1alpha1.PluginConfig, s conversion.Scope) error {
	return autoConvert_config_PluginConfig_To_v1alpha1_PluginConfig(in, out, s)
}

func autoConvert_v1alpha1_PluginSet_To_config_PluginSet(in *v1alpha1.PluginSet, out *config.PluginSet, s conversion.Scope) error {
	out.Enabled = *(*[]config.Plugin)(unsafe.Pointer(&in.Enabled))
	out.Disabled = *(*[]config.Plugin)(unsafe.Pointer(&in.Disabled))
	return nil
}

// Convert_v1alpha1_PluginSet_To_config_PluginSet is an autogenerated conversion function.
func Convert_v1alpha1_PluginSet_To_config_PluginSet(in *v1alpha1.PluginSet, out *config.PluginSet, s conversion.Scope) error {
	return autoConvert_v1alpha1_PluginSet_To_config_PluginSet(in, out, s)
}

func autoConvert_config_PluginSet_To_v1alpha1_PluginSet(in *config.PluginSet, out *v1alpha1.PluginSet, s conversion.Scope) error {
	out.Enabled = *(*[]v1alpha1.Plugin)(unsafe.Pointer(&in.Enabled))
	out.Disabled = *(*[]v1alpha1.Plugin)(unsafe.Pointer(&in.Disabled))
	return nil
}

// Convert_config_PluginSet_To_v1alpha1_PluginSet is an autogenerated conversion function.
func Convert_config_PluginSet_To_v1alpha1_PluginSet(in *config.PluginSet, out *v1alpha1.PluginSet, s conversion.Scope) error {
	return autoConvert_config_PluginSet_To_v1alpha1_PluginSet(in, out, s)
}

func autoConvert_v1alpha1_Plugins_To_config_Plugins(in *v1alpha1.Plugins, out *config.Plugins, s conversion.Scope) error {
	out.QueueSort = (*config.PluginSet)(unsafe.Pointer(in.QueueSort))
	out.PreFilter = (*config.PluginSet)(unsafe.Pointer(in.PreFilter))
	out.Filter = (*config.PluginSet)(unsafe.Pointer(in.Filter))
	out.PostFilter = (*config.PluginSet)(unsafe.Pointer(in.PostFilter))
	out.Score = (*config.PluginSet)(unsafe.Pointer(in.Score))
	out.NormalizeScore = (*config.PluginSet)(unsafe.Pointer(in.NormalizeScore))
	out.Reserve = (*config.PluginSet)(unsafe.Pointer(in.Reserve))
	out.Permit = (*config.PluginSet)(unsafe.Pointer(in.Permit))
	out.PreBind = (*config.PluginSet)(unsafe.Pointer(in.PreBind))
	out.Bind = (*config.PluginSet)(unsafe.Pointer(in.Bind))
	out.PostBind = (*config.PluginSet)(unsafe.Pointer(in.PostBind))
	out.Unreserve = (*config.PluginSet)(unsafe.Pointer(in.Unreserve))
	return nil
}

// Convert_v1alpha1_Plugins_To_config_Plugins is an autogenerated conversion function.
func Convert_v1alpha1_Plugins_To_config_Plugins(in *v1alpha1.Plugins, out *config.Plugins, s conversion.Scope) error {
	return autoConvert_v1alpha1_Plugins_To_config_Plugins(in, out, s)
}

func autoConvert_config_Plugins_To_v1alpha1_Plugins(in *config.Plugins, out *v1alpha1.Plugins, s conversion.Scope) error {
	out.QueueSort = (*v1alpha1.PluginSet)(unsafe.Pointer(in.QueueSort))
	out.PreFilter = (*v1alpha1.PluginSet)(unsafe.Pointer(in.PreFilter))
	out.Filter = (*v1alpha1.PluginSet)(unsafe.Pointer(in.Filter))
	out.PostFilter = (*v1alpha1.PluginSet)(unsafe.Pointer(in.PostFilter))
	out.Score = (*v1alpha1.PluginSet)(unsafe.Pointer(in.Score))
	out.NormalizeScore = (*v1alpha1.PluginSet)(unsafe.Pointer(in.NormalizeScore))
	out.Reserve = (*v1alpha1.PluginSet)(unsafe.Pointer(in.Reserve))
	out.Permit = (*v1alpha1.PluginSet)(unsafe.Pointer(in.Permit))
	out.PreBind = (*v1alpha1.PluginSet)(unsafe.Pointer(in.PreBind))
	out.Bind = (*v1alpha1.PluginSet)(unsafe.Pointer(in.Bind))
	out.PostBind = (*v1alpha1.PluginSet)(unsafe.Pointer(in.PostBind))
	out.Unreserve = (*v1alpha1.PluginSet)(unsafe.Pointer(in.Unreserve))
	return nil
}

// Convert_config_Plugins_To_v1alpha1_Plugins is an autogenerated conversion function.
func Convert_config_Plugins_To_v1alpha1_Plugins(in *config.Plugins, out *v1alpha1.Plugins, s conversion.Scope) error {
	return autoConvert_config_Plugins_To_v1alpha1_Plugins(in, out, s)
}

func autoConvert_v1alpha1_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(in *v1alpha1.SchedulerAlgorithmSource, out *config.SchedulerAlgorithmSource, s conversion.Scope) error {
	out.Policy = (*config.SchedulerPolicySource)(unsafe.Pointer(in.Policy))
	out.Provider = (*string)(unsafe.Pointer(in.Provider))
	return nil
}

// Convert_v1alpha1_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(in *v1alpha1.SchedulerAlgorithmSource, out *config.SchedulerAlgorithmSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(in, out, s)
}

func autoConvert_config_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(in *config.SchedulerAlgorithmSource, out *v1alpha1.SchedulerAlgorithmSource, s conversion.Scope) error {
	out.Policy = (*v1alpha1.SchedulerPolicySource)(unsafe.Pointer(in.Policy))
	out.Provider = (*string)(unsafe.Pointer(in.Provider))
	return nil
}

// Convert_config_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource is an autogenerated conversion function.
func Convert_config_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(in *config.SchedulerAlgorithmSource, out *v1alpha1.SchedulerAlgorithmSource, s conversion.Scope) error {
	return autoConvert_config_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(in, out, s)
}

func autoConvert_v1alpha1_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource(in *v1alpha1.SchedulerPolicyConfigMapSource, out *config.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource(in *v1alpha1.SchedulerPolicyConfigMapSource, out *config.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource(in, out, s)
}

func autoConvert_config_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(in *config.SchedulerPolicyConfigMapSource, out *v1alpha1.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_config_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource is an autogenerated conversion function.
func Convert_config_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(in *config.SchedulerPolicyConfigMapSource, out *v1alpha1.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	return autoConvert_config_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(in, out, s)
}

func autoConvert_v1alpha1_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource(in *v1alpha1.SchedulerPolicyFileSource, out *config.SchedulerPolicyFileSource, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_v1alpha1_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource(in *v1alpha1.SchedulerPolicyFileSource, out *config.SchedulerPolicyFileSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource(in, out, s)
}

func autoConvert_config_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(in *config.SchedulerPolicyFileSource, out *v1alpha1.SchedulerPolicyFileSource, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_config_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource is an autogenerated conversion function.
func Convert_config_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(in *config.SchedulerPolicyFileSource, out *v1alpha1.SchedulerPolicyFileSource, s conversion.Scope) error {
	return autoConvert_config_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(in, out, s)
}

func autoConvert_v1alpha1_SchedulerPolicySource_To_config_SchedulerPolicySource(in *v1alpha1.SchedulerPolicySource, out *config.SchedulerPolicySource, s conversion.Scope) error {
	out.File = (*config.SchedulerPolicyFileSource)(unsafe.Pointer(in.File))
	out.ConfigMap = (*config.SchedulerPolicyConfigMapSource)(unsafe.Pointer(in.ConfigMap))
	return nil
}

// Convert_v1alpha1_SchedulerPolicySource_To_config_SchedulerPolicySource is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerPolicySource_To_config_SchedulerPolicySource(in *v1alpha1.SchedulerPolicySource, out *config.SchedulerPolicySource, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerPolicySource_To_config_SchedulerPolicySource(in, out, s)
}

func autoConvert_config_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(in *config.SchedulerPolicySource, out *v1alpha1.SchedulerPolicySource, s conversion.Scope) error {
	out.File = (*v1alpha1.SchedulerPolicyFileSource)(unsafe.Pointer(in.File))
	out.ConfigMap = (*v1alpha1.SchedulerPolicyConfigMapSource)(unsafe.Pointer(in.ConfigMap))
	return nil
}

// Convert_config_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource is an autogenerated conversion function.
func Convert_config_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(in *config.SchedulerPolicySource, out *v1alpha1.SchedulerPolicySource, s conversion.Scope) error {
	return autoConvert_config_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(in, out, s)
}
