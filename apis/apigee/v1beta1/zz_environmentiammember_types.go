/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type EnvironmentIAMMemberObservation struct {
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	EnvID *string `json:"envId,omitempty" tf:"env_id,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type EnvironmentIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	EnvID *string `json:"envId,omitempty" tf:"env_id,omitempty"`

	// +kubebuilder:validation:Optional
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// EnvironmentIAMMemberSpec defines the desired state of EnvironmentIAMMember
type EnvironmentIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentIAMMemberParameters `json:"forProvider"`
}

// EnvironmentIAMMemberStatus defines the observed state of EnvironmentIAMMember.
type EnvironmentIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentIAMMember is the Schema for the EnvironmentIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type EnvironmentIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.envId)",message="envId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.member)",message="member is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.orgId)",message="orgId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.role)",message="role is a required parameter"
	Spec   EnvironmentIAMMemberSpec   `json:"spec"`
	Status EnvironmentIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentIAMMemberList contains a list of EnvironmentIAMMembers
type EnvironmentIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvironmentIAMMember `json:"items"`
}

// Repository type metadata.
var (
	EnvironmentIAMMember_Kind             = "EnvironmentIAMMember"
	EnvironmentIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EnvironmentIAMMember_Kind}.String()
	EnvironmentIAMMember_KindAPIVersion   = EnvironmentIAMMember_Kind + "." + CRDGroupVersion.String()
	EnvironmentIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(EnvironmentIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&EnvironmentIAMMember{}, &EnvironmentIAMMemberList{})
}
