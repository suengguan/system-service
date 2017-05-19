package ResourceModel

const (
	// version
	KUBE_API_VERSION = "v1"

	// resource type
	KUBE_RESOURCE_NAMESPACE             = "Namespace"
	KUBE_RESOURCE_REPLICATIONCONTROLLER = "ReplicationController"
	KUBE_RESOURCE_SERVICE               = "Service"
	KUBE_RESOURCE_RESOURCEQUOTA         = "ResourceQuota"

	// for replication controller
	KUBE_RC_RESTART_POLICY_ALWAYS    = "Always"
	KUBE_RC_RESTART_POLICY_NEVER     = "NEVER"
	KUBE_RC_RESTART_POLICY_ONFAILURE = "OnFailure"

	KUBE_RC_DNS_POLICY_DEFAULT      = "Default"
	KUBE_RC_DNS_POLICY_CLUSTERFIRST = "ClusterFirst"

	// for service
	KUBE_SVC_SPEC_TYPE_NODEPORT = "NodePort"
)
