package http

type PageInfo struct {
	Page     int64 `json:"page,omitempty"`
	PageSize int64 `json:"page_size,omitempty"`
	Total    int64 `json:"total,omitempty"`
}
