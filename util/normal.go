package util

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"time"
)

func GetCpuPercent() float64 {
	percent, _:= cpu.Percent(time.Second, false)
	var total float64
	var i int
	for _, v := range percent{
		total += v
		i++
	}
	return SetFloat(total/float64(i), 64)
}
func GetIoWait() uint64{
	infos, _ := cpu.Times(true)
	if infos != nil{
		// data, _ := json.MarshalIndent(infos[0], "", " ")
		return uint64(infos[0].Iowait)
	}
	return 0
	/*for _, info := range infos {
		data, _ := json.MarshalIndent(info, "", " ")
		fmt.Print(string(data))
	}*/
}
func GetPlatform() map[string]string{
	version, _ := host.KernelVersion()
	// fmt.Println(version)
	// var platform map[string]string
	platform := make(map[string]string)
	pf, family, version, _ := host.PlatformInformation()
	platform["platform"] = pf
	platform["version"] = version
	platform["family"] = family
	return platform
}
func GetMemPercent()float64 {
	memInfo, _ := mem.VirtualMemory()
	return SetFloat(memInfo.UsedPercent, 64)
}
