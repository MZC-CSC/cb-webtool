package common

// TB의 common.ConfigInfo -> TBConfigInfo로   ConfigInfo가 spider에도 있으므로  common package는 모두 TB를 붙이도록 한다
type TbConfigInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
