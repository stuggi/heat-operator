// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openstack-k8s-operators/heat-operator/pkg/apis/heat/v1.Heat":       schema_pkg_apis_heat_v1_Heat(ref),
		"github.com/openstack-k8s-operators/heat-operator/pkg/apis/heat/v1.HeatSpec":   schema_pkg_apis_heat_v1_HeatSpec(ref),
		"github.com/openstack-k8s-operators/heat-operator/pkg/apis/heat/v1.HeatStatus": schema_pkg_apis_heat_v1_HeatStatus(ref),
	}
}

func schema_pkg_apis_heat_v1_Heat(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Heat is the Schema for the heats API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openstack-k8s-operators/heat-operator/pkg/apis/heat/v1.HeatSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openstack-k8s-operators/heat-operator/pkg/apis/heat/v1.HeatStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openstack-k8s-operators/heat-operator/pkg/apis/heat/v1.HeatSpec", "github.com/openstack-k8s-operators/heat-operator/pkg/apis/heat/v1.HeatStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_heat_v1_HeatSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HeatSpec defines the desired state of Heat",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_heat_v1_HeatStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HeatStatus defines the observed state of Heat",
				Type:        []string{"object"},
			},
		},
	}
}