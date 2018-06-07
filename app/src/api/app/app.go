package app

import (
	"api/app/items"
	"api/app/documents"
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
	documents.Configure(r, db)
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
		loadSchemaDB(db)
		return db
	}

}

func loadSchemaDB(db *sql.DB) {
	schema, err := os.Open("/go/src/api/schema.sql")
	if err != nil {
		panic(err)
	}
	defer schema.Close()
	file, err := ioutil.ReadAll(schema)
	if err != nil {
	    panic(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			if err.Error() == "Error 1065: Query was empty" {
				fmt.Printf("a request is empty \n")
			} else {
				panic(err)
			}
	    }
	}
}
