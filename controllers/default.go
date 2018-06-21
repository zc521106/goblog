package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type MainController struct {
	beego.Controller
}

var o orm.Ormer

func init() {
	orm.RegisterDataBase("default", "mysql", "root:MysqlPsw1!@tcp(39.106.15.201:3306)/go_test?charset=utf8", 30)
	// 注册要在同步数据库之前 否则会报错，至少要注册一个model
	orm.RegisterModel(new(Article))
	orm.RunSyncdb("default", false, true)
	o = orm.NewOrm()
}

func (c *MainController) Get() {
	articles := getHotBlogs(10)
	if articles != nil {
		c.Data["firstBlog"] = articles[0]
		articles1 := [9]Article{}
		for i := 1; i < 10; i++ {
			articles1[i - 1] = articles[i]
		}
		c.Data["hotBlog"] = articles1
	} else {
		beego.Debug(" error index ")
	}
	c.TplName = "index.html"
}

type Article struct {
	Id            int       `json:id`
	Title         string    `json:title`
	UserId        int       `json:user_id`
	Author        string    `json:author`
	Class_        string    `json:class_`
	MainPic       string    `json:main_pic`
	SimpleContent string    `json:simple_content`
	Content       string    `json:content`
	CreateTime    time.Time `json:create_time`
	UpdateTime    time.Time `json:update_time`
	ScanNum       int       `json:scan_num`
	CommentNum    int       `json:comment_num`
	PraiseNum     int       `json:praise_num`
	Up            string    `json:up`
}

func getHotBlogs(size int) ([]Article) {
	sql := " SELECT " +
		"   a.id, a.title, a.class_, a.main_pic, a.check," +
		"   replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(a.simple_content, " +
		"   '&danyinhao_E&', '\\''), '&baifenhao&', '%'), '&danyinhao_C_L&', '‘'), '&danyinhao_C_R&', '’'), '&shuangyinhao_E&', '\\\"'), '&shuangyinhao_C_L&', '“'), " +
		"   '&shuangyinhao_C_R&', '”'), '&douhao_C&', '，'), '&douhao_E&', ','), '&fenhao_E&', ';'), '&fenhao_C&', '；'), '&eite&', '@'), '&xiaoyu&', '<'), '&dayu&', '>') simple_content, " +
		"   date_format(a.create_time, '%Y-%m-%d %H:%i:%s') create_time," +
		"   a.scan_num, a.comment_num, a.praise_num, a.up " +
		" FROM" +
		"   article a " +
		" WHERE" +
		"   a.check = '1' " +
		" ORDER BY " +
		"   (a.scan_num + (a.comment_num * 100)) DESC " +
		" LIMIT ? "
	var articles []Article
	num, err := o.Raw(sql, size).QueryRows(&articles)
	if err == nil && num > 0 {
		return articles
	} else {
		beego.Debug(" error => getHotBlogs ")
		return nil
	}
}