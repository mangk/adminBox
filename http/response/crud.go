package response

type CRUDResponse struct {
	Count     int64       `json:"count"`
	Page      int64         `json:"page"`
	PageCount int         `json:"pageCount"`
	PageSize  int64         `json:"pageSize"`
	List      interface{} `json:"list"`
}
