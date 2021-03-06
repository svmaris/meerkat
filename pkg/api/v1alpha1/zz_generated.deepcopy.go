// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnCertificateConfig) DeepCopyInto(out *OvpnCertificateConfig) {
	*out = *in
	out.Validity = in.Validity
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnCertificateConfig.
func (in *OvpnCertificateConfig) DeepCopy() *OvpnCertificateConfig {
	if in == nil {
		return nil
	}
	out := new(OvpnCertificateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnClient) DeepCopyInto(out *OvpnClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnClient.
func (in *OvpnClient) DeepCopy() *OvpnClient {
	if in == nil {
		return nil
	}
	out := new(OvpnClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OvpnClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnClientCertificate) DeepCopyInto(out *OvpnClientCertificate) {
	*out = *in
	out.OvpnCertificateConfig = in.OvpnCertificateConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnClientCertificate.
func (in *OvpnClientCertificate) DeepCopy() *OvpnClientCertificate {
	if in == nil {
		return nil
	}
	out := new(OvpnClientCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnClientCertificateConfig) DeepCopyInto(out *OvpnClientCertificateConfig) {
	*out = *in
	out.OvpnCertificateConfig = in.OvpnCertificateConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnClientCertificateConfig.
func (in *OvpnClientCertificateConfig) DeepCopy() *OvpnClientCertificateConfig {
	if in == nil {
		return nil
	}
	out := new(OvpnClientCertificateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnClientList) DeepCopyInto(out *OvpnClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OvpnClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnClientList.
func (in *OvpnClientList) DeepCopy() *OvpnClientList {
	if in == nil {
		return nil
	}
	out := new(OvpnClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OvpnClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnClientSpec) DeepCopyInto(out *OvpnClientSpec) {
	*out = *in
	out.Certificate = in.Certificate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnClientSpec.
func (in *OvpnClientSpec) DeepCopy() *OvpnClientSpec {
	if in == nil {
		return nil
	}
	out := new(OvpnClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnClientStatus) DeepCopyInto(out *OvpnClientStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnClientStatus.
func (in *OvpnClientStatus) DeepCopy() *OvpnClientStatus {
	if in == nil {
		return nil
	}
	out := new(OvpnClientStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnPKICertificateConfig) DeepCopyInto(out *OvpnPKICertificateConfig) {
	*out = *in
	out.OvpnCertificateConfig = in.OvpnCertificateConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnPKICertificateConfig.
func (in *OvpnPKICertificateConfig) DeepCopy() *OvpnPKICertificateConfig {
	if in == nil {
		return nil
	}
	out := new(OvpnPKICertificateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnPkiConfig) DeepCopyInto(out *OvpnPkiConfig) {
	*out = *in
	out.OvpnPKICertificateConfig = in.OvpnPKICertificateConfig
	out.DN = in.DN
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnPkiConfig.
func (in *OvpnPkiConfig) DeepCopy() *OvpnPkiConfig {
	if in == nil {
		return nil
	}
	out := new(OvpnPkiConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnPkiDnConfig) DeepCopyInto(out *OvpnPkiDnConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnPkiDnConfig.
func (in *OvpnPkiDnConfig) DeepCopy() *OvpnPkiDnConfig {
	if in == nil {
		return nil
	}
	out := new(OvpnPkiDnConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnSecurityConfig) DeepCopyInto(out *OvpnSecurityConfig) {
	*out = *in
	out.PKI = in.PKI
	out.Server = in.Server
	out.Clients = in.Clients
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnSecurityConfig.
func (in *OvpnSecurityConfig) DeepCopy() *OvpnSecurityConfig {
	if in == nil {
		return nil
	}
	out := new(OvpnSecurityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnServer) DeepCopyInto(out *OvpnServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnServer.
func (in *OvpnServer) DeepCopy() *OvpnServer {
	if in == nil {
		return nil
	}
	out := new(OvpnServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OvpnServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnServerAddress) DeepCopyInto(out *OvpnServerAddress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnServerAddress.
func (in *OvpnServerAddress) DeepCopy() *OvpnServerAddress {
	if in == nil {
		return nil
	}
	out := new(OvpnServerAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnServerCertificateConfig) DeepCopyInto(out *OvpnServerCertificateConfig) {
	*out = *in
	out.OvpnCertificateConfig = in.OvpnCertificateConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnServerCertificateConfig.
func (in *OvpnServerCertificateConfig) DeepCopy() *OvpnServerCertificateConfig {
	if in == nil {
		return nil
	}
	out := new(OvpnServerCertificateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnServerDeployment) DeepCopyInto(out *OvpnServerDeployment) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnServerDeployment.
func (in *OvpnServerDeployment) DeepCopy() *OvpnServerDeployment {
	if in == nil {
		return nil
	}
	out := new(OvpnServerDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnServerList) DeepCopyInto(out *OvpnServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OvpnServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnServerList.
func (in *OvpnServerList) DeepCopy() *OvpnServerList {
	if in == nil {
		return nil
	}
	out := new(OvpnServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OvpnServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnServerSecrets) DeepCopyInto(out *OvpnServerSecrets) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnServerSecrets.
func (in *OvpnServerSecrets) DeepCopy() *OvpnServerSecrets {
	if in == nil {
		return nil
	}
	out := new(OvpnServerSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnServerService) DeepCopyInto(out *OvpnServerService) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnServerService.
func (in *OvpnServerService) DeepCopy() *OvpnServerService {
	if in == nil {
		return nil
	}
	out := new(OvpnServerService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnServerSpec) DeepCopyInto(out *OvpnServerSpec) {
	*out = *in
	out.Network = in.Network
	in.Traffic.DeepCopyInto(&out.Traffic)
	out.Security = in.Security
	out.Secrets = in.Secrets
	in.Deployment.DeepCopyInto(&out.Deployment)
	in.Service.DeepCopyInto(&out.Service)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnServerSpec.
func (in *OvpnServerSpec) DeepCopy() *OvpnServerSpec {
	if in == nil {
		return nil
	}
	out := new(OvpnServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnServerStatus) DeepCopyInto(out *OvpnServerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnServerStatus.
func (in *OvpnServerStatus) DeepCopy() *OvpnServerStatus {
	if in == nil {
		return nil
	}
	out := new(OvpnServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvpnTrafficConfig) DeepCopyInto(out *OvpnTrafficConfig) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]SubnetMask, len(*in))
		copy(*out, *in)
	}
	if in.Nameservers != nil {
		in, out := &in.Nameservers, &out.Nameservers
		*out = make([]IPv4Address, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvpnTrafficConfig.
func (in *OvpnTrafficConfig) DeepCopy() *OvpnTrafficConfig {
	if in == nil {
		return nil
	}
	out := new(OvpnTrafficConfig)
	in.DeepCopyInto(out)
	return out
}
