package util

import (
	"github.com/shirou/gopsutil/disk"
)
type ioInfo struct {
	Write uint64 `json:"write"`
	Read  uint64 `json:"read"`
}
var lastIo ioInfo
func GetLastIo() ioInfo{
	return lastIo
}
func GetIoInfo()ioInfo{
	last := GetLastIo()
	i, _ := disk.IOCounters()
	var read uint64
	var write uint64
	for _, v := range i{
		read += v.ReadBytes
		write += v.WriteBytes
	}
	// fmt.Println(read, write)
	info := ioInfo{
		Read:  If(last.Read > 0, read-last.Read, uint64(0)).(uint64),
		Write: If(last.Write > 0, write-last.Write, uint64(0)).(uint64),
	}
	lastIo = ioInfo{
		Write: write,
		Read: read,
	}
	return info
}
func GetDiskPercent() (float64, uint64, uint64) {
	parts, _ := disk.Partitions(true)
	var total uint64
	var used uint64
	var free uint64
	for _, v := range parts{
		diskInfo, _ := disk.Usage(v.Mountpoint)
		total += diskInfo.Total
		used += diskInfo.Used
		free += diskInfo.Free
	}
	return float64(total / used), used/1024/1024/1024, free/1024/1024/1024
}
