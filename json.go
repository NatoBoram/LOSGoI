package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readDatabase(object *Database) error {

	// Read the JSON file
	file, err := ioutil.ReadFile(databasePath)
	if err != nil {
		return err
	}

	// Put the JSON in the object
	err = json.Unmarshal(file, &object)
	if err != nil {
		return err
	}

	return nil
}

func (database Database) write() error {

	// Create required directories
	os.MkdirAll(rootFolder, permPrivateDirectory)

	// From object to JSON
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

func writeTemplateDatabase() error {
	fmt.Println("Writing a new database configuration template...")
	return Database{
		User:     "LOSGoI",
		Address:  "localhost",
		Port:     3306,
		Database: "LOSGoI",
	}.write()
}
