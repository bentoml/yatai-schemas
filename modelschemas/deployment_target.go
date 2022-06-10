package modelschemas

import (
	"database/sql/driver"
	"encoding/json"
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
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
	GPU    string `json:"gpu,omitempty"`
}

func (in *DeploymentTargetResourceItem) DeepCopyInto(out *DeploymentTargetResourceItem) {
	*out = *in
	out.CPU = in.CPU
	out.Memory = in.Memory
	out.GPU = in.GPU
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

type DeploymentTargetHPAConf struct {
	CPU         *int32  `json:"cpu,omitempty"`
	GPU         *int32  `json:"gpu,omitempty"`
	Memory      *string `json:"memory,omitempty"`
	QPS         *int64  `json:"qps,omitempty"`
	MinReplicas *int32  `json:"min_replicas,omitempty"`
	MaxReplicas *int32  `json:"max_replicas,omitempty"`
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

type DeploymentTargetRunnerConfig struct {
	Resources *DeploymentTargetResources `json:"resources,omitempty"`
	HPAConf   *DeploymentTargetHPAConf   `json:"hpa_conf,omitempty"`
	Envs      *[]*LabelItemSchema        `json:"envs,omitempty"`
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

type DeploymentTargetConfig struct {
	KubeResourceUid     string                                  `json:"kubeResourceUid"`
	KubeResourceVersion string                                  `json:"kubeResourceVersion"`
	Resources           *DeploymentTargetResources              `json:"resources"`
	HPAConf             *DeploymentTargetHPAConf                `json:"hpa_conf,omitempty"`
	Envs                *[]*LabelItemSchema                     `json:"envs,omitempty"`
	Runners             map[string]DeploymentTargetRunnerConfig `json:"runners,omitempty"`
	EnableIngress       *bool                                   `json:"enable_ingress,omitempty"`
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
