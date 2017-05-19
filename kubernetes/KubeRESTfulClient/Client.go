package KubeRESTfulClient

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"config"

	"github.com/astaxie/beego"
)

// const (
// 	KUBE_BASE_URL = "http://192.168.0.22:8080/api/v1"
// )

type Client struct {
}

func (c *Client) Get(p *Parameter) (string, error) {
	res, err := http.Get(config.KUBE_BASE_URL + p.BuildPath())
	//beego.Debug(config.KUBE_BASE_URL + p.BuildPath())
	if err != nil {
		// handle error
		beego.Debug("erro : ", err)
		return "", err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		beego.Debug("erro : ", err)
		return "", err
	}

	return string(resBody), err
}

func (c *Client) Create(p *Parameter) (string, error) {
	//beego.Debug("request url  : ", config.KUBE_BASE_URL+p.BuildPath())
	//beego.Debug("request data : ", p.GetJson())
	res, err := http.Post(config.KUBE_BASE_URL+p.BuildPath(), "application/json", bytes.NewBuffer(([]byte)(p.GetJson())))
	if err != nil {
		// handle error
		beego.Debug("erro : ", err)

		return "", err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		beego.Debug("erro : ", err)

		return "", err
	}

	return string(resBody), err
}

func (c *Client) Delete(p *Parameter) (string, error) {
	client := http.Client{}
	req, _ := http.NewRequest("DELETE", config.KUBE_BASE_URL+p.BuildPath(), nil)

	res, err := client.Do(req)

	if err != nil {
		// handle error
		beego.Debug("erro : ", err)
		return "", err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		beego.Debug("erro : ", err)
		return "", err
	}

	return string(resBody), err
}

func (c *Client) Update(p *Parameter) (string, error) {
	client := http.Client{}

	beego.Debug("update:", config.KUBE_BASE_URL+p.BuildPath())
	req, _ := http.NewRequest("PUT", config.KUBE_BASE_URL+p.BuildPath(), bytes.NewBuffer(([]byte)(p.GetJson())))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		// handle error
		beego.Debug("erro : ", err)
		return "", err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		beego.Debug("erro : ", err)
		return "", err
	}

	return string(resBody), err
}
