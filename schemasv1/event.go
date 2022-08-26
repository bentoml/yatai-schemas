package schemasv1

import (
	"time"

	"github.com/bentoml/yatai-schemas/modelschemas"
)

type EventSchema struct {
	BaseSchema
	Resource        IResourceSchema          `json:"resource,omitempty"`
	Name            string                   `json:"name,omitempty"`
	Status          modelschemas.EventStatus `json:"status,omitempty"`
	OperationName   string                   `json:"operation_name,omitempty"`
	ApiTokenName    string                   `json:"api_token_name,omitempty"`
	Creator         *UserSchema              `json:"creator,omitempty"`
	CreatedAt       time.Time                `json:"created_at,omitempty"`
	ResourceDeleted bool                     `json:"resource_deleted,omitempty"`
}

type EventListSchema struct {
	BaseListSchema
	Items []*EventSchema `json:"items"`
}
