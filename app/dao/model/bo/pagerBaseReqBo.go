package bo

type PagerBaseReqBo struct {
	IsPager  int `json:"isPager"`
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}
