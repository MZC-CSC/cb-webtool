package dragonfly

// 멀티 클라우드 인프라 VM 온디맨드 모니터링 정보 결과용
type VmMonitoringOnDemandInfoByNetwork struct {
	NameSpaceID        string                             `json:"name"`
	VmMonitoringTags   VmMonitoringTagOnDemand            `json:"tags"`
	Time               string                             `json:"time"`
	VmMonitoringValues VmMonitoringValueOnDemandByNetwork `json:"values"`
}

type VmMonitoringValueOnDemandByNetwork struct {
	BytesIn  float64 `json:"bytes_in"`
	BytesOut float64 `json:"bytes_out"`
	PktsIn   float64 `json:"pkts_in"`
	PktsOut  float64 `json:"pkts_out"`
}
