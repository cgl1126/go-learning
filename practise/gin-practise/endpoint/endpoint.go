package endpoint

import (
	"context"

	"github.com/gin-gonic/gin"

	"go-learning/practise/gin-practise/hander"

	"go-learning/practise/gin-practise/log"
)

type GinResp struct {
	Code    int         `json:"code"`
	Resp    interface{} `json:"resp,omitempty"`
	Message string      `json:"message,omitempty"`
}

func (g *GinResp) SetCode(c int) {
	g.Code = c
}
func (g *GinResp) SetMessage(msg string) {
	g.Message = msg
}

func GetPractise(c *gin.Context) {
	r := GinResp{}

	log.Glog.Info("get GetPractise log")

	r.SetMessage("Get Practise Message")
	r.SetCode(200)

	c.JSON(200, r)
}

func PostPractise(c *gin.Context) {
	r := GinResp{}

	var gr hander.GinRequest
	var err error
	err = c.ShouldBindJSON(&gr)
	if err != nil {
		c.AbortWithStatusJSON(500, "bad request")
		return
	}

	if err = hander.Dohandler(context.TODO(), gr); err != nil {
		r.SetMessage(err.Error())
		c.AbortWithStatusJSON(500, "bad request")
		return
	}

	r.SetMessage("Get Practise Message")
	r.SetCode(200)

	log.Glog.Info("Post Practise log")

	r.Resp = gr
	c.JSON(200, r)
}