package cms

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"kandao_backend/models/appinfo"
)

type AddApp struct {
	beego.Controller
}

type AppList struct {
	beego.Controller
}

func (data *AppList) Get() {
    //获取软件列表
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
	num,err := query.Values(&maps,"Id","Name","NameEn")
    println(num)
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
		println(1111)
		data.Data["json"] = map[string]interface{}{"code":401,"msg":"Error"}//mystruct
		data.ServeJSON()
		return

	}
    //if state != ""{
    //	query = query.Filter("state")
	//}

}