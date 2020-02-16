package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/godror/godror"
)

var (
	db  *sql.DB //database
	ctx context.Context
)

func init() {

	e := godotenv.Load() //Load .env file

	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("%s/%s@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=tcp)(HOST=%s)(PORT=%s)))(CONNECT_DATA=(SID=%s)))", username, password, dbHost, dbPort,dbName)
	
	conn, err := sql.Open("godror", dbUri)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer db.Close()
	db = conn

}

//returns a handle to the DB object
func GetDB() *sql.DB {
	return db
}
func GetCTX() context.Context {
	return ctx
}
