package controller

import (
	//	"fmt"
	"html/template"
	"model"
	"net/http"
)

//定义一个分类的结构体.
type Category struct {
	Id            int
	Category_name string
	Is_delete     int
	Order         int
	Url           string
}

//定义一个文章的结构体.
type Article struct {
	Id          int
	Title       string
	Content     string
	Create_time int
	Read        int
	Is_delete   int
}

type IndexController struct {
}

//定义路径常量.
const IndexPath = "../template/html/index/"
const Layout = "../template/html/layout/"

//首页方法.
func (this *IndexController) IndexAction(resp http.ResponseWriter, req *http.Request) {

	//查询分类.
	rows := model.GetCategory()

	categorys := make([]*Category, 0)

	for rows.Next() {
		category := &Category{}
		rows.Scan(&category.Id, &category.Category_name, &category.Is_delete, &category.Order, &category.Url)
		categorys = append(categorys, category)
	}

	//查询文章.
	articleRow := model.GetArticle()

	//创建一个数组.
	articles := make([]*Article, 0)

	for articleRow.Next() {
		//获取结构体.
		article := &Article{}
		articleRow.Scan(&article.Id, &article.Title, &article.Content, &article.Create_time, &article.Read, &article.Is_delete)
		articles = append(articles, article)
	}

	t, _ := template.ParseFiles(Layout+"index.html", IndexPath+"right.html")
	t.Execute(resp, map[string]interface{}{"CategoryData": categorys, "articlesData": articles})
}

func (this *IndexController) DemoAction(resp http.ResponseWriter, req *http.Request) {

}
