package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"database/sql"

	_ "github.com/lib/pq"
)

var MPosDB *sql.DB
var MPosGORM *gorm.DB

var err error

func InitGormPostgres(){
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	port := os.Getenv("PORT")
	dbname := os.Getenv("DBNAME")
	host := os.Getenv("HOST")

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	MPosGORM, err = gorm.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}
}