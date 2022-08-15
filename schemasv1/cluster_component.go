package schemasv1

import (
	"time"

	"github.com/bentoml/yatai-schemas/modelschemas"
)

type ClusterComponentSchema struct {
	ResourceSchema
	Creator           *UserSchema                                  `json:"creator"`
	Cluster           *ClusterFullSchema                           `json:"cluster"`
	Description       string                                       `json:"description"`
	Version           string                                       `json:"version"`
	KubeNamespace     string                                       `json:"kube_namespace"`
	Manifest          *modelschemas.ClusterComponentManifestSchema `json:"manifest"`
	LatestInstalledAt *time.Time                                   `json:"latest_installed_at"`
	LatestHeartbeatAt *time.Time                                   `json:"latest_heartbeat_at"`
}

type RegisterClusterComponentSchema struct {
	Name           string            `json:"name"`
	Version        string            `json:"version"`
	KubeNamespace  string            `json:"kube_namespace"`
	SelectorLabels map[string]string `json:"selector_labels,omitempty"`
}
