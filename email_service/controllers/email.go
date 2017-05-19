package controllers

import (
	"encoding/json"
	"model"

	"system_service/email_service/models"
	"system_service/email_service/service"

	"github.com/astaxie/beego"
)

// Operations about Email
type EmailController struct {
	beego.Controller
}

// @Title Send
// @Description send emails
// @Param	body		body 	models.Email	true		"body for email content"
// @Success 200 {object} models.Response
// @Failure 403 body is empty
// @router /send/ [post]
func (this *EmailController) Send() {
	var err error
	var emails []*model.Email
	var response models.Response

	// unmarshal
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &emails)
	if err == nil {
		var svc service.EmailService
		err = svc.Send(emails)
		if err == nil {
			response.Status = model.MSG_RESULTCODE_SUCCESS
			response.Reason = "success"
			response.Result = ""
		}
	} else {
		beego.Debug("Unmarshal data failed")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}

	this.Data["json"] = &response

	this.ServeJSON()
}
