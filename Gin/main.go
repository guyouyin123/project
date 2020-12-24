package main

import (
	"gin-class/middlewares"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strconv"
)

type DududuShop struct {
	//gorm.Model
	ID        uint `gorm:"primary_key"`
	Spu_id	uint `gorm:"unique"`
	Logo_url string
	Title string
	Sub_title string
	Price float64
	Sold_num uint
	Goods_type uint
}


func main() {
	db, _ := gorm.Open("mysql", "root:admin123@/dududu?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&DududuShop{},)
	defer db.Close()


	r := gin.Default()
	r.Use(middlewares.Cors())


	r.GET("/:tag/:page", func(c *gin.Context) {
		tag := c.Param("tag")
		page,_ := strconv.Atoi(c.Param("page"))
		goods_type := 0
		if tag=="shoes" {
			goods_type = 0
		}else if tag=="cloth" {
			goods_type = 1
		}else if tag=="package"{
			goods_type = 3
		}else if tag=="watch"{
			goods_type = 100
		}else if tag=="lipstick"{
			goods_type = 5
		}
		var shop =make([]DududuShop,5)
		_ = c.ShouldBind(&shop)
		//db.Offset((page-1)*12).Limit(12).Find(&shop)
		db.Where("goods_type = ?",goods_type).Offset((page-1)*12).Limit(12).Find(&shop)
		c.JSON(200, gin.H{
			"msg": shop,
		})
	})



	_ = r.Run(":8080")
}
