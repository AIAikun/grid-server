package router

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"server/service/buildingposition"
	"server/service/cabledata"
	"server/service/laserout"
	"server/service/laseroutpoint"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/captchaImage", service.CaptchaImage) //生成验证码
	r.POST("/login", service.Login)              //jwt登录鉴权
	systemGroup := r.Group("/system")
	{
		cabledataGroup := systemGroup.Group("/cabledata")
		{
			cabledataGroup.GET("/list", cabledata.ListCabledata)  // 查询电缆数据列表
			cabledataGroup.GET("/:id", cabledata.GetCabledata)    // 查询电缆数据详细信息
			cabledataGroup.POST("", cabledata.AddCabledata)       // 新增电缆数据
			cabledataGroup.DELETE("/:id", cabledata.DelCabledata) // 删除电缆数据
		}
		buildingpositionGroup := systemGroup.Group("/buildingposition")
		{
			buildingpositionGroup.GET("/list", buildingposition.ListBuildingposition)  // 查询建筑位置信息列表
			buildingpositionGroup.GET("/:id", buildingposition.GetBuildingposition)    // 查询建筑位置信息详细信息
			buildingpositionGroup.POST("", buildingposition.AddBuildingposition)       // 新增建筑位置信息
			buildingpositionGroup.DELETE("/:id", buildingposition.DelBuildingposition) // 删除建筑位置信息
		}
		laserOutGroup := systemGroup.Group("/laserout")
		{
			laserOutGroup.GET("/list", laserout.ListLaserOut)  // 查询激光输出列表
			laserOutGroup.GET("/:id", laserout.GetLaserOut)    // 查询激光输出详细信息
			laserOutGroup.POST("", laserout.AddLaserOut)       // 新增激光输出信息
			laserOutGroup.DELETE("/:id", laserout.DelLaserOut) // 删除激光输出信息
		}
		laserOutPointGroup := systemGroup.Group("/laseroutpoint")
		{
			laserOutPointGroup.POST("", laseroutpoint.AddLaserOutPoint)       // 新增激光输出点信息
			laserOutPointGroup.GET("/list", laseroutpoint.ListLaserOutPoint)  // 查询激光输出点信息列表
			laserOutPointGroup.GET("/:id", laseroutpoint.GetLaserOutPoint)    // 查询激光输出点信息详细信息
			laserOutPointGroup.DELETE("/:id", laseroutpoint.DelLaserOutPoint) // 删除激光输出点信息
		}
	}
	r.GET("/circuit/state", cabledata.GetCircuitStatus) // 获取电路状态
	return r
}
