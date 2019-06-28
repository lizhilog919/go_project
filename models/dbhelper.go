package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type DbController struct {
	beego.Controller
}

type User struct {
	Id       int
	Nick     string
	PhoneNum string
	Pwd      string
}

var db *sql.DB

func init() {
	var err error
	db, err = connectDb()
	if err != nil {
		fmt.Println("connect error: ", err)
		return
	}
	err = createTable(db)
	if err != nil {
		fmt.Println("create table error: ", err)
	}
}

func connectDb() (*sql.DB, error) {
	fmt.Println("start connect database..")
	sqlName := "root"
	sqlPwd := "Qw112358"
	mysqlHost := "129.28.175.174"
	mysqlPort := "3306"
	driverName := "mysql"
	driverUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", sqlName, sqlPwd, mysqlHost, mysqlPort, DATABASE)
	db, err := sql.Open(driverName, driverUrl)
	if err != nil {
		return nil, err
	}
	db.Begin()
	return db, err
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(FUNC_CREATE_PHOTO_TABLE)
	if err != nil {
		return err
	}
	_, err = db.Exec(FUNC_CREATE_USER_TABLE)
	return err
}

const (
	DATABASE               = "photoalbum"
	FUNC_CREATE_USER_TABLE = "CREATE TABLE IF NOT EXISTS user (" +
		"id bigint(20) NOT NULL AUTO_INCREMENT, " +
		"phone char(20) NOT NULL," +
		"nick char(50) NOT NULL," +
		"pwd char(20) NOT NULL," +
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

func InsertUser(phoneNum string, nick string, pwd string) (int, error) {
	var count int64
	sql := fmt.Sprintf("SELECT count(*) FROM user where phone = ?")
	err := db.QueryRow(sql, phoneNum).Scan(&count)
	if err == nil && count == 0 { //查询不到会返回错误
		err = nil
		sql = fmt.Sprintf("INSERT INTO user (nick, phone, pwd) values (?, ?, ?)")
		_, err = db.Exec(sql, nick, phoneNum, pwd)
		if err == nil {
			sql = fmt.Sprintf("SELECT id FROM user where phone = ?")
			var id int
			err = db.QueryRow(sql, phoneNum).Scan(&id)
			if err == nil {
				fmt.Println(fmt.Sprintf("id: %d", id))
				return id, nil
			} else {
				return 0, errors.New("查询用户错误")
			}
		} else {
			return 0, errors.New("新建用户失败")
		}
	} else {
		if count > 0 {
			return 0, errors.New(fmt.Sprintf("%s已被使用", phoneNum))
		} else {
			return 0, errors.New("查询用户错误")
		}
	}
}

func Login(phoneNum string, pwd string) (*User, error) {
	sql := fmt.Sprintf("SELECT id,nick,pwd FROM user where phone = ?")
	rows, err := db.Query(sql, phoneNum)
	if err != nil {
		fmt.Println("err:" + err.Error())
		return nil, errors.New("登录错误")
	}
	defer rows.Close()
	var user User
	if !rows.Next() {
		return nil, errors.New("用户不存在")
	}
	err = rows.Scan(&user.Id, &user.Nick, &user.Pwd)
	if err != nil {
		return nil, err
	}
	if user.Pwd == pwd {
		return &user, nil
	}
	return nil, errors.New("密码错误")
}
