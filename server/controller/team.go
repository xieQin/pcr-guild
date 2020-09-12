package controller

import (
	"net/http"
	"pcr-guild/components/logger"
	models "pcr-guild/models"

	"github.com/gin-gonic/gin"

	"github.com/qbox/qvm-base/components/mysql"
)

type _Team struct{}

var TeamController *_Team

type DescribeTeamsResponse struct {
	Items []*models.Team `json:"items"`
}

type DescribeTeamCharactersResponse struct {
	Items []*models.TeamCharacter `json:"items"`
}

// type CreateTeamCharacterArg struct {
// 	Level           int  `json:"level"`      // 等级
// 	Rarity          int  `json:"rarity"`     // 星级
// 	Promotion       int  `json:"promotion"`  // rank
// 	LoveLevel       int  `json:"love_level"` // 好感度
// 	Slot1           bool `json:"slot1"`
// 	Slot2           bool `json:"slot2"`
// 	Slot3           bool `json:"slot3"`
// 	Slot4           bool `json:"slot4"`
// 	Slot5           bool `json:"slot5"`
// 	Slot6           bool `json:"slot6"`
// 	UniqueEquipRank int  `json:"unique_equip_rank"` // 专武等级
// 	TeamID          int  `json:"team_id"`           // 阵容ID
// 	CharacterID     int  `json:"character_id"`      // 角色ID
// }

func (_ *_Team) Index(c *gin.Context) {
	var (
		log      = logger.New(c)
		database = mysql.GetBiz(log.ReqID())
	)

	teams := []*models.Team{}
	// 增加角色关联关系
	database = database.Preload("Characters")

	// 增加角色关联关系
	database = database.Preload("Characters").Preload("Characters.CharacterInfo")

	if err := database.Limit(1000).Find(&teams).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &DescribeTeamsResponse{
		Items: teams,
	})
}

func (_ *_Team) IndexCharacter(c *gin.Context) {
	var (
		log      = logger.New(c)
		tid      = c.Param("tid")
		database = mysql.GetBiz(log.ReqID())
	)

	teamCharacters := []*models.TeamCharacter{}

	database = database.Preload("CharacterInfo")

	if err := database.Limit(1000).Where("team_characters.team_id = ?", tid).Find(&teamCharacters).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &DescribeTeamCharactersResponse{
		Items: teamCharacters,
	})
}

func (_ *_Team) Create(c *gin.Context) {
	var (
		log      = logger.New(c)
		database = mysql.GetBiz(log.ReqID())
		arg      = models.Team{}
	)

	if err := c.BindJSON(&arg); err != nil {
		log.Error("invalid args ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.Create(&arg).Error; err != nil {
		log.Error("create team failed ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, arg)
}

func (_ *_Team) Update(c *gin.Context) {
	var (
		log        = logger.New(c)
		tid        = c.Param("tid")
		updateData = map[string]interface{}{}
		team       = models.Team{}
	)

	if err := c.BindJSON(&updateData); err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := mysql.GetBiz(log.ReqID()).Model(&team).Select([]string{
		"boss_num",
		"boss_stage",
		"team_name",
		"target_damage",
		"context",
	})
	if err := database.Where("id = ?", tid).Updates(updateData).Error; err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (_ *_Team) CreateCharacter(c *gin.Context) {
	var (
		log      = logger.New(c)
		tid      = c.Param("tid")
		database = mysql.GetBiz(log.ReqID())
		args     = []models.TeamCharacter{}
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

	total := 0
	database = database.Model(&models.TeamCharacter{}).Where("team_characters.team_id = ?", tid).Count(&total)
	if total > (5 - len(args)) {
		log.Error("invalid team character count ")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid team character count"})
		return
	}

	for _, item := range args {
		if err := database.Create(&item).Error; err != nil {
			log.Error("create team character failed ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, args)
}

func (_ *_Team) UpdateCharacter(c *gin.Context) {
	var (
		log           = logger.New(c)
		id            = c.Param("id")
		updateData    = map[string]interface{}{}
		teamCharacter = models.TeamCharacter{}
	)

	if err := c.BindJSON(&updateData); err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := mysql.GetBiz(log.ReqID()).Model(&teamCharacter).Select([]string{
		"level",
		"rarity",
		"promotion",
		"love_level",
		"slot1",
		"slot2",
		"slot3",
		"slot4",
		"slot5",
		"slot6",
		"unique_equip_rank",
		"team_id",
		"character_id",
	})
	if err := database.Where("id = ?", id).Updates(updateData).Error; err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
