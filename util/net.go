package util

import "github.com/shirou/gopsutil/net"

type updownInfo struct {
	Download uint64 `json:"download"`
	Upload  uint64 `json:"upload"`
}
type netInfo struct {
	Upload uint64 `json:"upload"`
	Download uint64 `json:"download"`
	TotalUpload uint64 `json:"totalUpload"`
	TotalDownload uint64 `json:"totalDownload"`
}
var lastNet map[string]netInfo
func GetLastNet(name string) netInfo{
	val, ok := lastNet[name]
	if  ok {
		return val
	}
	return netInfo{
		Upload: uint64(0),
		Download: uint64(0),
		TotalDownload: uint64(0),
		TotalUpload: uint64(0),
	}
}
func GetNetInfo() (updownInfo, map[string]netInfo) {
	// var netArr []netInfo
	netArr := make(map[string]netInfo)
	nn, _ := net.IOCounters(true)
	var totalDownload uint64;
	var totalUpload uint64;
	// fmt.Println(nn)
	for _, v := range nn{
		last := GetLastNet(v.Name)
		netArr[v.Name] = netInfo{
			// name:          v.Name,
			TotalDownload: v.BytesRecv,
			TotalUpload:   v.BytesSent,
			Upload: If(last.TotalUpload > 0, v.BytesSent-last.TotalUpload, uint64(0)).(uint64),
			Download: If(last.TotalDownload > 0, v.BytesRecv-last.TotalDownload, uint64(0)).(uint64),
		}
		totalDownload += netArr[v.Name].Upload
		totalUpload += netArr[v.Name].Download
	}
	lastNet = netArr
	udi := updownInfo{
		Download: totalDownload,
		Upload: totalUpload,
	}
	return udi, netArr
}