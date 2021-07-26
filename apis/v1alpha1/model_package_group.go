// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ModelPackageGroupSpec defines the desired state of ModelPackageGroup.
//
// A group of versioned models in the model registry.
type ModelPackageGroupSpec struct {
	// A description for the model group.
	ModelPackageGroupDescription *string `json:"modelPackageGroupDescription,omitempty"`
	// The name of the model group.
	// +kubebuilder:validation:Required
	ModelPackageGroupName *string `json:"modelPackageGroupName"`
	// A list of key value pairs associated with the model group. For more information,
	// see Tagging AWS resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// in the AWS General Reference Guide.
	Tags []*Tag `json:"tags,omitempty"`
}

// ModelPackageGroupStatus defines the observed state of ModelPackageGroup
type ModelPackageGroupStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The status of the model group.
	ModelPackageGroupStatus *string `json:"modelPackageGroupStatus,omitempty"`
}

// ModelPackageGroup is the Schema for the ModelPackageGroups API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type=string,priority=0,JSONPath=`.status.modelPackageGroupStatus`
type ModelPackageGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ModelPackageGroupSpec   `json:"spec,omitempty"`
	Status            ModelPackageGroupStatus `json:"status,omitempty"`
}

// ModelPackageGroupList contains a list of ModelPackageGroup
// +kubebuilder:object:root=true
type ModelPackageGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ModelPackageGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ModelPackageGroup{}, &ModelPackageGroupList{})
}
