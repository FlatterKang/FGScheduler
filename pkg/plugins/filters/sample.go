package filters

import (
	"context"
	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const Name = "GPU Filter"

type GPUFilterPlugin struct {
	handle framework.Handle
}

func New(_ runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	return &GPUFilterPlugin{handle}, nil
}

func (g *GPUFilterPlugin) Name() string {
	return Name
}

func (g *GPUFilterPlugin) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.V(2).Infof("filter pod: %v, node: %v", pod.GetName(), nodeInfo.Node().GetName())
	return framework.NewStatus(framework.Success, "")
}
