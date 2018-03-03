package utils

import (
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego"
)

type T   beego.Controller

type JSONList struct {
	Code int             `json:"code"` //TODO 这个不能少
	Msg string           `json:"msg"`
	Result []orm.Params  `json:"result"`
	Count int64          `json:"count"`
	Page int             `json:"page"`
	Limit int            `json:"limit"`
}


type JSONObject struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Result map[string]interface{} `json:"result"`
}

func (data *T) Common_response( code int,count int64,page int,limit int,flag bool,r []orm.Params)  {

		msg := Codes[code]
		if flag {
			mystruct := &JSONList{code, msg, r, count, page, limit}
			data.Data["json"] = mystruct

		}else{
			mystruct := &JSONObject{code,msg,r[0]}
			data.Data["json"] = mystruct
		}

	    data.Ctx.Output.JSON(data.Data["json"], true, false)
		return
}
