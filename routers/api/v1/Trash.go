package v1

import (
	"SignInBoard/models"
	"github.com/gin-gonic/gin"
)

func GenerateTrash(c *gin.Context)  {
	for i:= 0;i < 100; i++ {
		err := models.CreateTrash()
		if err != nil {

			c.JSON(200,gin.H{
				"code":500,
				"data": nil,
				"msg":"获取垃圾信息失败",

			})
			return
		}
	}
	c.JSON(200,gin.H{
		"code":200,
		"data":nil,
		"msg":"生成垃圾成功",
	})


}
func GetAllTrash(c *gin.Context)  {

	data,err:=models.GetAllTrashData()
	if err != nil {
		c.JSON(200,gin.H{
			"code":500,
			"data": nil,
			"msg":"获取垃圾信息失败",

		})
		return
	}

	c.JSON(200,gin.H{
		"code":200,
		"data":data,
		"msg":"SUCCESS",
	})

}
func DeleteTrashById(c *gin.Context)  {
	trashId := c.Param("id")
	_,err1 := models.GetTrashById(trashId)
	if err1 != nil {
		c.JSON(200,gin.H{
			"code":500,
			"data":nil,
			"msg":"该条数据已经不存在，请删除其他垃圾数据",
		})
		return
	}
	err := models.DeleteTrash(trashId)
	if err != nil {
		c.JSON(200,gin.H{
			"code":500,
			"data":nil,
			"msg":"delete fail",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":200,
		"data":nil,
		"msg":"delete success",
	})

}
