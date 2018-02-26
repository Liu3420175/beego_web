package appinfo

/*
appinfo 模块
 */


 import (
 	"github.com/astaxie/beego/orm"
 	_ "github.com/go-sql-driver/mysql"
 	"kandao_backend/models/abstract"
 	"fmt"
	 "time"
 )

type AppStorage struct {
	abstract.ResourceStorage

}


func (app *AppStorage) TabelName() string {
	//重定义表名
	return "appinfo_appstorage"
}

func (app *AppStorage) TableUnique()[][] string{
	//定义唯一索引
	return [][]string{
		[]string{"Provider","Prefix"},
	}
}


type App struct {
	Id int64                  	`orm:"pk;auto"`
	Name string               	`orm:"size(70);null"`
	NaneEn string             	`orm:"size(128);null"`
	Title string              	`orm:"size(30);null"`
	Platform string           	`orm:"size(20);null"`
	DateCreated time.Time     	`orm:"auto_now_add;type(datetime)"`
	DateModified time.Time    	`orm:"auto_now;type(datetime)"`
	Creator string            	`orm:"size(128);null"`
	Lastmodifier string       	`orm:"size(128);null"`
	Description string        	`orm:"size(1024);null"`
	DescriptionEn string      	`orm:"size(1024);null"`
	IsOnline bool             	`orm:"default(false)"`
	IsActive bool 				`orm:"default(true)"`

	Versions []*AppVersion      `orm:"reverse(many)"`

}


func (app *App) TableName() string{
	return "appinfo_app"
}


type AppIcon struct {
	Id int64 					`orm:"pk;auto"`
	Title *App 					`orm:"rel(fk);null;on_delete(set_null)"`
	Storage *AppStorage 		`orm:"rel(fk);null;on_delete(set_null)"`
	Rid string 					`orm:"size(255);null"`
	RidSmall string 			`orm:"size(255);null"`
	Path string 				`orm:"size(255);null"`
	IsDefault bool 				`orm:"default(false)"`

}

func (app *AppIcon) TableName() string {
	return "appinfo_appicon"
}

func (app *AppIcon) TableIndex() [][]string{
	//添加索引
	return [][]string{
		[]string{"Title"},
	}

}



type AppVersion struct {
	Id int64 					`orm:"pk;auto"`
	Title *App 					`orm:"rel(fk);null;on_delete(set_null)"`
	VersionName string			`orm:"size(20);null"`
	VersionCode int             `orm:"null"`
	Status string       		`orm:"size(2);null;default(OF)"`
	DateCreated time.Time     	`orm:"auto_now_add;type(datetime)"`
	DateModified time.Time    	`orm:"auto_now;type(datetime)"`
	Creator string            	`orm:"size(128);null"`
	Lastmodifier string       	`orm:"size(128);null"`
    PkgName string              `orm:"size(128);null"`
	Description string        	`orm:"size(1024);null"`
	DescriptionEn string      	`orm:"size(1024);null"`
	Rid string 					`orm:"size(128);null"`
}



func (app *AppVersion) TableName() string {
	//重命名表名
	return "appinfo_appversion"
}


func (app *AppVersion) TableIndex() [][]string{
	return [][]string{
		[]string{"Title"},
	}
}

type AppDownloadPlatform struct {
	Id int64 					`orm:"pk;auto"`
	Title string                `orm:"size(32)"`
	Name string 				`orm:"size(64);null"`
	Description string        	`orm:"size(256);null"`
	DateCreated time.Time     	`orm:"auto_now_add;type(datetime)"`
	DateModified time.Time    	`orm:"auto_now;type(datetime)"`
	Creator string            	`orm:"size(128);null"`
	Lastmodifier string       	`orm:"size(128);null"`
	IsActive bool 				`orm:"default(true)"`
}


func (app *AppDownloadPlatform) TableName() string {
	return "appinfo_appdownloadplatform"
}


type AppDownloadAddress struct {
	Id int64 					`orm:"pk;auto"`
	Version *AppVersion         `orm:"rel(fk);null;on_delete(set_null)"`
	Platform *AppDownloadPlatform `orm:"rel(fk);null;on_delete(set_null)"`
	Uri string     				`orm:"size(256);null"`
}


func (app *AppDownloadAddress) TableName() string {
	return "appinfo_appdownloadaddress"
}


func (app *AppDownloadAddress) TableIndex() [][]string {
	return [][]string{
		[]string{"Version"},
	}
}

type AppLink struct {
	abstract.ResourceLink
	AppVersion *AppVersion        `orm:"rel(fk);null;on_delete(set_null)"`
	Title string                  `orm:"size(10);null;default(CN)"`
	Storage *AppStorage           `orm:"rel(fk);null"`
	Description string        	  `orm:"size(300);null"`
	DateCreated time.Time     	  `orm:"auto_now_add;type(datetime)"`

}


func (app *AppLink) TableName() string {
	return "appinfo_applink"
}


func (app *AppLink) TableIndex() [][]string{
	return [][]string{
		[]string{"AppVersion","Storage"},
	}
}

type AppVersionStat struct {
	abstract.ResourceStat
	AppVersion *AppVersion        `orm:"rel(fk);null;on_delete(set_null)"`
	Ndownload int                 `orm:"null;default(0)"`

}

func (app *AppVersionStat) TableName() string {
	return "appinfo_appversionstat"
}

func (app *AppVersionStat) TableIndex() [][]string{
	return [][]string{
		[]string{"AppVersion"},
	}
}


func init() {
	//这种设计支持分库
	host := "127.0.0.1"
    port := "3306"
    username := "root"
    password := "asdasd"
    db := "appinfo2"

    orm.RegisterDriver("mysql",orm.DRMySQL)

    coon := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8"
    orm.RegisterDataBase("appinfo","mysql",coon)
	fmt.Printf("数据库连接成功！%s\n", coon)

	orm.RegisterModel(new(AppStorage),
		              new(App),
		              new(AppIcon),
		              new(AppVersion),
		              new(AppDownloadAddress),
		              new(AppDownloadPlatform),
		              new(AppLink),
		              new(AppVersionStat),
	)
}