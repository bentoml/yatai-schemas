package modelschemas

import (
	"database/sql/driver"
	"encoding/json"
)

type YataiComponentName string

const (
	YataiComponentNameDeployment YataiComponentName = "deployment"
)

type YataiComponentManifestSchema struct {
	SelectorLabels map[string]string `json:"selector_labels,omitempty"`
}

func (c *YataiComponentManifestSchema) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), c)
}

func (c *YataiComponentManifestSchema) Value() (driver.Value, error) {
	if c == nil {
		return nil, nil
	}
	return json.Marshal(c)
}
