package main

import (
	"database/sql"
	"errors"
)

var db *sql.DB

// Errors
var (
	errHashEmptyFolder = errors.New("empty folder hash")
	errInvalidHash     = errors.New("invalid hash")
)
