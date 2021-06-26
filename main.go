package main

import (
	"encoding/json"
	"fmt"
	. "monitor/util"
	"time"
)

func main() {
	config := InitConfig("config.ini")
	timeTicker := time.NewTicker(time.Second * 10)
	i := 0
	for {

		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		i++
		<-timeTicker.C
		data := GetStatus(config["node"], config["sk"], fmt.Sprintf("%v", time.Now().Unix()))
		j, _ := json.Marshal(&data)
		url := fmt.Sprintf("http://%s/monitor/send/",config["proxy"] )
		fmt.Println(url)
		status, msg := SendHTTP(url, j, "POST")
		if status == 200{
		}else{
			fmt.Println(msg)
		}
		// data := make(map[string]string)
		fmt.Println(string(j))
	}
	// 清理计时器
	timeTicker.Stop()
}