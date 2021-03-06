package controller

import (
	"net/http"
	"pcr-guild/components/logger"
	models "pcr-guild/models"

	"github.com/gin-gonic/gin"
	"github.com/qbox/qvm-base/components/mysql"
)

type _Plan struct{}

var PlanController *_Plan

// DescribePlanResponse ...
type DescribePlanResponse struct {
	Items []*models.Plan `json:"items"`
}

func (_ *_Plan) Index(c *gin.Context) {
	var (
		log      = logger.New(c)
		database = mysql.GetBiz(log.ReqID())
	)

	plans := []*models.Plan{}

	// 增加阵容关联关系
	database = database.Preload("TeamInfo")
	database = database.Preload("TeamInfo").Preload("TeamInfo.Characters")
	database = database.Preload("TeamInfo").Preload("TeamInfo.Characters").Preload("TeamInfo.Characters.CharacterInfo")

	// 增加boss关联关系
	database = database.Preload("BossInfo")

	// 增加阵容关联关系
	database = database.Preload("AssistInfo")
	database = database.Preload("AssistInfo").Preload("AssistInfo.CharacterInfo")

	if err := database.Limit(1000).Find(&plans).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &DescribePlanResponse{
		Items: plans,
	})
}

func (_ *_Plan) Create(c *gin.Context) {
	var (
		log      = logger.New(c)
		database = mysql.GetBiz(log.ReqID())
		args     = []models.Plan{}
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

func (_ *_Plan) Update(c *gin.Context) {
	var (
		log        = logger.New(c)
		id         = c.Param("id")
		updateData = map[string]interface{}{}
		plan       = models.Plan{}
	)

	if err := c.BindJSON(&updateData); err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := mysql.GetBiz(log.ReqID()).Model(&plan).Select([]string{
		"battle_id",
		"cycle",
		"boss_num",
		"boss_stage",
		"damage",
		"day",
		"user_id",
		"team_id",
		"boss_id",
		"assist_id",
		"is_end",
		"is_continue",
	})
	if err := database.Where("id = ?", id).Updates(updateData).Error; err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
