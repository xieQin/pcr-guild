package controller

import (
	"net/http"
	"pcr-guild/components/logger"
	models "pcr-guild/models"

	"github.com/gin-gonic/gin"
	"github.com/qbox/qvm-base/components/mysql"
)

type _Boss struct{}

var BossController *_Boss

// DescribeBossResponse 短信模板返回
type DescribeBossResponse struct {
	Items []*models.Boss `json:"items"`
}

func (_ *_Boss) Index(c *gin.Context) {
	var (
		log      = logger.New(c)
		database = mysql.GetBiz(log.ReqID())
	)

	boss := []*models.Boss{}

	if err := database.Limit(1000).Find(&boss).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &DescribeBossResponse{
		Items: boss,
	})
}

func (_ *_Boss) Create(c *gin.Context) {
	var (
		log      = logger.New(c)
		database = mysql.GetBiz(log.ReqID())
		args     = []models.Boss{}
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

func (_ *_Boss) Update(c *gin.Context) {
	var (
		log        = logger.New(c)
		id         = c.Param("id")
		updateData = map[string]interface{}{}
		boss       = models.Boss{}
	)

	if err := c.BindJSON(&updateData); err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := mysql.GetBiz(log.ReqID()).Model(&boss).Select([]string{
		"battle_id",
		"battle",
		"boss_name",
		"cycle",
		"boss_num",
		"boss_stage",
		"total_hp",
		"current_hp",
		"remain_hp",
	})
	if err := database.Where("id = ?", id).Updates(updateData).Error; err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
