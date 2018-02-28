package routers

import (
	"kandao_backend/cms"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &cms.MainController{})
    beego.Router("/cms/v1/app/list/",&cms.Software{},"get:SoftwareList")
    beego.Router("/cms/v1/app/?:id",&cms.Software{},"get:GetSoftwareInfo")
    beego.Router("/cms/v1/app/add/",&cms.Software{},"post:SoftwareAdd")
}
