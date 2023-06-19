package bo

type MemberBo struct {
	UserID      string `json:"userId"`
	UserCode    string `json:"userCode"`
	UserName    string `json:"userName"`
	Level       int    `json:"level"`
	ExpValue    int64  `json:"expValue"`
	Age         int    `json:"age"`
	Gender      int    `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
}
