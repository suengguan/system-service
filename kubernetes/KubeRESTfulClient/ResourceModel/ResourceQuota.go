package ResourceModel

type ResourceQuotaSpecHard struct {
	Cpu                    string `json:"cpu"`
	Memory                 string `json:"memory"`
	PersistentVolumeClaims string `json:"persistentvolumeclaims"`
	Pods                   string `json:"pods"`
	ReplicationControllers string `json:"replicationcontrollers"`
	ResourceQuotas         string `json:"resourcequotas"`
	Secrets                string `json:"secrets"`
	Services               string `json:"services"`
}

type ResourceQuotaSpec struct {
	Hard *ResourceQuotaSpecHard `json:"hard"`
}

type ResourceQuota struct {
	ApiVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	MetaData   ResourceMetaData  `json:"metadata"`
	Spec       ResourceQuotaSpec `json:"spec"`
}

func (r *ResourceQuota) SetApiVersion(version string) {
	r.ApiVersion = version
}

func (r *ResourceQuota) GetApiVersion() string {
	return r.ApiVersion
}

func (r *ResourceQuota) SetKind(kind string) {
	r.Kind = kind
}

func (r *ResourceQuota) GetKind() string {
	return r.Kind
}

func (r *ResourceQuota) SetName(name string) {
	r.MetaData.Name = name
}

func (r *ResourceQuota) GetName() string {
	return r.MetaData.Name
}

func (r *ResourceQuota) SetNameSpace(namespace string) {
	r.MetaData.NameSpace = namespace
}

func (r *ResourceQuota) GetNameSpace() string {
	return r.MetaData.NameSpace
}

func (r *ResourceQuota) SetSpecHard(hard *ResourceQuotaSpecHard) {
	r.Spec.Hard = hard
}

func (r *ResourceQuota) GetSpecHard() *ResourceQuotaSpecHard {
	return r.Spec.Hard
}
