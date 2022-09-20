package v1

import (
	"SignInBoard/data"
	"SignInBoard/models"
	"github.com/gin-gonic/gin"
	"strconv"
)
func CreateSignIn(c *gin.Context)  {
	var signInParam data.SignInParam
	err1 := c.ShouldBind(&signInParam)
	if err1 != nil {
		c.JSON(
			200,gin.H{
				"code":500,
				"data":nil,
				"msg":"参数匹配失败，请检查参数的填写",
			})
		return
	}

	_, err := models.CreateSignIn(signInParam)
	if err != nil {
		c.JSON(
			200,gin.H{
				"code":500,
				"data":nil,
				"msg":"签到失败！",
			})
		return
	}

	result,err2 := models.GetSignByName(signInParam.Name)
	if err2 != nil {
		c.JSON(
			200,gin.H{
				"code":500,
				"data":nil,
				"msg":"签到失败！",
			})

		return
	}

	c.JSON(200,gin.H{
		"code":200,
		"data":result,
		"msg":"签到成功",
	})

}

func CurrentSign(c *gin.Context) {
	result,err := models.GetAllSign()

	if err != nil {
		c.JSON(200,gin.H{
			"code":500,
			"data":nil,
			"msg":"获取签到信息失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":200,
		"data":result,
		"msg":"SUCCESS",
	})
}

func UpdateSign(c *gin.Context) {
	var updateParam data.UpdateParam
	c.ShouldBind(&updateParam)
	_,err1:=models.GetSignIn(strconv.Itoa(updateParam.Id))
	if err1 != nil {
		c.JSON(200,gin.H{
			"code":500,
			"data":nil,
			"msg":"当前要更新的记录不存在，请核对id",
		})
		return

	}
	err := models.UpdateSignById(updateParam)
	if err != nil {
		c.JSON(200,gin.H{
			"code":500,
			"data":nil,
			"msg":"update fail",
		})
		return
	}

	result,err2 := models.GetSignIn(strconv.Itoa(updateParam.Id))
	if err2 != nil {
		c.JSON(200,gin.H{
			"code":500,
			"data":nil,
			"msg":"update fail",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":500,
		"data":result,
		"msg":"SUCCESS",
	})

}
