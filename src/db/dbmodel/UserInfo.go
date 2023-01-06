package dbmodel

type UserInfo struct {
	//MemberNo uuid.UUID `json:"memberNo" db:"member_no"`
	//MemberNo *uint64 `json:"memberNo" db:"member_no"`
	Id       string `json:"id" db:"id"`
	Passwd   string `json:"passwd" db:"passwd"`
	UserName string `json:"userName" db:"user_name"`
}

type ReqUserInfo struct {
	//MemberNo uuid.UUID `json:"memberNo" db:"member_no"`
	Id       string `json:"id" db:"id"`
	Passwd   string `json:"passwd" db:"passwd"`
	UserName string `json:"userName" db:"user_name"`
}
