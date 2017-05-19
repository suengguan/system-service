package ResourceModel

//============================================================================== container
type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Port struct {
	Name          string `json:"name"`
	ContainerPort int    `json:"containerPort"`
	HostPort      int    `json:"hostPort"`
	Protocol      string `json:"protocol"`
}

type VolumeMount struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	ReadOnly  bool   `json:"readOnly"`
}

type ResourceLimit struct {
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
}

type Resource struct {
	Limits *ResourceLimit `json:"limits"`
}

type SecurityContext struct {
	Privileged bool `json:"privileged"`
}

type Container struct {
	Name            string           `json:"name"`
	Image           string           `json:"image"`
	ImagePullPolicy string           `json:"imagePullPolicy"`
	Command         []string         `json:"command"`
	WorkDir         string           `json:"workDir"`
	VolumeMounts    []*VolumeMount   `json:"volumeMounts"`
	Ports           []*Port          `json:"ports"`
	Env             []*Env           `json:"env"`
	Resources       *Resource        `json:"resources"`
	SecurityContext *SecurityContext `json:"securityContext"`
	//Command         map[string]string `json:"command"`
}

//============================================================================== volume
type Path struct {
	Path string `json:"path"`
}

type Volume struct {
	Name      string `json:"name"`
	EmptyDidr string `json:"emptyDidr"`
	HostPath  *Path  `json:"hostPath"`
}

type ContainerHostVolume struct {
	Name          string
	ContainerPath string
	HostPath      string
	ReadOnly      bool
}

//==============================================================================

type PodSpecNodeSelector struct {
	Key string `json:"key"`
}

type PodSpecImagePullSecrets struct {
	Name string `json:"name"`
}

type PodSpec struct {
	Containers []*Container `json:"containers"`
	Volumes    []*Volume    `json:"volumes"`

	RestartPolicy string `json:"restartPolicy"`
	DnsPolicy     string `json:"dnsPolicy"`
	//NodeSelector  PodSpecNodeSelector `json:"nodeSelector"`
	//ImagePullSecrets PodSpecImagePullSecrets `json:"imagePullSecrets"`
}

type SpecSelector struct {
	Name string `json:"name"`
}

type ReplicationControllerSpecTemplate struct {
	MetaData ResourceMetaData `json:"metadata"`
	Spec     PodSpec          `json:"spec"`
}

type ReplicationControllerSpec struct {
	Replicas int                                `json:"replicas"`
	Selector *SpecSelector                      `json:"selector"`
	Template *ReplicationControllerSpecTemplate `json:"template"`
}

type ResourceMetaData struct {
	Name        string            `json:"name"`
	NameSpace   string            `json:"namespace"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type ReplicationController struct {
	ApiVersion string                    `json:"apiVersion"`
	Kind       string                    `json:"kind"`
	MetaData   ResourceMetaData          `json:"metadata"`
	Spec       ReplicationControllerSpec `json:"spec"`
}

func (r *ReplicationController) SetApiVersion(version string) {
	r.ApiVersion = version
}

func (r *ReplicationController) GetApiVersion() string {
	return r.ApiVersion
}

func (r *ReplicationController) SetKind(kind string) {
	r.Kind = kind
}

func (r *ReplicationController) GetKind() string {
	return r.Kind
}

func (r *ReplicationController) SetName(name string) {
	r.MetaData.Name = name
}

func (r *ReplicationController) GetName() string {
	return r.MetaData.Name
}

func (r *ReplicationController) SetNameSpace(namespace string) {
	r.MetaData.NameSpace = namespace
}

func (r *ReplicationController) GetNameSpace() string {
	return r.MetaData.NameSpace
}

func (r *ReplicationController) SetLabels(labels map[string]string) {
	r.MetaData.Labels = labels
}

func (r *ReplicationController) GetLabels() map[string]string {
	return r.MetaData.Labels
}

//func (r *ReplicationController) SetAnnotations(annotations map[string]string) {
//	r.MetaData.Annotations = annotations
//}

//func (r *ReplicationController) GetAnnotations() map[string]string {
//	return r.MetaData.Annotations
//}

func (r *ReplicationController) SetReplicas(replicas int) {
	r.Spec.Replicas = replicas
}

func (r *ReplicationController) GetReplicas() int {
	return r.Spec.Replicas
}

func (r *ReplicationController) SetSelector(selector *SpecSelector) {
	r.Spec.Selector = selector
}

func (r *ReplicationController) GetSelector() *SpecSelector {
	return r.Spec.Selector
}

func (r *ReplicationController) SetTemplate(template *ReplicationControllerSpecTemplate) {
	r.Spec.Template = template
}

func (r *ReplicationController) GetContainers() *ReplicationControllerSpecTemplate {
	return r.Spec.Template
}
