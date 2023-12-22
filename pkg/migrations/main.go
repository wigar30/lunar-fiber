package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

func main() {
	method := flag.String("method", "d", "migration method can be up or down")
	flag.Parse()

	config := viper.New()

	config.SetConfigName(".env")
	config.SetConfigType("dotenv")
	config.AddConfigPath("./")
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", config.GetString("DB_USERNAME"), config.GetString("DB_PASSWORD"), config.GetString("DB_HOST"), config.GetString("DB_PORT"), config.GetString("DB_DATABASE"))

	db, _ := sql.Open("mysql", conn)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/migrations/query",
		"mysql",
		driver)
	if err != nil {
		log.Fatal(err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		version = 0
		dirty = false
	}

	if dirty {
		m.Force(int(version))
	}

	var methodName string
	switch *method {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal(err)
			m.Force(int(version))
		}
		methodName = "up"
	case "down":
		if err := m.Steps(-1); err != nil {
			log.Fatal(err)
			m.Force(int(version))
		}
		methodName = "down"
	case "fresh":
		if err := m.Down(); err != nil {
			log.Fatal(err)
			m.Force(int(version))
		}
		if err := m.Up(); err != nil {
			log.Fatal(err)
			m.Force(int(version))
		}
		methodName = "fresh"
	default:
		log.Fatal("method not found")
		return
	}
	log.Printf("Successful database %s", methodName)
}
