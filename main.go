package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"quiz/global"
	"quiz/initialize"
	"time"
)

func main() {
	global.Global_DB = initialize.Gorm()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type GvaModel struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type MrtInfos struct {
	GvaModel
	MrtYearsId   uint   `json:"mrtYearsId" form:"mrtYearsId" gorm:"column:mrt_years_id;comment:年月id;size:19;"`
	OperatingDay string `json:"operatingDay" form:"operatingDay" gorm:"column:operating_day;comment:營運日;size:191;"`
	Week         string `json:"week" form:"week" gorm:"column:week;comment:星期;size:191;"`
	Total        string `json:"total" form:"total" gorm:"column:total;comment:總運量;size:191;"`
}

// TableName MrtInfos 表名
func (MrtInfos) TableName() string {
	return "mrt_infos"
}

func CreateMrtInfos(mrtInfos MrtInfos) (err error) {
	err = global.Global_DB.Create(&mrtInfos).Error
	return err
}
