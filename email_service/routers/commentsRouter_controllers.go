package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["system_service/email_service/controllers:EmailController"] = append(beego.GlobalControllerRouter["system_service/email_service/controllers:EmailController"],
		beego.ControllerComments{
			Method: "Send",
			Router: `/send/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
