package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sja-boiler-plate/config"

	_ "github.com/go-sql-driver/mysql" //get SQL driver
	_ "github.com/joho/godotenv/autoload"
)

var db *gorm.DB

func Init() {
	c := config.GetConfig()

	username := c.GetString("db.user")
	password := c.GetString("db.pass")
	dbName := c.GetString("db.name")
	dbHost := c.GetString("db.host")
	dbPort := c.GetString("db.port")
	charset := c.GetString("db.charset")
	parseTime := c.GetString("db.parse_time")
	timezone := c.GetString("db.timezone")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&&parseTime=%s&&loc=%s", username, password, dbHost, dbPort, dbName, charset, parseTime, timezone)
	conn, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})

	db = conn
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database on " + dbHost)

}

// GetDB .
func GetDB() *gorm.DB {
	return db
}
