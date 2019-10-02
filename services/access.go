package services

import (
	"fmt"
	"io/ioutil"
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v2"
	_"github.com/go-sql-driver/mysql"
)

//DBAccess ->
type DBAccess struct {
	SQLDB *sqlx.DB
}

// SQLDbAccess ->
var SQLDbAccess DBAccess

// DBConf ->
type DBConf struct {
	Name   string `yaml:"name"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Adress string `yaml:"dbIdbadress"`
	Port   string `yaml:"dbPort"`
}

// Configs ->
type Configs map[string]DBConf

// GetDBAccess ...
func GetDBAccess(env, path string) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("error opening configuration", err.Error())
	}

	var cs Configs

	err = yaml.Unmarshal(data, &cs)
	dbconf := cs[env]

	dbDriver := "mysql"
	dbUser := dbconf.User
	dbPass := dbconf.Pass
	dbName := dbconf.Name
	dbAdress := dbconf.Adress
	dbPort := dbconf.Port
	db, err := sqlx.Connect(dbDriver, dbUser+":"+dbPass+"@"+"tcp("+dbAdress+":"+dbPort+")/"+dbName)

	if err != nil {
		fmt.Println(err)
	}

	SQLDbAccess.SQLDB = db

	return err
}

// GetSQLDB ->
func (a DBAccess) GetSQLDB() *sqlx.DB {
	return a.SQLDB
}