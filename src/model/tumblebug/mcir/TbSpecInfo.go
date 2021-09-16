package mcir

type TbSpecInfo struct {
	AssociatedObjectList []string `json:"associatedObjectList"`

	ConnectionName string  `json:"connectionName"`
	CostPerHour    float32 `json:"costPerHour"`
	CspSpecName    string  `json:"cspSpecName"`
	Description    string  `json:"description"`
	EbsBwMbps      int     `json:"ebsBwMbps"`

	EvaluationScore01 float32 `json:"evaluationScore01"`
	EvaluationScore02 float32 `json:"evaluationScore02"`
	EvaluationScore03 float32 `json:"evaluationScore03"`
	EvaluationScore04 float32 `json:"evaluationScore04"`
	EvaluationScore05 float32 `json:"evaluationScore05"`
	EvaluationScore06 float32 `json:"evaluationScore06"`
	EvaluationScore07 float32 `json:"evaluationScore07"`
	EvaluationScore08 float32 `json:"evaluationScore08"`
	EvaluationScore09 float32 `json:"evaluationScore09"`
	EvaluationScore10 float32 `json:"evaluationScore10"`
	EvaluationStatus  string  `json:"evaluationStatus"`

	GpuModel           string `json:"gpuModel"`
	GpuP2p             string `json:"gpuP2p"`
	ID                 string `json:"id"`
	IsAutoGenerated    bool   `json:"isAutoGenerated"`
	MaxNumStorage      int    `json:"maxNumStorage"`
	MaxTotalStorageTiB int    `json:"maxTotalStorageTiB"`
	MemGiB             int    `json:"memGiB"`
	Name               string `json:"name"`
	Namespace          string `json:"namespace"`

	NetBwGbps             int    `json:"netBwGbps"`
	NumCore               int    `json:"numCore"`
	NumGpu                int    `json:"numGpu"`
	NumStorage            int    `json:"numStorage"`
	NumVCPU               int    `json:"numvCPU"`
	OrderInFilteredResult int    `json:"orderInFilteredResult"`
	OsType                string `json:"osType"`
	StorageGiB            int    `json:"storageGiB"`
}

type TbSpecInfos []TbSpecInfo