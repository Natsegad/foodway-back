package cfg

import (
	"foodway/internal/domain"
	"foodway/pkg/logger"
	"github.com/joho/godotenv"
	"os"
)

var Cfg domain.Config

func LoadEnv() {
	log := logger.GetLogger()

	err := godotenv.Load()
	if err != nil {
		log.Errorf("Error load env ! %s", err)
		return
	}
	log.Info("ENV loaded ! ")
}

func InitCfg() {
	log := logger.GetLogger()

	port, exst := os.LookupEnv("PORT")
	if !exst {
		log.Errorf("Error get PORT")
	}

	id, exst := os.LookupEnv("ID")
	if !exst {
		log.Errorf("Error get ID")
	}

	Cfg = domain.Config{
		Port: port,
		IP:   id,
	}
}
