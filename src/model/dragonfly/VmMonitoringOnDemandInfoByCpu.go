package dragonfly

// 멀티 클라우드 인프라 VM 온디맨드 모니터링 정보 결과용
type VmMonitoringOnDemandInfoByCpu struct {
	NameSpaceID        string                         `json:"name"`
	VmMonitoringTags   VmMonitoringTagOnDemand        `json:"tags"`
	Time               string                         `json:"time"`
	VmMonitoringValues VmMonitoringValueOnDemandByCpu `json:"values"`
}

type VmMonitoringValueOnDemandByCpu struct {
	CpuGuest       float64 `json:"cpu_guest"`
	CpuGuestNice   float64 `json:"cpu_guest_nice"`
	CpuHintr       float64 `json:"cpu_hintr"`
	CpuIdle        float64 `json:"cpu_idle"`
	CpuIowait      float64 `json:"cpu_iowait"`
	CpuNice        float64 `json:"cpu_nice"`
	CpuSintr       float64 `json:"cpu_sintr"`
	CpuSteal       float64 `json:"cpu_steal"`
	CpuSystem      float64 `json:"cpu_system"`
	CpuUser        float64 `json:"cpu_user"`
	CpuUtilization float64 `json:"cpu_utilization"`
}
