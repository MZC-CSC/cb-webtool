package dragonfly

// 멀티 클라우드 인프라 VM 온디맨드 모니터링 정보 결과용
type VmMonitoringOnDemandInfoByDisk struct {
	NameSpaceID        string                          `json:"name"`
	VmMonitoringTags   VmMonitoringTagOnDemand         `json:"tags"`
	Time               string                          `json:"time"`
	VmMonitoringValues VmMonitoringValueOnDemandByDisk `json:"values"`
}

type VmMonitoringValueOnDemandByDisk struct {
	Free        float64 `json:"free"`
	ReadBytes   float64 `json:"read_bytes"`
	ReadTime    float64 `json:"read_time"`
	Reads       float64 `json:"reads"`
	Total       float64 `json:"total"`
	Used        float64 `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	WriteBytes  float64 `json:"write_bytes"`
	WriteTime   float64 `json:"write_time"`
	Writes      float64 `json:"writes"`
}
