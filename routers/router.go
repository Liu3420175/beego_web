package routers

import (
	"kandao_backend/cms"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &cms.MainController{})
    beego.Router("/v1/app/list/",&cms.AppList{})
}
