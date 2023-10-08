package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *q