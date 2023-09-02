package data

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func openDb(dbType string, host string, dbname string, username string, password string, port string) (*gorm.DB, error) {
	var dialector gorm.Dialector

	if dbType == "mssql" {
		dialector = openSqlServerDb(host, dbname, username, password, port)
	} else if dbType == "mysql" {
		dialector = openMysqlDb(host, dbname, username, password, port)
	} else if dbType == "pgsql" {
		dialector = openPostgresqlDb(host, dbname, username, password, port)
	} else if dbType == "sqlite" {
		dialector = openSqliteDb()
	}

	return gorm.Open(dialector, &gorm.Config{})
}

func openSqlServerDb(host string, dbname string, username string, password string, port string) gorm.Dialector {
	if port == "" {
		port = "1433"
	}
	return sqlserver.New(sqlserver.Config{
		DSN: fmt.Sprintf("Server=%s;Database=%s;User Id=%s;Password=%s;", host, dbname, username, password),
	})
}

func openMysqlDb(host string, dbname string, username string, password string, port string) gorm.Dialector {
	if port == "" {
		port = "3306"
	}
	return mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname),
	})
}

func openPostgresqlDb(host string, dbname string, username string, password string, port string) gorm.Dialector {
	if port == "" {
		port = "9920"
	}
	return postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port),
	})
}

func openSqliteDb() gorm.Dialector {
	return sqlite.Open("test.db")
}
