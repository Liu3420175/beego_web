package utils

import (
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego"
)


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

func Common_response(data *beego.Controller ,code int,count int64,page int,limit int,r []orm.Params)  {

		msg := Codes[code]
		mystruct := &JSONList{code,msg,r,count,page,limit}
		data.Data["json"] = mystruct
		data.ServeJSON()
		return
}
