package router

import (
	"io"
	"os"
	"welfare/logger"
	"welfare/setting"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter() *gin.Engine {
	defer logger.Sync()
	r := gin.New() //不带中间件的路由
	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	//r := gin.Default()
	f, _ := os.OpenFile("welfare.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r.Use(gin.Logger())
	//r.Use(logger.LogGin()) //自己用zap封装的库
	r.Use(gin.Recovery()) //使用默认的恢复
	gin.SetMode(setting.ServerSetting.Model)
	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", getTags)
		//新建标签
		apiv1.POST("/tags", addTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", editTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", deleteTag)
	}
	logger.Info("router Ini", zap.String("activvity", "walfare"))
	return r
}
func getTags(c *gin.Context) {
	str := "test"
	c.JSON(200, gin.H{
		"message": str,
	})
}
func addTag(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "addTag",
	})
}
func editTag(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "editTag",
	})
}
func deleteTag(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "deleteTag",
	})
}
