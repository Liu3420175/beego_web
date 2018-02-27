package routers

import (
	"kandao_backend/cms"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &cms.MainController{})
    beego.Router("/cms/v1/app/list/",&cms.Software{},"get:SoftwareList")
}
