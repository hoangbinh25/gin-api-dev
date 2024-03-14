package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Golang/GinFw/gin-api-dev/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	// server.Use(middlewares.Authenticate) // Use auth middlewares
	server.GET("/", welcome)

	server.GET("/getQueryURL/:name/:age", getQueryUrl)
	server.POST("/postData", middlewares.Authenticate, middlewares.AddHeader, postData)

	// server.Run(":8080")

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// Custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// Use Basic authentication
	auth := gin.BasicAuth(gin.Accounts{
		"user":  "pass",
		"user2": "pass2",
		"user3": "pass3",
	})
	// End use Basic authentication

	// Group server
	admin := server.Group("/admin", auth)
	{
		admin.GET("/getData", getData)
	}

	client := server.Group("/client")
	{
		client.GET("/getQueryString", getQueryString)
	}
	// End group server

	// Custom HTTP configuration
	s := &http.Server{
		Addr:         ":8888",
		Handler:      server,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	s.ListenAndServe()
	// End custom HTTP configuration

}

// func welcome
func welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I am testing server Gin Framework",
	})
}

// end func welcome

// func getData
func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I am getting data from gin framework",
	})
}

// end func getData

// func postData
func postData(c *gin.Context) {
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)
	c.JSON(200, gin.H{
		"message":  "I am post data from gin framework",
		"bodyData": string(value),
	})
}

// end func postData

// QueryUrl : https://localhost:8080/getQueryUrl/Binh/20
func getQueryUrl(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"message": "Hi. I am getQueryURL Gin Framework",
		"name":    name,
		"age":     age,
	})
}

// end func getQueryUrl

// QueryString : https://localhost:8080/getQueryString?name=Binh&age=20
func getQueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"message": "Hi. I am from Gin Framework",
		"name":    name,
		"age":     age,
	})
}

// end QueryString
