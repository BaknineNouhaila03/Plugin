package plugin

import (
	 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime"
)

const (
    // Plugin name
    StackelbergPluginName = "StackelbergScheduler"
    
    // Annotations for tenant identification
    TenantLabelKey        = "tenant"
    TenantBudgetAnnotation = "stackelberg.scheduler/budget"
    TenantPreferencesAnnotation = "stackelberg.scheduler/preferences"
    
    // Default values
    DefaultBudget = 100.0
    DefaultCPUPreference = 1.0
    DefaultMemoryPreference = 1.0
)

type StackelbergArgs struct {
    metav1.TypeMeta `json:",inline"`
    
    APIEndpoint string `json:"apiEndpoint,omitempty"`
}

func (s *StackelbergArgs) DeepCopyObject() runtime.Object {
    if s == nil {
        return nil
    }
    return &StackelbergArgs{
        TypeMeta:    s.TypeMeta,
        APIEndpoint: s.APIEndpoint,
    }
}

