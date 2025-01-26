package buildingposition

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/db"
	"server/model"
	"strconv"
)

// 查询建筑位置信息列表
func ListBuildingposition(c *gin.Context) {
	var Rows []model.BuildingRow
	var total int64
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	err := db.DB.Table("building_position").Count(&total).Limit(pageSize).Offset((page - 1) * pageSize).Find(&Rows).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "error": err, "msg": "查询失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"total": total, "rows": Rows, "code": 200, "msg": "查询成功"})
}

// 查询建筑位置信息详细信息
func GetBuildingposition(c *gin.Context) {
	id := c.Param("id")
	var Row model.BuildingRow
	err := db.DB.Table("building_position").Take(&Row, id).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "error": err.Error(), "msg": "查询失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"row": Row, "code": 200, "msg": "查询成功"})
}

// 新增建筑位置信息
func AddBuildingposition(c *gin.Context) {
	var Row model.BuildingRow
	if err := c.ShouldBindJSON(&Row); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "error": err.Error(), "msg": "参数错误"})
		panic(err)
		return
	}
	err := db.DB.Table("building_position").Create(&Row).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "error": err.Error(), "msg": "新增失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "操作成功"})
}

// 删除建筑位置信息
func DelBuildingposition(c *gin.Context) {
	id := c.Param("id")
	err := db.DB.Table("building_position").Delete(&model.BuildingRow{}, id).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "操作失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "操作成功"})
}
