package util

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)
type msg struct {
	code int    `json:"code"`
	Action  string `json:"action"`
	Msg  string `json:"msg"`
}
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
func f2s(f float64)string{
	return strconv.FormatFloat(f, 'g', 4, 64)
}
func SetFloat(a float64, n int)float64{
	f, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", a), n)
	return f
}
func SendHTTP(url string, data []byte, method string)(int, msg){
	//method = If(method=="GET" || method=="POST", method, "GET").(string)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	var msg msg
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return 400, msg
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, &msg)
	return resp.StatusCode, msg
	/*fmt.Println("response Status:", resp.Status)

		fmt.Println("response Headers:", resp.Header)
		fmt.Println("response Body:", string(body))
	*/

}
func InitConfig(path string) map[string]string {
	config := make(map[string]string)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}