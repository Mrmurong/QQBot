package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/", func(context *gin.Context) {
		dataReader := context.Request.Body
		rawData, _ := ioutil.ReadAll(dataReader)
		postType := gjson.Get(string(rawData), "post_type").String()
		if postType == "message" {
			message := gjson.Get(string(rawData), "message").String()
			if message == "小爱" {
				context.JSON(http.StatusOK, gin.H{
					"reply": "我在呢,主人",
				})
			}
		}
	})

	_ = r.Run()
}
