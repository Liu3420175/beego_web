package cms

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"kandao_backend/models/appinfo"
)



type Software struct {
	beego.Controller

}

func (data *Software) SoftwareList() {
    //获取软件列表,自定义名
    page_ := data.GetString("page","1")
    q := data.GetString("q")
    //state := data.GetString("state")
    limit_ := data.GetString("limit","10")

    page_ = strings.Trim(page_," ")
    q = strings.Trim(q," ")
    //state = strings.Trim(state," ")
    limit_ = strings.Trim(limit_," ")

    page,page_err := strconv.Atoi(page_)
    limit,limit_err := strconv.Atoi(limit_)

    var maps []orm.Params

    if page_err != nil{
    	page = 1
	}

	if limit_err != nil{
		limit = 10
	}

	if q != ""{
		q = q[:128]
	}

    if page < 1{
    	page = 1
	}
	offset := (page -1) * limit

	o := orm.NewOrm()
	o.Using("appinfo")
	app := new(appinfo.App)
	query := o.QueryTable(app)
    query = query.Filter("is_active",1).OrderBy("-id")
    count,_ := query.Count()
    query = query.Limit(limit,offset)
	_,err := query.Values(&maps,"Id","Name","NameEn")

    type JSONStruct struct {
       Code int   `json:"code"`
       Msg string  `json:"msg"`
       Result []orm.Params `json:"result"`
       Count int64    `json:"count"`
       Page int       `json:"page"`
       Limit int      `json:"limit"`
	}

	if err == nil{

		mystruct := &JSONStruct{0,"Success",maps,count,page,limit}
		data.Data["json"] = mystruct
		data.ServeJSON()
		return
	}else{
		data.Data["json"] = map[string]interface{}{"code":401,"msg":"Error"}//mystruct
		data.ServeJSON()
		return

	}
    //if state != ""{
    //	query = query.Filter("state")
	//}

}

func (data *Software) GetSoftwareInfo(){
	 id_ := data.Ctx.Input.Param(":id")
	 id,id_err := strconv.Atoi(id_)

	 if id_err != nil{
	 	data.Data["json"] = map[string] interface{} {"code":10001,"msg":"Params Error"}
	 	data.ServeJSON()
	 	return
	 }

	 o := orm.NewOrm()
	 o.Using("appinfo")
	 app := appinfo.App{Id:int64(id)}
	 err := o.Read(&app)
	 if err == nil{
	 	
	 }else{
	 	data.Data["json"] = map[string]interface{}{"code":10101,"msg":"APP dose not exist"}
	 	data.ServeJSON()

	 }
}