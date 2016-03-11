package main

import (
  "gin-sandbox/controllers"
  "github.com/gin-gonic/gin"
  "net/http"
  //"github.com/k0kubun/pp"
)

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*")

  router.GET("/", func(c *gin.Context) {
    ctrl := task.NewTask()
    tasks := ctrl.GetAll()

    c.HTML(http.StatusOK, "home.tmpl", gin.H{
      "tasks": tasks,
    })
  })

  router.POST("/", func(c *gin.Context) {
    text := c.PostForm("text")

    ctrl := task.NewTask()
    ctrl.Create(text)

    tasks := ctrl.GetAll()

    c.HTML(http.StatusOK, "home.tmpl", gin.H{
      "tasks": tasks,
    })
  })

  router.Run(":8080")
}
