package main

import (
  "gin-sandbox/controllers"
  "github.com/gin-gonic/gin"
  "reflect"
  "strconv"
  "net/http"
)

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*")

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "home.tmpl", gin.H{
      "title": "Home",
    })
  })

  router.POST("/", func(c *gin.Context) {
    message := c.PostForm("message")
    c.HTML(http.StatusOK, "home.tmpl", gin.H{
      "title": "Home - After post",
      "message": message,
    })
  })

  // View User JSON
  router.GET("/:id", func(c *gin.Context) {
    n := c.Param("id")
    id, err := strconv.Atoi(n)
    if err != nil {
      c.JSON(400, err)
      return
    }

    if id <= 0 {
      c.JSON(400, gin.H{"status": "id should be bigger thsn 0"})
      return
    }

    ctrl := user.NewUser()
    result := ctrl.Get(id)
    if result == nil || reflect.ValueOf(result).IsNil() {
      c.JSON(404, gin.H{})
      return
    }

    c.JSON(200, result)
  })

  router.Run(":8080")
}
