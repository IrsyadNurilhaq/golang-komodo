package database

import (
	"database/sql"
	"log"

	"github.com/spf13/viper"
)

func Connect() *sql.DB {

	connection := viper.GetString("database.connection")
	user := viper.GetString("database.username")
	pwd := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	name := viper.GetString("database.name")
	db, err := sql.Open(connection, user+":"+pwd+"@tcp("+host+":"+port+")/"+name+"?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
