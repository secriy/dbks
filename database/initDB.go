package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"server/util"
)

var DB *sql.DB

// 初始化数据库
func InitDB(dbConf string) {
	var err error
	DB, err = sql.Open("mysql", dbConf)
	if err != nil {
		log.Panicln("Err:", err.Error())
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	createDB()
	defaultAdmin()
}

// 创建表
func createDB() {
	// 用户表
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS user(
			id MEDIUMINT(8) UNSIGNED  AUTO_INCREMENT,
			username VARCHAR(20) NOT NULL UNIQUE,
			password VARCHAR(20) NOT NULL,
			authority TINYINT NOT NULL,
			create_at TIMESTAMP NOT NULL,
			PRIMARY KEY(id)
			);`)
	if err != nil {
		util.Log().Panic("创建用户表失败", err)
	}

	// 基础数据表
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS basic_info(
			name VARCHAR(20) NOT NULL UNIQUE,
			address VARCHAR(50) NOT NULL UNIQUE,
			department VARCHAR(100) NOT NULL UNIQUE,
			phone VARCHAR(13) NOT NULL UNIQUE,
			email VARCHAR(20) NOT NULL UNIQUE,
			url VARCHAR(200) NOT NULL UNIQUE,
			create_at TIMESTAMP NOT NULL,
			PRIMARY KEY(name)
			);`)
	if err != nil {
		util.Log().Panic("创建基础数据表失败", err)
	}

	// 新闻表
	_, err = DB.Exec(`
			CREATE TABLE IF NOT EXISTS news(
			id MEDIUMINT(8) UNSIGNED  AUTO_INCREMENT,
			title VARCHAR(20) NOT NULL,
			content VARCHAR(255) NOT NULL ,
			create_at TIMESTAMP NOT NULL,
			PRIMARY KEY(id)
			);`)
	if err != nil {
		util.Log().Panic("创建新闻表失败", err)
	}

	// 产品宣传表
	_, err = DB.Exec(`
			CREATE TABLE IF NOT EXISTS products(
			id MEDIUMINT(8) UNSIGNED  AUTO_INCREMENT,
			title VARCHAR(20) NOT NULL,
			content VARCHAR(255) NOT NULL ,
			create_at TIMESTAMP NOT NULL,
			PRIMARY KEY(id)
			);`)
	if err != nil {
		util.Log().Panic("创建产品宣传表失败", err)
	}

	// 企业招聘表
	_, err = DB.Exec(`
			CREATE TABLE IF NOT EXISTS offers(
			id MEDIUMINT(8) UNSIGNED  AUTO_INCREMENT,
			title VARCHAR(255) NOT NULL,
			content VARCHAR(255) NOT NULL ,
			create_at TIMESTAMP NOT NULL,
			PRIMARY KEY(id)
			);`)
	if err != nil {
		util.Log().Panic("创建企业招聘表失败", err)
	}
}

// 创建默认管理员
func defaultAdmin() {
	// 查询是否已经存在默认管理员
	var username string
	row := DB.QueryRow(`
			SELECT username FROM user WHERE username = ?;`, "admin")
	_ = row.Scan(&username)
	// 不存在则创建默认管理员
	if username == "" {
		_, err := DB.Exec(`INSERT INTO dbks.user(username,password,authority,create_at )
		VALUES (?,?,?,?)`, os.Getenv("ADMIN"), os.Getenv("ADMIN_PASSWORD"), 1, time.Now())
		if err != nil {
			util.Log().Panic("创建默认管理员失败", err)
		}

	}
}
