/*
Copyright 2023.

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

package v1alpha1

import (
	"github.com/keptn/lifecycle-toolkit/metrics-operator/api/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConvertTo converts this KeptnMetric to the Hub version (v1alpha2)
func (src *KeptnMetric) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.KeptnMetric)

	// Spec
	dst.Spec = v1alpha2.KeptnMetricSpec{
		Provider: v1alpha2.ProviderRef{
			Name: src.Spec.Provider.Name,
		},
		Query:                src.Spec.Query,
		FetchIntervalSeconds: src.Spec.FetchIntervalSeconds,
	}

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Status
	dst.Status = v1alpha2.KeptnMetricStatus{
		Value:       src.Status.Value,
		RawValue:    src.Status.RawValue,
		LastUpdated: src.Status.LastUpdated,
	}

	return nil
}

// ConvertFrom converts from the Hub version (v1alpha2) to this version.
func (dst *KeptnMetric) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.KeptnMetric)

	// Spec
	dst.Spec = KeptnMetricSpec{
		Provider: ProviderRef{
			Name: src.Spec.Provider.Name,
		},
		Query:                src.Spec.Query,
		FetchIntervalSeconds: src.Spec.FetchIntervalSeconds,
	}

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Status
	dst.Status = KeptnMetricStatus{
		Value:       src.Status.Value,
		RawValue:    src.Status.RawValue,
		LastUpdated: src.Status.LastUpdated,
	}

	return nil
}

// KeptnMetricSpec defines the desired state of KeptnMetric
type KeptnMetricSpec struct {
	// Provider represents the provider object
	Provider ProviderRef `json:"provider"`
	// Query represents the query to be run
	Query string `json:"query"`
	// FetchIntervalSeconds represents the update frequency in seconds that is used to update the metric
	FetchIntervalSeconds uint `json:"fetchIntervalSeconds"`
}

// KeptnMetricStatus defines the observed state of KeptnMetric
type KeptnMetricStatus struct {
	// Value represents the resulting value
	Value string `json:"value"`
	// RawValue represents the resulting value in raw format
	RawValue []byte `json:"rawValue"`
	// LastUpdated represents the time when the status data was last updated
	LastUpdated metav1.Time `json:"lastUpdated"`
}

// ProviderRef represents the provider object
type ProviderRef struct {
	// Name of the provider
	Name string `json:"name"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Provider",type=string,JSONPath=`.spec.provider.name`
//+kubebuilder:printcolumn:name="Query",type=string,JSONPath=`.spec.query`
//+kubebuilder:printcolumn:name="Value",type=string,JSONPath=`.status.value`

// KeptnMetric is the Schema for the keptnmetrics API
type KeptnMetric struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeptnMetricSpec   `json:"spec,omitempty"`
	Status KeptnMetricStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KeptnMetricList contains a list of KeptnMetric
type KeptnMetricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeptnMetric `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeptnMetric{}, &KeptnMetricList{})
}

func (s *KeptnMetric) IsStatusSet() bool {
	return s.Status.Value != ""
}
