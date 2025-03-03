package laseroutpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/db"
	"server/model"
	"strconv"
)

// 新增激光输出点数据
func AddLaserOutPoint(c *gin.Context) {
	var Row model.LaserOutPoint
	if err := c.ShouldBindJSON(&Row); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "error": err.Error(), "msg": "参数错误"})
		panic(err)
		return
	}
	err := db.DB.Table("laser_out_point").Create(&Row).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "error": err.Error(), "msg": "新增失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "操作成功"})
}

// 分页查询激光输出点信息列表
func ListLaserOutPointByPage(c *gin.Context) {
	var Rows []model.LaserOutPoint
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	d := db.DB.Table("86_system_grid.laser_out_point").Limit(pageSize).Offset((page - 1) * pageSize).Find(&Rows)
	if d.Error != nil {
		c.JSON(500, gin.H{"code": 500, "error": d.Error.Error(), "msg": "查询失败"})
		panic(d.Error.Error())
		return
	}
	total := d.RowsAffected
	c.JSON(200, gin.H{"total": total, "rows": Rows, "code": 200, "msg": "查询成功"})
}

// 查询激光输出点信息详细信息
func GetLaserOutPoint(c *gin.Context) {
	id := c.Param("id")
	var Row model.LaserOutPoint
	err := db.DB.Table("laser_out_point").Take(&Row, id).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "error": err.Error(), "msg": "查询失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"row": Row, "code": 200, "msg": "查询成功"})
}

// 删除激光输出点信息
func DelLaserOutPoint(c *gin.Context) {
	id := c.Param("id")
	err := db.DB.Table("laser_out_point").Where("id = ?", id).Delete(&model.LaserOutPoint{}).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "error": err.Error(), "msg": "删除失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "删除成功"})
}

// 查询激光输出点信息列表
func ListLaserOutPoint(c *gin.Context) {
	var Rows []model.LaserOutPoint
	d := db.DB.Table("86_system_grid.laser_out_point").Find(&Rows)
	if d.Error != nil {
		c.JSON(500, gin.H{"code": 500, "error": d.Error.Error(), "msg": "查询失败"})
		panic(d.Error.Error())
		return
	}
	total := d.RowsAffected
	c.JSON(200, gin.H{"total": total, "rows": Rows, "code": 200, "msg": "查询成功"})
}
