package dragonfly

type VmMonitoringTagOnDemand struct {
	CspType string `json:"cspType"`
	OsType  string `json:"osType"`
	McisId  string `json:"mcisId"`
	NsId    string `json:"nsId"`
	VmId    string `json:"vmId"`
}
