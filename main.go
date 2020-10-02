package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

func main(){
	hash:=""
	r:=gin.Default()
	r.LoadHTMLGlob("assets/*.html")
	r.Static("/style","./style")
	r.GET("/",func(c* gin.Context){
		c.Redirect(http.StatusFound,"/hash")
	})
	r.GET("/hash",func(c*gin.Context){
		c.HTML(200,"index.html",hash)
		hash=""
	})
	r.POST("/hash/process",func(c* gin.Context){
		strToHash:=c.PostForm("str")
		if strToHash=="" {
			hash=""
			c.Redirect(http.StatusFound,"/hash")
			return
		}
		data,_:=bcrypt.GenerateFromPassword([]byte(strToHash),3)
		hash = string(data)
		c.Redirect(http.StatusFound,"/hash")
	})

	r.Run()
}