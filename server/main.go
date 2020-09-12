package main

import (
	"flag"
	"fmt"
	"os"
	"pcr-guild/components/config"
	"pcr-guild/controller"
	"pcr-guild/models"

	"github.com/gin-gonic/gin"
	"github.com/qbox/qvm-base/components/mysql"
)

// InitConfig 初始化配置
func InitConfig(srcPath string) (err error) {
	config.SetConfigType("yaml")

	f, err := os.Open(srcPath)
	if err != nil {
		return
	}

	defer f.Close()

	err = config.ReadConfig(f)
	if err != nil {
		return
	}

	return
}

func init() {
	var confPath string
	flag.StringVar(&confPath, "f", "", "-f=/path/to/config")
	flag.Parse()

	err := InitConfig(confPath)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	// 初始化数据库配置
	var mysqlConfigs []*mysql.ConfigStructure
	config.UnmarshalKey("mysqls", &mysqlConfigs)
	mysql.Init(mysqlConfigs...)

	mysql.AutoMigrateBiz(
		&models.Character{},
		&models.TeamCharacter{},
		&models.Team{},
	)
}

func main() {
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	api := engine.Group("/api")

	api.GET("/team", controller.TeamController.Index)
	api.POST("/team", controller.TeamController.Create)
	api.GET("/team/:tid/character", controller.TeamController.IndexCharacter)
	api.POST("/team/:tid/character", controller.TeamController.CreateCharacter)

	api.GET("/character", controller.CharacterController.Index)
	api.POST("/character", controller.CharacterController.Create)
	api.PUT("/character/:id", controller.CharacterController.Update)

	engine.Run(":9099")
}
