package kubernetes

import (
	"strings"

	"encoding/json"
	"system-service/kubernetes/KubeRESTfulClient"
	"system-service/kubernetes/KubeRESTfulClient/ResourceModel"

	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type Service struct {
}

func (this *Service) CreateResource(namespace string, resourceType string, jsonContent string) error {
	var err error
	var client KubeRESTfulClient.Client
	var param KubeRESTfulClient.Parameter
	var resp string

	param.SetNamespace(namespace)
	param.SetResourceType(resourceType)
	param.SetJson(jsonContent)

	resp, err = client.Create(&param)
	if err != nil {
		beego.Debug(err, resp)
	}
	beego.Debug(resp)

	return err
}

func (this *Service) updateResource(namespace string, resourceType string, resourceName string, jsonContent string) error {
	var err error
	var client KubeRESTfulClient.Client
	var param KubeRESTfulClient.Parameter
	var resp string

	param.SetNamespace(namespace)
	param.SetResourceType(resourceType)
	param.SetJson(jsonContent)
	param.SetSubPath(resourceName)

	resp, err = client.Update(&param)
	if err != nil {
		beego.Debug(err, resp)
	}
	beego.Debug(resp)

	return err
}

func (this *Service) CreateResourceQuota(namespace string, name string, cpu string, memory string) error {
	var err error

	var jsonStr string
	resourceQuota := new(ResourceModel.ResourceQuota)
	resourceQuota.SetApiVersion(ResourceModel.KUBE_API_VERSION)
	resourceQuota.SetKind(ResourceModel.KUBE_RESOURCE_RESOURCEQUOTA)
	resourceQuota.SetName(name)
	resourceQuota.SetNameSpace(namespace)

	var hard ResourceModel.ResourceQuotaSpecHard
	hard.Cpu = cpu
	hard.Memory = memory
	hard.PersistentVolumeClaims = "50"
	hard.Pods = "50"
	hard.ReplicationControllers = "50"
	hard.ResourceQuotas = "1"
	hard.Secrets = "50"
	hard.Services = "50"
	resourceQuota.SetSpecHard(&hard)

	resultByte, err := json.Marshal(resourceQuota)
	jsonStr = string(resultByte)

	//beego.Debug(jsonStr)

	err = this.CreateResource(namespace, KubeRESTfulClient.RESOURCEQUOTAS, jsonStr)
	if err != nil {
		beego.Debug(err)
	}

	return err
}

func (this *Service) UpdateResourceQuota(namespace string, name string, cpu string, memory string) error {
	var err error

	var jsonStr string
	resourceQuota := new(ResourceModel.ResourceQuota)
	resourceQuota.SetApiVersion(ResourceModel.KUBE_API_VERSION)
	resourceQuota.SetKind(ResourceModel.KUBE_RESOURCE_RESOURCEQUOTA)
	resourceQuota.SetName(name)
	resourceQuota.SetNameSpace(namespace)

	var hard ResourceModel.ResourceQuotaSpecHard
	hard.Cpu = cpu
	hard.Memory = memory
	hard.PersistentVolumeClaims = "50"
	hard.Pods = "50"
	hard.ReplicationControllers = "50"
	hard.ResourceQuotas = "1"
	hard.Secrets = "50"
	hard.Services = "50"
	resourceQuota.SetSpecHard(&hard)

	resultByte, err := json.Marshal(resourceQuota)
	jsonStr = string(resultByte)

	beego.Debug(jsonStr)

	err = this.updateResource(namespace, KubeRESTfulClient.RESOURCEQUOTAS, name, jsonStr)
	if err != nil {
		beego.Debug(err)
	}

	return err
}

func (this *Service) CreateNamespace(ns string) error {
	var err error

	var jsonStr string
	namespace := new(ResourceModel.NameSpace)
	namespace.SetApiVersion(ResourceModel.KUBE_API_VERSION)
	namespace.SetKind(ResourceModel.KUBE_RESOURCE_NAMESPACE)
	namespace.SetName(ns)
	resultByte, err := json.Marshal(namespace)
	jsonStr = string(resultByte)

	err = this.CreateResource("", KubeRESTfulClient.NAMESPACES, jsonStr)
	if err != nil {
		beego.Debug(err)
	}

	return err
}

func (this *Service) GetPodsByNamespace(namespace string) ([]string, error) {
	var err error

	var pods []*ResourceModel.Pod
	var podsName []string

	pods, err = this.getPodsByNamespace(namespace)
	if err == nil {
		// get pod name
		for _, p := range pods {
			podsName = append(podsName, p.MetaData.Name)
		}
	}

	return podsName, err
}

func (this *Service) getPodsByNamespace(namespace string) ([]*ResourceModel.Pod, error) {
	var client KubeRESTfulClient.Client
	var param KubeRESTfulClient.Parameter
	var err error
	var response ResourceModel.Response

	// get pods under namespace
	param.SetNamespace(namespace)
	param.SetResourceType(KubeRESTfulClient.PODS)

	resp, err := client.Get(&param)
	if err == nil {
		err = json.Unmarshal(([]byte)(resp), &response)
		if err == nil {

		} else {
			beego.Debug("json Unmarshal data failed")
		}
	} else {
		beego.Debug("kubernetes get pods by namespace failed")
		beego.Debug(err)
	}

	return response.Items, err
}

func (this *Service) CreateRc(namespace string, label string, image string, containerPort int,
	limitCpu, limitMemory string, envList []*ResourceModel.Env, hostPathList []*ResourceModel.ContainerHostVolume) error {
	var jsonContent string
	var err error

	var rc ResourceModel.ReplicationController
	rc.SetApiVersion(ResourceModel.KUBE_API_VERSION)
	rc.SetKind(ResourceModel.KUBE_RESOURCE_REPLICATIONCONTROLLER)
	rc.SetName(label)
	rc.SetNameSpace(namespace)
	labels := make(map[string]string)
	labels["name"] = label
	rc.SetLabels(labels)
	rc.SetReplicas(1)

	var selector ResourceModel.SpecSelector
	selector.Name = label
	rc.SetSelector(&selector)

	// set template
	var template ResourceModel.ReplicationControllerSpecTemplate
	// template metadata
	template.MetaData.Labels = make(map[string]string)
	template.MetaData.Labels["name"] = label

	// template spec
	var container ResourceModel.Container
	container.Name = label
	container.Image = image

	// resource
	if len(limitCpu) > 0 && len(limitMemory) > 0 {
		var resourceLimit ResourceModel.ResourceLimit
		resourceLimit.Cpu = limitCpu
		resourceLimit.Memory = limitMemory

		var resource ResourceModel.Resource
		resource.Limits = &resourceLimit
		container.Resources = &resource
	}

	// env
	for _, e := range envList {
		container.Env = append(container.Env, e)
	}

	// host path
	for _, h := range hostPathList {
		var containerMountPath ResourceModel.VolumeMount
		containerMountPath.Name = h.Name
		containerMountPath.MountPath = h.ContainerPath
		containerMountPath.ReadOnly = h.ReadOnly
		container.VolumeMounts = append(container.VolumeMounts, &containerMountPath)

		var hostPathVolume ResourceModel.Volume
		hostPathVolume.Name = h.Name
		var hostPath ResourceModel.Path
		hostPath.Path = h.HostPath
		hostPathVolume.HostPath = &hostPath
		template.Spec.Volumes = append(template.Spec.Volumes, &hostPathVolume)
	}

	// port
	var port ResourceModel.Port
	port.ContainerPort = containerPort
	container.Ports = append(container.Ports, &port)

	// securityContext
	var securityContext ResourceModel.SecurityContext
	securityContext.Privileged = true
	container.SecurityContext = &securityContext

	template.Spec.Containers = append(template.Spec.Containers, &container)

	rc.SetTemplate(&template)

	// convert to json
	resultByte, err := json.Marshal(rc)
	if err != nil {
		beego.Debug(err)
		return err
	} else {
		jsonContent = string(resultByte)
	}

	err = this.CreateResource(namespace, KubeRESTfulClient.REPLICATIONCONTROLLERS, jsonContent)
	if err == nil {
		beego.Debug("wating for create rc complete")
		err = this.watingForRCComplete(namespace, label)
		if err == nil {
			//beego.Debug("waiting 20s")
			//time.Sleep(time.Second * 20)
			beego.Debug("create rc success")
		} else {
			beego.Debug("wating for create rc complete failed")
			beego.Debug(err)
		}
	} else {
		beego.Debug("create rc failed")
		beego.Debug(err)
	}

	return err
}

func (this *Service) watingForRCComplete(namespace string, label string) error {
	// todo
	var err error
	var client KubeRESTfulClient.Client
	var param KubeRESTfulClient.Parameter
	var resp string
	var response ResourceModel.Response

	var podName string
	var status string

	param.SetNamespace(namespace)
	param.SetResourceType(KubeRESTfulClient.PODS)

	timeOutCnt := 0
	for {
		resp, err = client.Get(&param)
		if err != nil {
			beego.Debug(err, resp)
		}

		var waitingReason string
		err = json.Unmarshal(([]byte)(resp), &response)
		for j := 0; j < len(response.Items); j++ {
			podName = response.Items[j].MetaData.Name
			status = response.Items[j].Status.Phase

			if strings.Contains(podName, label) {
				//beego.Debug("=============================================")
				// phase
				//beego.Debug(podName, "phase", status)
				containerCnt := len(response.Items[j].Status.ContainerStatuses)
				for k := 0; k < containerCnt; k++ {
					if response.Items[j].Status.ContainerStatuses[k].State.Waiting != nil {
						waitingReason = response.Items[j].Status.ContainerStatuses[k].State.Waiting.Reason
						beego.Debug(podName, "watingReason:", waitingReason)
						if waitingReason == "ContainerCreating" {
							if status == "Running" {
								beego.Debug(podName, "is", status)
								return nil
							}
						} else {
							err = fmt.Errorf("%s", waitingReason)
							return err
						}
					} else {
						if status == "Running" {
							beego.Debug(podName, "is", status)
							return nil
						}
					}
				}

			}
		}

		time.Sleep(time.Second * 5)
		timeOutCnt++
		if timeOutCnt >= 1000 {
			err = fmt.Errorf("%s", "create rc time out")
			return err
		}
	}

	err = fmt.Errorf("%s", "create pod failed")

	return err
}

func (this *Service) GetPodName(namespace, label string) (string, error) {
	var podName string
	var err error

	podName, err = this.getPodnameByLabel(namespace, label)

	return podName, err
}

func (this *Service) getPodnameByLabel(namespace string, label string) (string, error) {
	var podName string
	var err error

	var pods []*ResourceModel.Pod

	pods, err = this.getPodsByNamespace(namespace)
	if err == nil {
		// get pod name
		for i := 0; i < len(pods); i++ {
			podName = pods[i].MetaData.Name
			if strings.Contains(podName, label+"-") {
				return podName, nil
			}
		}
	}

	err = fmt.Errorf("%s %s", label, "pod is not existed")
	beego.Debug(err)

	return podName, err
}

func (this *Service) CreateSvc(namespace string, name string, label string, containerPort int) error {
	var jsonContent string
	var err error

	var svc ResourceModel.Service
	svc.SetApiVersion(ResourceModel.KUBE_API_VERSION)
	svc.SetKind(ResourceModel.KUBE_RESOURCE_SERVICE)

	// meta data
	svc.SetName(name)
	svc.SetNameSpace(namespace)
	labels := make(map[string]string)
	labels["name"] = name
	svc.SetLabels(labels)

	// spec
	var selector ResourceModel.SpecSelector
	selector.Name = label
	svc.Spec.Selector = &selector

	var port ResourceModel.ServicePort
	port.Port = containerPort
	svc.Spec.Ports = append(svc.Spec.Ports, &port)

	resultByte, err := json.Marshal(svc)
	if err != nil {
		beego.Debug(err)
		return err
	} else {
		jsonContent = string(resultByte)
	}

	err = this.CreateResource(namespace, KubeRESTfulClient.SERVICES, jsonContent)
	if err == nil {
		beego.Debug("wating for create svc complete")
		err = this.watingForServiceComplete(namespace)
		if err == nil {
			beego.Debug("create svc success")
		} else {
			beego.Debug("wating for create svc complete failed")
			beego.Debug(err)
		}
	} else {
		beego.Debug("create svc failed")
		beego.Debug(err)
	}

	return err
}

func (this *Service) watingForServiceComplete(namespace string) error {
	// todo
	var err error
	var client KubeRESTfulClient.Client
	var param KubeRESTfulClient.Parameter
	var resp string
	//var response ResourceModel.Response

	param.SetNamespace(namespace)
	param.SetResourceType(KubeRESTfulClient.SERVICES)

	// get service list
	resp, err = client.Get(&param)
	if err != nil {
		beego.Debug(err, resp)
	}

	// check service is exit

	//beego.Debug(string(resp))

	return err
}

func (this *Service) DeleteRc(namespace string, label string) error {
	var err error
	var podName string

	err = this.deleteResource(namespace, KubeRESTfulClient.REPLICATIONCONTROLLERS, label)
	if err == nil {
		podName, err = this.getPodnameByLabel(namespace, label)
		err = this.deleteResource(namespace, KubeRESTfulClient.PODS, podName)
		if err != nil {
			beego.Debug(err)
		}
	} else {
		beego.Debug(err)
	}

	return err
}

func (this *Service) deleteResource(namespace string, resourceType string, label string) error {
	var err error
	var client KubeRESTfulClient.Client
	var param KubeRESTfulClient.Parameter
	var resp string

	// delete rc
	param.SetNamespace(namespace)
	param.SetResourceType(resourceType)
	param.SetName(label)

	resp, err = client.Delete(&param)
	if err != nil {
		beego.Debug(err, resp)
	}
	//beego.Debug(resp)

	return err
}

func (this *Service) DeleteSvc(namespace string, label string) error {
	var err error

	err = this.deleteResource(namespace, KubeRESTfulClient.SERVICES, label)
	if err != nil {
		beego.Debug(err)
	}

	return err
}

func (this *Service) getNodeList() (*ResourceModel.ResponseNode, error) {
	var err error
	var client KubeRESTfulClient.Client
	var param KubeRESTfulClient.Parameter
	var resp string
	var response ResourceModel.ResponseNode

	param.SetResourceType(KubeRESTfulClient.NODES)

	resp, err = client.Get(&param)
	if err != nil {
		beego.Debug(err, resp)
		return nil, err
	}

	err = json.Unmarshal(([]byte)(resp), &response)
	if err != nil {
		beego.Debug("json Unmarshal data failed")
		return nil, err
	}

	return &response, err
}

// return cpu,memory,error
func (this *Service) GetTotalCpuAndMemory() (float64, float64, error) {
	var totalCpu float64
	var totalMemory float64
	var err error

	// get cpu and memory kubernetes
	var nodeList *ResourceModel.ResponseNode

	nodeList, err = this.getNodeList()
	var tmpTotalMemory float64
	for _, n := range nodeList.Items {
		// check condition
		for _, c := range n.Status.Conditions {
			if c.Type == "Ready" && c.Status == "True" {
				// cpu
				var nodeCpu float64
				nodeCpu, err = strconv.ParseFloat(n.Status.Allocatable.Cpu, 32)
				totalCpu += nodeCpu

				// memory
				var nodeMemory float64
				memory := n.Status.Allocatable.Memory
				nodeMemory, err = strconv.ParseFloat(memory[:len(memory)-2], 32)
				tmpTotalMemory += nodeMemory
			}
		}
	}

	totalMemory = tmpTotalMemory / 1024.0 / 1024.0

	return totalCpu, totalMemory, err
}

func (this *Service) DeleteNamespace(namespace string) error {
	var err error

	err = this.deleteResource("", KubeRESTfulClient.NAMESPACES, namespace)
	if err != nil {
		beego.Debug(err)
	}

	return err
}
