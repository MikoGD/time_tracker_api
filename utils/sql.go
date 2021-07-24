package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB;

func ConnectToDatabaase() {
	loadEnvFile()

	cfg := mysql.Config{
		Addr: "192.168.1.32:3306",
		Net: "tcp",
		User: getEnvVariable("USERNAME"),
		Passwd: getEnvVariable("PASSWORD"),
		DBName: getEnvVariable("DATABASE"),
		AllowNativePasswords: true,
	}

	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
}

type Statement struct {
	table string
	columns string
	conditions string
	values string
}

func (statement *Statement) CreateSelectStatement() string {
	if statement.conditions == "" {
		return fmt.Sprintf("SELECT %s FROM %s", statement.columns, statement.table)
	}

	return fmt.Sprintf("SELECT %s FROM %s WHERE %s", statement.columns, statement.table, statement.conditions)
}

func (statement *Statement) CreateInsertStatement() string {
	return fmt.Sprintf("INSERT INTO %s %s VALUES %s", statement.table, statement.columns, statement.values)
}

func (statement *Statement) CreateDeleteStatement() string {
	return fmt.Sprintf("DELETE FROM %s WHERE %s", statement.table, statement.conditions)
}

func (statement *Statement) CreateUpdateStatement() string {
	return fmt.Sprintf("UPDATE %s SET %s WHERE %s", statement.table, statement.columns, statement.conditions)
}

func SelectFromTable(table string, columns string, conditions string) (*sql.Rows, error) {
	statement := Statement{table, columns, conditions, ""}
	selectStatement := statement.CreateSelectStatement()

	rows, err := db.Query(selectStatement)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func InsertToTable(table string, columns string, values string) (int64, error) {
	statement := Statement{table, columns, "", values}
	insertStatement := statement.CreateInsertStatement()

	result, err := db.Exec(insertStatement)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteFromTable(table string, conditions string) (int64, error) {
	statement := Statement{table, "", conditions, ""}
	deleteStatement := statement.CreateDeleteStatement()

	result, err := db.Exec(deleteStatement)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func UpdateRowInTable(db *sql.DB, table string, columns string, conditions string) (int64, error) {
	statement := Statement{table, columns, conditions, ""}
	updateStatement := statement.CreateInsertStatement()

	result, err := db.Exec(updateStatement)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func loadEnvFile() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func getEnvVariable(key string) string {
	return os.Getenv(key)
}