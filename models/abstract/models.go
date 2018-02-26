package abstract

/*
抽象类
 */




type ResourceStorage struct {
	Id int64              `orm:"pk;auto"`
	Title string          `orm:"size(62);unique;null"`
	Provider string       `orm:"size(1);null;default(Q)"`
	Prefix string         `orm:"size(254);null"`
	IsA bool              `orm:"null"`
	NA int                `orm:"null"`
	SA string             `orm:"size(15);null"`

}

type ResourceLink struct {
	Id int64              `orm:"pk;auto"`
	Rid string            `orm:"size(255);null"`
	Path string           `orm:"size(255);null"`
	IsA bool              `orm:"null"`
	IsB bool              `orm:"null"`
	NA int                `orm:"null"`
	NB int                `orm:"null"`
	SA string             `orm:"size(7);null"`
	SB string             `orm:"size(7);null"`
}


type ResourceStat struct {
	Id int64              `orm:"pk;auto"`
	IsA bool              `orm:"null"`
	NA int                `orm:"null"`
	NB int                `orm:"null"`
	NC int                `orm:"null"`
	SA string             `orm:"size(15);null"`
	SB string             `orm:"size(15);null"`
}