package controller

import (
	"net/http"
	"pcr-guild/components/logger"
	models "pcr-guild/models"

	"github.com/gin-gonic/gin"
	"github.com/qbox/qvm-base/components/mysql"
)

type _Box struct{}

var BoxController *_Box

// DescribeBoxResponse 短信模板返回
type DescribeBoxResponse struct {
	Items []*models.Box `json:"items"`
}

func (_ *_Box) Index(c *gin.Context) {
	var (
		log      = logger.New(c)
		database = mysql.GetBiz(log.ReqID())
	)

	box := []*models.Box{}

	if err := database.Limit(1000).Find(&box).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &DescribeBoxResponse{
		Items: box,
	})
}

func (_ *_Box) Create(c *gin.Context) {
	var (
		log      = logger.New(c)
		database = mysql.GetBiz(log.ReqID())
		args     = []models.Box{}
	)

	if err := c.BindJSON(&args); err != nil {
		log.Error("invalid args ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(args) < 1 {
		log.Error("invalid length ")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid length"})
		return
	}

	for _, item := range args {
		if err := database.Create(&item).Error; err != nil {
			log.Error("create character failed ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, args)
}

func (_ *_Box) Update(c *gin.Context) {
	var (
		log        = logger.New(c)
		id         = c.Param("id")
		updateData = map[string]interface{}{}
		box        = models.Box{}
	)

	if err := c.BindJSON(&updateData); err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := mysql.GetBiz(log.ReqID()).Model(&box).Select([]string{
		"user_id",
		"server",
		"unit_id",
		"unit_name",
		"trace",
		"level",
		"rarity",
		"promotion",
		"love_level",
		"pieces",
		"slot1",
		"slot2",
		"slot3",
		"slot4",
		"slot5",
		"slot6",
		"unique_equip_rank",
		"icon",
		"target_rarity",
		"target_promotion",
		"target_love_level",
		"target_unique_equip_rank",
		"is_finished",
		"position",
		"search_area_width",
	})
	if err := database.Where("id = ?", id).Updates(updateData).Error; err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
