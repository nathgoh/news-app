package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	news "news/api"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Search struct {
	Query      string
	NextPage   int
	TotalPages int
	Results    *news.Results
}

func router(newsApi *news.Client) {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello news!",
		})
	})

	router.GET("/search", newsHandler(newsApi))

	router.Run(":" + os.Getenv("PORT"))
}

func newsHandler(newsApi *news.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := url.Parse(c.Request.URL.String())
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		params := u.Query()

		searchQuery := params.Get("topic")
		page := params.Get("page")
		if page == "" {
			page = "1"
		}

		results, err := newsApi.FetchEverything(searchQuery, page)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		nextPage, err := strconv.Atoi(page)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		// Instance of Search struct created with relevant fields
		search := &Search{
			searchQuery,
			nextPage,
			int(results.TotalResults / newsApi.PageSize),
			results,
		}

		fmt.Printf("%v", search)

		c.JSON(http.StatusOK, search)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	myClient := &http.Client{Timeout: 10 * time.Second}
	newsApi := news.NewClient(myClient, os.Getenv("NEWS_API_KEY"), 20)
	router(newsApi)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
