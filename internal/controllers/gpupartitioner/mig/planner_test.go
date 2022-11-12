package mig_test

import (
	"context"
	"fmt"
	partitioner_mig "github.com/nebuly-ai/nebulnetes/internal/controllers/gpupartitioner/mig"
	"github.com/nebuly-ai/nebulnetes/internal/controllers/gpupartitioner/state"
	"github.com/nebuly-ai/nebulnetes/pkg/api/n8s.nebuly.ai/v1alpha1"
	"github.com/nebuly-ai/nebulnetes/pkg/constant"
	"github.com/nebuly-ai/nebulnetes/pkg/gpu/mig"
	"github.com/nebuly-ai/nebulnetes/pkg/test/factory"
	scheduler_mock "github.com/nebuly-ai/nebulnetes/pkg/test/mocks/scheduler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"strconv"
	"testing"
)

func TestPlanner__Plan(t *testing.T) {
	testCases := []struct {
		name                     string
		snapshotNodes            []v1.Node
		candidatePods            []v1.Pod
		schedulerPreFilterStatus *framework.Status
		schedulerFilterStatus    *framework.Status

		expectedOverallPartitioning []state.GPUPartitioning
		expectedErr                 bool
	}{
		{
			name:                     "Empty snapshot, no candidates",
			snapshotNodes:            make([]v1.Node, 0),
			candidatePods:            make([]v1.Pod, 0),
			schedulerPreFilterStatus: framework.NewStatus(framework.Success),
			schedulerFilterStatus:    framework.NewStatus(framework.Success),

			expectedOverallPartitioning: make([]state.GPUPartitioning, 0),
			expectedErr:                 false,
		},
		{
			name:          "Empty snapshot, many candidates",
			snapshotNodes: make([]v1.Node, 0),
			candidatePods: []v1.Pod{
				factory.BuildPod("ns-1", "pd-1").Get(),
				factory.BuildPod("ns-2", "pd-2").Get(),
			},
			schedulerPreFilterStatus: framework.NewStatus(framework.Success),
			schedulerFilterStatus:    framework.NewStatus(framework.Success),

			expectedOverallPartitioning: make([]state.GPUPartitioning, 0),
			expectedErr:                 false,
		},
		{
			name: "Cluster geometry cannot be changed for pending Pods",
			snapshotNodes: []v1.Node{
				factory.BuildNode("node-1").
					WithAnnotations(map[string]string{
						fmt.Sprintf(v1alpha1.AnnotationUsedMigStatusFormat, 0, mig.Profile4g20gb): "1", // node provides required MIG resource, but it's used
					}).
					WithLabels(map[string]string{
						constant.LabelNvidiaProduct: string(mig.GPUModel_A30),
					}).
					WithAllocatableResources(v1.ResourceList{
						mig.Profile4g24gb.AsResourceName(): *resource.NewQuantity(1, resource.DecimalSI),
					}).
					Get(),
				factory.BuildNode("node-2").
					WithAnnotations(map[string]string{
						fmt.Sprintf(v1alpha1.AnnotationFreeMigStatusFormat, 0, mig.Profile1g5gb): "1",
					}).
					WithLabels(map[string]string{
						constant.LabelNvidiaProduct: string(mig.GPUModel_A30),
					}).
					WithAllocatableResources(v1.ResourceList{
						mig.Profile1g5gb.AsResourceName(): *resource.NewQuantity(1, resource.DecimalSI),
					}).
					Get(),
			},
			candidatePods: []v1.Pod{
				factory.BuildPod("ns-1", "pd-1").Get(), // not requesting any MIG resource
				factory.BuildPod("ns-1", "pd-1").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile4g20gb.AsResourceName(), 1).
						WithCPUMilliRequest(100).
						Get(),
				).Get(),
			},
			schedulerPreFilterStatus: framework.NewStatus(framework.Success),
			schedulerFilterStatus:    framework.NewStatus(framework.Success),
			expectedOverallPartitioning: []state.GPUPartitioning{
				{
					Resources: map[v1.ResourceName]int{
						mig.Profile4g20gb.AsResourceName(): 1,
					},
				},
				{
					Resources: map[v1.ResourceName]int{
						mig.Profile1g5gb.AsResourceName(): 1,
					},
				},
			},
			expectedErr: false,
		},
		{
			name: "Cluster geometry can be changed, but pod scheduling PreFilter fails",
			snapshotNodes: []v1.Node{
				factory.BuildNode("node-1").
					WithAnnotations(map[string]string{
						fmt.Sprintf(v1alpha1.AnnotationFreeMigStatusFormat, 0, mig.Profile4g24gb): "1",
					}).
					WithLabels(map[string]string{
						constant.LabelNvidiaProduct: string(mig.GPUModel_A30),
					}).
					WithAllocatableResources(v1.ResourceList{
						mig.Profile4g24gb.AsResourceName(): *resource.NewQuantity(1, resource.DecimalSI),
					}).
					Get(),
				factory.BuildNode("node-2").
					WithLabels(map[string]string{
						constant.LabelNvidiaProduct: string(mig.GPUModel_A30),
					}).
					Get(),
			},
			candidatePods: []v1.Pod{
				factory.BuildPod("ns-1", "pd-2").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile1g6gb.AsResourceName(), 1).
						Get(),
				).Get(),
				factory.BuildPod("ns-1", "pd-1").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile1g6gb.AsResourceName(), 1).
						WithCPUMilliRequest(100).
						Get(),
				).Get(),
				factory.BuildPod("ns-2", "pd-1").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile2g12gb.AsResourceName(), 1).
						Get(),
				).Get(),
			},
			schedulerPreFilterStatus: framework.NewStatus(framework.Error),
			schedulerFilterStatus:    framework.NewStatus(framework.Success),
			expectedOverallPartitioning: []state.GPUPartitioning{
				{
					GPUIndex: 0,
					Resources: map[v1.ResourceName]int{
						mig.Profile4g24gb.AsResourceName(): 1,
					},
				},
			},
			expectedErr: false,
		},
		{
			name: "Cluster geometry can be changed, but pod scheduling Filter fails",
			snapshotNodes: []v1.Node{
				factory.BuildNode("node-1").
					WithAnnotations(map[string]string{
						fmt.Sprintf(v1alpha1.AnnotationFreeMigStatusFormat, 0, mig.Profile4g24gb): "1",
					}).
					WithLabels(map[string]string{
						constant.LabelNvidiaProduct: string(mig.GPUModel_A30),
					}).
					WithAllocatableResources(v1.ResourceList{
						mig.Profile4g24gb.AsResourceName(): *resource.NewQuantity(1, resource.DecimalSI),
					}).
					Get(),
				factory.BuildNode("node-2").
					WithLabels(map[string]string{
						constant.LabelNvidiaProduct: string(mig.GPUModel_A30),
					}).
					Get(),
			},
			candidatePods: []v1.Pod{
				factory.BuildPod("ns-1", "pd-2").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile1g6gb.AsResourceName(), 1).
						Get(),
				).Get(),
				factory.BuildPod("ns-1", "pd-1").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile1g6gb.AsResourceName(), 1).
						WithCPUMilliRequest(100).
						Get(),
				).Get(),
				factory.BuildPod("ns-2", "pd-1").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile2g12gb.AsResourceName(), 1).
						Get(),
				).Get(),
			},
			schedulerPreFilterStatus: framework.NewStatus(framework.Success),
			schedulerFilterStatus:    framework.NewStatus(framework.Error),
			expectedOverallPartitioning: []state.GPUPartitioning{
				{
					Resources: map[v1.ResourceName]int{
						mig.Profile4g24gb.AsResourceName(): 1,
					},
				},
			},
			expectedErr: false,
		},
		{
			name: "Cluster geometry gets changed by splitting up MIG profile and creating new ones from spare MIG capacity",
			snapshotNodes: []v1.Node{
				factory.BuildNode("node-1").
					WithAnnotations(map[string]string{
						fmt.Sprintf(v1alpha1.AnnotationFreeMigStatusFormat, 0, mig.Profile4g24gb): "1",
					}).
					WithLabels(map[string]string{
						constant.LabelNvidiaProduct: string(mig.GPUModel_A30),
						constant.LabelNvidiaCount:   strconv.Itoa(1),
					}).
					WithAllocatableResources(v1.ResourceList{
						mig.Profile4g24gb.AsResourceName(): *resource.NewQuantity(1, resource.DecimalSI),
					}).
					Get(),
				factory.BuildNode("node-2").
					WithLabels(map[string]string{
						constant.LabelNvidiaProduct: string(mig.GPUModel_A30),
						constant.LabelNvidiaCount:   strconv.Itoa(1),
					}).
					WithAllocatableResources(v1.ResourceList{
						constant.ResourceNvidiaGPU: *resource.NewQuantity(1, resource.DecimalSI),
					}).
					Get(),
			},
			candidatePods: []v1.Pod{
				factory.BuildPod("ns-1", "pd-2").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile1g6gb.AsResourceName(), 1).
						Get(),
				).Get(),
				factory.BuildPod("ns-1", "pd-1").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile1g6gb.AsResourceName(), 1).
						Get(),
				).Get(),
				factory.BuildPod("ns-2", "pd-1").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile1g6gb.AsResourceName(), 1).
						Get(),
				).Get(),
				factory.BuildPod("ns-2", "pd-2").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile1g6gb.AsResourceName(), 1).
						Get(),
				).Get(),
				factory.BuildPod("ns-2", "pd-3").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile1g6gb.AsResourceName(), 1).
						Get(),
				).Get(),
			},
			schedulerPreFilterStatus: framework.NewStatus(framework.Success),
			schedulerFilterStatus:    framework.NewStatus(framework.Success),
			expectedOverallPartitioning: []state.GPUPartitioning{
				{
					Resources: map[v1.ResourceName]int{
						mig.Profile1g6gb.AsResourceName():  2,
						mig.Profile2g12gb.AsResourceName(): 1,
					},
				},
				{
					Resources: map[v1.ResourceName]int{
						mig.Profile1g6gb.AsResourceName(): 4,
					},
				},
			},
			expectedErr: false,
		},
		{
			name: "Cluster geometry gets updated by grouping small unused MIG profiles into a larger one",
			snapshotNodes: []v1.Node{
				factory.BuildNode("node-1").
					WithAnnotations(map[string]string{
						fmt.Sprintf(v1alpha1.AnnotationFreeMigStatusFormat, 0, mig.Profile1g6gb): "4",
						fmt.Sprintf(v1alpha1.AnnotationFreeMigStatusFormat, 1, mig.Profile1g6gb): "4",
					}).
					WithLabels(map[string]string{
						constant.LabelNvidiaProduct: string(mig.GPUModel_A30),
						constant.LabelNvidiaCount:   strconv.Itoa(2),
					}).
					WithAllocatableResources(v1.ResourceList{
						mig.Profile1g6gb.AsResourceName(): *resource.NewQuantity(8, resource.DecimalSI),
					}).
					Get(),
				factory.BuildNode("node-2").
					WithAnnotations(map[string]string{
						fmt.Sprintf(v1alpha1.AnnotationFreeMigStatusFormat, 0, mig.Profile1g5gb):  "5",
						fmt.Sprintf(v1alpha1.AnnotationFreeMigStatusFormat, 0, mig.Profile2g10gb): "1",
					}).
					WithLabels(map[string]string{
						constant.LabelNvidiaProduct: string(mig.GPUModel_A100_SMX4_40GB),
						constant.LabelNvidiaCount:   strconv.Itoa(1),
					}).
					WithAllocatableResources(v1.ResourceList{
						mig.Profile1g5gb.AsResourceName():  *resource.NewQuantity(5, resource.DecimalSI),
						mig.Profile2g10gb.AsResourceName(): *resource.NewQuantity(1, resource.DecimalSI),
					}).
					Get(),
			},
			candidatePods: []v1.Pod{
				factory.BuildPod("ns-1", "pd-2").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile3g20gb.AsResourceName(), 1).
						Get(),
				).Get(),
				factory.BuildPod("ns-1", "pd-1").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile4g24gb.AsResourceName(), 1).
						Get(),
				).Get(),
				factory.BuildPod("ns-1", "pd-1").WithContainer(
					factory.BuildContainer("test", "test").
						WithScalarResourceRequest(mig.Profile2g12gb.AsResourceName(), 1).
						Get(),
				).Get(),
			},
			schedulerPreFilterStatus: framework.NewStatus(framework.Success),
			schedulerFilterStatus:    framework.NewStatus(framework.Success),
			expectedOverallPartitioning: []state.GPUPartitioning{
				{
					Resources: map[v1.ResourceName]int{
						mig.Profile3g20gb.AsResourceName(): 2,
					},
				},
				{
					Resources: map[v1.ResourceName]int{
						mig.Profile4g24gb.AsResourceName(): 1,
					},
				},
				{
					Resources: map[v1.ResourceName]int{
						mig.Profile2g12gb.AsResourceName(): 2,
					},
				},
			},
			expectedErr: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockedScheduler := scheduler_mock.NewFramework(t)
			mockedScheduler.On(
				"RunPreFilterPlugins",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).Return(nil, tt.schedulerPreFilterStatus).Maybe()
			mockedScheduler.On(
				"RunFilterPlugins",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).Return(framework.PluginToStatus{"": tt.schedulerFilterStatus}).Maybe()

			planner := partitioner_mig.NewPlanner(mockedScheduler)
			snapshot := newSnapshotFromNodes(tt.snapshotNodes)
			plan, err := planner.Plan(context.Background(), snapshot, tt.candidatePods)

			// Compute overall partitioning ignoring GPU index
			overallGpuPartitioning := make([]state.GPUPartitioning, 0)
			for _, nodePartitioning := range plan {
				for _, g := range nodePartitioning.GPUs {
					g.GPUIndex = 0
					overallGpuPartitioning = append(overallGpuPartitioning, g)
				}
			}
			for i := range tt.expectedOverallPartitioning {
				gpuPartitioning := &tt.expectedOverallPartitioning[i]
				gpuPartitioning.GPUIndex = 0
			}

			// Run assertions
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tt.expectedOverallPartitioning, overallGpuPartitioning)
			}
		})
	}
}

func newSnapshotFromNodes(nodes []v1.Node) state.ClusterSnapshot {
	nodeInfos := make(map[string]framework.NodeInfo)
	for _, node := range nodes {
		n := node
		ni := framework.NewNodeInfo()
		ni.Requested = framework.NewResource(v1.ResourceList{})
		ni.Allocatable = framework.NewResource(v1.ResourceList{})
		ni.SetNode(&n)
		nodeInfos[n.Name] = *ni
	}
	snapshot := state.NewClusterSnapshot(nodeInfos)
	return snapshot
}