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

type RouteObservation struct {

	// an identifier for the resource with format projects/{{project}}/global/routes/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// URL to a Network that should handle matching packets.
	// URL to a Network that should handle matching packets.
	NextHopNetwork *string `json:"nextHopNetwork,omitempty" tf:"next_hop_network,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type RouteParameters struct {

	// An optional description of this resource. Provide this property
	// when you create the resource.
	// An optional description of this resource. Provide this property
	// when you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The destination range of outgoing packets that this route applies to.
	// Only IPv4 is supported.
	// The destination range of outgoing packets that this route applies to.
	// Only IPv4 is supported.
	// +kubebuilder:validation:Required
	DestRange *string `json:"destRange" tf:"dest_range,omitempty"`

	// The network that this route applies to.
	// The network that this route applies to.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.Network
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// URL to a gateway that should handle matching packets.
	// Currently, you can only specify the internet gateway, using a full or
	// partial valid URL:
	// URL to a gateway that should handle matching packets.
	// Currently, you can only specify the internet gateway, using a full or
	// partial valid URL:
	// * 'https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway'
	// * 'projects/project/global/gateways/default-internet-gateway'
	// * 'global/gateways/default-internet-gateway'
	// * The string 'default-internet-gateway'.
	// +kubebuilder:validation:Optional
	NextHopGateway *string `json:"nextHopGateway,omitempty" tf:"next_hop_gateway,omitempty"`

	// Network IP address of an instance that should handle matching packets.
	// Network IP address of an instance that should handle matching packets.
	// +kubebuilder:validation:Optional
	NextHopIP *string `json:"nextHopIp,omitempty" tf:"next_hop_ip,omitempty"`

	// The IP address or URL to a forwarding rule of type
	// loadBalancingScheme=INTERNAL that should handle matching
	// packets.
	// With the GA provider you can only specify the forwarding
	// rule as a partial or full URL. For example, the following
	// are all valid values:
	// The IP address or URL to a forwarding rule of type
	// loadBalancingScheme=INTERNAL that should handle matching
	// packets.
	//
	// With the GA provider you can only specify the forwarding
	// rule as a partial or full URL. For example, the following
	// are all valid values:
	// * 10.128.0.56
	// * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	// * regions/region/forwardingRules/forwardingRule
	//
	// When the beta provider, you can also specify the IP address
	// of a forwarding rule from the same VPC or any peered VPC.
	//
	// Note that this can only be used when the destinationRange is
	// a public (non-RFC 1918) IP CIDR range.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.ForwardingRule
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NextHopIlb *string `json:"nextHopIlb,omitempty" tf:"next_hop_ilb,omitempty"`

	// +kubebuilder:validation:Optional
	NextHopIlbRef *v1.Reference `json:"nextHopIlbRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NextHopIlbSelector *v1.Selector `json:"nextHopIlbSelector,omitempty" tf:"-"`

	// URL to an instance that should handle matching packets.
	// You can specify this as a full or partial URL. For example:
	// URL to an instance that should handle matching packets.
	// You can specify this as a full or partial URL. For example:
	// * 'https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance'
	// * 'projects/project/zones/zone/instances/instance'
	// * 'zones/zone/instances/instance'
	// * Just the instance name, with the zone in 'next_hop_instance_zone'.
	// +kubebuilder:validation:Optional
	NextHopInstance *string `json:"nextHopInstance,omitempty" tf:"next_hop_instance,omitempty"`

	// The zone of the instance specified in next_hop_instance. Omit if next_hop_instance is specified as a URL.
	// +kubebuilder:validation:Optional
	NextHopInstanceZone *string `json:"nextHopInstanceZone,omitempty" tf:"next_hop_instance_zone,omitempty"`

	// URL to a VpnTunnel that should handle matching packets.
	// URL to a VpnTunnel that should handle matching packets.
	// +kubebuilder:validation:Optional
	NextHopVPNTunnel *string `json:"nextHopVpnTunnel,omitempty" tf:"next_hop_vpn_tunnel,omitempty"`

	// The priority of this route. Priority is used to break ties in cases
	// where there is more than one matching route of equal prefix length.
	// In the case of two routes with equal prefix length, the one with the
	// lowest-numbered priority value wins.
	// Default value is 1000. Valid range is 0 through 65535.
	// The priority of this route. Priority is used to break ties in cases
	// where there is more than one matching route of equal prefix length.
	//
	// In the case of two routes with equal prefix length, the one with the
	// lowest-numbered priority value wins.
	//
	// Default value is 1000. Valid range is 0 through 65535.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A list of instance tags to which this route applies.
	// A list of instance tags to which this route applies.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteParameters `json:"forProvider"`
}

// RouteStatus defines the observed state of Route.
type RouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route is the Schema for the Routes API. Represents a Route resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteSpec   `json:"spec"`
	Status            RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Repository type metadata.
var (
	Route_Kind             = "Route"
	Route_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Route_Kind}.String()
	Route_KindAPIVersion   = Route_Kind + "." + CRDGroupVersion.String()
	Route_GroupVersionKind = CRDGroupVersion.WithKind(Route_Kind)
)

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
