package laserout

import (
	"github.com/gin-gonic/gin"
	"server/db"
	"server/model"
	"strconv"
)

// 分页查询激光输出列表
func ListLaserOutByPage(c *gin.Context) {
	var Rows []model.LaserOut
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	d := db.DB.Table("86_system_grid.laser_out").Limit(pageSize).Offset((page - 1) * pageSize).Find(&Rows)
	if d.Error != nil {
		c.JSON(500, gin.H{"code": 500, "error": d.Error.Error(), "msg": "查询失败"})
		panic(d.Error.Error())
		return
	}
	total := d.RowsAffected
	c.JSON(200, gin.H{"total": total, "rows": Rows, "code": 200, "msg": "查询成功"})
}

// 查询激光输出详细信息
func GetLaserOut(c *gin.Context) {
	id := c.Param("id")
	var Row model.LaserOut
	err := db.DB.Table("86_system_grid.laser_out").Take(&Row, id).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "error": err.Error(), "msg": "查询失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"row": Row, "code": 200, "msg": "查询成功"})
}

// 新增激光输出信息
func AddLaserOut(c *gin.Context) {
	var Row model.LaserOut
	if err := c.ShouldBindJSON(&Row); err != nil {
		c.JSON(400, gin.H{"code": 400, "error": err.Error(), "msg": "参数错误"})
		panic(err)
		return
	}
	err := db.DB.Table("86_system_grid.laser_out").Create(&Row).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "error": err.Error(), "msg": "新增失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "操作成功"})
}

// 删除激光输出信息
func DelLaserOut(c *gin.Context) {
	id := c.Param("id")
	err := db.DB.Table("86_system_grid.laser_out").Delete(&model.LaserOut{}, id).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "error": err.Error(), "msg": "删除失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "删除成功"})
}

// 查询激光输出列表
func ListLaserOut(c *gin.Context) {
	var Rows []model.LaserOut
	err := db.DB.Table("86_system_grid.laser_out").Find(&Rows).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "error": err.Error(), "msg": "查询失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"rows": Rows, "code": 200, "msg": "查询成功"})
}
