package app

import (
	"foodway/internal/cfg"
	"foodway/internal/delivery/http/authorization"
	"foodway/internal/delivery/http/register"
	"github.com/gin-gonic/gin"
)

func Start() {
	cfg.LoadEnv()
	cfg.InitCfg()

	r := gin.Default()

	r.Handle("POST", "/registration", register.Register)
	r.Handle("POST", "/authorization", authorization.Autho)

	r.Run(cfg.Cfg.IP + ":" + cfg.Cfg.Port)
}
