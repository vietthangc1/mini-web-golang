package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetTestURI() string {
	return "root:Chaugn@rs2@tcp(127.0.0.1:3306)/mini_golang_project_v2_test?charset=utf8&parseTime=True&loc=Local"
}

func ConnectDatabaseORMTest(dropTable bool) (*gorm.DB, error) {
	uri := GetTestURI()

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if dropTable {
		dropDB(db)
	}

	migrateDB()

	return db, nil
}

func dropDB(db *gorm.DB) {
	var tables []string
	db.Raw("SHOW TABLES").Scan(&tables)

	if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0;").Error; err != nil {
		log.Fatalf(err.Error())
	}
	
	for _, table := range tables {
		if err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error; err != nil {
			log.Fatalf(err.Error())
		}
	}

	// if err := db.Exec("SET FOREIGN_KEY_CHECKS = 1;").Error; err != nil {
	// 	log.Fatalf(err.Error())
	// }
}

func migrateDB() {
	uri := GetTestURI()
	dsn := fmt.Sprintf(
		"mysql://%s",
		uri,
	)

	migrateDir := "/Users/lap02804/Documents/Data/golang-project/mini-web-golang/migrations/mysql"
	migrateDir = fmt.Sprintf("file://%s", migrateDir)

	m, err := migrate.New(
		migrateDir,
		dsn,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalln(err)
	}
}
