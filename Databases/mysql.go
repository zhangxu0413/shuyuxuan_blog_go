package Mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type Config struct {
	Mysql DatabaseConfig `yaml:"mysql"`
}
var DB *gorm.DB

// GetDB 初始化数据库
func GetDB() *gorm.DB  {
	var conf Config
	conf.getConf()
	databaseConfig := conf.Mysql
	databaseConfigStr := databaseConfig.User + ":" + databaseConfig.Password + "@tcp(" + databaseConfig.Host + ":" + databaseConfig.Port + ")/" + databaseConfig.Name
	db, err := gorm.Open("mysql", databaseConfigStr)
	db.SingularTable(true)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if db.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	return db
}

//读取Yaml配置文件,并转换成conf对象
func (c *Config) getConf() *Config {
	var confUrl string
	if os.Getenv("GO_ENV") == "release"  {
		confUrl = "/build/config/config.release.yaml"
	} else {
		confUrl = "./Config/config.dev.yaml"
	}
	//应该是 绝对地址
	yamlFile, err :=ioutil.ReadFile(confUrl)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		fmt.Println(err.Error())
	}

	return c
}