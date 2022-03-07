package modelschemas

type LabelItemSchema struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (in *LabelItemSchema) DeepCopyInto(out *LabelItemSchema) {
	out.Key = in.Key
	out.Value = in.Value
}

type LabelItemsSchema []LabelItemSchema
