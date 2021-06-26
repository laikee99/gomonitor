package util
type reportData struct {
	Node string `json:"node"`
	Sk string `json:"sk"`
	Now string `json:"now"`
	Cpu float64 `json:"cpu"`
	Mem float64 `json:"mem"`
	IoWait uint64 `json:"ioWait"`
	Disk float64 `json:"disk"`
	DiskSpace uint64 `json:"diskSpace"`
	Net updownInfo `json:"net"`
	NetInfo map[string]netInfo `json:"netInfo"`
	IoInfo ioInfo `json:"ioInfo"`
}
func GetStatus(node string, sk string, now string)reportData{
	cpu := GetCpuPercent()
	mem := GetMemPercent()
	ioWait := GetIoWait()
	disk, _, diskSpace := GetDiskPercent()
	udi, netInfo := GetNetInfo()
	ioInfo := GetIoInfo()
	return reportData{
		Node: node,
		Sk: sk,
		Now: now,
		Cpu: cpu,
		Mem: mem,
		IoWait: ioWait,
		Disk: disk,
		DiskSpace: diskSpace,
		NetInfo: netInfo,
		Net: udi,
		IoInfo: ioInfo,
	}
}