package db

import (
	"fmt"
	"log"
	"reflect"
	"sync"

	"github.com/bagusandrian/mini-api/src/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	conf     *config.Config
	mtxDBMap = &sync.Mutex{}
	dbMap    map[string]*sqlx.DB
	err      error
)

const (
	DriverPostgres = "postgres"
	DriverSQLMock  = "sqlmock"
)

func Init(config *config.Config) {

	if config == nil {
		log.Fatalln("Postgre Init Failed! Exiting...")
	}

	conf = config
	dbMap = make(map[string]*sqlx.DB)
	return
}

func Get(conn string) *sqlx.DB {
	mtxDBMap.Lock()
	if dbMap[conn] != nil {
		err = dbMap[conn].Ping()
		if err == nil {
			mtxDBMap.Unlock()
			return dbMap[conn]
		}
		log.Println("[mini-api][db] dbMap connection:", conn, "disconnected. Creating new connection..", err)
	}
	mtxDBMap.Unlock()

	strConn := fmt.Sprint(reflect.ValueOf(conf.Database).FieldByName(conn))
	if strConn == "" {
		log.Println("[mini-api][db] no dbMap conn", conn)
		return nil
	}

	var db *sqlx.DB
	strDriver := DriverPostgres
	if conf.Environment == "testing" {
		strDriver = DriverSQLMock
	}

	db, err = sqlx.Connect(strDriver, strConn)
	if err != nil {
		log.Printf("[db] problem establishing connection sqlx.Connect:%+v\n", err)
		return nil
	}

	db.SetMaxIdleConns(conf.Database.MaxIdleConn)
	db.SetMaxOpenConns(conf.Database.MaxOpenConn)

	mtxDBMap.Lock()
	dbMap[conn] = db
	mtxDBMap.Unlock()

	return dbMap[conn]
}

func PutMock(conn string, mockDb *sqlx.DB) {
	if len(dbMap) == 0 {
		dbMap = make(map[string]*sqlx.DB)
	}
	mtxDBMap.Lock()
	dbMap[conn] = mockDb
	mtxDBMap.Unlock()
}
