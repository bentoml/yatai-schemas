package modelschemas

import (
	"database/sql/driver"
	"encoding/json"
)

type ClusterComponentManifestSchema struct {
	SelectorLabels map[string]string `json:"selector_labels,omitempty"`
}

func (c *ClusterComponentManifestSchema) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), c)
}

func (c *ClusterComponentManifestSchema) Value() (driver.Value, error) {
	if c == nil {
		return nil, nil
	}
	return json.Marshal(c)
}
