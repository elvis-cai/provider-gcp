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

type GarbageCollectionPolicyObservation struct {

	// The name of the column family.
	ColumnFamily *string `json:"columnFamily,omitempty" tf:"column_family,omitempty"`

	// The deletion policy for the GC policy.
	// Setting ABANDON allows the resource to be abandoned rather than deleted. This is useful for GC policy as it cannot be deleted in a replicated instance.
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// Serialized JSON object to represent a more complex GC policy. Conflicts with mode, max_age and max_version. Conflicts with mode, max_age and max_version.
	GcRules *string `json:"gcRules,omitempty" tf:"gc_rules,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Bigtable instance.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// GC policy that applies to all cells older than the given age.
	MaxAge []MaxAgeObservation `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// GC policy that applies to all versions of a cell except for the most recent.
	MaxVersion []MaxVersionObservation `json:"maxVersion,omitempty" tf:"max_version,omitempty"`

	// If multiple policies are set, you should choose between UNION OR INTERSECTION.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the table.
	Table *string `json:"table,omitempty" tf:"table,omitempty"`
}

type GarbageCollectionPolicyParameters struct {

	// The name of the column family.
	// +kubebuilder:validation:Optional
	ColumnFamily *string `json:"columnFamily,omitempty" tf:"column_family,omitempty"`

	// The deletion policy for the GC policy.
	// Setting ABANDON allows the resource to be abandoned rather than deleted. This is useful for GC policy as it cannot be deleted in a replicated instance.
	// +kubebuilder:validation:Optional
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// Serialized JSON object to represent a more complex GC policy. Conflicts with mode, max_age and max_version. Conflicts with mode, max_age and max_version.
	// +kubebuilder:validation:Optional
	GcRules *string `json:"gcRules,omitempty" tf:"gc_rules,omitempty"`

	// The name of the Bigtable instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigtable/v1beta1.Instance
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Reference to a Instance in bigtable to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameRef *v1.Reference `json:"instanceNameRef,omitempty" tf:"-"`

	// Selector for a Instance in bigtable to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameSelector *v1.Selector `json:"instanceNameSelector,omitempty" tf:"-"`

	// GC policy that applies to all cells older than the given age.
	// +kubebuilder:validation:Optional
	MaxAge []MaxAgeParameters `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// GC policy that applies to all versions of a cell except for the most recent.
	// +kubebuilder:validation:Optional
	MaxVersion []MaxVersionParameters `json:"maxVersion,omitempty" tf:"max_version,omitempty"`

	// If multiple policies are set, you should choose between UNION OR INTERSECTION.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigtable/v1beta1.Table
	// +kubebuilder:validation:Optional
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// Reference to a Table in bigtable to populate table.
	// +kubebuilder:validation:Optional
	TableRef *v1.Reference `json:"tableRef,omitempty" tf:"-"`

	// Selector for a Table in bigtable to populate table.
	// +kubebuilder:validation:Optional
	TableSelector *v1.Selector `json:"tableSelector,omitempty" tf:"-"`
}

type MaxAgeObservation struct {

	// Number of days before applying GC policy.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Duration before applying GC policy (ex. "8h"). This is required when days isn't set
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`
}

type MaxAgeParameters struct {

	// Number of days before applying GC policy.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Duration before applying GC policy (ex. "8h"). This is required when days isn't set
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`
}

type MaxVersionObservation struct {

	// Number of version before applying the GC policy.
	Number *float64 `json:"number,omitempty" tf:"number,omitempty"`
}

type MaxVersionParameters struct {

	// Number of version before applying the GC policy.
	// +kubebuilder:validation:Required
	Number *float64 `json:"number" tf:"number,omitempty"`
}

// GarbageCollectionPolicySpec defines the desired state of GarbageCollectionPolicy
type GarbageCollectionPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GarbageCollectionPolicyParameters `json:"forProvider"`
}

// GarbageCollectionPolicyStatus defines the observed state of GarbageCollectionPolicy.
type GarbageCollectionPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GarbageCollectionPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GarbageCollectionPolicy is the Schema for the GarbageCollectionPolicys API. Creates a Google Cloud Bigtable GC Policy inside a family.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type GarbageCollectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.columnFamily)",message="columnFamily is a required parameter"
	Spec   GarbageCollectionPolicySpec   `json:"spec"`
	Status GarbageCollectionPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GarbageCollectionPolicyList contains a list of GarbageCollectionPolicys
type GarbageCollectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GarbageCollectionPolicy `json:"items"`
}

// Repository type metadata.
var (
	GarbageCollectionPolicy_Kind             = "GarbageCollectionPolicy"
	GarbageCollectionPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GarbageCollectionPolicy_Kind}.String()
	GarbageCollectionPolicy_KindAPIVersion   = GarbageCollectionPolicy_Kind + "." + CRDGroupVersion.String()
	GarbageCollectionPolicy_GroupVersionKind = CRDGroupVersion.WithKind(GarbageCollectionPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&GarbageCollectionPolicy{}, &GarbageCollectionPolicyList{})
}
