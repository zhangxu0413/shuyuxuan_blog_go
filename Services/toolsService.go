package Services

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"os"
	"shuyuxuan_blog_go/Types"
	"shuyuxuan_blog_go/Utils"
	"strconv"
)

type ToolsService struct {}
// TranslateConfig 百度翻译数据库配置
type TranslateConfig struct {
	Appid     string `yaml:"appid"`
	Secret    string `yaml:"secret"`
}
type Config struct {
	Translate TranslateConfig `yaml:"translate"`
}

func (d ToolsService) Translate(query Types.TranslateReq) map[string]interface{} {
	// 拉配置
	var conf Config
	conf.getConf()
	translateConf := conf.Translate
	salt := strconv.Itoa(rand.Intn(1e10-1e9) + 1e9)
	params := url.Values{}
	Url, _:= url.Parse("https://fanyi-api.baidu.com/api/trans/vip/translate")
	params.Set("appid",translateConf.Appid)
	params.Set("q",query.Query)
	params.Set("from",query.From)
	params.Set("to",query.To)
	params.Set("salt",salt)
	params.Set("sign",Utils.Md5(translateConf.Appid + query.Query + salt + translateConf.Secret))
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	log.Println(urlPath)
	res := Utils.HttpGet(urlPath)
	return res
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