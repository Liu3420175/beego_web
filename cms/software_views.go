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
    page_ := data.GetString("page","10")
    q := data.GetString("q")
    //state := data.GetString("state")
    limit_ := data.GetString("limit","1")

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

	//offset := (page -1) * limit
	o := orm.NewOrm()
	o.Using("appinfo")
	app := new(appinfo.App)
	query := o.QueryTable(app)
    query = query.Filter("is_active",1).OrderBy("-id")
    count,_ := query.Count()
    //query = query.Limit(limit,offset)
	num,err := query.Values(&maps,"Id","Name","NameEn")
    println(maps[0]["Id"])
    type JSONStruct struct {
       code int
       msg string
       result []orm.Params
       count int64
       page int
       limit int
	}

	if err == nil{
		println(2222)
		println("num==",num)
		mystruct := &JSONStruct{0,"Success",maps,count,page,limit}
		data.Data["json"] = mystruct
		data.ServeJSON()
	}else{
		println(1111)
		mystruct := &JSONStruct{401,"Faile",nil,count,page,limit}
		data.Data["json"] = mystruct
		data.ServeJSON()

	}
    //if state != ""{
    //	query = query.Filter("state")
	//}

}