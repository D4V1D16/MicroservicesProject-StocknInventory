package DB

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=david password=abc123 dbname=ecommerceinventory port=4500"
var DB *gorm.DB

func DBConnection(){
	var ERROR error;
	DB,ERROR = gorm.Open(postgres.Open(DSN),&gorm.Config{})
	if ERROR != nil {
		log.Fatal(ERROR)
	} else{
		log.Println("DB Connected")
	}
}