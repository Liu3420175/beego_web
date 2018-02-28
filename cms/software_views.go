package cms

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"kandao_backend/models/appinfo"
	"kandao_backend/utils"
	"reflect"
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

     maps := []orm.Params{} //

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
	_,err := query.Values(&maps,"Id","Name","Title","Description","Platform") // TODO信息不完全

	if err == nil{
        code := 0
        msg := utils.Codes[code]
		mystruct := &utils.JSONList{code,msg,maps,count,page,limit}
		data.Data["json"] = mystruct
		data.ServeJSON()
		return

	}else{
		code := 10000
		msg := utils.Codes[code]
		mystruct := &utils.JSONObject{code,msg,make(map[string]interface{})}
		data.Data["json"] = mystruct
		data.ServeJSON()
		return

	}
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
	 	result := map[string]interface{}{
	 		"obj_id":app.Id,
			"name": app.Name,
			"title": app.Title,
			"description": app.Description,
			"creator": app.Creator,
			"lastmodifier": app.Lastmodifier,
			//"date_created": cdate,
			//"date_modified": mdate,
			"platform": app.Platform,
			"name_en":app.NameEn,
			"is_online":app.IsOnline,
			"description_en":app.DescriptionEn,
			//"icon_uri":obj.icon_uri(),
			/label":reflect.TypeOf(app),
		}
		 code := 0
		 msg := utils.Codes[code]
		 mystruct := &utils.JSONObject{code,msg,result}
		 data.Data["json"] = mystruct
		 data.ServeJSON()
		 return

	 }else{
		 code := 10701
		 msg := utils.Codes[code]
		 mystruct := &utils.JSONObject{code,msg,make(map[string]interface{})}
		 data.Data["json"] = mystruct
		 data.ServeJSON()
		 return
	 }
}