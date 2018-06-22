package models

// 获取新闻
func GetNews(newType string, size int) []News {
	sql := "SELECT " +
		"   s.news_title, s.news_type, s.news_typename, s.main_pic, s.news_content, " +
		"   s.doc_url, date_format(s.create_time, '%m-%d') publish_time " +
		" FROM " +
		"   news s " +
		" WHERE " +
		"   s.news_type = ? " +
		" ORDER BY " +
		"   s.create_time DESC " +
		" LIMIT ? "
	var news []News
	if num, err := o.Raw(sql, newType, size).QueryRows(&news); err == nil && num > 0 {
		return news
	} else {
		return nil
	}
}
