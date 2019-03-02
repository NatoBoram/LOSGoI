package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Database hosts the database configuration.
type Database struct {
	User     string
	Password string
	Address  string
	Port     int
	Database string
}

// IsEmpty returns `true` if one field from a `Database` object is empty.
func (database Database) IsEmpty() bool {
	return database.Address == "" ||
		database.Database == "" ||
		database.Password == "" ||
		database.Port == 0 ||
		database.User == ""
}

func (database *Database) read() error {

	// From File to JSON
	file, err := ioutil.ReadFile(databasePath)
	if err != nil {
		return err
	}

	// From JSON to Database
	err = json.Unmarshal(file, &database)
	if err != nil {
		return err
	}

	return nil
}

func (database Database) write() error {

	// Create required directories
	os.MkdirAll(rootFolder, permPrivateDirectory)

	// From Database to JSON
	json, err := json.Marshal(database)
	if err != nil {
		return err
	}

	// From JSON to File
	err = ioutil.WriteFile(databasePath, json, permPrivateFile)
	if err != nil {
		return err
	}

	return nil
}

func (database Database) template() error {
	fmt.Println("Writing a new database configuration template...")
	return Database{
		User:     "LOSGoI",
		Address:  "localhost",
		Port:     3306,
		Database: "LOSGoI",
	}.write()
}

func (database Database) string() string {
	return database.User + ":" + database.Password + "@tcp(" + database.Address + ":" + strconv.Itoa(database.Port) + ")/" + database.Database + "?parseTime=true"
}

func selectLatest() (rows *sql.Rows, err error) {
	return db.Query("SELECT `device`, `date`, `datetime`, `filename`, `filepath`, `sha1`, `sha256`, `size`, `type`, `version`, `ipfs` FROM `builds_latest`;")
}

func selectOld() (rows *sql.Rows, err error) {
	return db.Query("SELECT `device`, `date`, `datetime`, `filename`, `filepath`, `sha1`, `sha256`, `size`, `type`, `version`, `ipfs` FROM `builds` WHERE `ipfs` NOT IN( SELECT `ipfs` FROM `builds_latest` );")
}
