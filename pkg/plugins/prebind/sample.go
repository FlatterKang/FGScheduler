package prebind

import (
	"context"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const Name = "GPU PreBinder"

type GPUFilterPlugin struct {
	handle framework.Handle
}

func (g *GPUFilterPlugin) PreBind(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) *framework.Status {
	klog.V(2).Infof("filter pod: %v, node: %v", p.GetName(), nodeName)
	return framework.NewStatus(framework.Success)
}

func New(_ runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	return &GPUFilterPlugin{handle}, nil
}

func (g *GPUFilterPlugin) Name() string {
	return Name
}
