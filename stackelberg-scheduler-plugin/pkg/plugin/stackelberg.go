package plugin

import (
	"context"
	"fmt"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"stackelberg-scheduler-plugin/pkg/client"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const Name = "StackelbergPrefilter"

type StackelbergPlugin struct {
	handle framework.Handle
	client *stackelberg.Client
}

func New(obj runtime.Object, h framework.Handle) (framework.Plugin, error) {
	client := stackelberg.NewClient("http://localhost:5000") // adjust if needed
	return &StackelbergPlugin{
		handle: h,
		client: client,
	}, nil
}


func (sp *StackelbergPlugin) Name() string {
	return Name
}

func (sp *StackelbergPlugin) PreFilter(ctx context.Context, cycleState *framework.CycleState, pod *v1.Pod) (*framework.PreFilterResult, *framework.Status) {
	params := map[string]interface{}{
		"web_app_budget": 800,
		"web_app_min_cpu": 2,
		"web_app_min_memory": 8,
	}

	resp, err := sp.client.Allocate(110, 300, params)
	if err != nil {
		return nil, framework.NewStatus(framework.Error, fmt.Sprintf("Stackelberg call failed: %v", err))
	}

	fmt.Printf("Stackelberg PreFilter - Allocations: %+v\n", resp.Allocations)

	// You can store results in CycleState if needed for later phases
	// cycleState.Write("stackelberg-results", resp)

	return nil, framework.NewStatus(framework.Success, "PreFilter passed")
}