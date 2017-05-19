package kubernetes

import (
	"strings"

	"encoding/json"
	"system-service/kubernetes/KubeRESTfulClient"
	"system-service/kubernetes/KubeRESTfulClient/ResourceModel"

	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type KubernetesService struct {
}

func (this *KubernetesService) getPodsByNamespace(namespace string) ([]*ResourceModel.Pod, error) {
	var client KubeRESTfulClient.Client
	var param KubeRESTfulClient.Parameter
	var err error

	var response ResourceModel.Response

	// get pods under namespace
	param.SetNamespace(namespace)
	param.SetResourceType(KubeRESTfulClient.PODS)

	//beego.Debug("kubernetes get pods by namespace")
	resp, err := client.Get(&param)
	if err == nil {
		//beego.Debug("kubernetes get pods by namespace success")
		//beego.Debug("RESP :", resp)
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

func (this *KubernetesService) getLogByPodname(namespace string, podname string) (string, error) {
	var result string
	var err error
	var client KubeRESTfulClient.Client
	var param KubeRESTfulClient.Parameter

	// get pod namespace
	param.SetNamespace(namespace)
	param.SetResourceType(KubeRESTfulClient.PODS)
	param.SetName(podname)
	param.SetSubPath("log")

	result, err = client.Get(&param)
	if err != nil {
		beego.Debug(err)
	}

	return result, err
}

func (this *KubernetesService) GetPodnameByLabel(namespace string, label string) (string, error) {
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

func (this *KubernetesService) GetPodsByNamespace(namespace string) ([]string, error) {
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

func (this *KubernetesService) GetPodLog(namespace string, labelName string) (string, error) {
	var result string
	var err error
	var podName string

	beego.Debug("get pod name by label")
	podName, err = this.GetPodnameByLabel(namespace, labelName)
	if err == nil {
		beego.Debug("get pod name by label success")
		// get log by pod name
		beego.Debug(podName)

		beego.Debug("get log by pod name")
		result, err = this.getLogByPodname(namespace, podName)
		if err == nil {
			beego.Debug("get log by pod name success")
		} else {
			beego.Debug("get log by pod name failed")
			beego.Debug(err)
		}
	} else {
		beego.Debug("get pod name by label failed")
	}

	return result, err
}

func (this *KubernetesService) WatingForRCComplete(namespace string, label string) error {
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

// todo only for test
func (this *KubernetesService) printResponse(resp string) {
	// todo
	// for i := 0; i < 60; i++ {
	// 	resp, err = client.Get(&param)
	// 	if err != nil {
	// 		beego.Debug(err, resp)
	// 	}

	// 	//beego.Debug(label)
	// 	err = json.Unmarshal(([]byte)(resp), &response)
	// 	for j := 0; j < len(response.Items); j++ {
	// 		podName = response.Items[j].MetaData.Name
	// 		status = response.Items[j].Status.Phase
	// 		//beego.Debug(podName)
	// 		//beego.Debug(status)

	// 		if strings.Contains(podName, label) {
	// 			beego.Debug("=============================================")
	// 			beego.Debug("podName:", podName)
	// 			beego.Debug("label:", label)
	// 			// phase
	// 			beego.Debug(podName, "phase", status)
	// 			beego.Debug(podName, "hostIP", response.Items[j].Status.HostIP)
	// 			beego.Debug(podName, "podIP", response.Items[j].Status.PodIP)
	// 			beego.Debug(podName, "startTime", response.Items[j].Status.StartTime)
	// 			// conditions
	// 			conditionCnt := len(response.Items[j].Status.Conditions)
	// 			for k := 0; k < conditionCnt; k++ {
	// 				//response.Items[j].Status.Conditions[j]
	// 				beego.Debug(podName, "condition:Type", response.Items[j].Status.Conditions[k].Type)
	// 				beego.Debug(podName, "condition:Status", response.Items[j].Status.Conditions[k].Status)
	// 				if response.Items[j].Status.Conditions[k].LastProbeTime != nil {
	// 					beego.Debug(podName, "condition:LastProbeTime", *response.Items[j].Status.Conditions[k].LastProbeTime)
	// 				}
	// 				beego.Debug(podName, "condition:LastTransitionTime", response.Items[j].Status.Conditions[k].LastTransitionTime)
	// 				beego.Debug(podName, "condition:Reason", response.Items[j].Status.Conditions[k].Reason)
	// 				beego.Debug(podName, "condition:Message", response.Items[j].Status.Conditions[k].Message)
	// 			}
	// 			// container
	// 			containerCnt := len(response.Items[j].Status.ContainerStatuses)
	// 			for k := 0; k < containerCnt; k++ {
	// 				beego.Debug(podName, "container:name", response.Items[j].Status.ContainerStatuses[k].Name)
	// 				if response.Items[j].Status.ContainerStatuses[k].State.Running != nil {
	// 					beego.Debug(podName, "container:Running:StartedAt", response.Items[j].Status.ContainerStatuses[k].State.Running.StartedAt)
	// 				}
	// 				if response.Items[j].Status.ContainerStatuses[k].State.Waiting != nil {
	// 					beego.Debug(podName, "container:Waiting:Reason", response.Items[j].Status.ContainerStatuses[k].State.Waiting.Reason)
	// 					beego.Debug(podName, "container:Waiting:Message", response.Items[j].Status.ContainerStatuses[k].State.Waiting.Message)
	// 				}
	// 				beego.Debug(podName, "container:ready", response.Items[j].Status.ContainerStatuses[k].Ready)
	// 				beego.Debug(podName, "container:RestartCount", response.Items[j].Status.ContainerStatuses[k].RestartCount)
	// 				beego.Debug(podName, "container:Image", response.Items[j].Status.ContainerStatuses[k].Image)
	// 				beego.Debug(podName, "container:ImageID", response.Items[j].Status.ContainerStatuses[k].ImageID)
	// 				beego.Debug(podName, "container:ContainerID", response.Items[j].Status.ContainerStatuses[k].ContainerID)

	// 			}

	// 			if utility.StringEqual(status, "Running") {
	// 				beego.Debug(podName, "is", status)
	// 				return nil
	// 			}
	// 		}
	// 	}

	// 	time.Sleep(time.Second * 2)
	// }
}

func (this *KubernetesService) WatingForServiceComplete(namespace string) error {
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

func (this *KubernetesService) CreateResource(namespace string, resourceType string, jsonContent string) error {
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

func (this *KubernetesService) pdateResource(namespace string, resourceType string, resourceName string, jsonContent string) error {
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

func (this *KubernetesService) DeleteResource(namespace string, resourceType string, label string) error {
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

func (this *KubernetesService) GetNodeList() (*ResourceModel.ResponseNode, error) {
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
