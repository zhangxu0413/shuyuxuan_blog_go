package Utils
import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
)
// HttpGet Get 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func HttpGet(url string) map[string]interface{} {
	resp,_ := http.Get(url)
	defer resp.Body.Close()
	return Transformation(resp)
}

func Transformation(response *http.Response)  map[string]interface{}{
	var result map[string]interface{}
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		json.Unmarshal([]byte(string(body)), &result)
	}
	return result
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
