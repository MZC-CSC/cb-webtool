package dbmodel

type NamespaceInfo struct {
	NsId   string `json:"nsId" db:"ns_id"`
	NsName string `json:"nsName" db:"ns_name"`
	//OwnerNo uuid.UUID `json:"owner_no" db:"owner_no"`
	OwnerNo *uint64 `json:"ownerNo" db:"owner_no"`
}
