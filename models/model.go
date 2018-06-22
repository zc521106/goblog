package models

/**
	实体类
 */

import "time"

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

type KeyValue struct {
	Key   string
	Value string
	Order int
	Desc  string
}

type News struct {
	Id           int
	NewsId       string    //新闻id
	NewsTitle    string    //新闻标题
	NewsType     string    //新闻类型
	NewsTypename string    //新闻类型名称
	MainPic      string    //新闻封面图片
	NewsContent  string    //新闻内容
	DocUrl       string    //原文链接
	PublishTime  time.Time //发布时间
	CreateTime   time.Time //收录时间
}
