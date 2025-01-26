package cabledata

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"server/db"
	"server/model"
	"strconv"
	"time"
)

// 查询电缆数据列表
func ListCabledata(c *gin.Context) {
	var Rows []model.CableRow
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	d := db.DB.Table("cable_data").Limit(pageSize).Offset((page - 1) * pageSize).Find(&Rows)
	if d.Error != nil {
		c.JSON(500, gin.H{"code": 500, "error": d.Error.Error(), "msg": "查询失败"})
		panic(d.Error.Error())
		return
	}
	total := d.RowsAffected
	c.JSON(200, gin.H{"total": total, "rows": Rows, "code": 200, "msg": "查询成功"})
}

// 查询电缆数据详细信息
func GetCabledata(c *gin.Context) {
	id := c.Param("id")
	var Row model.CableRow
	err := db.DB.Table("cable_data").Take(&Row, id).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "error": err.Error(), "msg": "查询失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"row": Row, "code": 200, "msg": "查询成功"})
}

// 新增电缆数据
func AddCabledata(c *gin.Context) {
	var Row model.CableRow
	if err := c.ShouldBindJSON(&Row); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "error": err.Error(), "msg": "参数错误"})
		panic(err)
		return
	}
	err := db.DB.Table("cable_data").Create(&Row).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "error": err.Error(), "msg": "新增失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "操作成功"})
}

// 删除电缆数据
func DelCabledata(c *gin.Context) {
	id := c.Param("id")
	err := db.DB.Table("cable_data").Delete(&model.CableRow{}, id).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "操作失败"})
		panic(err)
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "操作成功"})
}

// 获取电路状态
func GetCircuitStatus(c *gin.Context) {
	num := 200
	normalProbability := 0.8
	rand.Seed(time.Now().UnixNano())
	var list []model.CircuitState
	for i := 0; i < num; i++ {
		circuitState := model.CircuitState{
			Id: int64(i),
		}
		p := rand.Float64()
		if p < normalProbability {
			circuitState.Normal = true
			circuitState.FaultProbability = rand.Float64()
		} else {
			circuitState.Normal = false
			circuitState.FaultProbability = 1.0
		}
		list = append(list, circuitState)
	}
	c.JSON(http.StatusOK, gin.H{
		"total": num,
		"rows":  list,
		"code":  200,
		"msg":   "查询成功",
	})
}
