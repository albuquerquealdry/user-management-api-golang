package config

import (
	"fmt"
	"log"
	"os"
	"user-management/src/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

func LoadDBConfig() (*DBConfig, error) {
	_ = godotenv.Load("../../.env")

	dbConfig := &DBConfig{
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}
	fmt.Println(os.Getenv("DB_USER"))
	if dbConfig.User == "" || dbConfig.Pass == "" {
		return nil, fmt.Errorf("variáveis de ambiente do banco não configuradas corretamente")
	}

	return dbConfig, nil
}

func ConnectDatabase() {

	dbConfig, err := LoadDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(dbConfig.User)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("[Connection Database Error]", err)
		os.Exit(1)
		return
	}
	fmt.Println("Connect to mysql database")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Error during migration:", err)
	}
}

func IsDatabaseReady() bool {
	sqlDB, err := DB.DB()
	if err != nil {
		return false
	}
	if err := sqlDB.Ping(); err != nil {
		return false
	}
	return true
}
