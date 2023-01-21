package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitConnectionDB() (*sql.DB, error) {

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("USER_DB"),
		os.Getenv("PASSWORD_DB"),
		os.Getenv("HOST_DB"),
		os.Getenv("NAME_DB"),
	)

	db, err := sql.Open("mysql", url)

	if err != nil {
		fmt.Println("ERROR IN CONNECTION DB")
		log.Panic(err)
	}
	fmt.Print("connection db successful")

	return db, err
}
