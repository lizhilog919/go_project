package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)
import _ "github.com/go-sql-driver/mysql"

type DatabaseController struct {
	beego.Controller
}

const (
	DATABASE               = "photoalbum"
	FUNC_CREATE_USER_TABLE = "CREATE TABLE IF NOT EXISTS user (" +
		"id bigint(20) NOT NULL AUTO_INCREMENT, " +
		"phone char(20) NOT NULL," +
		"nick char(50) NOT NULL," +
		"PRIMARY KEY (id)" +
		") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;"
	FUNC_CREATE_PHOTO_TABLE = "CREATE TABLE IF NOT EXISTS photo (" +
		"id bigint(20) NOT NULL AUTO_INCREMENT, " +
		"userId bigint(20) NOT NULL," +
		"photoAdr char(100) NOT NULL," +
		"uploadTime int(20) NOT NULL," +
		"PRIMARY KEY (id)" +
		") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;"
)

func init() {
	db, err := connectDb()
	if err != nil {
		fmt.Println("connect error: ", err)
		return
	}
	err = createTable(db)
	if err != nil {
		fmt.Println("create table error: ", err)
	}
}

func connectDb() (orm.Ormer, error) {
	fmt.Println("start connect database..")
	sqlName := "root"
	sqlPwd := "Qw112358"
	mysqlHost := "129.28.175.174"
	mysqlPort := "3306"
	hostUrl := fmt.Sprintf("tcp(%s:%s)", mysqlHost, mysqlPort)
	DriverUrl := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", sqlName, sqlPwd, hostUrl, DATABASE)
	err := orm.RegisterDataBase("default", "mysql", DriverUrl, 100)
	err = orm.RegisterDataBase(DATABASE, "mysql", DriverUrl, 100)
	if err != nil {
		return nil, err
	}
	db, err := GetDB()
	if err != nil {
		fmt.Println("use database fail")
		return db, err
	}
	err = db.Begin()
	return db, err
}

func GetDB() (orm.Ormer, error) {
	o := orm.NewOrm()
	err := o.Using(DATABASE)
	return o, err
}

func createTable(db orm.Ormer) error {
	_, err := db.Raw(FUNC_CREATE_PHOTO_TABLE).Exec()
	if err != nil {
		return err
	}
	_, err = db.Raw(FUNC_CREATE_USER_TABLE).Exec()
	return err
}
