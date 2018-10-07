package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{", "}}")

	r.LoadHTMLGlob("./templates/*.gohtml")

	r.GET("/", func(c *gin.Context) {
		var posts []string

		files, err := ioutil.ReadDir("./markdown/")
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
			posts = append(posts, file.Name())
		}

		c.HTML(http.StatusOK, "index.gohtml", gin.H{
			"posts": posts,
		})
	})
	r.Run()

	r.GET("/:postName", func(c *gin.Context) {
		postName := c.Param("postName")

		mdfile, err := ioutil.ReadFile("./markdown/" + postName)

		if err != nil {
			fmt.Println(err)
			c.HTML(http.StatusNotFound, "error.gohtml", nil)
			c.Abort()
			return
		}
		postHTML := template.HTML(blackfriday.MarkdownCommon([]byte(mdfile)))

		type Post struct {
			Title   string
			Content template.HTML
		}

		post := Post{Title: postName, Content: postHTML}

		c.HTML(http.StatusOK, "post.gohtml", gin.H{
			"Title":   post.Title,
			"Content": post.Content,
		})
	})

}
