package main

func InitializeRoutes() {
	router.GET("/", ShowIndexPage)
	router.GET("/article/view/:id", ShowArticlePage)
}
