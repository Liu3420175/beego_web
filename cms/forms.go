package cms


//表单验证器

type SoftwareForm struct {
	Name string         `valid:"Required;MaxSize(70)"`
	Title string        `valid:"Required;MaxSize(30)"`
	Platform string     `valid:"Required"`
    Description string   `valid:"MaxSize(1000)"`
    NameEn string        `valid:"MaxSize(128)"`
    DescriptionEn string  `valid:"MaxSize(1024)"`
    IsOnline  bool
}

