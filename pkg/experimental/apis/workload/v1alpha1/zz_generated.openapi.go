// +build !ignore_autogenerated

/*
Copyright 2022 The Kubernetes Authors.

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"k8s.io/ingress-gce/pkg/experimental/apis/workload/v1alpha1.Condition":               schema_experimental_apis_workload_v1alpha1_Condition(ref),
		"k8s.io/ingress-gce/pkg/experimental/apis/workload/v1alpha1.ExternalWorkloadAddress": schema_experimental_apis_workload_v1alpha1_ExternalWorkloadAddress(ref),
		"k8s.io/ingress-gce/pkg/experimental/apis/workload/v1alpha1.Workload":                schema_experimental_apis_workload_v1alpha1_Workload(ref),
		"k8s.io/ingress-gce/pkg/experimental/apis/workload/v1alpha1.WorkloadSpec":            schema_experimental_apis_workload_v1alpha1_WorkloadSpec(ref),
		"k8s.io/ingress-gce/pkg/experimental/apis/workload/v1alpha1.WorkloadStatus":          schema_experimental_apis_workload_v1alpha1_WorkloadStatus(ref),
	}
}

func schema_experimental_apis_workload_v1alpha1_Condition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of condition in CamelCase or in foo.example.com/CamelCase. Many .condition.type values are consistent across resources like Available, but because arbitrary conditions can be useful (see .node.status.conditions), the ability to deconflict is important.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the condition, one of True, False, Unknown.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"observedGeneration": {
						SchemaProps: spec.SchemaProps{
							Description: "If set, this represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.condition[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Last time the condition transitioned from one status to another. This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "The reason for the condition's last transition in CamelCase. The specific API may choose whether or not this field is considered a guaranteed API. This field may not be empty.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human readable message indicating details about the transition. This field may be empty.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "status", "lastTransitionTime", "reason", "message"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_experimental_apis_workload_v1alpha1_ExternalWorkloadAddress(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExternalWorkloadAddress represents an address used by an ExternalWorkload",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"address": {
						SchemaProps: spec.SchemaProps{
							Description: "Address is the address of the workload exposed to the cluster.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"addressType": {
						SchemaProps: spec.SchemaProps{
							Description: "AddressType specifies the address type of the external workload.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"address", "addressType"},
			},
		},
	}
}

func schema_experimental_apis_workload_v1alpha1_Workload(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/ingress-gce/pkg/experimental/apis/workload/v1alpha1.WorkloadSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/ingress-gce/pkg/experimental/apis/workload/v1alpha1.WorkloadStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "k8s.io/ingress-gce/pkg/experimental/apis/workload/v1alpha1.WorkloadSpec", "k8s.io/ingress-gce/pkg/experimental/apis/workload/v1alpha1.WorkloadStatus"},
	}
}

func schema_experimental_apis_workload_v1alpha1_WorkloadSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WorkloadSpec is the spec for a Workload resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"enableHeartbeat": {
						SchemaProps: spec.SchemaProps{
							Description: "EnableHeartbeat indicates whether Heartbeat condition is enabled on this ExternalWorkload.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"enablePing": {
						SchemaProps: spec.SchemaProps{
							Description: "EnablePing indicates whether Ping condition is enabled on this ExternalWorkload.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"hostname": {
						SchemaProps: spec.SchemaProps{
							Description: "Hostname is the hostname of this workload.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"addresses": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Addresses specifies the addresses that can be used to access the workload from the cluster.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/ingress-gce/pkg/experimental/apis/workload/v1alpha1.ExternalWorkloadAddress"),
									},
								},
							},
						},
					},
				},
				Required: []string{"enableHeartbeat", "enablePing", "addresses"},
			},
		},
		Dependencies: []string{
			"k8s.io/ingress-gce/pkg/experimental/apis/workload/v1alpha1.ExternalWorkloadAddress"},
	}
}

func schema_experimental_apis_workload_v1alpha1_WorkloadStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WorkloadStatus is the status for a Workload resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"type",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/ingress-gce/pkg/experimental/apis/workload/v1alpha1.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/ingress-gce/pkg/experimental/apis/workload/v1alpha1.Condition"},
	}
}
