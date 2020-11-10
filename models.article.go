package main

type article struct {
	ID    int    `json: id`
	Title string `json: title`
	Body  string `json: body`
}

var articleList = []article{
	article{
		ID: 1, Title: "Article 1", Body: "Article body 1",
	},
	article{
		ID: 2, Title: "Article 2", Body: "Article body 2",
	},
}

func getAllArticles() []article {
	return articleList
}

func getArticleById(id int) *article {

	for _, v := range articleList {
		if id == v.ID {
			return &v
		}
	}
	return nil
}

// type article struct {
// 	ID      int    `json:id`
// 	Title   string `json:title`
// 	Content string `json:content`
// }

// var articleList = []article{
// 	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
// 	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
// }

// func getAllArticles() []article {
// 	return articleList
// }

// func getArticleByID(id int) (*article, error) {
// 	for _, a := range articleList {
// 		if a.ID == id {
// 			return &a, nil
// 		}
// 	}
// 	return nil, errors.New("Article not found")
// }
