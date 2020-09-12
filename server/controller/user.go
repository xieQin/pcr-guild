package controller

import (
	"net/http"
	"pcr-guild/components/logger"
	models "pcr-guild/models"

	"github.com/gin-gonic/gin"
	"github.com/qbox/qvm-base/components/mysql"
)

type _User struct{}

var UserController *_User

// DescribeUserResponse 短信模板返回
type DescribeUserResponse struct {
	Items []*models.User `json:"items"`
}

func (_ *_User) Index(c *gin.Context) {
	var (
		log      = logger.New(c)
		database = mysql.GetBiz(log.ReqID())
	)

	user := []*models.User{}

	if err := database.Limit(1000).Find(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &DescribeUserResponse{
		Items: user,
	})
}

func (_ *_User) Create(c *gin.Context) {
	var (
		log      = logger.New(c)
		database = mysql.GetBiz(log.ReqID())
		args     = []models.User{}
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

func (_ *_User) Update(c *gin.Context) {
	var (
		log        = logger.New(c)
		id         = c.Param("id")
		updateData = map[string]interface{}{}
		user       = models.User{}
	)

	if err := c.BindJSON(&updateData); err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := mysql.GetBiz(log.ReqID()).Model(&user).Select([]string{
		"user_name",
		"user_uid",
		"qq",
		"box_id",
	})
	if err := database.Where("id = ?", id).Updates(updateData).Error; err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
