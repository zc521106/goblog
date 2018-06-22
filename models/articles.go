package models

/**
	blog文章相关操作
 */
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

func init() {
	orm.RegisterDataBase("default", "mysql", "root:MysqlPsw1!@tcp(39.106.15.201:3306)/go_test?charset=utf8", 30)
	// 注册要在同步数据库之前 否则会报错，至少要注册一个model
	orm.RegisterModel(new(Article))
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
	o = orm.NewOrm()
}

// 获取最火文章
func GetHotArticles(size int) ([]Article) {
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
	return queryArticles(sql, size)
}

// 获取最新文章
func GetNewArticles(size int) []Article {
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
		"   a.create_time DESC " +
		" LIMIT ? "
	return queryArticles(sql, size)
}

// 获取置顶文章
func GetUpArticles(size int) []Article {
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
		"   a.up DESC ,a.scan_num DESC " +
		" LIMIT ? "
	return queryArticles(sql, size)
}

// 执行数据库 查询文章信息
func queryArticles(sql string, size int) []Article {
	var articles []Article
	num, err := o.Raw(sql, size).QueryRows(&articles)
	if err == nil && num > 0 {
		return articles
	} else {
		beego.Debug(" error => GetNewArticles")
		return nil
	}
}

// 加载文章分类信息
func GetArticleClasses() []KeyValue {
	sql := " SELECT " +
		"   s.name `key`, s.value `value`, s.order `order` " +
		" FROM " +
		"   article_class s " +
		" WHERE " +
		"   s.status = '1' " +
		" ORDER BY " +
		"   s.order "
	var classes []KeyValue
	num, err := o.Raw(sql).QueryRows(&classes)
	if err == nil && num > 0 {
		return classes
	} else {
		return nil
	}
}

// 加载文章分类信息
func GetArticleClassesWithType(articleType string) []KeyValue {
	sql := " SELECT " +
		"   s.name `key`, s.value `value`, " +
		"   CASE WHEN s.value = ? then 100 ELSE s.order END `order` " +
		" FROM " +
		"   article_class s " +
		" WHERE " +
		"   s.status = '1' " +
		" ORDER BY " +
		"   s.order"
	var classes []KeyValue
	num, err := o.Raw(sql, articleType).QueryRows(&classes)
	if err == nil && num > 0 {
		return classes
	} else {
		return nil
	}
}

// 加载友情链接
func GetAppUrls() []KeyValue {
	sql := " SELECT " +
		"   a.apply_url_name `key`, a.apply_url `value` " +
		" FROM " +
		"   apply_url a " +
		" WHERE " +
		"   a.check = '1' " +
		" ORDER BY " +
		"   a.order DESC, " +
		"   a.create_time DESC "
	var urls []KeyValue
	num, err := o.Raw(sql).QueryRows(&urls)
	if err == nil && num > 0 {
		return urls
	} else {
		return nil
	}
}

// 获取公告
func GetAnnouncement(page string) string {
	sql := "SELECT a.content FROM announcement a WHERE a.page = ? AND a.end_time > now() AND a.`status` = '1'"
	var s string
	err := o.Raw(sql, page).QueryRow(&s)
	if err == nil {
		return s
	} else {
		sql := "SELECT a.content FROM announcement a WHERE a.page = 'all' AND a.end_time > now() AND a.`status` = '1'"
		err := o.Raw(sql).QueryRow(&s)
		if err == nil {
			return s
		} else {
			return "暂无公告！"
		}
	}
}