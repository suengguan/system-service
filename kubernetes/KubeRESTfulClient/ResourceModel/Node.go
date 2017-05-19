package ResourceModel

//=============================================================================
type NodeMetaDataLabel struct {
	Arch     string `json:"beta.kubernetes.io/arch"`
	Os       string `json:"beta.kubernetes.io/os"`
	HostName string `json:"kubernetes.io/hostname"`
}

type NodeMetaDataAnnotations struct {
	Detach string `json:"volumes.kubernetes.io/controller-managed-attach-detach"`
}

type NodeMetaData struct {
	Name              string                   `json:"name"`
	SelfLink          string                   `json:"selfLink"`
	Uid               string                   `json:"uid"`
	ResourceVersion   string                   `json:"resourceVersion"`
	CreationTimestamp string                   `json:"creationTimestamp"`
	Labels            *NodeMetaDataLabel       `json:"labels"`
	Annotations       *NodeMetaDataAnnotations `json:"annotations"`
}

type NodeSpec struct {
	ExternalID string `json:"externalID"`
}

type NodeStatusCapacity struct {
	NvidiaGpu string `json:"alpha.kubernetes.io/nvidia-gpu"`
	Cpu       string `json:"cpu"`
	Memory    string `json:"memory"`
	Pods      string `json:"pods"`
}

type NodeStatusCondition struct {
	Type               string `json:"Type"`
	Status             string `json:"status"`
	LastHeartbeatTime  string `json:"lastHeartbeatTime"`
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Message            string `json:"message"`
}

type NodeStatusAddress struct {
	Type    string `json:"Type"`
	Address string `json:"address"`
}

type NodeStatusDaemonEndpointsKubeleteEndpoint struct {
	Port int `json:"Port"`
}

type NodeStatusDaemonEndpoints struct {
	KubeletEndpoint *NodeStatusDaemonEndpointsKubeleteEndpoint `json:"kubeletEndpoint"`
}

type NodeStatusNodeInfo struct {
	MachineID               string `json:"machineID"`
	SystemUUID              string `json:"systemUUID"`
	BootID                  string `json:"bootID"`
	KernelVersion           string `json:"kernelVersion"`
	OsImage                 string `json:"osImage"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	OperatingSystem         string `json:"operatingSystem"`
	Architecture            string `json:"architecture"`
}

type NodeStatusImage struct {
	Names     []string `json:"names"`
	SizeBytes int64    `json:"sizeBytes"`
}

type NodeStatus struct {
	Capacity        *NodeStatusCapacity        `json:"capacity"`
	Allocatable     *NodeStatusCapacity        `json:"allocatable"`
	Conditions      []*NodeStatusCondition     `json:"conditions"`
	Addresses       []*NodeStatusAddress       `json:"addresses"`
	DaemonEndpoints *NodeStatusDaemonEndpoints `json:"daemonEndpoints"`
	NodeInfo        *NodeStatusNodeInfo        `json:"nodeInfo"`
	Images          []*NodeStatusImage         `json:"images"`
}

type Node struct {
	Metadata *NodeMetaData `json:"metadata"`
	Spec     *NodeSpec     `json:"spec"`
	Status   *NodeStatus   `json:"status"`
}

// type MetaData struct {
// 	SelfLink        string `json:"selfLink"`
// 	ResourceVersion string `json:"resourceVersion"`
// }

// type NodeList struct {
// 	ApiVersion string    `json:"apiVersion"`
// 	Kind       string    `json:"kind"`
// 	MetaData   *MetaData `json:"metadata"`
// 	Items      []*Item   `json:"items"`
// }
