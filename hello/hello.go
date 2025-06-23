
package main

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func index_handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "message": "this is the index" })
}

func hello_handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "message": "This is a, HELLO." })
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/", index_handler)
	r.GET("/hello", hello_handler)
	r.Run()
}
