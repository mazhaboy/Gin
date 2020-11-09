package main

func initializeRoutes() {
	router.GET("/", ShowIndexPage)
	router.GET("/article/view/:id", ShowArticlePage)
}

// func initializeRoutes() {
// 	router.GET("/", showIndexPage)
// 	router.GET("/article/view/:article_id", getArticle)
// }
