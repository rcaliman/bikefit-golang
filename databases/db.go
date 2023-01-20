package databases

import (
	"bikefit/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBanco() {
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Panic("erro ao conectar ao banco de dados: ", err)
	}
	DB.AutoMigrate(&models.Medidas{})
	DB.AutoMigrate(&models.Mensagem{})
}
