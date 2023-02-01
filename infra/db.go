package infra

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	constants "telkom-haioo/domain/constants/general"
	"telkom-haioo/domain/model/general"

	log "github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// IDatabase is interface for database
type Database interface {
	ConnectDB(dbConn *general.DBDetail)
	Close()

	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error

	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	// DriverName() string

	Begin() (*sql.Tx, error)
	In(query string, params ...interface{}) (string, []interface{}, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row
	// QueryRowSqlx(query string, args ...interface{}) *sqlx.Row
	// QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	// GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type DatabaseList struct {
	Backend DatabaseType
}

type DatabaseType struct {
	Read  Database
	Write Database
}

// DBHandler - Database struct.
type DBHandler struct {
	DB  *sqlx.DB
	Err error
	log *log.Logger
}

func NewDB(log *log.Logger) DBHandler {
	return DBHandler{
		log: log,
	}
}

// ConnectDB - function for connect DB.

func (d *DBHandler) ConnectDB(dbConn *general.DBDetail) {
	prdctDb := fmt.Sprintf("user=" + dbConn.Username + " password=" + dbConn.Password + " sslmode=" + dbConn.SSLMode + " dbname=" + dbConn.DBName + " host=" + dbConn.URL + " port=" + dbConn.Port + " connect_timeout=" + dbConn.Timeout)

	dataSource := fmt.Sprintf("%s", prdctDb)

	dbs, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		log.Error(constants.ConnectDBFail + " | " + err.Error())
		d.Err = err
	}

	d.DB = dbs

	err = d.DB.Ping()
	if err != nil {
		log.Error(constants.ConnectDBFail, err.Error())
		d.Err = err
	}

	d.log.Info(constants.ConnectDBSuccess + dbConn.DBName)
	// d.log.Info(dbs)
	// }
	d.DB.SetConnMaxLifetime(time.Duration(dbConn.MaxLifeTime))
}

// Close - function for connection lost.
func (d *DBHandler) Close() {
	if err := d.DB.Close(); err != nil {
		d.log.Println(constants.ClosingDBFailed + " | " + err.Error())
	} else {
		d.log.Println(constants.ClosingDBSuccess)
	}
}

func (d *DBHandler) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := d.DB.Exec(query, args...)
	return result, err
}

func (d *DBHandler) Query(query string, args ...interface{}) (*sql.Rows, error) {
	result, err := d.DB.Query(query, args...)
	return result, err
}

func (d *DBHandler) Select(dest interface{}, query string, args ...interface{}) error {
	err := d.DB.Select(dest, query, args...)
	return err
}

func (d *DBHandler) Get(dest interface{}, query string, args ...interface{}) error {
	err := d.DB.Get(dest, query, args...)
	return err
}

func (d *DBHandler) Rebind(query string) string {
	return d.DB.Rebind(query)
}

func (d *DBHandler) In(query string, params ...interface{}) (string, []interface{}, error) {
	query, args, err := sqlx.In(query, params...)
	return query, args, err
}

func (d *DBHandler) Begin() (*sql.Tx, error) {
	return d.DB.Begin()
}

func (d *DBHandler) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return d.DB.QueryRowContext(ctx, query, args...)
}

func (d *DBHandler) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	err := d.DB.GetContext(ctx, dest, query, args...)
	return err
}

func (d *DBHandler) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := d.DB.ExecContext(ctx, query, args...)
	return result, err
}
