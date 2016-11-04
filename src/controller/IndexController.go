package controller

import (
	"html/template"
	"model"
	"net/http"
	"strconv"
	"time"
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
	Create_time string
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
	categorys := GetCategory()

	//查询文章.
	articles := GetArticle()

	t, _ := template.ParseFiles(Layout+"index.html", IndexPath+"right.html")
	t.Execute(resp, map[string]interface{}{"CategoryData": categorys, "articlesData": articles})
}

func (this *IndexController) DemoAction(resp http.ResponseWriter, req *http.Request) {

}

//查询分类.
func GetCategory() []*Category {
	rows := model.GetCategory()

	categorys := make([]*Category, 0)

	for rows.Next() {
		category := &Category{}
		rows.Scan(&category.Id, &category.Category_name, &category.Is_delete, &category.Order, &category.Url)
		categorys = append(categorys, category)
	}
	return categorys
}

//查询文章.
func GetArticle() []*Article {

	//查询文章.
	articleRow := model.GetArticle()

	articles := make([]*Article, 0)

	for articleRow.Next() {

		//获取结构体.
		article := &Article{}
		articleRow.Scan(&article.Id, &article.Title, &article.Content, &article.Create_time, &article.Read, &article.Is_delete)

		//将string类型的时间戳转换成int64.
		t, _ := strconv.ParseInt(article.Create_time, 10, 64)

		//将int64位的时间戳转换成时间格式.
		article.Create_time = time.Unix(t, 0).Format("2006-01-02 15:04:05")

		articles = append(articles, article)
	}
	return articles
}
