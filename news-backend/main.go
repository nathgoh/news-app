package main

import (
	news "chat/api"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

		fmt.Printf("%+v", results)
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
