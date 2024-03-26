package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response for letting all api return the same format
type Response struct {
	Status int         `json:"code"` // 0: success , 1: fail
	Meta   Meta        `json:"meta"`
	Data   interface{} `json:"data"`
}
type Meta struct {
	Count int `json:"count"`
}

func RespErr(ctx *gin.Context, err string) {
	ctx.JSON(
		http.StatusOK,
		Response{
			Status: 1,
			Meta: Meta{
				Count: 0,
			},
			Data: err,
		})
	return
}
func Resp(ctx *gin.Context, data interface{}) {

	ctx.JSON(
		http.StatusOK,
		Response{
			Status: 0,
			Meta: Meta{
				Count: 1,
			},
			Data: data,
		})
	return

}

func RespWithCount(ctx *gin.Context, data interface{}, count int) {
	ctx.JSON(
		http.StatusOK,
		Response{
			Status: 0,
			Meta: Meta{
				Count: count,
			},
			Data: data,
		})
	return
}

/*
{
    "code": 0,
    "meta": {
        "count": 1
    },
    "data": [
        {
            "id": 0,
            "state": 1,
            "owner": 1,
            "stamp": 1716711837000,
            "location": "Travel Prepare 5",
            "together": [
                2,
                3,
                4,
                5,
                6
            ],
            "logs": "I can fly~",
            "img": "www.baidu.com"
        }
    ]
}

*/
