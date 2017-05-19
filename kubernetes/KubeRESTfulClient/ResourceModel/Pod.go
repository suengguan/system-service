package ResourceModel

type PodMetaDataLabels struct {
	Name string `json:"name"`
}

type PodMetaDataAnnotations struct {
	Anonotation string `json:"kubernetes.io/created-by"`
}

type PodMetaData struct {
	Name              string                  `json:"name"`
	GenerateName      string                  `json:"generateName"`
	Namespace         string                  `json:"namespace"`
	SelfLink          string                  `json:"selfLink"`
	Uid               string                  `json:"uid"`
	ResourceVersion   string                  `json:"resourceVersion"`
	CreationTimestamp string                  `json:"creationTimestamp"`
	Labels            *PodMetaDataLabels      `json:"labels"`
	Annotations       *PodMetaDataAnnotations `json:"annotations"`
}

type RespPodSpec struct {
}

type PodStatusConditionLastProbeTime struct {
}

type PodStatusCondition struct {
	Type               string                           `json:"type"`
	Status             string                           `json:"status"`
	LastProbeTime      *PodStatusConditionLastProbeTime `json:"lastProbeTime"`
	LastTransitionTime string                           `json:"lastTransitionTime"`
	Reason             string                           `json:"reason"`
	Message            string                           `json:"message"`
}

type PodStatusContainerStatusStateRunning struct {
	StartedAt string `json:"startedAt"`
}

type PodStatusContainerStatusStateWaiting struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

type PodStatusContainerStatusState struct {
	Running *PodStatusContainerStatusStateRunning `json:"running"`
	Waiting *PodStatusContainerStatusStateWaiting `json:"waiting"`
}

type PodStatusContainerStatusLastState struct {
}

type PodStatusContainerStatus struct {
	Name         string                             `json:"name"`
	State        *PodStatusContainerStatusState     `json:"state"`
	LastState    *PodStatusContainerStatusLastState `json:"lastState"`
	Ready        bool                               `json:"ready"`
	RestartCount int                                `json:"restartCount"`
	Image        string                             `json:"image"`
	ImageID      string                             `json:"imageID"`
	ContainerID  string                             `json:"containerID"`
}

type PodStatus struct {
	Phase             string                      `json:"phase"`
	Conditions        []*PodStatusCondition       `json:"conditions"`
	HostIP            string                      `json:"hostIP"`
	PodIP             string                      `json:"podIP"`
	StartTime         string                      `json:"startTime"`
	ContainerStatuses []*PodStatusContainerStatus `json:"containerStatuses"`
}

type Pod struct {
	MetaData *PodMetaData `json:"metadata"`
	Spec     *RespPodSpec `json:"spec"`
	Status   *PodStatus   `json:"status"`
}
