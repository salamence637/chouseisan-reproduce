package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	// This function will run immediately after backend is started, used to initialize database connection

	var err error

	// Connect to mysql database
	DBMS := "mysql"
	// USER := "root"
	USER := "user"
	// PASS := ""
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "chouseisan"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}

	// check if database connection is correctly established
	if DB == nil {
		log.Println("ERROR: db is nil!")
	}

	// Migrate the schema
	DB.AutoMigrate(&Event{})
}

// func InitDB() {
// 	// Initialize database connection
// 	cfg := mysql.Config{
// 		User:                 "user",
// 		Passwd:               "password",
// 		Net:                  "tcp",
// 		Addr:                 "mysql:3306",
// 		DBName:               "chouseisan",
// 		AllowNativePasswords: true,
// 	}

// 	var err error
// 	DB, err = sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Test the connection
// 	if err := DB.Ping(); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Connected!")
// }

//a
