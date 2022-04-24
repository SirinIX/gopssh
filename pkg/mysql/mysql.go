package mysql

import (
	"fmt"

	"cmd-scaffold/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLConfig struct {
	MysqlHost     string
	MysqlPort     int
	MysqlUsername string
	MysqlPassword string
	MysqlDatbase  string
}

func (c *MySQLConfig) NewMysqlConnection() (*sqlx.DB, error) {
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		c.MysqlUsername, c.MysqlPassword, c.MysqlHost, c.MysqlPort, c.MysqlDatbase)
	log.Info("connect to mysql: %s", addr)

	// Connect to MySQL
	db, err := sqlx.Connect("mysql", addr)
	trycount := 0
	if err != nil {
		return nil, err
	}
	for {
		if trycount >= 3 {
			break
		}
		err := db.Ping()
		if err != nil {
			trycount++
			continue
		}
		break
	}

	return db, err
}

func ExecSql(DB *sqlx.DB, execSql string) error {
	_, err := DB.Exec(execSql)
	if err != nil {
		return err
	}
	return nil
}
