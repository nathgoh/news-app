package main

import (
	"log"
	"math"
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

func enableCors(w gin.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
}

func newsHandler(newsApi *news.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		enableCors(c.Writer)

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

		// Instance of Search struct populated with relevant fields
		search := &Search{
			searchQuery,
			nextPage,
			int(math.Ceil(float64(results.TotalResults) / float64(newsApi.PageSize))),
			results,
		}

		c.JSON(http.StatusOK, search)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// tr := &http.Transport{
	// 	MaxIdleConns:       10,
	// 	IdleConnTimeout:    30 * time.Second,
	// 	WriteBufferSize: 128 * 1024,

	// }
	// client := &http.Client{Transport: tr}

	client := &http.Client{Timeout: 30 * time.Second}
	newsApi := news.NewClient(client, os.Getenv("NEWS_API_KEY"), 20)
	router(newsApi)
}
