package Mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

// Init 初始化数据库
func Init()  {
	var conf Config
	conf.getConf()
	databaseConfig := conf.Mysql
	databaseConfigStr := databaseConfig.User + ":" + databaseConfig.Password + "@tcp(" + databaseConfig.Host + ":" + databaseConfig.Port + ")/" + databaseConfig.Name
	var err error
	DB, err = gorm.Open("mysql", databaseConfigStr)
	DB.SingularTable(true)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
	defer DB.Close()

	//if !DB.HasTable(&Like{}) {
	//	if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
	//		panic(err)
	//	}
	//}
	//DB.AutoMigrate(&Like {})
	//l1 := &Like{2, "192.168.150.18","Test", "测试",1231,time.Now() }
	//DB.Create(&l1)
}

//读取Yaml配置文件,并转换成conf对象
func (c *Config) getConf() *Config {
	//应该是 绝对地址
	yamlFile, err :=ioutil.ReadFile("./Config/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		fmt.Println(err.Error())
	}

	return c
}