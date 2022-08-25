package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tanweiquan/webmake/dao/db"
	"github.com/tanweiquan/webmake/models/article"
	"github.com/tanweiquan/webmake/models/user"
	"github.com/tanweiquan/webmake/router"
	"github.com/tanweiquan/webmake/util"
)

func init() {
	vc := art{}
	router.GET("/article_author", vc.JsonList)
	router.GET("/article_content", vc.JsonArt)
	router.POST("/article_edit", vc.Edit)
}

type art struct {
	A string
}

/*
* 函数绑定一个结构体，防止重名的函数影响
* 业务逻辑
 */

func (t *art) JsonList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	data := make(map[string]interface{})
	msg := "succeed"
	code := 0
	for {
		record := new(user.User).ListAll()

		data["userAbout"] = record
		data["time"] = time.Now().String()
		break
	}
	response := make(map[string]interface{})
	response["msg"] = msg
	response["code"] = code
	response["data"] = data
	c.JSON(http.StatusOK, response)
}
func (t *art) JsonArt(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	data := make(map[string]interface{})
	msg := "succeed"
	code := 0
	for {
		record := new(article.Article).ListAll()
		data["article"] = record
		break
	}
	response := make(map[string]interface{})
	response["msg"] = msg
	response["code"] = code
	response["data"] = data
	// 设置响应头包含：Access-Control-Allow-Origin"，官方解决解决跨域问题
	c.Header("Access-Control-Allow-Origin", "*") // 这里的*是通配符，代表所有的静态页面都可以跨域请求服务端数据
	c.JSON(http.StatusOK, response)
}
func (t *art) Edit(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	data := make(map[string]interface{})
	msg := "succeed"
	code := 0
	for {
		// c.BindJSON();c.BindXML();c.BindYAML();c.BindQuery();c.BindUri()；c.Bind()
		articleStr := c.PostForm("article")
		authorStr := c.PostForm("author")
		record := article.Article{}
		record.Id = util.GetUUID()

		record.Article = articleStr
		record.Author = authorStr
		effect, err := db.GetDbengine().Insert(record)
		if err != nil {
			code = 100
			msg = err.Error()
			log.Println(err)
			break
		}
		effectStr := "上传成功，共上传" + strconv.Itoa(int(effect)) + "条"
		data["effect"] = effectStr
		break
	}
	response := make(map[string]interface{})
	response["msg"] = msg
	response["code"] = code
	response["data"] = data
	c.JSON(http.StatusOK, response)
}
