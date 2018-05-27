package app

import (
	"api/app/items"
	"api/app/gdrive"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"golang.org/x/net/context"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/drive/v3"
)

var (
	r *gin.Engine
)

const (
	port string = ":3000"
)

// StartApp ...
func StartApp() {
	r = gin.Default()
	db := configDataBase()
	items.Configure(r, db)
	gdriveService := configureGdriveService()
	gdrive.Configure(r, gdriveService)
	r.Run(port)
}

func configDataBase() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "user", "userpwd", "db", "db"))
	if err != nil {
		panic("Could not connect to the db")
	} 

	for {
		err := db.Ping()
		if err != nil {
			time.Sleep(1*time.Second)
			continue
		}
		// This is bad practice... You should create a schema.sql with all the definitions
		createTable(db)
		return db
	}

}

func createTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS items (
		id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
		name TEXT,
		description TEXT
	);`
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func configureGdriveService() *drive.Service {
	ctx := context.TODO()

    client, err := google.DefaultClient(ctx, drive.DriveScope)
    if err != nil {
		return nil
    }
	driveService, err := drive.New(client)
	if err != nil {
		return nil
	}
	return driveService
}
