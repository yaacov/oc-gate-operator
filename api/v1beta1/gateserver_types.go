/*
Copyright 2021 Yaacov Zamir.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GateServerSpec defines the desired state of GateServer
type GateServerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// api-url is the k8s API url.
	// Defalut value is "https://kubernetes.default.svc".
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Type="string"
	// +kubebuilder:validation:Pattern="^(http|https)://.*"
	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:default:="https://kubernetes.default.svc"
	APIURL string `json:"api-url,omitempty"`

	// route is the the gate proxy server.
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type="string"
	// +kubebuilder:validation:Pattern="^([a-z0-9-_])+[.]([a-z0-9-_])+[.]([a-z0-9-._])+$"
	// +kubebuilder:validation:MaxLength=226
	Route string `json:"route,omitempty"`

	// admin-role is the verbs athorization role of the service (reader/admin)
	// if service is role is reader, clients getting tokens to use this service
	// will be able to excute get, watch and list verbs.
	// if service is role is admin, clients getting tokens to use this service
	// will be able to excute get, watch, list, patch, creat and delete verbs.
	// Defalut value is "reader".
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Type="string"
	// +kubebuilder:validation:Pattern="^(reader|admin)$"
	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:default:="reader"
	AdminRole string `json:"admin-role,omitempty"`

	// admin-resources is a comma separated list of resources athorization role of the service,
	// if left empty service could access any resource.
	// Defalut value is "".
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Type="string"
	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:default:=""
	AdminResources string `json:"admin-resources,omitempty"`

	// admin-namespaced determain if the athorization role of the service is namespaced.
	// Defalut value is false.
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Type="boolean"
	// +kubebuilder:default:=false
	AdminNamespaced bool `json:"admin-namespaced,omitempty"`

	// passthrough determain if the tokens acquired from OAuth2 server directly to k8s API.
	// Defalut value is false.
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Type="boolean"
	// +kubebuilder:default:=false
	PassThrough bool `json:"passthrough,omitempty"`

	// image is the oc gate proxy image to use.
	// Defalut value is "quay.io/yaacov/oc-gate:latest".
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Type="string"
	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:default:="quay.io/yaacov/oc-gate:latest"
	Image string `json:"image,omitempty"`

	// web-app-image is the oc gate proxy web application image to use,
	// It's an image including the static web application to be served together
	// with k8s API.
	// The static web application should be in the directory "/data/web/public/"
	// and it will be copied to the proxy servers "/web/public/" directory on pproxy
	// startup. If left empty, the proxies default web application will not be replaced.
	// Defalut value is "" (use default web application).
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Type="string"
	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:default:=""
	WebAppImage string `json:"web-app-image,omitempty"`

	// gnerate-secret determain if a secrete with public and private kes will be automatically
	// generated when the oc-gate server is created.
	// Defalut value is true.
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Type="boolean"
	// +kubebuilder:default:=true
	GenerateSecret bool `json:"gnerate-secret,omitempty"`
}

// GateServerStatus defines the observed state of GateServer
type GateServerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Conditions represent the latest available observations of an object's state
	Conditions []metav1.Condition `json:"conditions"`

	// Token generation phase (ready|error)
	Phase string `json:"phase"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// GateServer is the Schema for the gateservers API
type GateServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GateServerSpec   `json:"spec,omitempty"`
	Status GateServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GateServerList contains a list of GateServer
type GateServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GateServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GateServer{}, &GateServerList{})
}
