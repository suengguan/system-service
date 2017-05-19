package ResourceModel

type NameSpace struct {
	ApiVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   map[string]string `json:"metadata"`
}

func (n *NameSpace) SetApiVersion(version string) {
	n.ApiVersion = version
}

func (n *NameSpace) GetApiVersion() string {
	return n.ApiVersion
}

func (n *NameSpace) SetKind(kind string) {
	n.Kind = kind
}

func (n *NameSpace) GetKind() string {
	return n.Kind
}

func (n *NameSpace) SetName(name string) {
	if len(n.Metadata) == 0 {
		n.Metadata = make(map[string]string)
	}
	n.Metadata["name"] = name
}

func (n *NameSpace) GetName() string {
	return n.Metadata["name"]
}
