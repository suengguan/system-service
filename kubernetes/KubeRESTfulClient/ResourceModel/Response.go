package ResourceModel

type ResponseMetaData struct {
	SelfLink        string `json:"selfLink"`
	ResourceVersion string `json:"resourceVersion"`
}

// pod
type Response struct {
	Kind       string           `json:"kind"`
	ApiVersion string           `json:"apiVersion"`
	MetaData   ResponseMetaData `json:"metadata"`
	Items      []*Pod           `json:"items"`
}

// node
type ResponseNode struct {
	Kind       string           `json:"kind"`
	ApiVersion string           `json:"apiVersion"`
	MetaData   ResponseMetaData `json:"metadata"`
	Items      []*Node          `json:"items"`
}
