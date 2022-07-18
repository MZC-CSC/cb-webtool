package dragonfly

// 멀티 클라우드 인프라 VM 온디맨드 모니터링 정보 결과용
type VmMonitoringOnDemandInfoByMemory struct {
	NameSpaceID        string                            `json:"name"`
	VmMonitoringTags   VmMonitoringTagOnDemand           `json:"tags"`
	Time               string                            `json:"time"`
	VmMonitoringValues VmMonitoringValueOnDemandByMemory `json:"values"`
}

type VmMonitoringValueOnDemandByMemory struct {
	MemCached      float64 `json:"mem_cached"`
	MemFree        float64 `json:"mem_free"`
	MemShared      float64 `json:"mem_shared"`
	MemTotal       float64 `json:"mem_total"`
	MemUsed        float64 `json:"mem_used"`
	MemUtilization float64 `json:"mem_utilization"`
}
