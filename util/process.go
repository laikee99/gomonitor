package util

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
)

type processData struct{
	name    string
	cpu     string
	mem     string
	ioRead  string
	ioWrite string
}
/*
 遍历进程，获取其CPU、内存和IO占用
*/
func proc() []processData {
	var data []processData= make([]processData, 0);
	//var s3 []int = make([]int,0)
	//index := 0;
	processes, _ := process.Processes()
	for _, p := range processes {
		c, _ := p.CPUPercent()
		i, _ := p.IOCounters()
		r := ""
		w := ""
		if i == nil{
		}else{
			r = fmt.Sprintf("%v", (*i).ReadBytes)
			w = fmt.Sprintf("%v", (*i).WriteBytes)
		}
		n, _ := p.Name()
		m, _ := p.MemoryPercent()
		// fmt.Sprintf(""+n+":"+cpu)
		//fmt.Println(i.ReadBytes)
		data = append(data, processData{
			name:    n,
			mem:     fmt.Sprintf("%.3f", m),
			cpu:     fmt.Sprintf("%.3f", c),
			ioRead:  r,
			ioWrite: w,
		})
	}
	for _, v := range data{
		fmt.Println(v)
	}
	return data
}
