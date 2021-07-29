package Utils
import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
)
// HttpGet Get 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func HttpGet(url string) string {
	resp,_ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
