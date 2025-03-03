package polygonpoint

import (
	"github.com/gin-gonic/gin"
	"server/db"
	"server/model"
)

// 查询多边形点信息列表
func ListPolygonPoint(c *gin.Context) {
	var Rows []model.PolygonPoint
	var total int64
	err := db.DB.Table("polygon_point").Find(&Rows).Count(&total).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "error": err, "msg": "查询失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"total": total, "rows": Rows, "code": 200, "msg": "查询成功"})
}

// 查询多边形点信息详细信息
func GetPolygonPoint(c *gin.Context) {
	id := c.Param("id")
	var Row model.PolygonPoint
	d := db.DB.Table("polygon_point").Take(&Row, id).Error
	if d != nil {
		c.JSON(500, gin.H{"code": 500, "error": d.Error(), "msg": "查询失败"})
		panic(d)
		return
	}
	c.JSON(200, gin.H{"row": Row, "code": 200, "msg": "查询成功"})
}

// 新增多边形点信息
func AddPolygonPoint(c *gin.Context) {
	var Row model.PolygonPoint
	if err := c.ShouldBindJSON(&Row); err != nil {
		c.JSON(400, gin.H{"code": 400, "error": err.Error(), "msg": "参数错误"})
		panic(err)
		return
	}
	err := db.DB.Table("polygon_point").Create(&Row).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "error": err.Error(), "msg": "新增失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "操作成功"})
}

// 删除多边形点信息
func DelPolygonPoint(c *gin.Context) {
	id := c.Param("id")
	err := db.DB.Table("polygon_point").Delete(&model.PolygonPoint{}, id).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "操作失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "操作成功"})
}
