package ResourceModel

type ServiceStatusLoadBalancerIngress struct {
	Ip       string `json:"ip"`
	HostName string `json:"hostname"`
}

type ServiceStatusLoadBalancer struct {
	Ingress ServiceStatusLoadBalancerIngress `json:"ingress"`
}

type ServiceStatus struct {
	LoadBalancer ServiceStatusLoadBalancer `json:"loadBalancer"`
}

type ServicePort struct {
	Name       string `json:"name"`
	Port       int    `json:"port"`
	TargetPort int    `json:"targetPort"`
	NodePort   int    `json:"nodePort"`
	Protocol   string `json:"protocol"`
}

type ServiceSpec struct {
	Selector        *SpecSelector  `json:"selector"`
	Type            string         `json:"type"`
	ClusterIP       string         `json:"clusterIP"`
	SessionAffinity string         `json:"sessionAffinity"`
	Ports           []*ServicePort `json:"ports"`
}

//=============================================================================
type Service struct {
	ApiVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	MetaData   ResourceMetaData `json:"metadata"`
	Spec       ServiceSpec      `json:"spec"`
	//Status     ServiceStatus    `json:"status"`
}

func (s *Service) SetApiVersion(version string) {
	s.ApiVersion = version
}

func (s *Service) GetApiVersion() string {
	return s.ApiVersion
}

func (s *Service) SetKind(kind string) {
	s.Kind = kind
}

func (s *Service) GetKind() string {
	return s.Kind
}

func (s *Service) SetName(name string) {
	s.MetaData.Name = name
}

func (s *Service) GetName() string {
	return s.MetaData.Name
}

func (s *Service) SetNameSpace(namespace string) {
	s.MetaData.NameSpace = namespace
}

func (s *Service) GetNameSpace() string {
	return s.MetaData.NameSpace
}

func (s *Service) SetLabels(labels map[string]string) {
	s.MetaData.Labels = labels
}

func (s *Service) GetLabels() map[string]string {
	return s.MetaData.Labels
}
