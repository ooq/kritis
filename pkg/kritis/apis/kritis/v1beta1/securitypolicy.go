/*
Copyright 2018 Google LLC

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

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ImageSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ImageSecurityPolicySpec `json:"spec"`
}

// ImageSecurityPolicySpec is the spec for a ImageSecurityPolicy resource
type ImageSecurityPolicySpec struct {
	ImageAllowlist                   []string                         `json:"imageAllowlist"`
	PackageVulnerabilityRequirements PackageVulnerabilityRequirements `json:"packageVulnerabilityRequirements"`
	AttestorNames                    []string                         `json:"attestorNames"`
}

// PackageVulnerabilityRequirements is the requirements for package vulnz for an ImageSecurityPolicy
type PackageVulnerabilityRequirements struct {
	// CVE's with fixes.
	MaximumSeverity string `json:"maximumSeverity"`
	// CVE's without fixes.
	MaximumFixUnavailableSeverity string   `json:"maximumFixNotAvailableSeverity"`
	AllowlistCVEs                 []string `json:"allowlistCVEs"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ImageSecurityPolicyList is a list of ImageSecurityPolicy resources
type ImageSecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ImageSecurityPolicy `json:"items"`
}
