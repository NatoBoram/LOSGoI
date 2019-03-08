package main

import (
	"database/sql"
	"errors"
)

var db *sql.DB

// Errors
var (
	errHashEmptyFolder = errors.New("this is the hash of an empty folder")
)
