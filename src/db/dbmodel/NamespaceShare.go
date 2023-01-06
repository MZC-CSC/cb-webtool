package dbmodel

type NamespaceShare struct {
	NsID string `json:"nsId" db:"ns_id"`
	//MemberNo uuid.UUID `json:"member_no" db:"member_no"`
	MemberNo *uint64 `json:"memberNo" db:"member_no"`
}
