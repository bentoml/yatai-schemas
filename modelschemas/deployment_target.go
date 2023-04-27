package modelschemas

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	corev1 "k8s.io/api/core/v1"

	servingv2alpha1 "github.com/bentoml/yatai-deployment/apis/serving/v2alpha1"
	resourcesv1alpha1 "github.com/bentoml/yatai-image-builder/apis/resources/v1alpha1"
)

type DeploymentTargetType string

const (
	DeploymentTargetTypeStable DeploymentTargetType = "stable"
	DeploymentTargetTypeCanary DeploymentTargetType = "canary"
)

var DeploymentTargetTypeAddrs = map[DeploymentTargetType]string{
	DeploymentTargetTypeStable: "stb",
	DeploymentTargetTypeCanary: "cnr",
}

type DeploymentTargetResourceItem struct {
	CPU    string            `json:"cpu,omitempty"`
	Memory string            `json:"memory,omitempty"`
	GPU    string            `json:"gpu,omitempty"`
	Custom map[string]string `json:"custom,omitempty"`
}

func (in *DeploymentTargetResourceItem) DeepCopyInto(out *DeploymentTargetResourceItem) {
	*out = *in
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

type DeploymentTargetResources struct {
	Requests *DeploymentTargetResourceItem `json:"requests,omitempty"`
	Limits   *DeploymentTargetResourceItem `json:"limits,omitempty"`
}

func (in *DeploymentTargetResources) DeepCopy() (out *DeploymentTargetResources) {
	if in == nil {
		return nil
	}
	out = new(DeploymentTargetResources)
	in.DeepCopyInto(out)
	return
}

func (in *DeploymentTargetResources) DeepCopyInto(out *DeploymentTargetResources) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(DeploymentTargetResourceItem)
		(*in).DeepCopyInto(*out)
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(DeploymentTargetResourceItem)
		(*in).DeepCopyInto(*out)
	}
}

type HPAMetricType string

const (
	HPAMetricTypeCPU   HPAMetricType = "cpu"
	HPAMetricTypeGPU   HPAMetricType = "gpu"
	HPAMetricTypeQueue HPAMetricType = "queue"
)

type HPAMetric struct {
	Type  HPAMetricType `json:"type"`
	Value int64         `json:"value"`
}

type HPAScaleBehavior string

const (
	HPAScaleBehaviorDisabled HPAScaleBehavior = "disabled"
	HPAScaleBehaviorStable   HPAScaleBehavior = "stable"
	HPAScaleBehaviorFast     HPAScaleBehavior = "fast"
)

type HPAPolicy struct {
	Metrics           []HPAMetric      `json:"metrics"`
	ScaleDownBehavior HPAScaleBehavior `json:"scale_down_behavior"`
	ScaleUpBehavior   HPAScaleBehavior `json:"scale_up_behavior"`
}

type DeploymentTargetHPAConf struct {
	CPU         *int32     `json:"cpu,omitempty"`
	GPU         *int32     `json:"gpu,omitempty"`
	Memory      *string    `json:"memory,omitempty"`
	QPS         *int64     `json:"qps,omitempty"`
	MinReplicas *int32     `json:"min_replicas,omitempty"`
	MaxReplicas *int32     `json:"max_replicas,omitempty"`
	Policy      *HPAPolicy `json:"policy,omitempty"`
}

func (in *DeploymentTargetHPAConf) DeepCopy() (out *DeploymentTargetHPAConf) {
	if in == nil {
		return nil
	}
	out = new(DeploymentTargetHPAConf)
	in.DeepCopyInto(out)
	return
}

func (in *DeploymentTargetHPAConf) DeepCopyInto(out *DeploymentTargetHPAConf) {
	*out = *in
	if in.CPU != nil {
		out.CPU = new(int32)
		*out.CPU = *in.CPU
	}
	if in.GPU != nil {
		out.GPU = new(int32)
		*out.GPU = *in.GPU
	}
	if in.Memory != nil {
		out.Memory = new(string)
		*out.Memory = *in.Memory
	}
	if in.QPS != nil {
		out.QPS = new(int64)
		*out.QPS = *in.QPS
	}
	if in.MinReplicas != nil {
		out.MinReplicas = new(int32)
		*out.MinReplicas = *in.MinReplicas
	}
	if in.MaxReplicas != nil {
		out.MaxReplicas = new(int32)
		*out.MaxReplicas = *in.MaxReplicas
	}
}

type BentoRequestOverrides struct {
	ImageBuildTimeout *time.Duration `json:"imageBuildTimeout,omitempty"`

	ImageBuilderExtraPodMetadata   *resourcesv1alpha1.ExtraPodMetadata `json:"imageBuilderExtraPodMetadata,omitempty"`
	ImageBuilderExtraPodSpec       *resourcesv1alpha1.ExtraPodSpec     `json:"imageBuilderExtraPodSpec,omitempty"`
	ImageBuilderExtraContainerEnv  []corev1.EnvVar                     `json:"imageBuilderExtraContainerEnv,omitempty"`
	ImageBuilderContainerResources *corev1.ResourceRequirements        `json:"imageBuilderContainerResources,omitempty"`

	DockerConfigJSONSecretName string `json:"dockerConfigJsonSecretName,omitempty"`

	DownloaderContainerEnvFrom []corev1.EnvFromSource `json:"downloaderContainerEnvFrom,omitempty"`
}

type ApiServerBentoDeploymentOverrides struct {
	MonitorExporter  *servingv2alpha1.MonitorExporterSpec `json:"monitorExporter,omitempty"`
	ExtraPodMetadata *servingv2alpha1.ExtraPodMetadata    `json:"extraPodMetadata,omitempty"`
	ExtraPodSpec     *servingv2alpha1.ExtraPodSpec        `json:"extraPodSpec,omitempty"`
}

type RunnerBentoDeploymentOverrides struct {
	ExtraPodMetadata *servingv2alpha1.ExtraPodMetadata `json:"extraPodMetadata,omitempty"`
	ExtraPodSpec     *servingv2alpha1.ExtraPodSpec     `json:"extraPodSpec,omitempty"`
}

type DeploymentTargetRunnerConfig struct {
	Resources                              *DeploymentTargetResources      `json:"resources,omitempty"`
	HPAConf                                *DeploymentTargetHPAConf        `json:"hpa_conf,omitempty"`
	Envs                                   *[]*LabelItemSchema             `json:"envs,omitempty"`
	EnableStealingTrafficDebugMode         *bool                           `json:"enable_stealing_traffic_debug_mode,omitempty"`
	EnableDebugMode                        *bool                           `json:"enable_debug_mode,omitempty"`
	EnableDebugPodReceiveProductionTraffic *bool                           `json:"enable_debug_pod_receive_production_traffic,omitempty"`
	DeploymentStrategy                     *DeploymentStrategy             `json:"deployment_strategy,omitempty"`
	BentoDeploymentOverrides               *RunnerBentoDeploymentOverrides `json:"bento_deployment_overrides,omitempty"`
}

func (in *DeploymentTargetRunnerConfig) DeepCopy() (out *DeploymentTargetRunnerConfig) {
	if in == nil {
		return nil
	}
	out = new(DeploymentTargetRunnerConfig)
	in.DeepCopyInto(out)
	return
}

func (in *DeploymentTargetRunnerConfig) DeepCopyInto(out *DeploymentTargetRunnerConfig) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(DeploymentTargetResources)
		(*in).DeepCopyInto(*out)
	}
	if in.HPAConf != nil {
		in, out := &in.HPAConf, &out.HPAConf
		*out = new(DeploymentTargetHPAConf)
		(*in).DeepCopyInto(*out)
	}
	if in.Envs != nil {
		out.Envs = new([]*LabelItemSchema)
		for _, item := range *in.Envs {
			newItem := new(LabelItemSchema)
			item.DeepCopyInto(newItem)
			*out.Envs = append(*out.Envs, newItem)
		}
	}
}

type DeploymentStrategy string

const (
	DeploymentStrategyRollingUpdate               DeploymentStrategy = "RollingUpdate"
	DeploymentStrategyRecreate                    DeploymentStrategy = "Recreate"
	DeploymentStrategyRampedSlowRollout           DeploymentStrategy = "RampedSlowRollout"
	DeploymentStrategyBestEffortControlledRollout DeploymentStrategy = "BestEffortControlledRollout"
)

type RequestQueueConfig struct {
	Enabled               *bool  `json:"enabled,omitempty"`
	MaxConsumeConcurrency *int32 `json:"max_consume_concurrency,omitempty"`
}

type DeploymentTargetConfig struct {
	KubeResourceUid                        string                                  `json:"kubeResourceUid"`
	KubeResourceVersion                    string                                  `json:"kubeResourceVersion"`
	Resources                              *DeploymentTargetResources              `json:"resources"`
	HPAConf                                *DeploymentTargetHPAConf                `json:"hpa_conf,omitempty"`
	Envs                                   *[]*LabelItemSchema                     `json:"envs,omitempty"`
	Runners                                map[string]DeploymentTargetRunnerConfig `json:"runners,omitempty"`
	EnableIngress                          *bool                                   `json:"enable_ingress,omitempty"`
	EnableStealingTrafficDebugMode         *bool                                   `json:"enable_stealing_traffic_debug_mode,omitempty"`
	EnableDebugMode                        *bool                                   `json:"enable_debug_mode,omitempty"`
	EnableDebugPodReceiveProductionTraffic *bool                                   `json:"enable_debug_pod_receive_production_traffic,omitempty"`
	DeploymentStrategy                     *DeploymentStrategy                     `json:"deployment_strategy,omitempty"`
	BentoDeploymentOverrides               *ApiServerBentoDeploymentOverrides      `json:"bento_deployment_overrides,omitempty"`
	BentoRequestOverrides                  *BentoRequestOverrides                  `json:"bento_request_overrides,omitempty"`
	RequestQueue                           *RequestQueueConfig                     `json:"request_queue,omitempty"`
}

func (in *DeploymentTargetConfig) DeepCopy() (out *DeploymentTargetConfig) {
	if in == nil {
		return nil
	}
	out = new(DeploymentTargetConfig)
	in.DeepCopyInto(out)
	return
}

func (in *DeploymentTargetConfig) DeepCopyInto(out *DeploymentTargetConfig) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(DeploymentTargetResources)
		(*in).DeepCopyInto(*out)
	}
	if in.HPAConf != nil {
		in, out := &in.HPAConf, &out.HPAConf
		*out = new(DeploymentTargetHPAConf)
		(*in).DeepCopyInto(*out)
	}
	if in.Envs != nil {
		out.Envs = new([]*LabelItemSchema)
		for _, item := range *in.Envs {
			newItem := new(LabelItemSchema)
			item.DeepCopyInto(newItem)
			*out.Envs = append(*out.Envs, newItem)
		}
	}
}

func (c *DeploymentTargetConfig) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	return json.Unmarshal([]byte(value.(string)), c)
}

func (c *DeploymentTargetConfig) Value() (driver.Value, error) {
	if c == nil {
		return nil, nil
	}
	return json.Marshal(c)
}
