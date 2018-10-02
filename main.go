package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()
	router.Run(":8081")
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/golang", func(c *gin.Context) {
		c.String(200, "Hello from Golang Service")
	})

	r.GET("/java", func(c *gin.Context) {
		resp, err := http.Get("http://java-telepresence-service:8080/java")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		c.String(200, string(contents))
	})

	return r
}
