package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func router() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello news!",
		})
	})

	router.GET("/search", newsSearchHandler)

	router.Run(":" + os.Getenv("PORT"))
}

func newsSearchHandler(c *gin.Context) {
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

	fmt.Println("Search Query is: ", searchQuery)
	fmt.Println("Page is: ", page)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	router()

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
