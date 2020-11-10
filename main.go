package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", ShowIndexPage)
	router.GET("/article/view/:id", ShowArticlePage)

	router.Run()

}

// func ShowIndexPage(c *gin.Context) {

// 	articles := getAllArticles()

// 	render(c, gin.H{
// 		"title":   "Home Page",
// 		"payload": articles}, "index.html")
// }

// func ShowArticlePage(c *gin.Context) {

// 	if id, err := strconv.Atoi(c.Param("id")); err == nil {
// 		art := getArticleById(id)

// 		render(c, gin.H{
// 			"title":   art.Title,
// 			"payload": art}, "article.html")
// 	} else {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}

// }
// func render(c *gin.Context, data gin.H, templateName string) {

// 	switch c.Request.Header.Get("Accept") {
// 	case "application/json":
// 		// Respond with JSON
// 		c.JSON(http.StatusOK, data["payload"])
// 	case "application/xml":
// 		// Respond with XML
// 		c.XML(http.StatusOK, data["payload"])
// 	default:
// 		// Respond with HTML
// 		c.HTML(http.StatusOK, templateName, data)
// 	}

// }

// type article struct {
// 	ID    int    `json: id`
// 	Title string `json: title`
// 	Body  string `json: body`
// }

// var articleList = []article{
// 	article{
// 		ID: 1, Title: "Article 1", Body: "Article body 1",
// 	},
// 	article{
// 		ID: 2, Title: "Article 2", Body: "Article body 2",
// 	},
// }

// func getAllArticles() []article {
// 	return articleList
// }

// func getArticleById(id int) *article {

// 	for _, v := range articleList {
// 		if id == v.ID {
// 			return &v
// 		}
// 	}
// 	return nil
// }

// var router *gin.Engine

// func main() {

// 	router = gin.Default()

// 	router.LoadHTMLGlob("templates/*")

// 	initializeRoutes()

// 	router.Run()
// }
