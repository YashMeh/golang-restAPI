package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/m/v2/db"
	"example.com/m/v2/store"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func viperEnvVariable(key string) string {

	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

func main() {
	//Initialised the database
	initDb()
	//Creating the table
	db.DBCon.Query("create table if not exists books(id serial,isbn varchar(8),title text not null,price NUMBERIC(10,2) NOT NULL DEFAULT 0.00,CONSTRAINT books_pkey PRIMARY KEY (id)")
	//Initialised the mux router
	r := store.NewRouter()

	//Enabling CORS
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	//Adding CORS and logging to the router
	corsEnabledRouter := handlers.CORS(allowedOrigins, allowedMethods)(r)
	loggedRouter := handlers.LoggingHandler(os.Stdout, corsEnabledRouter)

	log.Fatal(http.ListenAndServe(":8000", loggedRouter))

}
func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db.DBCon, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.DBCon.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database successfully connected!")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)

	conf[dbhost] = viperEnvVariable(dbhost)
	conf[dbport] = viperEnvVariable(dbport)
	conf[dbuser] = viperEnvVariable(dbuser)
	conf[dbpass] = viperEnvVariable(dbpass)
	conf[dbname] = viperEnvVariable(dbname)
	return conf
}
