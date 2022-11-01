package app

import (
	"foodway/internal/cfg"
	"foodway/internal/delivery/http/authorization"
	"foodway/internal/delivery/http/register"
	"foodway/pkg/db"
	"github.com/gin-gonic/gin"
)

func Start() {
	cfg.LoadEnv()
	cfg.InitCfg()
	db.InitDb()
	db.AutoMigrate()

	r := gin.Default()

	r.Handle("POST", "/registration", register.Register)
	r.Handle("POST", "/authorization", authorization.Autho)

	r.Run(cfg.Cfg.IP + ":" + cfg.Cfg.Port)
}
