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
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
	GPU    string `json:"gpu"`
}

func (in *DeploymentTargetResourceItem) DeepCopyInto(out *DeploymentTargetResourceItem) {
	*out = *in
	out.CPU = in.CPU
	out.Memory = in.Memory
	out.GPU = in.GPU
}

type DeploymentTargetResources struct {
	Requests *DeploymentTargetResourceItem `json:"requests"`
	Limits   *DeploymentTargetResourceItem `json:"limits"`
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

type DeploymentTargetConfig struct {
	KubeResourceUid string                     `json:"kubeResourceUid"`
	Resources       *DeploymentTargetResources `json:"resources"`
	HPAConf         *DeploymentTargetHPAConf   `json:"hpa_conf,omitempty"`
	Envs            *[]*LabelItemSchema        `json:"envs,omitempty"`
}

func (in *DeploymentTargetConfig) DeepCopyInto(out *DeploymentTargetConfig) {
	*out = *in
	if in.Resources != nil {
		out.Resources = new(DeploymentTargetResources)
		(*in).Resources.DeepCopyInto(out.Resources)
	}
	if in.HPAConf != nil {
		out.HPAConf = new(DeploymentTargetHPAConf)
		(*in).HPAConf.DeepCopyInto(out.HPAConf)
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
