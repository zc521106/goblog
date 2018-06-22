package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	articles := models.GetHotArticles(10)
	c.Data["firstArticle"] = articles[0]
	articles1 := [9]models.Article{}
	for i := 1; i < 10; i++ {
		articles1[i - 1] = articles[i]
	}
	c.Data["hotArticles"] = articles1
	c.Data["newArticles"] = models.GetNewArticles(22)
	c.Data["upArticles"] = models.GetUpArticles(10)
	c.Data["news"] = models.GetNews("tech", 12)
	c.Data["classes"] = models.GetArticleClasses()
	c.Data["urls"] = models.GetAppUrls()
	c.Data["announcement"] = models.GetAnnouncement("index")
	c.TplName = "index.html"
}
